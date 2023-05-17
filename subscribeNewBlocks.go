package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

/*
*
Apply for Network API.
https://www.quicknode.com/

// in the main.go
monitorBlocks()
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

func monitorBlocks() {
	subclient, err := ethclient.Dial("wss://cool-muddy-butterfly.discover.quiknode.pro/0e41f42d5a7c9611f30ef800444bfcb93d3ae9a6/")
	if err != nil {
		log.Printf("subclient failed to dial: %v", err)
		return
	}
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
