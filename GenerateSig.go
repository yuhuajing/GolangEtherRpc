package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
)

func main() {

	// Parse private key
	privateKey, err := crypto.HexToECDSA("797391c7bd2e156e52329ceb6471496798e0c125ef35c4c3393329bd2a64f3f5")
	if err != nil {
		log.Fatal(err)
	}
	// Sign message
	signature, err := PersonalSign("Hello, World!", privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signature)
}

func PersonalSign(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	mes := solsha3.SoliditySHA3(
		// types
		[]string{"string"},

		// values
		[]interface{}{
			message,
		},
	)
	hasht := solsha3.SoliditySHA3WithPrefix(mes)
	signatureBytes, err := crypto.Sign(hasht, privateKey)

	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), nil
}
