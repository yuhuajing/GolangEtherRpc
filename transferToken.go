package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

/**
When we transfer Token by a transaction, we need to target the Token_Contract_Address、To_Address、GasPrice、GasLimit、Nonce and input data including(method、input parameters).
1.import the from address by private key or generate a new address
2.get the address newest nonce value from the pending pool
3. optinally target the transfer value with wei
4.target the gas price and gas limit
The go-ethereum client provides the SuggestGasPrice function for getting the average gas price based on x number of previous blocks
only Target the gas price will result:
 (1) gas price = min(gasPrice,basefee+priorityFee)
 if we just provide the gas price
  (2)priorityFee = gasPrice - basefee
5.target the token contract address
6. construct the input data including(methodID、input parameters)
  (1) methodID is the front 8 characters of the keccak-256 hash of the function
7.Construct the tx including all necessary parameters
8. sign the tx by the private key
9. send the tx
**/

func transferToken() {
	//fromAddress := generateNode()
	//or Import the from address
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
	var data []byte
	data = append(data, methodID...)
	data = append(data, common.LeftPadBytes(amount.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(toAddress.Bytes(), 32)...)
	tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
	chainId, _ := client.ChainID(context.Background())
	signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(chainId), priKey)
	err := client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
