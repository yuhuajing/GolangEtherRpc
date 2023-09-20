// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wallet

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// WalletMetaData contains all meta data concerning the Wallet contract.
var WalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInitialManager\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCodeInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidEmailInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQAInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotManagerAuthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwnerAuthorized\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"inputQA\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"storedQA\",\"type\":\"string\"}],\"name\":\"QAerror\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"inputemail\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"storedemail\",\"type\":\"string\"}],\"name\":\"emailerror\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_bytes32\",\"type\":\"bytes32\"}],\"name\":\"convertByte32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_calldata\",\"type\":\"bytes\"}],\"name\":\"executeCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"initialManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"resetManaget\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newaddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_code\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_mixed_question\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_mixed_answer\",\"type\":\"string\"}],\"name\":\"resetOrforgetPassword\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_code\",\"type\":\"uint256\"}],\"name\":\"verifycode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611ca4806100206000396000f3fe60806040526004361061007f5760003560e01c80638da5cb5b1161004e5780638da5cb5b146101425780639e5d4c491461016d578063b16865581461019d578063f6bae2b1146101c657610086565b806343f93c5a14610088578063481c6a75146100b15780635ebf87db146100dc578063814494b11461011957610086565b3661008657005b005b34801561009457600080fd5b506100af60048036038101906100aa91906111e4565b6101ef565b005b3480156100bd57600080fd5b506100c6610320565b6040516100d39190611220565b60405180910390f35b3480156100e857600080fd5b5061010360048036038101906100fe9190611271565b610346565b604051610110919061132e565b60405180910390f35b34801561012557600080fd5b50610140600480360381019061013b91906114bb565b61041c565b005b34801561014e57600080fd5b506101576109f2565b6040516101649190611220565b60405180910390f35b610187600480360381019061018291906115fc565b610aaa565b60405161019491906116c5565b60405180910390f35b3480156101a957600080fd5b506101c460048036038101906101bf91906111e4565b610b2e565b005b3480156101d257600080fd5b506101ed60048036038101906101e891906116e7565b610ccc565b005b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610276576040517f7c90ebca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036102dc576040517fb4fa3fb300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000602067ffffffffffffffff8111156103655761036461135a565b5b6040519080825280601f01601f1916602001820160405280156103975781602001600182028036833780820191505090505b50905060005b6020811015610412578381602081106103b9576103b8611743565b5b1a60f81b8282815181106103d0576103cf611743565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350808061040a906117a1565b91505061039d565b5080915050919050565b600360159054906101000a900460ff166104575760008060008061043e610d9e565b935093509350935061045284848484610e55565b505050505b60008060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206040518060a00160405290816000820180546104b290611818565b80601f01602080910402602001604051908101604052809291908181526020018280546104de90611818565b801561052b5780601f106105005761010080835404028352916020019161052b565b820191906000526020600020905b81548152906001019060200180831161050e57829003601f168201915b50505050508152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820180546105a490611818565b80601f01602080910402602001604051908101604052809291908181526020018280546105d090611818565b801561061d5780601f106105f25761010080835404028352916020019161061d565b820191906000526020600020905b81548152906001019060200180831161060057829003601f168201915b5050505050815260200160048201805461063690611818565b80601f016020809104026020016040519081016040528092919081815260200182805461066290611818565b80156106af5780601f10610684576101008083540402835291602001916106af565b820191906000526020600020905b81548152906001019060200180831161069257829003601f168201915b50505050508152505090506106c8858260000151610f8e565b6107295780600001516040516106de9190611885565b6040518091039020856040516106f49190611885565b60405180910390207f52411f6568cfae5fce8e0047a965f329ee985a982c44b19791dead75a8a20d3160405160405180910390a35b6107338585610fb6565b610769576040517fd7653c9800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6107806107768484610ffd565b8260800151610f8e565b6107ea5780608001516040516107969190611885565b60405180910390206107a88484610ffd565b6040516107b59190611885565b60405180910390207fedd01afb553064ce077219e98147c10ff1bf205311942202767414ee744b77bf60405160405180910390a35b8381602001818152505085816040019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050806000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000190816108809190611a48565b506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030190816108e79190611a48565b5060808201518160040190816108fd9190611a48565b509050506000808873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600080820160006109509190611115565b60018201600090556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560038201600061098f9190611115565b60048201600061099f9190611115565b50506001856040516109b19190611885565b9081526020016040518091039020600085815260200190815260200160002060006101000a81549060ff02191690556109e986611029565b50505050505050565b6000600360149054906101000a900460ff1615610a3357600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610aa7565b6000602067ffffffffffffffff811115610a5057610a4f61135a565b5b6040519080825280601f01601f191660200182016040528015610a825781602001600182028036833780820191505090505b5090506020602d60208301303c80806020019051810190610aa39190611b58565b9150505b90565b6060610ab46109f2565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b18576040517f0f380e1b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610b2485858585611088565b9050949350505050565b610b366109f2565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b9a576040517f0f380e1b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610c00576040517fb4fa3fb300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610c88576040517fe840bea600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d53576040517f7c90ebca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60018083604051610d649190611885565b9081526020016040518091039020600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600060608060606000608067ffffffffffffffff811115610dc257610dc161135a565b5b6040519080825280601f01601f191660200182016040528015610df45781602001600182028036833780820191505090505b5090506080602d60208301303c600080600083806020019051810190610e1a9190611b9a565b809450819550829650839b5050505050610e3383610346565b9650610e3e82610346565b9550610e4981610346565b94505050505090919293565b60006040518060a00160405280858152602001600081526020018673ffffffffffffffffffffffffffffffffffffffff168152602001848152602001838152509050806000808773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000820151816000019081610eeb9190611a48565b506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506060820151816003019081610f529190611a48565b506080820151816004019081610f689190611a48565b509050506001600360156101000a81548160ff0219169083151502179055505050505050565b600081518351148015610fae575081805190602001208380519060200120145b905092915050565b6000600183604051610fc89190611885565b9081526020016040518091039020600083815260200190815260200160002060009054906101000a900460ff16905092915050565b60608282604051602001611012929190611c01565b604051602081830303815290604052905092915050565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600360146101000a81548160ff02191690831515021790555050565b606060008573ffffffffffffffffffffffffffffffffffffffff168585856040516110b4929190611c55565b60006040518083038185875af1925050503d80600081146110f1576040519150601f19603f3d011682016040523d82523d6000602084013e6110f6565b606091505b5080935081925050508061110c57815160208301fd5b50949350505050565b50805461112190611818565b6000825580601f106111335750611152565b601f0160209004906000526020600020908101906111519190611155565b5b50565b5b8082111561116e576000816000905550600101611156565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006111b182611186565b9050919050565b6111c1816111a6565b81146111cc57600080fd5b50565b6000813590506111de816111b8565b92915050565b6000602082840312156111fa576111f961117c565b5b6000611208848285016111cf565b91505092915050565b61121a816111a6565b82525050565b60006020820190506112356000830184611211565b92915050565b6000819050919050565b61124e8161123b565b811461125957600080fd5b50565b60008135905061126b81611245565b92915050565b6000602082840312156112875761128661117c565b5b60006112958482850161125c565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156112d85780820151818401526020810190506112bd565b60008484015250505050565b6000601f19601f8301169050919050565b60006113008261129e565b61130a81856112a9565b935061131a8185602086016112ba565b611323816112e4565b840191505092915050565b6000602082019050818103600083015261134881846112f5565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611392826112e4565b810181811067ffffffffffffffff821117156113b1576113b061135a565b5b80604052505050565b60006113c4611172565b90506113d08282611389565b919050565b600067ffffffffffffffff8211156113f0576113ef61135a565b5b6113f9826112e4565b9050602081019050919050565b82818337600083830152505050565b6000611428611423846113d5565b6113ba565b90508281526020810184848401111561144457611443611355565b5b61144f848285611406565b509392505050565b600082601f83011261146c5761146b611350565b5b813561147c848260208601611415565b91505092915050565b6000819050919050565b61149881611485565b81146114a357600080fd5b50565b6000813590506114b58161148f565b92915050565b60008060008060008060c087890312156114d8576114d761117c565b5b60006114e689828a016111cf565b96505060206114f789828a016111cf565b955050604087013567ffffffffffffffff81111561151857611517611181565b5b61152489828a01611457565b945050606061153589828a016114a6565b935050608087013567ffffffffffffffff81111561155657611555611181565b5b61156289828a01611457565b92505060a087013567ffffffffffffffff81111561158357611582611181565b5b61158f89828a01611457565b9150509295509295509295565b600080fd5b600080fd5b60008083601f8401126115bc576115bb611350565b5b8235905067ffffffffffffffff8111156115d9576115d861159c565b5b6020830191508360018202830111156115f5576115f46115a1565b5b9250929050565b600080600080606085870312156116165761161561117c565b5b6000611624878288016111cf565b9450506020611635878288016114a6565b935050604085013567ffffffffffffffff81111561165657611655611181565b5b611662878288016115a6565b925092505092959194509250565b600081519050919050565b600082825260208201905092915050565b600061169782611670565b6116a1818561167b565b93506116b18185602086016112ba565b6116ba816112e4565b840191505092915050565b600060208201905081810360008301526116df818461168c565b905092915050565b600080604083850312156116fe576116fd61117c565b5b600083013567ffffffffffffffff81111561171c5761171b611181565b5b61172885828601611457565b9250506020611739858286016114a6565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006117ac82611485565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036117de576117dd611772565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061183057607f821691505b602082108103611843576118426117e9565b5b50919050565b600081905092915050565b600061185f8261129e565b6118698185611849565b93506118798185602086016112ba565b80840191505092915050565b60006118918284611854565b915081905092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026118fe7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826118c1565b61190886836118c1565b95508019841693508086168417925050509392505050565b6000819050919050565b600061194561194061193b84611485565b611920565b611485565b9050919050565b6000819050919050565b61195f8361192a565b61197361196b8261194c565b8484546118ce565b825550505050565b600090565b61198861197b565b611993818484611956565b505050565b5b818110156119b7576119ac600082611980565b600181019050611999565b5050565b601f8211156119fc576119cd8161189c565b6119d6846118b1565b810160208510156119e5578190505b6119f96119f1856118b1565b830182611998565b50505b505050565b600082821c905092915050565b6000611a1f60001984600802611a01565b1980831691505092915050565b6000611a388383611a0e565b9150826002028217905092915050565b611a518261129e565b67ffffffffffffffff811115611a6a57611a6961135a565b5b611a748254611818565b611a7f8282856119bb565b600060209050601f831160018114611ab25760008415611aa0578287015190505b611aaa8582611a2c565b865550611b12565b601f198416611ac08661189c565b60005b82811015611ae857848901518255600182019150602085019450602081019050611ac3565b86831015611b055784890151611b01601f891682611a0e565b8355505b6001600288020188555050505b505050505050565b6000611b2582611186565b9050919050565b611b3581611b1a565b8114611b4057600080fd5b50565b600081519050611b5281611b2c565b92915050565b600060208284031215611b6e57611b6d61117c565b5b6000611b7c84828501611b43565b91505092915050565b600081519050611b9481611245565b92915050565b60008060008060808587031215611bb457611bb361117c565b5b6000611bc287828801611b43565b9450506020611bd387828801611b85565b9350506040611be487828801611b85565b9250506060611bf587828801611b85565b91505092959194509250565b6000611c0d8285611854565b9150611c198284611854565b91508190509392505050565b600081905092915050565b6000611c3c8385611c25565b9350611c49838584611406565b82840190509392505050565b6000611c62828486611c30565b9150819050939250505056fea2646970667358221220a95acb06c6b2fee7d99512bd37b179e52c91450650fdce2014f2fdb8b118f46e64736f6c63430008120033",
}

// WalletABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletMetaData.ABI instead.
var WalletABI = WalletMetaData.ABI

// WalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WalletMetaData.Bin instead.
var WalletBin = WalletMetaData.Bin

// DeployWallet deploys a new Ethereum contract, binding an instance of Wallet to it.
func DeployWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Wallet, error) {
	parsed, err := WalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// Wallet is an auto generated Go binding around an Ethereum contract.
type Wallet struct {
	WalletCaller     // Read-only binding to the contract
	WalletTransactor // Write-only binding to the contract
	WalletFilterer   // Log filterer for contract events
}

// WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSession struct {
	Contract     *Wallet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletCallerSession struct {
	Contract *WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletTransactorSession struct {
	Contract     *WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletRaw struct {
	Contract *Wallet // Generic contract binding to access the raw methods on
}

// WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletCallerRaw struct {
	Contract *WalletCaller // Generic read-only contract binding to access the raw methods on
}

// WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletTransactorRaw struct {
	Contract *WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWallet creates a new instance of Wallet, bound to a specific deployed contract.
func NewWallet(address common.Address, backend bind.ContractBackend) (*Wallet, error) {
	contract, err := bindWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}, WalletFilterer: WalletFilterer{contract: contract}}, nil
}

// NewWalletCaller creates a new read-only instance of Wallet, bound to a specific deployed contract.
func NewWalletCaller(address common.Address, caller bind.ContractCaller) (*WalletCaller, error) {
	contract, err := bindWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCaller{contract: contract}, nil
}

