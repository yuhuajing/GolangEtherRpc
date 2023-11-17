package main

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	privateKey, err := crypto.HexToECDSA("797391c7bd2e156e52329ceb6471496798e0c125ef35c4c3393329bd2a64f3f5")
	publicKeyECDSA, _ := privateKey.Public().(*ecdsa.PublicKey)
	signer := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	log.Println("signer: ", signer)

	ethHash := packetWithEth("Hello, World!")
	//log.Println("eth hash: ", ethHash.Hex())

	signature, err := crypto.Sign(ethHash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("signature: ", hexutil.Encode(signature))
	log.Println("verify: ", verifySig(signer, ethHash.Bytes(), signature))
}

func packetWithEth(str string) common.Hash {
	hash := crypto.Keccak256Hash([]byte(str))
	return crypto.Keccak256Hash(
		[]byte("\x19Ethereum Signed Message:\n32"),
		hash.Bytes(),
	)
}

func verifySig(signer string, msg []byte, sig []byte) bool {

	pubKey, err := crypto.Ecrecover(msg, sig)
	if err != nil {
		log.Fatal(err)
	}
	hash := sha3.NewLegacyKeccak256()
	hash.Write(pubKey[1:])
	straddr := common.HexToAddress(hexutil.Encode(hash.Sum(nil)[12:]))
	return straddr == common.HexToAddress(signer)
}
