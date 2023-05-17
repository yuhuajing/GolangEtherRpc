package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

/**
When a node receives a Tx message, it encodes the Tx and propagates the transaction to the network
**/

func getRawTx() {
	priKey, _ := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	publicKey := priKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromaddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	nonce, _ := client.PendingNonceAt(context.Background(), common.HexToAddress(fromaddress))
	value := big.NewInt(0) // in wei (1 eth) optionally given
	gasLimit := uint64(21000)
	//gasPrice := big.NewInt(30000000000) // in wei (30 gwei)
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	amount := big.NewInt(1000000000000000000)
	tokenAddress := common.HexToAddress("0x28b149020d2152179873ec60bed6bf7cd705775d")
	//construct the input data structure
	//padding the input parameters to 256 bit including to address and the token amount
	transferFn := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFn)
	methodID := hash.Sum(nil)[:4]
	int256 := new(big.Int)
	int256.SetBytes(methodID)
	fmt.Printf("0x%x\n", int256) //0xa9059cbb
	var data []byte
	data = append(data, methodID...)
	data = append(data, common.LeftPadBytes(amount.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(toAddress.Bytes(), 32)...)

	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	chainId, _ := client.ChainID(context.Background())
	signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(chainId), priKey)
	fmt.Println(signedTx)
	rawBytes, _ := signedTx.MarshalBinary()
	fmt.Println(hexutil.Encode(rawBytes))
	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())
	retx := new(types.Transaction)
	retx.UnmarshalBinary(rawBytes)
	//	rlp.DecodeBytes(rawBytes, &retx)
	fmt.Println(retx)
}

//0xf8a9038509023dfa578252089428b149020d2152179873ec60bed6bf7cd705775d80b844a9059cbb0000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d25a002e03b2d4aa6e5228d56fee8f248ee13600dcff522238680d1e66256f8fc57fba00eda770be6b1bb0824d152fef073fd13af8fa5030b6eb81cc4abc9e5983d34e8
