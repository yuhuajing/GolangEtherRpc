package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// var (
// 	client  *ethclient.Client
// 	account common.Address
// )

// func init() {
// 	client = getConn()
// 	account = common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
// }
// func getConn() *ethclient.Client {
// 	client, err := ethclient.Dial("https://cloudflare-eth.com")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("we have a connection")
// 	return client
// }

type Bloom [256]byte
type BlockNonce [8]byte
type Header struct {
	ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address `json:"miner"`
	Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
	Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
	Number      *big.Int       `json:"number"           gencodec:"required"`
	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
	Time        uint64         `json:"timestamp"        gencodec:"required"`
	Extra       []byte         `json:"extraData"        gencodec:"required"`
	MixDigest   common.Hash    `json:"mixHash"`
	Nonce       BlockNonce     `json:"nonce"`

	// BaseFee was added by EIP-1559 and is ignored in legacy headers.
	BaseFee *big.Int `json:"baseFeePerGas" rlp:"optional"`

	// WithdrawalsHash was added by EIP-4895 and is ignored in legacy headers.
	WithdrawalsHash *common.Hash `json:"withdrawalsRoot" rlp:"optional"`

	// ExcessDataGas was added by EIP-4844 and is ignored in legacy headers.
	ExcessDataGas *big.Int `json:"excessDataGas" rlp:"optional"`

	/*
		TODO (MariusVanDerWijden) Add this field once needed
		// Random was added during the merge and contains the BeaconState randomness
		Random common.Hash `json:"random" rlp:"optional"`
	*/
}

func getLatestBlockHeader() {
	header, _ := client.HeaderByNumber(context.Background(), nil) //nil for the latest block header
	fmt.Println(header.BaseFee)
}

func getTargetBlockHeader(number *big.Int) {
	header, _ := client.HeaderByNumber(context.Background(), number) //nil for the latest block header
	fmt.Println(header.ExcessDataGas)

}

func getLatestBlock() {
	block, _ := client.BlockByNumber(context.Background(), nil) //nil for the latest block

	test := block.Transactions()[:1]
	chainID, _ := client.ChainID(context.Background())
	fmt.Println(chainID)
	for _, tx := range test { //block.Transactions() {
		fmt.Println(tx.Hash())
		signer := types.NewLondonSigner(big.NewInt(1))
		from, _ := types.Sender(signer, tx)
		fmt.Println(from)
	}
	tint, _ := client.TransactionCount(context.Background(), block.Hash())
	fmt.Println(tint)
}

func getTargetBlock(number *big.Int) {
	block, _ := client.BlockByNumber(context.Background(), number)
	test := block.Transactions()[:1]
	fmt.Println(test)
}

// type TxData interface {
// 	txType() byte // returns the type ID
// 	copy() TxData // creates a deep copy and initializes all fields

// 	chainID() *big.Int
// 	accessList() AccessList
// 	data() []byte
// 	gas() uint64
// 	gasPrice() *big.Int
// 	gasTipCap() *big.Int
// 	gasFeeCap() *big.Int
// 	value() *big.Int
// 	nonce() uint64
// 	to() *common.Address

// 	rawSignatureValues() (v, r, s *big.Int)
// 	setSignatureValues(chainID, v, r, s *big.Int)

// 	// effectiveGasPrice computes the gas price paid by the transaction, given
// 	// the inclusion block baseFee.
// 	//
// 	// Unlike other TxData methods, the returned *big.Int should be an independent
// 	// copy of the computed value, i.e. callers are allowed to mutate the result.
// 	// Method implementations can use 'dst' to store the result.
// 	effectiveGasPrice(dst *big.Int, baseFee *big.Int) *big.Int
// }
