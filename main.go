package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
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

func main() {
	//	fmt.Println(account.Hex())
	getAddrBalFronLatestBlock()
	getLatestBlockHeader()
	fmt.Println("hello world")

}
