package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func getTranctionsFromlatestBlock() []string {
	block, _ := client.BlockByNumber(context.Background(), nil) //nil for the latest block
	txNum, _ := client.TransactionCount(context.Background(), block.Hash())
	strArray := make([]string, txNum)
	for idx := uint(0); idx < txNum; idx++ {
		tx, _ := client.TransactionInBlock(context.Background(), block.Hash(), idx)
		from, _ := client.TransactionSender(context.Background(), tx, block.Hash(), idx)
		str := ParseTransactionBaseInfo(tx, from)
		//contractABI, _ := abi.JSON(strings.NewReader(GetLocalABI()))
		//DecodeTransactionInputData(&contractABI, tx.Data())
		strArray = append(strArray, str)
	}
	return strArray
}

func getTranctionsFromTargetBlock(number *big.Int) []string {
	block, _ := client.BlockByNumber(context.Background(), number) //nil for the latest block
	txNum, _ := client.TransactionCount(context.Background(), block.Hash())
	strArray := make([]string, txNum)
	for idx := uint(0); idx < txNum; idx++ {
		tx, _ := client.TransactionInBlock(context.Background(), block.Hash(), idx)
		from, _ := client.TransactionSender(context.Background(), tx, block.Hash(), idx)
		str := ParseTransactionBaseInfo(tx, from)
		//contractABI, _ := abi.JSON(strings.NewReader(GetLocalABI()))
		//DecodeTransactionInputData(&contractABI, tx.Data())
		strArray = append(strArray, str)
		//contractABI, _ := abi.JSON(strings.NewReader(GetLocalABI()))
		//DecodeTransactionInputData(&contractABI, tx.Data())
	}
	return strArray
}

func getReciptFromTxHash(hash string) *types.Receipt {
	// block, _ := client.BlockByNumber(context.Background(), nil) //nil for the latest block
	// tx := block.Transactions()
	// for _, txinfo := range tx {
	// 	fmt.Println(txinfo.Hash().Hex())
	receipt, _ := client.TransactionReceipt(context.Background(), common.HexToHash(hash)) //txinfo.Hash())
	//receStr := ParseTransactionReceBaseInfo(receipt)
	// fmt.Println(receipt.GasUsed)
	return receipt
	// }
}

func getTxFromTxHash(hash string) string {
	tx, _, _ := client.TransactionByHash(context.Background(), common.HexToHash(hash)) //txinfo.Hash())
	from := getTranctionsSenderFronRSV(tx)
	txinfo := ParseTransactionBaseInfo(tx, from)
	receipt, _ := client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	t := receipt.BlockNumber
	TxBlockHeader := getTargetBlockHeader(t)
	basefee := TxBlockHeader.BaseFee

	receStr := ParseTransactionReceBaseInfo(basefee, receipt)
	return txinfo + " " + receStr
	//fmt.Println(txinfo)
}

func getTranctionsSenderFronRSV(tx *types.Transaction) common.Address {
	// block, _ := client.BlockByNumber(context.Background(), nil) //nil for the latest block
	// test := block.Transactions()[:1]
	// chainID, _ := client.ChainID(context.Background())
	// fmt.Println(chainID)
	// for _, tx := range test { //block.Transactions() {
	// 	fmt.Println(tx.Hash())
	signer := types.NewLondonSigner(big.NewInt(1))
	from, _ := types.Sender(signer, tx)
	return from
	//fmt.Println(from)
	// }
}

