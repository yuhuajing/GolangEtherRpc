package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func getBytecode(address common.Address, blockNum int64) {
	res, _ := client.CodeAt(context.Background(), address, big.NewInt(blockNum))
	fmt.Println(hex.EncodeToString(res))
}
