package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func getTranctionsFromlatestBlock() {
	block, _ := client.BlockByNumber(context.Background(), nil) //nil for the latest block
	txNum, _ := client.TransactionCount(context.Background(), block.Hash())
	for idx := uint(0); idx < txNum; idx++ {
		tx, _ := client.TransactionInBlock(context.Background(), block.Hash(), idx)
		from, _ := client.TransactionSender(context.Background(), tx, block.Hash(), idx)
		ParseTransactionBaseInfo(tx, from)
		//contractABI, _ := abi.JSON(strings.NewReader(GetLocalABI()))
		//DecodeTransactionInputData(&contractABI, tx.Data())
	}
}

func getTranctionsFromTargetBlock(number *big.Int) {
	block, _ := client.BlockByNumber(context.Background(), number) //nil for the latest block
	txNum, _ := client.TransactionCount(context.Background(), block.Hash())
	for idx := uint(0); idx < txNum; idx++ {
		tx, _ := client.TransactionInBlock(context.Background(), block.Hash(), idx)
		from, _ := client.TransactionSender(context.Background(), tx, block.Hash(), idx)
		ParseTransactionBaseInfo(tx, from)
		//contractABI, _ := abi.JSON(strings.NewReader(GetLocalABI()))
		//DecodeTransactionInputData(&contractABI, tx.Data())
	}
}

// 	return fmt.Sprintf("Tx type: %v nonce: %v gasPrice: %v maxPriorFee: %v gas: %v value: %v input: %v to: %v hash: %v", txType, nonce, gasPrice, maxPriorFee, gas, value, input, to, hash)

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

func ParseTransactionBaseInfo(tx *types.Transaction, from common.Address) {
	fmt.Printf("Type: %d\n", tx.Type())
	fmt.Printf("Hash: %s\n", tx.Hash().Hex())
	fmt.Printf("ChainId: %d\n", tx.ChainId())
	fmt.Printf("Value: %s\n", tx.Value().String())
	fmt.Printf("From: %s\n", from) // from field is not inside of transation
	fmt.Printf("To: %s\n", tx.To().Hex())
	fmt.Printf("Gas: %d\n", tx.Gas())
	fmt.Printf("maxPriorFee: %d\n", tx.GasTipCap())
	fmt.Printf("Gas Price: %d\n", tx.GasPrice().Uint64())
	fmt.Printf("Nonce: %d\n", tx.Nonce())
	fmt.Printf("Transaction Data in hex: %s\n", hex.EncodeToString(tx.Data()))
	fmt.Print("\n")
}

func GetLocalABI() string {
	path := "abis/boredape.json"
	abiFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer abiFile.Close()

	result, err := io.ReadAll(abiFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(result)
}

func DecodeTransactionInputData(contractABI *abi.ABI, data []byte) {
	// The first 4 bytes of the t represent the ID of the method in the ABI
	// https://docs.soliditylang.org/en/v0.5.3/abi-spec.html#function-selector
	methodSigData := data[:4]
	method, err := contractABI.MethodById(methodSigData)
	if err != nil {
		log.Fatal(err)
	}

	inputsSigData := data[4:]
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Method Name: %s\n", method.Name)
	fmt.Printf("Method inputs: %v\n", inputsMap)
}
