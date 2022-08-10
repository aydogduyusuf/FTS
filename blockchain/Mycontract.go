// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

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
)

// MycontractData is an auto generated low-level Go binding around an user-defined struct.
type MycontractData struct {
	DataSize      *big.Int
	Time          *big.Int
	WalletAddress common.Address
	Price         *big.Int
}

// BlockchainMetaData contains all meta data concerning the Blockchain contract.
var BlockchainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dataSize\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_uniqueCode\",\"type\":\"string\"}],\"name\":\"addData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"basePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"data\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dataSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uniqueCode\",\"type\":\"string\"}],\"name\":\"getData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"dataSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"walletAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"internalType\":\"structMycontract.Data\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052662386f26fc1000060015534801561001b57600080fd5b5061003861002d61003d60201b60201c565b61004560201b60201c565b610109565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610c9f806101186000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063a2b40d191161005b578063a2b40d1914610104578063ae55c88814610120578063c7876ea414610150578063f2fde38b1461016e57610088565b806304f6748c1461008d57806342413315146100c0578063715018a6146100dc5780638da5cb5b146100e6575b600080fd5b6100a760048036038101906100a2919061077d565b61018a565b6040516100b79493929190610820565b60405180910390f35b6100da60048036038101906100d59190610891565b6101f0565b005b6100e46102b3565b005b6100ee6102c7565b6040516100fb91906108ed565b60405180910390f35b61011e60048036038101906101199190610908565b6102f0565b005b61013a6004803603810190610135919061077d565b610302565b60405161014791906109a8565b60405180910390f35b610158610412565b60405161016591906109c3565b60405180910390f35b61018860048036038101906101839190610a0a565b610418565b005b6002818051602081018201805184825260208301602085012081835280955050505050506000915090508060000154908060010154908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154905084565b60405180608001604052808381526020014281526020013373ffffffffffffffffffffffffffffffffffffffff1681526020016001548152506002826040516102399190610ab1565b9081526020016040518091039020600082015181600001556020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550606082015181600301559050505050565b6102bb61049b565b6102c56000610519565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6102f861049b565b8060018190555050565b61030a6105e5565b600060028360405161031c9190610ab1565b9081526020016040518091039020600101540361036e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036590610b4b565b60405180910390fd5b60028260405161037e9190610ab1565b908152602001604051809103902060405180608001604052908160008201548152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820154815250509050919050565b60015481565b61042061049b565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361048f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048690610bdd565b60405180910390fd5b61049881610519565b50565b6104a36105dd565b73ffffffffffffffffffffffffffffffffffffffff166104c16102c7565b73ffffffffffffffffffffffffffffffffffffffff1614610517576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050e90610c49565b60405180910390fd5b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600033905090565b60405180608001604052806000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61068a82610641565b810181811067ffffffffffffffff821117156106a9576106a8610652565b5b80604052505050565b60006106bc610623565b90506106c88282610681565b919050565b600067ffffffffffffffff8211156106e8576106e7610652565b5b6106f182610641565b9050602081019050919050565b82818337600083830152505050565b600061072061071b846106cd565b6106b2565b90508281526020810184848401111561073c5761073b61063c565b5b6107478482856106fe565b509392505050565b600082601f83011261076457610763610637565b5b813561077484826020860161070d565b91505092915050565b6000602082840312156107935761079261062d565b5b600082013567ffffffffffffffff8111156107b1576107b0610632565b5b6107bd8482850161074f565b91505092915050565b6000819050919050565b6107d9816107c6565b82525050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061080a826107df565b9050919050565b61081a816107ff565b82525050565b600060808201905061083560008301876107d0565b61084260208301866107d0565b61084f6040830185610811565b61085c60608301846107d0565b95945050505050565b61086e816107c6565b811461087957600080fd5b50565b60008135905061088b81610865565b92915050565b600080604083850312156108a8576108a761062d565b5b60006108b68582860161087c565b925050602083013567ffffffffffffffff8111156108d7576108d6610632565b5b6108e38582860161074f565b9150509250929050565b60006020820190506109026000830184610811565b92915050565b60006020828403121561091e5761091d61062d565b5b600061092c8482850161087c565b91505092915050565b61093e816107c6565b82525050565b61094d816107ff565b82525050565b6080820160008201516109696000850182610935565b50602082015161097c6020850182610935565b50604082015161098f6040850182610944565b5060608201516109a26060850182610935565b50505050565b60006080820190506109bd6000830184610953565b92915050565b60006020820190506109d860008301846107d0565b92915050565b6109e7816107ff565b81146109f257600080fd5b50565b600081359050610a04816109de565b92915050565b600060208284031215610a2057610a1f61062d565b5b6000610a2e848285016109f5565b91505092915050565b600081519050919050565b600081905092915050565b60005b83811015610a6b578082015181840152602081019050610a50565b83811115610a7a576000848401525b50505050565b6000610a8b82610a37565b610a958185610a42565b9350610aa5818560208601610a4d565b80840191505092915050565b6000610abd8284610a80565b915081905092915050565b600082825260208201905092915050565b7f5468697320756e6971756520636f646520646f65736e27742062656c6f6e672060008201527f746f20616e792064617461000000000000000000000000000000000000000000602082015250565b6000610b35602b83610ac8565b9150610b4082610ad9565b604082019050919050565b60006020820190508181036000830152610b6481610b28565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000610bc7602683610ac8565b9150610bd282610b6b565b604082019050919050565b60006020820190508181036000830152610bf681610bba565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000610c33602083610ac8565b9150610c3e82610bfd565b602082019050919050565b60006020820190508181036000830152610c6281610c26565b905091905056fea26469706673582212201b13a7dba364e921e177c57d24e73b146b626f08b97324b94c472cd9a4cae33164736f6c634300080f0033",
}

// BlockchainABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockchainMetaData.ABI instead.
var BlockchainABI = BlockchainMetaData.ABI

// BlockchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlockchainMetaData.Bin instead.
var BlockchainBin = BlockchainMetaData.Bin

// DeployBlockchain deploys a new Ethereum contract, binding an instance of Blockchain to it.
func DeployBlockchain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Blockchain, error) {
	parsed, err := BlockchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlockchainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Blockchain{BlockchainCaller: BlockchainCaller{contract: contract}, BlockchainTransactor: BlockchainTransactor{contract: contract}, BlockchainFilterer: BlockchainFilterer{contract: contract}}, nil
}

// Blockchain is an auto generated Go binding around an Ethereum contract.
type Blockchain struct {
	BlockchainCaller     // Read-only binding to the contract
	BlockchainTransactor // Write-only binding to the contract
	BlockchainFilterer   // Log filterer for contract events
}

// BlockchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockchainSession struct {
	Contract     *Blockchain       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlockchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockchainCallerSession struct {
	Contract *BlockchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BlockchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockchainTransactorSession struct {
	Contract     *BlockchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BlockchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockchainRaw struct {
	Contract *Blockchain // Generic contract binding to access the raw methods on
}

// BlockchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockchainCallerRaw struct {
	Contract *BlockchainCaller // Generic read-only contract binding to access the raw methods on
}

// BlockchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockchainTransactorRaw struct {
	Contract *BlockchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockchain creates a new instance of Blockchain, bound to a specific deployed contract.
func NewBlockchain(address common.Address, backend bind.ContractBackend) (*Blockchain, error) {
	contract, err := bindBlockchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blockchain{BlockchainCaller: BlockchainCaller{contract: contract}, BlockchainTransactor: BlockchainTransactor{contract: contract}, BlockchainFilterer: BlockchainFilterer{contract: contract}}, nil
}

// NewBlockchainCaller creates a new read-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainCaller(address common.Address, caller bind.ContractCaller) (*BlockchainCaller, error) {
	contract, err := bindBlockchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainCaller{contract: contract}, nil
}

