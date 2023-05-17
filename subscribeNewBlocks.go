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
DoSubscribe("Block") -- Get all pending blocks
DoSubscribe("TxHash") -- Get all pending transaction hash

outPut example:
Pending:true of tx: {"type":"0x0","nonce":"0x4","gasPrice":"0x3b9aca00","maxPriorityFeePerGas":null,"maxFeePerGas":null,"gas":"0x7530","value":"0x19ebac6047ff","input":"0x","v":"0x26","r":"0x25ce543705f138288a7a8c728870fb9cda6c8e356a566900e1c800d1962f5fa3","s":"0xf45eaba2ad5ade4d41323bc073c9ebdd20c21a36cf046b6da1b686b6b84fa1a","to":"0xa03ae5eddc99e65dd7e1315d4b238e51f9baa6bc","hash":"0xa385817662a5008c6b742868ec37355138245ba013b4609c2ceaa533f98a70cd"}

Need change the default client Connections to ws
//	client, err := ethclient.Dial("wss://cool-muddy-butterfly.discover.quiknode.pro/xx/")

// Header represents a block header in the Ethereum blockchain.

	type Header struct {
		ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
		UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
		Coinbase    common.Address `json:"miner"`
		Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
		TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
		ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
		Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
		Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
		Number      *big.Int       `json:"number"           gencodec:"required"`
		GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
		GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
		Time        uint64         `json:"timestamp"        gencodec:"required"`
		Extra       []byte         `json:"extraData"        gencodec:"required"`
		MixDigest   common.Hash    `json:"mixHash"`
		Nonce       BlockNonce     `json:"nonce"`

		// BaseFee was added by EIP-1559 and is ignored in legacy headers.
		BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`

		// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
		WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

		// ExcessDataGas was added by EIP-4844 and is ignored in legacy headers.
		ExcessDataGas *big.Int `json:"excessDataGas" rlp:"optional"`

}
*
*/
var (
	subclient  *ethclient.Client
	subgclient *gethclient.Client
	err        error
	wss        string = "wss://cool-muddy-butterfly.discover.quiknode.pro/0e41f42d5a7c9611f30ef800444bfcb93d3ae9a6/"
)

func initClient() {
	//Retrieve pending blocks
	subclient, err = ethclient.Dial(wss)
	if err != nil {
		log.Printf("subclient failed to dial: %v", err)
		return
	}

	rpcCli, err := rpc.Dial(wss)
	if err != nil {
		log.Printf("rpcDial failed to dial: %v", err)
		return
	}
	//Reterive pending transactions
	subgclient = gethclient.New(rpcCli)
}

func monitorBlocks() {
	headers := make(chan *types.Header)
	sub, err := subclient.SubscribeNewHead(context.Background(), headers)

	if err != nil {
		fmt.Printf("Error:%v\n", err)
		//log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			fmt.Printf("Error:%v\n", err)
			//log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())
			// block, _ := client.BlockByNumber(context.Background(), header.Number)
			// for _, tx := range block.Transactions() {
			// 	msg := tx.To()
			// }
		}
	}
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
			fmt.Println(tx.To())
			// data, _ := tx.MarshalJSON()
			// fmt.Println(string(data))
		}
	}
}

func moitorPendingTxHash() {
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
	case "Blocks":
		go monitorBlocks()
	case "Tx":
		go moitorPendingTx()
	case "TxHash":
		go moitorPendingTxHash()
	}
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan
}
