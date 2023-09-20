# 代理合约

abigen --bin=./wallet.bin --abi=./wallet.abi --pkg=wallet --out=wallet.go

factory.sol: 通过写入字节吗的方式实现代理，将收到的所有数据通过fallback函数 delegate 到代理中去

Wallet.sol: 实现合约，实现工厂合约的实际业务逻辑

测试：
1. 部署 Wallet.sol
2. 部署 factory.sol
3. 调用 factory.sol 中的 createAccount 函数，输入 Wallet.sol 的地址和其他参数，生成新的钱包地址
4. 通过  Wallet.sol 的 abi 和 bytecode 生成 对 Wallet.sol 合约的调用 golang代码 （abigen --bin=./wallet.bin --abi=./wallet.abi --pkg=wallet --out=wallet.go）
5. 将 对合约 Wallet.sol 的调用全部转到 新的钱包地址，然后再 fallback 回 Wallet.sol 合约

测试读数据代码：
```golang
func Test_wallet(t *testing.T) {
	//priKey, _ := crypto.HexToECDSA("d34772897e5d6bc952cff1094945d6d05bca81decd773f0cfb3575fbc4a73493")
	client := ethconn.ConnBlockchain("http://localhost:8545")
	// publicKey := priKey.Public()
	// publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	// fromaddress := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println(fromaddress)
	address := common.HexToAddress("0xbB56d2FAd712F29D74241e9C45936Ac9fd049640") // 钱包地址
	instance, err := wallet.NewWallet(address, client)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}

	number, _ := instance.Gettestnum(nil)
	fmt.Println(number)

	// bal, _ := client.BalanceAt(context.Background(), common.HexToAddress(fromaddress), nil)
	// fmt.Println(bal)
}
```

测试写数据：
```golang
func Test_wallet_write(t *testing.T) {
	client := ethconn.ConnBlockchain("http://localhost:8545")
	privateKey, _ := crypto.HexToECDSA("d34772897e5d6bc952cff1094945d6d05bca81decd773f0cfb3575fbc4a73493")
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//fmt.Println(fromaddress)
	address := common.HexToAddress("0xbB56d2FAd712F29D74241e9C45936Ac9fd049640")
	instance, err := wallet.NewWallet(address, client)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	newnumber := big.NewInt(99)
	tx, err := instance.Settestnum(auth, newnumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	time.Sleep(10 * time.Second)
	number, _ := instance.Gettestnum(nil)
	fmt.Println("new Number: ", number)
}
```

发起转账交易
```golang
func Test_wallet_write(t *testing.T) {
	client := ethconn.ConnBlockchain("http://localhost:8545")
	privateKey, _ := crypto.HexToECDSA("d34772897e5d6bc952cff1094945d6d05bca81decd773f0cfb3575fbc4a73493")
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//fmt.Println(fromaddress)
	address := common.HexToAddress("0x72eB05f0D23173dA7A5Db61876Bec0D9702ccAb8")
	to_address := common.HexToAddress("0xf5fBB766074124A574fc9aFaF9c9f139e7efB981")
	instance, err := wallet.NewWallet(address, client)
	if err != nil {
		fmt.Println("error creating instance")
		log.Fatal(err)
	}
	value := big.NewInt(99)
	tx, err := instance.ExecuteCall(auth, to_address, value, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870
	time.Sleep(10 * time.Second)
	// number, _ := instance.Gettestnum(nil)
	// fmt.Println("new Number: ", number)
}
```
