# EtherRpcBook
Learn how to deploy the Ethereum rpc functions

The learning page in the https://goethereumbook.org/account-balance/
# getSCSlotData

```bash
Usage of ./main:
  -address string
        The smart contract address to get storage.
  -arrayslot "1 2 3 4 5"
        The specific slot to get storage like "1 2 3 4 5"
  -blockNum int
        The blocknum to get storage, default: the latest block.
  -chain string
        The public Ethereum server to connect to, default: Ethereum.
  -highslot int
        The contiounus highest slot to get storage.
  -lowslot int
        The contiounus lowest slot to get storage, default: 0.
  -slot int
        The singal slot to get storage, default: 0.
```
1. chain configure

Now support ethereum、bsc、polygon、optimisim、arbitrum
```golang
./main --chain="xx"

switch xx{
    case "bsc":
			Get data from BSC
	case "polygon":
			Get data from Polygon
	case "optimism":
			Get data from Optimism
	case "arbitrum":
			Get data from Arbitrum
    case "ethereum":
			Get data from Ethereum
}
```

# Get smart contract slot data from Ethereum

1. Get single slot data
./main --address="0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" --slot=1 --blockNum=17269518

Example
```bash
0x5745544800000000000000000000000000000000000000000000000000000008
```

2. Get continuous slot data, from 0 to the highslot data. The default lowslot is 0.

 ./main --address="0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" --highslot=1 --blockNum=17269518

 Example
```bash
0x577261707065642045746865720000000000000000000000000000000000001a
0x5745544800000000000000000000000000000000000000000000000000000008
``` 

3. Get continuous slot data, from lowslot to the highslot data.

./main --address="0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" --lowslot=1 --highslot=2 --blockNum=17269518

 Example
```bash
0x5745544800000000000000000000000000000000000000000000000000000008
0x12
``` 

4. get specific slot data by providing string slot.

 ./main --address="0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2" --arrayslot="0 1 2" --blockNum=17269518

 Example
```bash
0x577261707065642045746865720000000000000000000000000000000000001a
0x5745544800000000000000000000000000000000000000000000000000000008
0x12
``` 

# Get smart contract slot data from BSC

./main --chain="bsc" --address="0x245E0BA4562E96CC3873b7C1594954AaD877a4dE" --lowslot=0 --highslot=10 --blockNum=28312630

 Example
```bash
0x55a3b37957bfbd3345bed9968e7e8dd56d67066
0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
0x10000000000
0x10000000000000000
0x100000000000000000000000000000000
0x0
0x0
0x55a3b37957bfbd3345bed9968e7e8dd56d67066
0x4946c0e9f43f4dee607b0ef1fa1c
0x0
0x0
```

# Get smart contract slot data from polygon

./main --chain="polygon" --address="0x53fe4D1aB48363f1bFf8BFB5d96148E738f4ed6B" --lowslot=0 --highslot=10 --blockNum=42846967

 Example
```bash
0x5069204e465420322e3000000000000000000000000000000000000000000014
0x5049000000000000000000000000000000000000000000000000000000000004
0x0
0x0
0x0
0x0
0x0
0x0
0x2c86
0x0
0x142b27b469bb2bb1257d8e077efd6d8b666cf1e
```

# Get smart contract slot data from optimism

./main --chain="optimism" --address="0x86Bb63148d17d445Ed5398ef26Aa05Bf76dD5b59" --highslot=10 

Q: 无法指定blockNum

 Example
```bash
0x9cff694c1689f5c4bc4482b75b013e9634319b7b
0x0
0x0
0x0
0x0
0x1
0x6c
0x0
0x0
0x0
```

# Get smart contract slot data from arbitrum

./main --chain="arbitrum" --address="0x15b2fb8f08E4Ac1Ce019EADAe02eE92AeDF06851" --highslot=10 --blockNum=91881435

 Example
```bash
0x3c27899d2f495928d00000
0x0
0x0
0x233001101d35d7a5cab7d51bcfbab3679ae4dc1501
0x0
0x0
0x0
0x0
0x0
0x0
0x0
```