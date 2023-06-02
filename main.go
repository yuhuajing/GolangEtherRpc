package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	client *ethclient.Client
	//account common.Address
)

const (
	ethServer  = "https://cloudflare-eth.com"
	bscRpc     = "https://bsc-mainnet.nodereal.io/v1/64a9df0874fb4a93b9d0a3849de012d3"
	polygonRpc = "https://polygon-mainnet.nodereal.io/v1/f510fc4d083b49d1ab383d25246cc7de"
	opRpc      = "https://opt-mainnet.nodereal.io/v1/1fd7be3e976444759d636dd367aae9ac"
	arbitrum   = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/arbitrum-nitro/"
	avalanch   = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/avalanche-c/ext/bc/C/avax"
	salana     = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/solana/"
	near       = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/near/"
	fantom     = "https://open-platform.nodereal.io/1fd7be3e976444759d636dd367aae9ac/fantom/"
)

func init() {
	// server := "https://cloudflare-eth.com"
	client = getConn(ethServer)

}

func getConn(str string) *ethclient.Client {
	nclient, err := ethclient.Dial(str)
	//https://cool-muddy-butterfly.discover.quiknode.pro/0e41f42d5a7c9611f30ef800444bfcb93d3ae9a6/
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("we have a connection")
	return nclient
}

var (
	address                  string
	slot, highslot, lowslot  int
	arrayslot                string
	chain, txHash, analyHash string
	blockNum                 int64
)

func main() {
	//different chain
	//different RPC
	flag.StringVar(&chain, "chain", "", "The public Ethereum server to connect to")
	flag.StringVar(&address, "address", "", "The smart contract address to get storage")
	flag.IntVar(&slot, "slot", 0, "The singal slot to get storage")
	flag.IntVar(&highslot, "highslot", 0, "The contiounus highest slot to get storage")
	flag.Int64Var(&blockNum, "blockNum", 0, "The blocknum to get storage")
	flag.IntVar(&lowslot, "lowslot", 0, "The contiounus lowest slot to get storage")
	flag.StringVar(&arrayslot, "arrayslot", "", "The specific slot to get storage like `1 2 3 4 5` ")
	flag.StringVar(&txHash, "txhash", "", "Transactions infrmation of the target tx hash")
	flag.StringVar(&analyHash, "analyhash", "", "Transactions infrmation of the target tx hash")
	flag.Parse()

	DoSubscribe("Tx")
	// anaTxInfo := analysis(analyHash)
	// fmt.Println(anaTxInfo)

	// txInfo := getTxFromTxHash(txHash)
	//fmt.Println(txInfo)
	if chain != "" {
		switch chain {
		case "bsc":
			client = getConn(bscRpc)
		case "polygon":
			client = getConn(polygonRpc)
		case "optimism":
			client = getConn(opRpc)
		case "arbitrum":
			client = getConn(arbitrum)
		case "ethereum":
			client = getConn(ethServer)
			// case "avalanch":
			// 	client = getConn(avalanch)
			// case "solana":
			// 	client = getConn(salana)
			// case "near":
			// 	client = getConn(near)
			// case "fantom":
			// 	client = getConn(fantom)
		}
	}

	// if address == "" || !checkContractAddress(address) {
	// 	fmt.Println("--address should be provided or the address should be a smart contract address")
	// 	return
	// }

	naddress := common.HexToAddress(address)
	//	fmt.Println(naddress)
	if slot > 0 {
		//	fmt.Printf("signal Slot provided, get slot %d of the address %s\n", slot, naddress.Hex())
		getSCstorage(naddress, slot, blockNum)
		return
	} else if highslot > 0 && lowslot >= 0 {
		for i := lowslot; i <= highslot; i++ {
			//fmt.Printf("highSlot provided, get slot %d of the address %s\n", i, naddress.Hex())
			getSCstorage(naddress, i, blockNum)
		}
		return
	} else if arrayslot != "" {
		strArray := strings.Split(arrayslot, " ")
		for i := 0; i < len(strArray); i++ {
			num, err := strconv.Atoi(strArray[i])
			if err != nil {
				fmt.Println("error:", err)
				return
			}
			//	fmt.Printf("arraySlot provided, get slot %d of the address %s\n", num, naddress.Hex())
			getSCstorage(naddress, num, blockNum)
		}
		return
	} // else {
	// 	//	fmt.Printf("nothing provided,get slot 0 of the address %s\n", naddress.Hex())
	// 	getSCstorage(naddress, slot, blockNum)
	// }
	//getLatestBlockHeader()

}
