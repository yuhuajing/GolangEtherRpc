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

func init() {
	server := "https://cloudflare-eth.com"
	client = getConn(server)
	//account = common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
}

func getConn(str string) *ethclient.Client {
	client, err := ethclient.Dial(str)
	//https://cool-muddy-butterfly.discover.quiknode.pro/0e41f42d5a7c9611f30ef800444bfcb93d3ae9a6/
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	return client
}

func main() {

	flag.StringVar(&ethServer, "ethServer", "", "The public Ethereum server to connect to")
	flag.StringVar(&address, "address", "", "The smart contract address to get storage")
	flag.IntVar(&slot, "slot", 0, "The singal slot to get storage")
	flag.IntVar(&highslot, "highslot", 0, "The contiounus highest slot to get storage")
	flag.Int64Var(&blockNum, "blockNum", 0, "The blocknum to get storage")
	flag.IntVar(&lowslot, "lowslot", 0, "The contiounus lowest slot to get storage")
	flag.StringVar(&arrayslot, "arrayslot", "", "The specific slot to get storage like `1 2 3 4 5` ")
	flag.Parse()
	if address == "" || !checkContractAddress(address) {
		fmt.Println("--address should be provided or the address should be a smart contract address")
		return
	}

	if ethServer != "" {
		client = getConn(ethServer)
	}
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
	} else {
		//	fmt.Printf("nothing provided,get slot 0 of the address %s\n", naddress.Hex())
		getSCstorage(naddress, slot, blockNum)
	}
	//getLatestBlockHeader()

}