// NewWalletTransactor creates a new write-only instance of Wallet, bound to a specific deployed contract.
func NewWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletTransactor, error) {
	contract, err := bindWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletTransactor{contract: contract}, nil
}

// NewWalletFilterer creates a new log filterer instance of Wallet, bound to a specific deployed contract.
func NewWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletFilterer, error) {
	contract, err := bindWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletFilterer{contract: contract}, nil
}

// bindWallet binds a generic wrapper to an already deployed contract.
func bindWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transact(opts, method, params...)
}

// ConvertByte32ToString is a free data retrieval call binding the contract method 0x5ebf87db.
//
// Solidity: function convertByte32ToString(bytes32 _bytes32) pure returns(string)
func (_Wallet *WalletCaller) ConvertByte32ToString(opts *bind.CallOpts, _bytes32 [32]byte) (string, error) {
	var out []interface{}
	err := _Wallet.contract.Call(opts, &out, "convertByte32ToString", _bytes32)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ConvertByte32ToString is a free data retrieval call binding the contract method 0x5ebf87db.
//
// Solidity: function convertByte32ToString(bytes32 _bytes32) pure returns(string)
func (_Wallet *WalletSession) ConvertByte32ToString(_bytes32 [32]byte) (string, error) {
	return _Wallet.Contract.ConvertByte32ToString(&_Wallet.CallOpts, _bytes32)
}

// ConvertByte32ToString is a free data retrieval call binding the contract method 0x5ebf87db.
//
// Solidity: function convertByte32ToString(bytes32 _bytes32) pure returns(string)
func (_Wallet *WalletCallerSession) ConvertByte32ToString(_bytes32 [32]byte) (string, error) {
	return _Wallet.Contract.ConvertByte32ToString(&_Wallet.CallOpts, _bytes32)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Wallet *WalletCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wallet.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Wallet *WalletSession) Manager() (common.Address, error) {
	return _Wallet.Contract.Manager(&_Wallet.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Wallet *WalletCallerSession) Manager() (common.Address, error) {
	return _Wallet.Contract.Manager(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wallet *WalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wallet *WalletSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Wallet *WalletCallerSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes _calldata) payable returns(bytes)
func (_Wallet *WalletTransactor) ExecuteCall(opts *bind.TransactOpts, to common.Address, value *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "executeCall", to, value, _calldata)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes _calldata) payable returns(bytes)
func (_Wallet *WalletSession) ExecuteCall(to common.Address, value *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteCall(&_Wallet.TransactOpts, to, value, _calldata)
}

// ExecuteCall is a paid mutator transaction binding the contract method 0x9e5d4c49.
//
// Solidity: function executeCall(address to, uint256 value, bytes _calldata) payable returns(bytes)
func (_Wallet *WalletTransactorSession) ExecuteCall(to common.Address, value *big.Int, _calldata []byte) (*types.Transaction, error) {
	return _Wallet.Contract.ExecuteCall(&_Wallet.TransactOpts, to, value, _calldata)
}

// InitialManager is a paid mutator transaction binding the contract method 0xb1686558.
//
// Solidity: function initialManager(address _manager) returns()
func (_Wallet *WalletTransactor) InitialManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "initialManager", _manager)
}