// NewBlockchainTransactor creates a new write-only instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockchainTransactor, error) {
	contract, err := bindBlockchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainTransactor{contract: contract}, nil
}

// NewBlockchainFilterer creates a new log filterer instance of Blockchain, bound to a specific deployed contract.
func NewBlockchainFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockchainFilterer, error) {
	contract, err := bindBlockchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockchainFilterer{contract: contract}, nil
}

// bindBlockchain binds a generic wrapper to an already deployed contract.
func bindBlockchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockchainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.BlockchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.BlockchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blockchain *BlockchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blockchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blockchain *BlockchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blockchain *BlockchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blockchain.Contract.contract.Transact(opts, method, params...)
}

// BasePrice is a free data retrieval call binding the contract method 0xc7876ea4.
//
// Solidity: function basePrice() view returns(uint256)
func (_Blockchain *BlockchainCaller) BasePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "basePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BasePrice is a free data retrieval call binding the contract method 0xc7876ea4.
//
// Solidity: function basePrice() view returns(uint256)
func (_Blockchain *BlockchainSession) BasePrice() (*big.Int, error) {
	return _Blockchain.Contract.BasePrice(&_Blockchain.CallOpts)
}

// BasePrice is a free data retrieval call binding the contract method 0xc7876ea4.
//
// Solidity: function basePrice() view returns(uint256)
func (_Blockchain *BlockchainCallerSession) BasePrice() (*big.Int, error) {
	return _Blockchain.Contract.BasePrice(&_Blockchain.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x04f6748c.
//
// Solidity: function data(string ) view returns(uint256 dataSize, uint256 time, address walletAddress, uint256 price)
func (_Blockchain *BlockchainCaller) Data(opts *bind.CallOpts, arg0 string) (struct {
	DataSize      *big.Int
	Time          *big.Int
	WalletAddress common.Address
	Price         *big.Int
}, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "data", arg0)

	outstruct := new(struct {
		DataSize      *big.Int
		Time          *big.Int
		WalletAddress common.Address
		Price         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DataSize = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Time = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.WalletAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Data is a free data retrieval call binding the contract method 0x04f6748c.
//
// Solidity: function data(string ) view returns(uint256 dataSize, uint256 time, address walletAddress, uint256 price)
func (_Blockchain *BlockchainSession) Data(arg0 string) (struct {
	DataSize      *big.Int
	Time          *big.Int
	WalletAddress common.Address
	Price         *big.Int
}, error) {
	return _Blockchain.Contract.Data(&_Blockchain.CallOpts, arg0)
}

// Data is a free data retrieval call binding the contract method 0x04f6748c.
//
// Solidity: function data(string ) view returns(uint256 dataSize, uint256 time, address walletAddress, uint256 price)
func (_Blockchain *BlockchainCallerSession) Data(arg0 string) (struct {
	DataSize      *big.Int
	Time          *big.Int
	WalletAddress common.Address
	Price         *big.Int
}, error) {
	return _Blockchain.Contract.Data(&_Blockchain.CallOpts, arg0)
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string _uniqueCode) view returns((uint256,uint256,address,uint256))
func (_Blockchain *BlockchainCaller) GetData(opts *bind.CallOpts, _uniqueCode string) (MycontractData, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "getData", _uniqueCode)

	if err != nil {
		return *new(MycontractData), err
	}

	out0 := *abi.ConvertType(out[0], new(MycontractData)).(*MycontractData)

	return out0, err

}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string _uniqueCode) view returns((uint256,uint256,address,uint256))
func (_Blockchain *BlockchainSession) GetData(_uniqueCode string) (MycontractData, error) {
	return _Blockchain.Contract.GetData(&_Blockchain.CallOpts, _uniqueCode)
}

// GetData is a free data retrieval call binding the contract method 0xae55c888.
//
// Solidity: function getData(string _uniqueCode) view returns((uint256,uint256,address,uint256))
func (_Blockchain *BlockchainCallerSession) GetData(_uniqueCode string) (MycontractData, error) {
	return _Blockchain.Contract.GetData(&_Blockchain.CallOpts, _uniqueCode)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockchain *BlockchainCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blockchain.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockchain *BlockchainSession) Owner() (common.Address, error) {
	return _Blockchain.Contract.Owner(&_Blockchain.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blockchain *BlockchainCallerSession) Owner() (common.Address, error) {
	return _Blockchain.Contract.Owner(&_Blockchain.CallOpts)
}

// AddData is a paid mutator transaction binding the contract method 0x42413315.
//
// Solidity: function addData(uint256 _dataSize, string _uniqueCode) returns()
func (_Blockchain *BlockchainTransactor) AddData(opts *bind.TransactOpts, _dataSize *big.Int, _uniqueCode string) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "addData", _dataSize, _uniqueCode)
}

// AddData is a paid mutator transaction binding the contract method 0x42413315.
//
// Solidity: function addData(uint256 _dataSize, string _uniqueCode) returns()
func (_Blockchain *BlockchainSession) AddData(_dataSize *big.Int, _uniqueCode string) (*types.Transaction, error) {
	return _Blockchain.Contract.AddData(&_Blockchain.TransactOpts, _dataSize, _uniqueCode)
}

// AddData is a paid mutator transaction binding the contract method 0x42413315.
//
// Solidity: function addData(uint256 _dataSize, string _uniqueCode) returns()
func (_Blockchain *BlockchainTransactorSession) AddData(_dataSize *big.Int, _uniqueCode string) (*types.Transaction, error) {
	return _Blockchain.Contract.AddData(&_Blockchain.TransactOpts, _dataSize, _uniqueCode)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(uint256 newPrice) returns()
func (_Blockchain *BlockchainTransactor) ChangePrice(opts *bind.TransactOpts, newPrice *big.Int) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "changePrice", newPrice)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(uint256 newPrice) returns()
func (_Blockchain *BlockchainSession) ChangePrice(newPrice *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.ChangePrice(&_Blockchain.TransactOpts, newPrice)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xa2b40d19.
//
// Solidity: function changePrice(uint256 newPrice) returns()
func (_Blockchain *BlockchainTransactorSession) ChangePrice(newPrice *big.Int) (*types.Transaction, error) {
	return _Blockchain.Contract.ChangePrice(&_Blockchain.TransactOpts, newPrice)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockchain *BlockchainTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockchain *BlockchainSession) RenounceOwnership() (*types.Transaction, error) {
	return _Blockchain.Contract.RenounceOwnership(&_Blockchain.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blockchain *BlockchainTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Blockchain.Contract.RenounceOwnership(&_Blockchain.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockchain *BlockchainTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Blockchain.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockchain *BlockchainSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.TransferOwnership(&_Blockchain.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blockchain *BlockchainTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blockchain.Contract.TransferOwnership(&_Blockchain.TransactOpts, newOwner)
}

// BlockchainOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Blockchain contract.
type BlockchainOwnershipTransferredIterator struct {
	Event *BlockchainOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlockchainOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainOwnershipTransferred)
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
		it.Event = new(BlockchainOwnershipTransferred)
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
func (it *BlockchainOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainOwnershipTransferred represents a OwnershipTransferred event raised by the Blockchain contract.
type BlockchainOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockchain *BlockchainFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlockchainOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockchain.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainOwnershipTransferredIterator{contract: _Blockchain.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockchain *BlockchainFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockchainOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blockchain.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainOwnershipTransferred)
				if err := _Blockchain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blockchain *BlockchainFilterer) ParseOwnershipTransferred(log types.Log) (*BlockchainOwnershipTransferred, error) {
	event := new(BlockchainOwnershipTransferred)
	if err := _Blockchain.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
