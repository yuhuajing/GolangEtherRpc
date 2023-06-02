package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

/*
*
Apply for Network API.
https://www.quicknode.com/

// in the main.go
DoSubscribe("Tx") -- Get all pending transactions
DoSubscribe("TxHash") -- Get all pending transaction hash
outPut example:
Pending:true of tx: {"type":"0x0","nonce":"0x4","gasPrice":"0x3b9aca00","maxPriorityFeePerGas":null,"maxFeePerGas":null,"gas":"0x7530","value":"0x19ebac6047ff","input":"0x","v":"0x26","r":"0x25ce543705f138288a7a8c728870fb9cda6c8e356a566900e1c800d1962f5fa3","s":"0xf45eaba2ad5ade4d41323bc073c9ebdd20c21a36cf046b6da1b686b6b84fa1a","to":"0xa03ae5eddc99e65dd7e1315d4b238e51f9baa6bc","hash":"0xa385817662a5008c6b742868ec37355138245ba013b4609c2ceaa533f98a70cd"}

	type Transaction struct {
		inner TxData    // Consensus contents of a transaction
		time  time.Time // Time first seen locally (spam avoidance)

		// caches
		hash atomic.Value
		size atomic.Value
		from atomic.Value
	}

	type TxData interface {
		txType() byte // returns the type ID
		copy() TxData // creates a deep copy and initializes all fields

		chainID() *big.Int
		accessList() AccessList
		data() []byte
		gas() uint64
		gasPrice() *big.Int
		gasTipCap() *big.Int
		gasFeeCap() *big.Int
		value() *big.Int
		nonce() uint64
		to() *common.Address

		rawSignatureValues() (v, r, s *big.Int)
		setSignatureValues(chainID, v, r, s *big.Int)

		// effectiveGasPrice computes the gas price paid by the transaction, given
		// the inclusion block baseFee.
		//
		// Unlike other TxData methods, the returned *big.Int should be an independent
		// copy of the computed value, i.e. callers are allowed to mutate the result.
		// Method implementations can use 'dst' to store the result.
		effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int
	}

*
*/
var (
	subgclient *gethclient.Client
	wss        string = "wss://cool-muddy-butterfly.discover.quiknode.pro/0e41f42d5a7c9611f30ef800444bfcb93d3ae9a6/"
)

func initClient() {
	//Retrieve pending blocks

	rpcCli, err := rpc.Dial(wss)
	if err != nil {
		log.Printf("rpcDial failed to dial: %v", err)
		return
	}
	//Reterive pending transactions
	subgclient = gethclient.New(rpcCli)
}

func moitorPendingTx() {
	pendingTx := make(chan *types.Transaction)
	sub, err := subgclient.SubscribeFullPendingTransactions(context.Background(), pendingTx)
	if err != nil {
		fmt.Printf("Error:%v\n", err)
		//log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			fmt.Printf("Error:%v\n", err)
			//log.Fatal(err)
		case tx := <-pendingTx:
			//fmt.Println(tx.To())
			//fmt.Println(hex.EncodeToString(tx.Data()))
			data, _ := tx.MarshalJSON()
			fmt.Println(string(data))
			/**
			{"type":"0x2","chainId":"0x1","nonce":"0x2","to":"0x417a5538c0af25ecea6a7eb87e66d553b34ad9ab","gas":"0x5208","gasPrice":null,"maxPriorityFeePerGas":"0x5f5e100","maxFeePerGas":"0x8a63c4190","value":"0xfd9caec58e1cce","input":"0x","accessList":[],"v":"0x0","r":"0xa878da21c2227d29bb4ae28d19238a80957880a2a04d04467f9aa3bde7dacc24","s":"0x3c267fb8d0348c2cec77d88179635db78055a16ee9da25a2c5d8beb51d8c2460","hash":"0xf3b2eb14180d1876f067c21397684874d22f1fc5b89219fb64868fad56712dec"}
			**/
		}
	}
}

func moitorPendingTxHash() {
	subclient, err := ethclient.Dial(wss)
	if err != nil {
		log.Printf("subclient failed to dial: %v", err)
		return
	}
	pendingTxHash := make(chan common.Hash)
	sub, err := subgclient.SubscribePendingTransactions(context.Background(), pendingTxHash)
	if err != nil {
		fmt.Printf("Error:%v\n", err)
		//log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			fmt.Printf("Error:%v\n", err)
			//log.Fatal(err)
		case txhash := <-pendingTxHash:
			tx, isPending, _ := subclient.TransactionByHash(context.Background(), txhash)
			data, _ := tx.MarshalJSON()
			log.Printf("Pending:%v of tx: %v", isPending, string(data))
		}
	}
}

func DoSubscribe(str string) {
	initClient()
	switch str {
	case "Tx":
		go moitorPendingTx()
	case "TxHash":
		go moitorPendingTxHash()
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