// InitialManager is a paid mutator transaction binding the contract method 0xb1686558.
//
// Solidity: function initialManager(address _manager) returns()
func (_Wallet *WalletSession) InitialManager(_manager common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.InitialManager(&_Wallet.TransactOpts, _manager)
}

// InitialManager is a paid mutator transaction binding the contract method 0xb1686558.
//
// Solidity: function initialManager(address _manager) returns()
func (_Wallet *WalletTransactorSession) InitialManager(_manager common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.InitialManager(&_Wallet.TransactOpts, _manager)
}

// ResetManaget is a paid mutator transaction binding the contract method 0x43f93c5a.
//
// Solidity: function resetManaget(address _manager) returns()
func (_Wallet *WalletTransactor) ResetManaget(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "resetManaget", _manager)
}

// ResetManaget is a paid mutator transaction binding the contract method 0x43f93c5a.
//
// Solidity: function resetManaget(address _manager) returns()
func (_Wallet *WalletSession) ResetManaget(_manager common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.ResetManaget(&_Wallet.TransactOpts, _manager)
}

// ResetManaget is a paid mutator transaction binding the contract method 0x43f93c5a.
//
// Solidity: function resetManaget(address _manager) returns()
func (_Wallet *WalletTransactorSession) ResetManaget(_manager common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.ResetManaget(&_Wallet.TransactOpts, _manager)
}

