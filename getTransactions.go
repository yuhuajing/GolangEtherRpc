package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	client  *ethclient.Client
	account common.Address
)

func init() {
	client = getConn()
	account = common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
}
func getConn() *ethclient.Client {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	return client
}

func getTranctionsFromlatestBlock() {
	block, _ := client.BlockByNumber(context.Background(), nil) //nil for the latest block
	txNum, _ := client.TransactionCount(context.Background(), block.Hash())
	for idx := uint(0); idx < txNum; idx++ {
		tx, _ := client.TransactionInBlock(context.Background(), block.Hash(), idx)
		from, _ := client.TransactionSender(context.Background(), tx, block.Hash(), idx)
		fmt.Println(tx.Hash().Hex())
		fmt.Println(from)
	}
}

func getReciptFromTxHash() {
	block, _ := client.BlockByNumber(context.Background(), nil) //nil for the latest block
	tx := block.Transactions()
	for _, txinfo := range tx {
		fmt.Println(txinfo.Hash().Hex())
		receipt, _ := client.TransactionReceipt(context.Background(), txinfo.Hash())
		fmt.Println(receipt.Logs)
	}
}

func getTranctionsSenderFronRSV() {
	block, _ := client.BlockByNumber(context.Background(), nil) //nil for the latest block
	test := block.Transactions()[:1]
	chainID, _ := client.ChainID(context.Background())
	fmt.Println(chainID)
	for _, tx := range test { //block.Transactions() {
		fmt.Println(tx.Hash())
		signer := types.NewLondonSigner(big.NewInt(1))
		from, _ := types.Sender(signer, tx)
		fmt.Println(from)
	}
}
