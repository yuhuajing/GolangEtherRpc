package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

/**
When we transfer ETH by a transaction, we need to target the To、GasPrice、GasLimit、Nonce and optionally data.
1.import the from address by private key or generate a new address
2.get the address newest nonce value from the pending pool
3. target the transfer value with wei
4.target the gas price and gas limit
The go-ethereum client provides the SuggestGasPrice function for getting the average gas price based on x number of previous blocks
only Target the gas price will result:
 (1) gas price = min(gasPrice,basefee+priorityFee)
 if we just provide the gas price
  (2)priorityFee = gasPrice - basefee
5.target the to address
6.Construct the tx including all necessary parameters
7. sign the tx by the private key
8. send the tx
**/

func transferEth() {
	//fromAddress := generateNode()
	//or Import the from address
	priKey, _ := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	publicKey := priKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromaddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	nonce, _ := client.PendingNonceAt(context.Background(), common.HexToAddress(fromaddress))
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(21000)
	//gasPrice := big.NewInt(30000000000) // in wei (30 gwei)
	gasPrice, _ := client.SuggestGasPrice(context.Background())
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)

	chainId, _ := client.ChainID(context.Background())
	signedTx, _ := types.SignTx(tx, types.NewEIP155Signer(chainId), priKey)
	// err := client.SendTransaction(context.Background(), signedTx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
