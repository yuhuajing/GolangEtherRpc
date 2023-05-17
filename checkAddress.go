package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

// check the address whether it is a valid  address
func checkAddress(addr string) bool {
	// 16 hex 0-f
	re := regexp.MustCompile("0x[0-9a-fA-F]{40}$")
	return re.MatchString(addr)
}

// check the address whether is a smart contract address
func checkContractAddress(addr string) bool {
	if !checkAddress(addr) {
		return false
	}
	address := common.HexToAddress(addr)
	bytecode, err := client.CodeAt(context.Background(), address, nil) //nil is the latest block
	if err != nil {
		panic(err)
	}
	isContract := len(bytecode) > 0
	if isContract {
		//fmt.Println("SC address")
		return true
	}
	fmt.Println("This is normal address, but we want a smart contract address")
	return false
}