func ParseTransactionBaseInfo(tx *types.Transaction, from common.Address) string {
	//tx.EffectiveGasTip(base)
	//return fmt.Sprintf("Tx type: %d nonce: %d gasPrice: %d maxPriorFee: %d gas: %v value: %v  from %s to: %s hash: %s input: %s", tx.Type(), tx.Nonce(), tx.GasPrice().Uint64(), tx.GasTipCap().Uint64(), tx.Gas(), tx.Value(), from.Hex(), tx.To().Hex(), tx.Hash().Hex(), hex.EncodeToString(tx.Data()))
	return fmt.Sprintf("Tx type: %d nonce: %d gasPrice: %d maxPriorFee: %d gas: %v value: %v  from %s to: %s hash: %s", tx.Type(), tx.Nonce(), tx.GasPrice().Uint64(), tx.GasTipCap().Uint64(), tx.Gas(), tx.Value(), from.Hex(), tx.To().Hex(), tx.Hash().Hex())
	// fmt.Printf("Type: %d\n", tx.Type())
	// fmt.Printf("Hash: %s\n", tx.Hash().Hex())
	// fmt.Printf("ChainId: %d\n", tx.ChainId())
	// fmt.Printf("Value: %s\n", tx.Value().String())
	// fmt.Printf("From: %s\n", from) // from field is not inside of transation
	// fmt.Printf("To: %s\n", tx.To().Hex())
	// fmt.Printf("Gas: %d\n", tx.Gas())
	// fmt.Printf("maxPriorFee: %d\n", tx.GasTipCap())
	// fmt.Printf("Gas Price: %d\n", tx.GasPrice().Uint64())
	// fmt.Printf("Nonce: %d\n", tx.Nonce())
	// fmt.Printf("Transaction Data in hex: %s\n", hex.EncodeToString(tx.Data()))
	// fmt.Print("\n")
}
func ParseTransactionReceBaseInfo(basefee *big.Int, tx *types.Receipt) string {
	gasUsed := tx.GasUsed
	effecGasPrice := tx.EffectiveGasPrice.Uint64()
	effertips := effecGasPrice - basefee.Uint64()
	//	effecGasTipCap := tx.EffectiveGasTipCap.Uint64()
	fee := big.NewInt(int64(gasUsed) * int64(effecGasPrice))
	feeEth := calcuBalanceToEth(fee)
	return fmt.Sprintf("gasUsed: %d, basePrice %v, effectipsPrice %v, effecGasPrice: %v,feeEth: %v\n", gasUsed, basefee, effertips, effecGasPrice, feeEth)
}

func GetLocalABI() string {
	path := "abis/boredape.json"
	abiFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer abiFile.Close()

	result, err := io.ReadAll(abiFile)
	if err != nil {
		log.Fatal(err)
	}
	return string(result)
}

func DecodeTransactionInputData(contractABI *abi.ABI, data []byte) {
	// The first 4 bytes of the t represent the ID of the method in the ABI
	// https://docs.soliditylang.org/en/v0.5.3/abi-spec.html#function-selector
	methodSigData := data[:4]
	method, err := contractABI.MethodById(methodSigData)
	if err != nil {
		log.Fatal(err)
	}
	inputsSigData := data[4:]
	inputsMap := make(map[string]interface{})
	if err := method.Inputs.UnpackIntoMap(inputsMap, inputsSigData); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Method Name: %s\n", method.Name)
	fmt.Printf("Method inputs: %v\n", inputsMap)
}

func analysis(hash string) string {
	tx, _, _ := client.TransactionByHash(context.Background(), common.HexToHash(hash)) //txinfo.Hash())
	//from := getTranctionsSenderFronRSV(tx)
	txGasprice := tx.GasPrice().Int64()
	gasFeecap := tx.GasFeeCap().Int64()
	tipsCap := tx.GasTipCap().Int64()
	receipt, _ := client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	gasused := receipt.GasUsed
	t := receipt.BlockNumber
	TxBlockHeader := getTargetBlockHeader(t)
	basefee := TxBlockHeader.BaseFee
	effectGasPrice := receipt.EffectiveGasPrice.Int64()
	effectTipPrice := effectGasPrice - basefee.Int64()
	fee := int64(gasused) * effectGasPrice
	feeEth := calcuBalanceToEth(big.NewInt(fee))
	//percent
	var gaspricePercent float64
	var tipspercent float64
	fmt.Println()
	if effectGasPrice <= txGasprice {
		gaspricePercent = (float64(txGasprice) - float64(effectGasPrice)) / float64(txGasprice)
	}
	if effectTipPrice <= tipsCap {
		tipspercent = (float64(tipsCap) - float64(effectTipPrice)) / float64(tipsCap)
	}
	return fmt.Sprintf("TxgasPrice: %d, effectGasPrice: %d, percentPrice: %f, MaxFeePerGas: %d, MaxPriorityFeePerGas: %d,EffectMaxPriorityFeePerGas: %d, percentTipPrice: %f, gasUsed: %d, txFee:%v",
		txGasprice, effectGasPrice, gaspricePercent, gasFeecap, tipsCap, effectTipPrice, 1-tipspercent, gasused, feeEth)
	//fmt.Println(txinfo)
}
