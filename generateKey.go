package main

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func generateNode() {
	priveKey, _ := crypto.GenerateKey()
	//privateKeyBytes := crypto.FromECDSA(priveKey)
	//fmt.Println(hexutil.Encode(privateKeyBytes))
	publicKey := priveKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	//fmt.Println(hexutil.Encode(publicKeyBytes))
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}