// ResetOrforgetPassword is a paid mutator transaction binding the contract method 0x814494b1.
//
// Solidity: function resetOrforgetPassword(address _address, address _newaddress, string _email, uint256 _code, string _mixed_question, string _mixed_answer) returns()
func (_Wallet *WalletTransactor) ResetOrforgetPassword(opts *bind.TransactOpts, _address common.Address, _newaddress common.Address, _email string, _code *big.Int, _mixed_question string, _mixed_answer string) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "resetOrforgetPassword", _address, _newaddress, _email, _code, _mixed_question, _mixed_answer)
}

// ResetOrforgetPassword is a paid mutator transaction binding the contract method 0x814494b1.
//
// Solidity: function resetOrforgetPassword(address _address, address _newaddress, string _email, uint256 _code, string _mixed_question, string _mixed_answer) returns()
func (_Wallet *WalletSession) ResetOrforgetPassword(_address common.Address, _newaddress common.Address, _email string, _code *big.Int, _mixed_question string, _mixed_answer string) (*types.Transaction, error) {
	return _Wallet.Contract.ResetOrforgetPassword(&_Wallet.TransactOpts, _address, _newaddress, _email, _code, _mixed_question, _mixed_answer)
}

// ResetOrforgetPassword is a paid mutator transaction binding the contract method 0x814494b1.
//
// Solidity: function resetOrforgetPassword(address _address, address _newaddress, string _email, uint256 _code, string _mixed_question, string _mixed_answer) returns()
func (_Wallet *WalletTransactorSession) ResetOrforgetPassword(_address common.Address, _newaddress common.Address, _email string, _code *big.Int, _mixed_question string, _mixed_answer string) (*types.Transaction, error) {
	return _Wallet.Contract.ResetOrforgetPassword(&_Wallet.TransactOpts, _address, _newaddress, _email, _code, _mixed_question, _mixed_answer)
}

