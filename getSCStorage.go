package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	address                 string
	slot, highslot, lowslot int
	arrayslot               string
	//client                  *ethclient.Client
	//account                 common.Address
	ethServer string
	blockNum  int64
)

func getSCstorage(address common.Address, slot int, blockNum int64) {
	t := common.BigToHash(big.NewInt(int64(slot)))
	int256 := new(big.Int)
	if blockNum != 0 {
		//fmt.Printf("get slot %d of the address %s in the block %d\n", slot, address.Hex(), blockNum)
		blocknumBigInt := big.NewInt(int64(blockNum))
		res, _ := client.StorageAt(context.Background(), address, t, blocknumBigInt)
		//	fmt.Println(res)
		int256.SetBytes(res)
	} else {
		//fmt.Printf("get slot %d of the address %s in the latest block\n", slot, address.Hex())
		res, _ := client.StorageAt(context.Background(), address, t, nil)
		//	fmt.Println(res)
		int256.SetBytes(res)
	}
	//fmt.Println()
	fmt.Printf("0x%x\n", int256)
	// fmt.Printf("uint256: %v\n", int256)
}
