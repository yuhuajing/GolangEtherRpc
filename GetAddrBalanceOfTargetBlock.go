package main

/**
There are serveral Ethereum client getting ways
func (*ethclient.Client).BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
1. target the address
2. target the block(if nil, default the latest block)
3. get the balance
func (*ethclient.Client).PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error)
1. target the address
 the pending balance returns the undealt pending balance if there is one
 the pending balance equals to the current balance if there is no pending tx

 It is recommended to use the BalanceAt function to obtain the actual balance when querying the account balance, and then use the PendingBalanceAt function to query the pending balance to obtain more accurate balance information.
**/

import (
	"context"
	"fmt"
	"log"
	"math"
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

func getAddrBalFronLatestBlock() {
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
}

func getAddrBalFronTargetBlock() {
	balance, err := client.BalanceAt(context.Background(), account, big.NewInt(17263443))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
	calcuBalanceToEth(balance)
}

func calcuBalanceToEth(bal *big.Int) {
	fbalance := new(big.Float)
	fbalance.SetString(bal.String())
	fbalance = fbalance.Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(fbalance)
}

func getPendingBalance() {
	balance, _ := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(balance)
}
