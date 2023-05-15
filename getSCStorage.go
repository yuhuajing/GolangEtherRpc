package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	client  *ethclient.Client
	account common.Address
)

func init() {
	client = getConn()
	account = common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2")
}
func getConn() *ethclient.Client {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	return client
}

func getSCstorage() {

	for i := int64(0); i < 4; i++ {
		t := common.BigToHash(big.NewInt(i))
		// fmt.Println(t)
		res, _ := client.StorageAt(context.Background(), account, t, nil)
		// StorageAt
		fmt.Println(i)
		fmt.Println(res)
	}

}