// Verifycode is a paid mutator transaction binding the contract method 0xf6bae2b1.
//
// Solidity: function verifycode(string _email, uint256 _code) returns()
func (_Wallet *WalletTransactor) Verifycode(opts *bind.TransactOpts, _email string, _code *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "verifycode", _email, _code)
}

// Verifycode is a paid mutator transaction binding the contract method 0xf6bae2b1.
//
// Solidity: function verifycode(string _email, uint256 _code) returns()
func (_Wallet *WalletSession) Verifycode(_email string, _code *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Verifycode(&_Wallet.TransactOpts, _email, _code)
}

// Verifycode is a paid mutator transaction binding the contract method 0xf6bae2b1.
//
// Solidity: function verifycode(string _email, uint256 _code) returns()
func (_Wallet *WalletTransactorSession) Verifycode(_email string, _code *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Verifycode(&_Wallet.TransactOpts, _email, _code)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Wallet *WalletTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Wallet.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Wallet *WalletSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Wallet.Contract.Fallback(&_Wallet.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Wallet *WalletTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Wallet.Contract.Fallback(&_Wallet.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Wallet *WalletTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Wallet *WalletSession) Receive() (*types.Transaction, error) {
	return _Wallet.Contract.Receive(&_Wallet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Wallet *WalletTransactorSession) Receive() (*types.Transaction, error) {
	return _Wallet.Contract.Receive(&_Wallet.TransactOpts)
}

// WalletQAerrorIterator is returned from FilterQAerror and is used to iterate over the raw logs and unpacked data for QAerror events raised by the Wallet contract.
type WalletQAerrorIterator struct {
	Event *WalletQAerror // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletQAerrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletQAerror)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletQAerror)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletQAerrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletQAerrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletQAerror represents a QAerror event raised by the Wallet contract.
type WalletQAerror struct {
	InputQA  common.Hash
	StoredQA common.Hash
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterQAerror is a free log retrieval operation binding the contract event 0xedd01afb553064ce077219e98147c10ff1bf205311942202767414ee744b77bf.
//
// Solidity: event QAerror(string indexed inputQA, string indexed storedQA)
func (_Wallet *WalletFilterer) FilterQAerror(opts *bind.FilterOpts, inputQA []string, storedQA []string) (*WalletQAerrorIterator, error) {

	var inputQARule []interface{}
	for _, inputQAItem := range inputQA {
		inputQARule = append(inputQARule, inputQAItem)
	}
	var storedQARule []interface{}
	for _, storedQAItem := range storedQA {
		storedQARule = append(storedQARule, storedQAItem)
	}

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "QAerror", inputQARule, storedQARule)
	if err != nil {
		return nil, err
	}
	return &WalletQAerrorIterator{contract: _Wallet.contract, event: "QAerror", logs: logs, sub: sub}, nil
}

// WatchQAerror is a free log subscription operation binding the contract event 0xedd01afb553064ce077219e98147c10ff1bf205311942202767414ee744b77bf.
//
// Solidity: event QAerror(string indexed inputQA, string indexed storedQA)
func (_Wallet *WalletFilterer) WatchQAerror(opts *bind.WatchOpts, sink chan<- *WalletQAerror, inputQA []string, storedQA []string) (event.Subscription, error) {

	var inputQARule []interface{}
	for _, inputQAItem := range inputQA {
		inputQARule = append(inputQARule, inputQAItem)
	}
	var storedQARule []interface{}
	for _, storedQAItem := range storedQA {
		storedQARule = append(storedQARule, storedQAItem)
	}

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "QAerror", inputQARule, storedQARule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletQAerror)
				if err := _Wallet.contract.UnpackLog(event, "QAerror", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseQAerror is a log parse operation binding the contract event 0xedd01afb553064ce077219e98147c10ff1bf205311942202767414ee744b77bf.
//
// Solidity: event QAerror(string indexed inputQA, string indexed storedQA)
func (_Wallet *WalletFilterer) ParseQAerror(log types.Log) (*WalletQAerror, error) {
	event := new(WalletQAerror)
	if err := _Wallet.contract.UnpackLog(event, "QAerror", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WalletEmailerrorIterator is returned from FilterEmailerror and is used to iterate over the raw logs and unpacked data for Emailerror events raised by the Wallet contract.
type WalletEmailerrorIterator struct {
	Event *WalletEmailerror // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *WalletEmailerrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WalletEmailerror)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(WalletEmailerror)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *WalletEmailerrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WalletEmailerrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WalletEmailerror represents a Emailerror event raised by the Wallet contract.
type WalletEmailerror struct {
	Inputemail  common.Hash
	Storedemail common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterEmailerror is a free log retrieval operation binding the contract event 0x52411f6568cfae5fce8e0047a965f329ee985a982c44b19791dead75a8a20d31.
//
// Solidity: event emailerror(string indexed inputemail, string indexed storedemail)
func (_Wallet *WalletFilterer) FilterEmailerror(opts *bind.FilterOpts, inputemail []string, storedemail []string) (*WalletEmailerrorIterator, error) {

	var inputemailRule []interface{}
	for _, inputemailItem := range inputemail {
		inputemailRule = append(inputemailRule, inputemailItem)
	}
	var storedemailRule []interface{}
	for _, storedemailItem := range storedemail {
		storedemailRule = append(storedemailRule, storedemailItem)
	}

	logs, sub, err := _Wallet.contract.FilterLogs(opts, "emailerror", inputemailRule, storedemailRule)
	if err != nil {
		return nil, err
	}
	return &WalletEmailerrorIterator{contract: _Wallet.contract, event: "emailerror", logs: logs, sub: sub}, nil
}

// WatchEmailerror is a free log subscription operation binding the contract event 0x52411f6568cfae5fce8e0047a965f329ee985a982c44b19791dead75a8a20d31.
//
// Solidity: event emailerror(string indexed inputemail, string indexed storedemail)
func (_Wallet *WalletFilterer) WatchEmailerror(opts *bind.WatchOpts, sink chan<- *WalletEmailerror, inputemail []string, storedemail []string) (event.Subscription, error) {

	var inputemailRule []interface{}
	for _, inputemailItem := range inputemail {
		inputemailRule = append(inputemailRule, inputemailItem)
	}
	var storedemailRule []interface{}
	for _, storedemailItem := range storedemail {
		storedemailRule = append(storedemailRule, storedemailItem)
	}

	logs, sub, err := _Wallet.contract.WatchLogs(opts, "emailerror", inputemailRule, storedemailRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WalletEmailerror)
				if err := _Wallet.contract.UnpackLog(event, "emailerror", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEmailerror is a log parse operation binding the contract event 0x52411f6568cfae5fce8e0047a965f329ee985a982c44b19791dead75a8a20d31.
//
// Solidity: event emailerror(string indexed inputemail, string indexed storedemail)
func (_Wallet *WalletFilterer) ParseEmailerror(log types.Log) (*WalletEmailerror, error) {
	event := new(WalletEmailerror)
	if err := _Wallet.contract.UnpackLog(event, "emailerror", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
