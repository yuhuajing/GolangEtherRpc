package main

/**
There are serveral Ethereum client getting ways
1. local server
client, err := ethclient.Dial("http://localhost:8545")
OR
client, err := ethclient.Dial("/home/user/.ethereum/geth.ipc")
2. cloudflare
client, err := ethclient.Dial("https://cloudflare-eth.com")
3. mainnet
client, err := ethclient.Dial("https://mainnet.infura.io")
**/

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func initConn() {
	//client, err := ethclient.Dial("https://cloudflare-eth.com")
	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	_ = client
}
