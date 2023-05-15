package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// check the address whether it is a valid  address
func checkAddress(addr string) {
	// 16 hex 0-f
	re := regexp.MustCompile("0x[0-9a-fA-F]{40}$")
	if !re.MatchString(addr) {
		panic("invalid address")
	}
}

// check the address whether is a smart contract address
func checkContractAddress(addr string) {
	address := common.HexToAddress(addr)
	client, _ := ethclient.Dial("https://cloudflare-eth.com")
	bytecode, err := client.CodeAt(context.Background(), address, nil) //nil is the latest block
	if err != nil {
		panic(err)
	}
	isContract := len(bytecode) > 0
	if isContract {
		fmt.Println("SC address")
	} else {
		fmt.Println("Normal address")
	}
}
