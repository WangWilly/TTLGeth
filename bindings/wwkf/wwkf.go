// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wwkf

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

// WwkfMetaData contains all meta data concerning the Wwkf contract.
var WwkfMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numTokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WwkfABI is the input ABI used to generate the binding from.
// Deprecated: Use WwkfMetaData.ABI instead.
var WwkfABI = WwkfMetaData.ABI

// Wwkf is an auto generated Go binding around an Ethereum contract.
type Wwkf struct {
	WwkfCaller     // Read-only binding to the contract
	WwkfTransactor // Write-only binding to the contract
	WwkfFilterer   // Log filterer for contract events
}

// WwkfCaller is an auto generated read-only Go binding around an Ethereum contract.
type WwkfCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WwkfTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WwkfTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WwkfFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WwkfFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WwkfSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WwkfSession struct {
	Contract     *Wwkf             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WwkfCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WwkfCallerSession struct {
	Contract *WwkfCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WwkfTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WwkfTransactorSession struct {
	Contract     *WwkfTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WwkfRaw is an auto generated low-level Go binding around an Ethereum contract.
type WwkfRaw struct {
	Contract *Wwkf // Generic contract binding to access the raw methods on
}

// WwkfCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WwkfCallerRaw struct {
	Contract *WwkfCaller // Generic read-only contract binding to access the raw methods on
}

// WwkfTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WwkfTransactorRaw struct {
	Contract *WwkfTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWwkf creates a new instance of Wwkf, bound to a specific deployed contract.
func NewWwkf(address common.Address, backend bind.ContractBackend) (*Wwkf, error) {
	contract, err := bindWwkf(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wwkf{WwkfCaller: WwkfCaller{contract: contract}, WwkfTransactor: WwkfTransactor{contract: contract}, WwkfFilterer: WwkfFilterer{contract: contract}}, nil
}

// NewWwkfCaller creates a new read-only instance of Wwkf, bound to a specific deployed contract.
func NewWwkfCaller(address common.Address, caller bind.ContractCaller) (*WwkfCaller, error) {
	contract, err := bindWwkf(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WwkfCaller{contract: contract}, nil
}

// NewWwkfTransactor creates a new write-only instance of Wwkf, bound to a specific deployed contract.
func NewWwkfTransactor(address common.Address, transactor bind.ContractTransactor) (*WwkfTransactor, error) {
	contract, err := bindWwkf(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WwkfTransactor{contract: contract}, nil
}

// NewWwkfFilterer creates a new log filterer instance of Wwkf, bound to a specific deployed contract.
func NewWwkfFilterer(address common.Address, filterer bind.ContractFilterer) (*WwkfFilterer, error) {
	contract, err := bindWwkf(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WwkfFilterer{contract: contract}, nil
}

// bindWwkf binds a generic wrapper to an already deployed contract.
func bindWwkf(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WwkfMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wwkf *WwkfRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wwkf.Contract.WwkfCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wwkf *WwkfRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wwkf.Contract.WwkfTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wwkf *WwkfRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wwkf.Contract.WwkfTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wwkf *WwkfCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wwkf.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wwkf *WwkfTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wwkf.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wwkf *WwkfTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wwkf.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_Wwkf *WwkfCaller) Allowance(opts *bind.CallOpts, owner common.Address, delegate common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Wwkf.contract.Call(opts, &out, "allowance", owner, delegate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_Wwkf *WwkfSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _Wwkf.Contract.Allowance(&_Wwkf.CallOpts, owner, delegate)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address delegate) view returns(uint256)
func (_Wwkf *WwkfCallerSession) Allowance(owner common.Address, delegate common.Address) (*big.Int, error) {
	return _Wwkf.Contract.Allowance(&_Wwkf.CallOpts, owner, delegate)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_Wwkf *WwkfCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Wwkf.contract.Call(opts, &out, "balanceOf", tokenOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_Wwkf *WwkfSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Wwkf.Contract.BalanceOf(&_Wwkf.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) view returns(uint256)
func (_Wwkf *WwkfCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Wwkf.Contract.BalanceOf(&_Wwkf.CallOpts, tokenOwner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Wwkf *WwkfCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Wwkf.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Wwkf *WwkfSession) Decimals() (uint8, error) {
	return _Wwkf.Contract.Decimals(&_Wwkf.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Wwkf *WwkfCallerSession) Decimals() (uint8, error) {
	return _Wwkf.Contract.Decimals(&_Wwkf.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wwkf *WwkfCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Wwkf.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wwkf *WwkfSession) Name() (string, error) {
	return _Wwkf.Contract.Name(&_Wwkf.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Wwkf *WwkfCallerSession) Name() (string, error) {
	return _Wwkf.Contract.Name(&_Wwkf.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wwkf *WwkfCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Wwkf.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wwkf *WwkfSession) Symbol() (string, error) {
	return _Wwkf.Contract.Symbol(&_Wwkf.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Wwkf *WwkfCallerSession) Symbol() (string, error) {
	return _Wwkf.Contract.Symbol(&_Wwkf.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wwkf *WwkfCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wwkf.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wwkf *WwkfSession) TotalSupply() (*big.Int, error) {
	return _Wwkf.Contract.TotalSupply(&_Wwkf.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Wwkf *WwkfCallerSession) TotalSupply() (*big.Int, error) {
	return _Wwkf.Contract.TotalSupply(&_Wwkf.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_Wwkf *WwkfTransactor) Approve(opts *bind.TransactOpts, delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Wwkf.contract.Transact(opts, "approve", delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_Wwkf *WwkfSession) Approve(delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Wwkf.Contract.Approve(&_Wwkf.TransactOpts, delegate, numTokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address delegate, uint256 numTokens) returns(bool)
func (_Wwkf *WwkfTransactorSession) Approve(delegate common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Wwkf.Contract.Approve(&_Wwkf.TransactOpts, delegate, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_Wwkf *WwkfTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Wwkf.contract.Transact(opts, "transfer", receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_Wwkf *WwkfSession) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Wwkf.Contract.Transfer(&_Wwkf.TransactOpts, receiver, numTokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address receiver, uint256 numTokens) returns(bool)
func (_Wwkf *WwkfTransactorSession) Transfer(receiver common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Wwkf.Contract.Transfer(&_Wwkf.TransactOpts, receiver, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_Wwkf *WwkfTransactor) TransferFrom(opts *bind.TransactOpts, owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Wwkf.contract.Transact(opts, "transferFrom", owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_Wwkf *WwkfSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Wwkf.Contract.TransferFrom(&_Wwkf.TransactOpts, owner, buyer, numTokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address owner, address buyer, uint256 numTokens) returns(bool)
func (_Wwkf *WwkfTransactorSession) TransferFrom(owner common.Address, buyer common.Address, numTokens *big.Int) (*types.Transaction, error) {
	return _Wwkf.Contract.TransferFrom(&_Wwkf.TransactOpts, owner, buyer, numTokens)
}

// WwkfApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Wwkf contract.
type WwkfApprovalIterator struct {
	Event *WwkfApproval // Event containing the contract specifics and raw log

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
func (it *WwkfApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WwkfApproval)
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
		it.Event = new(WwkfApproval)
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
func (it *WwkfApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WwkfApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WwkfApproval represents a Approval event raised by the Wwkf contract.
type WwkfApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Wwkf *WwkfFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*WwkfApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Wwkf.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WwkfApprovalIterator{contract: _Wwkf.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Wwkf *WwkfFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WwkfApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Wwkf.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WwkfApproval)
				if err := _Wwkf.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Wwkf *WwkfFilterer) ParseApproval(log types.Log) (*WwkfApproval, error) {
	event := new(WwkfApproval)
	if err := _Wwkf.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WwkfTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Wwkf contract.
type WwkfTransferIterator struct {
	Event *WwkfTransfer // Event containing the contract specifics and raw log

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
func (it *WwkfTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WwkfTransfer)
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
		it.Event = new(WwkfTransfer)
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
func (it *WwkfTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WwkfTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WwkfTransfer represents a Transfer event raised by the Wwkf contract.
type WwkfTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Wwkf *WwkfFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WwkfTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Wwkf.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WwkfTransferIterator{contract: _Wwkf.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Wwkf *WwkfFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WwkfTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Wwkf.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WwkfTransfer)
				if err := _Wwkf.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Wwkf *WwkfFilterer) ParseTransfer(log types.Log) (*WwkfTransfer, error) {
	event := new(WwkfTransfer)
	if err := _Wwkf.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}