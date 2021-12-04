// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fomo

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

// MixinProcessMetaData contains all meta data concerning the MixinProcess contract.
var MixinProcessMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"MixinEvent\",\"inputs\":[{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MixinTransaction\",\"inputs\":[{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"AMOUNT\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint128\",\"name\":\"\",\"internalType\":\"uint128\"}],\"name\":\"ASSET\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"BIDDER\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"END\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint64\",\"name\":\"\",\"internalType\":\"uint64\"}],\"name\":\"GRACE\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint64\",\"name\":\"\",\"internalType\":\"uint64\"}],\"name\":\"NONCE\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint128\",\"name\":\"\",\"internalType\":\"uint128\"}],\"name\":\"PID\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint64\",\"name\":\"\",\"internalType\":\"uint64\"}],\"name\":\"TIMESTAMP\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"custodian\",\"inputs\":[{\"type\":\"uint128\",\"name\":\"\",\"internalType\":\"uint128\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes\",\"name\":\"\",\"internalType\":\"bytes\"}],\"name\":\"members\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"mixin\",\"inputs\":[{\"type\":\"bytes\",\"name\":\"raw\",\"internalType\":\"bytes\"}]}]",
}

// MixinProcessABI is the input ABI used to generate the binding from.
// Deprecated: Use MixinProcessMetaData.ABI instead.
var MixinProcessABI = MixinProcessMetaData.ABI

// MixinProcess is an auto generated Go binding around an Ethereum contract.
type MixinProcess struct {
	MixinProcessCaller     // Read-only binding to the contract
	MixinProcessTransactor // Write-only binding to the contract
	MixinProcessFilterer   // Log filterer for contract events
}

// MixinProcessCaller is an auto generated read-only Go binding around an Ethereum contract.
type MixinProcessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MixinProcessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MixinProcessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MixinProcessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MixinProcessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MixinProcessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MixinProcessSession struct {
	Contract     *MixinProcess     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MixinProcessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MixinProcessCallerSession struct {
	Contract *MixinProcessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MixinProcessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MixinProcessTransactorSession struct {
	Contract     *MixinProcessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MixinProcessRaw is an auto generated low-level Go binding around an Ethereum contract.
type MixinProcessRaw struct {
	Contract *MixinProcess // Generic contract binding to access the raw methods on
}

// MixinProcessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MixinProcessCallerRaw struct {
	Contract *MixinProcessCaller // Generic read-only contract binding to access the raw methods on
}

// MixinProcessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MixinProcessTransactorRaw struct {
	Contract *MixinProcessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMixinProcess creates a new instance of MixinProcess, bound to a specific deployed contract.
func NewMixinProcess(address common.Address, backend bind.ContractBackend) (*MixinProcess, error) {
	contract, err := bindMixinProcess(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MixinProcess{MixinProcessCaller: MixinProcessCaller{contract: contract}, MixinProcessTransactor: MixinProcessTransactor{contract: contract}, MixinProcessFilterer: MixinProcessFilterer{contract: contract}}, nil
}

// NewMixinProcessCaller creates a new read-only instance of MixinProcess, bound to a specific deployed contract.
func NewMixinProcessCaller(address common.Address, caller bind.ContractCaller) (*MixinProcessCaller, error) {
	contract, err := bindMixinProcess(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MixinProcessCaller{contract: contract}, nil
}

// NewMixinProcessTransactor creates a new write-only instance of MixinProcess, bound to a specific deployed contract.
func NewMixinProcessTransactor(address common.Address, transactor bind.ContractTransactor) (*MixinProcessTransactor, error) {
	contract, err := bindMixinProcess(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MixinProcessTransactor{contract: contract}, nil
}

// NewMixinProcessFilterer creates a new log filterer instance of MixinProcess, bound to a specific deployed contract.
func NewMixinProcessFilterer(address common.Address, filterer bind.ContractFilterer) (*MixinProcessFilterer, error) {
	contract, err := bindMixinProcess(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MixinProcessFilterer{contract: contract}, nil
}

// bindMixinProcess binds a generic wrapper to an already deployed contract.
func bindMixinProcess(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MixinProcessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MixinProcess *MixinProcessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MixinProcess.Contract.MixinProcessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MixinProcess *MixinProcessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MixinProcess.Contract.MixinProcessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MixinProcess *MixinProcessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MixinProcess.Contract.MixinProcessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MixinProcess *MixinProcessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MixinProcess.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MixinProcess *MixinProcessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MixinProcess.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MixinProcess *MixinProcessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MixinProcess.Contract.contract.Transact(opts, method, params...)
}

// AMOUNT is a free data retrieval call binding the contract method 0xd1789176.
//
// Solidity: function AMOUNT() view returns(uint256)
func (_MixinProcess *MixinProcessCaller) AMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MixinProcess.contract.Call(opts, &out, "AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AMOUNT is a free data retrieval call binding the contract method 0xd1789176.
//
// Solidity: function AMOUNT() view returns(uint256)
func (_MixinProcess *MixinProcessSession) AMOUNT() (*big.Int, error) {
	return _MixinProcess.Contract.AMOUNT(&_MixinProcess.CallOpts)
}

// AMOUNT is a free data retrieval call binding the contract method 0xd1789176.
//
// Solidity: function AMOUNT() view returns(uint256)
func (_MixinProcess *MixinProcessCallerSession) AMOUNT() (*big.Int, error) {
	return _MixinProcess.Contract.AMOUNT(&_MixinProcess.CallOpts)
}

// ASSET is a free data retrieval call binding the contract method 0x4800d97f.
//
// Solidity: function ASSET() view returns(uint128)
func (_MixinProcess *MixinProcessCaller) ASSET(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MixinProcess.contract.Call(opts, &out, "ASSET")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ASSET is a free data retrieval call binding the contract method 0x4800d97f.
//
// Solidity: function ASSET() view returns(uint128)
func (_MixinProcess *MixinProcessSession) ASSET() (*big.Int, error) {
	return _MixinProcess.Contract.ASSET(&_MixinProcess.CallOpts)
}

// ASSET is a free data retrieval call binding the contract method 0x4800d97f.
//
// Solidity: function ASSET() view returns(uint128)
func (_MixinProcess *MixinProcessCallerSession) ASSET() (*big.Int, error) {
	return _MixinProcess.Contract.ASSET(&_MixinProcess.CallOpts)
}

// BIDDER is a free data retrieval call binding the contract method 0x7a45bcc4.
//
// Solidity: function BIDDER() view returns(address)
func (_MixinProcess *MixinProcessCaller) BIDDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MixinProcess.contract.Call(opts, &out, "BIDDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BIDDER is a free data retrieval call binding the contract method 0x7a45bcc4.
//
// Solidity: function BIDDER() view returns(address)
func (_MixinProcess *MixinProcessSession) BIDDER() (common.Address, error) {
	return _MixinProcess.Contract.BIDDER(&_MixinProcess.CallOpts)
}

// BIDDER is a free data retrieval call binding the contract method 0x7a45bcc4.
//
// Solidity: function BIDDER() view returns(address)
func (_MixinProcess *MixinProcessCallerSession) BIDDER() (common.Address, error) {
	return _MixinProcess.Contract.BIDDER(&_MixinProcess.CallOpts)
}

// END is a free data retrieval call binding the contract method 0xefe7a504.
//
// Solidity: function END() view returns(bool)
func (_MixinProcess *MixinProcessCaller) END(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MixinProcess.contract.Call(opts, &out, "END")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// END is a free data retrieval call binding the contract method 0xefe7a504.
//
// Solidity: function END() view returns(bool)
func (_MixinProcess *MixinProcessSession) END() (bool, error) {
	return _MixinProcess.Contract.END(&_MixinProcess.CallOpts)
}

// END is a free data retrieval call binding the contract method 0xefe7a504.
//
// Solidity: function END() view returns(bool)
func (_MixinProcess *MixinProcessCallerSession) END() (bool, error) {
	return _MixinProcess.Contract.END(&_MixinProcess.CallOpts)
}

// GRACE is a free data retrieval call binding the contract method 0x137d29d9.
//
// Solidity: function GRACE() view returns(uint64)
func (_MixinProcess *MixinProcessCaller) GRACE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MixinProcess.contract.Call(opts, &out, "GRACE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GRACE is a free data retrieval call binding the contract method 0x137d29d9.
//
// Solidity: function GRACE() view returns(uint64)
func (_MixinProcess *MixinProcessSession) GRACE() (uint64, error) {
	return _MixinProcess.Contract.GRACE(&_MixinProcess.CallOpts)
}

// GRACE is a free data retrieval call binding the contract method 0x137d29d9.
//
// Solidity: function GRACE() view returns(uint64)
func (_MixinProcess *MixinProcessCallerSession) GRACE() (uint64, error) {
	return _MixinProcess.Contract.GRACE(&_MixinProcess.CallOpts)
}

// NONCE is a free data retrieval call binding the contract method 0xe091dd1a.
//
// Solidity: function NONCE() view returns(uint64)
func (_MixinProcess *MixinProcessCaller) NONCE(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MixinProcess.contract.Call(opts, &out, "NONCE")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NONCE is a free data retrieval call binding the contract method 0xe091dd1a.
//
// Solidity: function NONCE() view returns(uint64)
func (_MixinProcess *MixinProcessSession) NONCE() (uint64, error) {
	return _MixinProcess.Contract.NONCE(&_MixinProcess.CallOpts)
}

// NONCE is a free data retrieval call binding the contract method 0xe091dd1a.
//
// Solidity: function NONCE() view returns(uint64)
func (_MixinProcess *MixinProcessCallerSession) NONCE() (uint64, error) {
	return _MixinProcess.Contract.NONCE(&_MixinProcess.CallOpts)
}

// PID is a free data retrieval call binding the contract method 0x5eaec0e4.
//
// Solidity: function PID() view returns(uint128)
func (_MixinProcess *MixinProcessCaller) PID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MixinProcess.contract.Call(opts, &out, "PID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PID is a free data retrieval call binding the contract method 0x5eaec0e4.
//
// Solidity: function PID() view returns(uint128)
func (_MixinProcess *MixinProcessSession) PID() (*big.Int, error) {
	return _MixinProcess.Contract.PID(&_MixinProcess.CallOpts)
}

// PID is a free data retrieval call binding the contract method 0x5eaec0e4.
//
// Solidity: function PID() view returns(uint128)
func (_MixinProcess *MixinProcessCallerSession) PID() (*big.Int, error) {
	return _MixinProcess.Contract.PID(&_MixinProcess.CallOpts)
}

// TIMESTAMP is a free data retrieval call binding the contract method 0x749e480d.
//
// Solidity: function TIMESTAMP() view returns(uint64)
func (_MixinProcess *MixinProcessCaller) TIMESTAMP(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MixinProcess.contract.Call(opts, &out, "TIMESTAMP")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// TIMESTAMP is a free data retrieval call binding the contract method 0x749e480d.
//
// Solidity: function TIMESTAMP() view returns(uint64)
func (_MixinProcess *MixinProcessSession) TIMESTAMP() (uint64, error) {
	return _MixinProcess.Contract.TIMESTAMP(&_MixinProcess.CallOpts)
}

// TIMESTAMP is a free data retrieval call binding the contract method 0x749e480d.
//
// Solidity: function TIMESTAMP() view returns(uint64)
func (_MixinProcess *MixinProcessCallerSession) TIMESTAMP() (uint64, error) {
	return _MixinProcess.Contract.TIMESTAMP(&_MixinProcess.CallOpts)
}

// Custodian is a free data retrieval call binding the contract method 0x2d459b6b.
//
// Solidity: function custodian(uint128 ) view returns(uint256)
func (_MixinProcess *MixinProcessCaller) Custodian(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MixinProcess.contract.Call(opts, &out, "custodian", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Custodian is a free data retrieval call binding the contract method 0x2d459b6b.
//
// Solidity: function custodian(uint128 ) view returns(uint256)
func (_MixinProcess *MixinProcessSession) Custodian(arg0 *big.Int) (*big.Int, error) {
	return _MixinProcess.Contract.Custodian(&_MixinProcess.CallOpts, arg0)
}

// Custodian is a free data retrieval call binding the contract method 0x2d459b6b.
//
// Solidity: function custodian(uint128 ) view returns(uint256)
func (_MixinProcess *MixinProcessCallerSession) Custodian(arg0 *big.Int) (*big.Int, error) {
	return _MixinProcess.Contract.Custodian(&_MixinProcess.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bytes)
func (_MixinProcess *MixinProcessCaller) Members(opts *bind.CallOpts, arg0 common.Address) ([]byte, error) {
	var out []interface{}
	err := _MixinProcess.contract.Call(opts, &out, "members", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bytes)
func (_MixinProcess *MixinProcessSession) Members(arg0 common.Address) ([]byte, error) {
	return _MixinProcess.Contract.Members(&_MixinProcess.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bytes)
func (_MixinProcess *MixinProcessCallerSession) Members(arg0 common.Address) ([]byte, error) {
	return _MixinProcess.Contract.Members(&_MixinProcess.CallOpts, arg0)
}

// Mixin is a paid mutator transaction binding the contract method 0x5cae8005.
//
// Solidity: function mixin(bytes raw) returns(bool)
func (_MixinProcess *MixinProcessTransactor) Mixin(opts *bind.TransactOpts, raw []byte) (*types.Transaction, error) {
	return _MixinProcess.contract.Transact(opts, "mixin", raw)
}

// Mixin is a paid mutator transaction binding the contract method 0x5cae8005.
//
// Solidity: function mixin(bytes raw) returns(bool)
func (_MixinProcess *MixinProcessSession) Mixin(raw []byte) (*types.Transaction, error) {
	return _MixinProcess.Contract.Mixin(&_MixinProcess.TransactOpts, raw)
}

// Mixin is a paid mutator transaction binding the contract method 0x5cae8005.
//
// Solidity: function mixin(bytes raw) returns(bool)
func (_MixinProcess *MixinProcessTransactorSession) Mixin(raw []byte) (*types.Transaction, error) {
	return _MixinProcess.Contract.Mixin(&_MixinProcess.TransactOpts, raw)
}

// MixinProcessMixinEventIterator is returned from FilterMixinEvent and is used to iterate over the raw logs and unpacked data for MixinEvent events raised by the MixinProcess contract.
type MixinProcessMixinEventIterator struct {
	Event *MixinProcessMixinEvent // Event containing the contract specifics and raw log

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
func (it *MixinProcessMixinEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MixinProcessMixinEvent)
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
		it.Event = new(MixinProcessMixinEvent)
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
func (it *MixinProcessMixinEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MixinProcessMixinEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MixinProcessMixinEvent represents a MixinEvent event raised by the MixinProcess contract.
type MixinProcessMixinEvent struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMixinEvent is a free log retrieval operation binding the contract event 0xd81b55010c63687df072fb24ad086c9e75c10adfe0526499b87c43ab75e25742.
//
// Solidity: event MixinEvent(bytes arg0)
func (_MixinProcess *MixinProcessFilterer) FilterMixinEvent(opts *bind.FilterOpts) (*MixinProcessMixinEventIterator, error) {

	logs, sub, err := _MixinProcess.contract.FilterLogs(opts, "MixinEvent")
	if err != nil {
		return nil, err
	}
	return &MixinProcessMixinEventIterator{contract: _MixinProcess.contract, event: "MixinEvent", logs: logs, sub: sub}, nil
}

// WatchMixinEvent is a free log subscription operation binding the contract event 0xd81b55010c63687df072fb24ad086c9e75c10adfe0526499b87c43ab75e25742.
//
// Solidity: event MixinEvent(bytes arg0)
func (_MixinProcess *MixinProcessFilterer) WatchMixinEvent(opts *bind.WatchOpts, sink chan<- *MixinProcessMixinEvent) (event.Subscription, error) {

	logs, sub, err := _MixinProcess.contract.WatchLogs(opts, "MixinEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MixinProcessMixinEvent)
				if err := _MixinProcess.contract.UnpackLog(event, "MixinEvent", log); err != nil {
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

// ParseMixinEvent is a log parse operation binding the contract event 0xd81b55010c63687df072fb24ad086c9e75c10adfe0526499b87c43ab75e25742.
//
// Solidity: event MixinEvent(bytes arg0)
func (_MixinProcess *MixinProcessFilterer) ParseMixinEvent(log types.Log) (*MixinProcessMixinEvent, error) {
	event := new(MixinProcessMixinEvent)
	if err := _MixinProcess.contract.UnpackLog(event, "MixinEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MixinProcessMixinTransactionIterator is returned from FilterMixinTransaction and is used to iterate over the raw logs and unpacked data for MixinTransaction events raised by the MixinProcess contract.
type MixinProcessMixinTransactionIterator struct {
	Event *MixinProcessMixinTransaction // Event containing the contract specifics and raw log

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
func (it *MixinProcessMixinTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MixinProcessMixinTransaction)
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
		it.Event = new(MixinProcessMixinTransaction)
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
func (it *MixinProcessMixinTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MixinProcessMixinTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MixinProcessMixinTransaction represents a MixinTransaction event raised by the MixinProcess contract.
type MixinProcessMixinTransaction struct {
	Arg0 []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMixinTransaction is a free log retrieval operation binding the contract event 0xdb53e751d28ed0d6e3682814bf8d23f7dd7b29c94f74a56fbb7f88e9dca9f39b.
//
// Solidity: event MixinTransaction(bytes arg0)
func (_MixinProcess *MixinProcessFilterer) FilterMixinTransaction(opts *bind.FilterOpts) (*MixinProcessMixinTransactionIterator, error) {

	logs, sub, err := _MixinProcess.contract.FilterLogs(opts, "MixinTransaction")
	if err != nil {
		return nil, err
	}
	return &MixinProcessMixinTransactionIterator{contract: _MixinProcess.contract, event: "MixinTransaction", logs: logs, sub: sub}, nil
}

// WatchMixinTransaction is a free log subscription operation binding the contract event 0xdb53e751d28ed0d6e3682814bf8d23f7dd7b29c94f74a56fbb7f88e9dca9f39b.
//
// Solidity: event MixinTransaction(bytes arg0)
func (_MixinProcess *MixinProcessFilterer) WatchMixinTransaction(opts *bind.WatchOpts, sink chan<- *MixinProcessMixinTransaction) (event.Subscription, error) {

	logs, sub, err := _MixinProcess.contract.WatchLogs(opts, "MixinTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MixinProcessMixinTransaction)
				if err := _MixinProcess.contract.UnpackLog(event, "MixinTransaction", log); err != nil {
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

// ParseMixinTransaction is a log parse operation binding the contract event 0xdb53e751d28ed0d6e3682814bf8d23f7dd7b29c94f74a56fbb7f88e9dca9f39b.
//
// Solidity: event MixinTransaction(bytes arg0)
func (_MixinProcess *MixinProcessFilterer) ParseMixinTransaction(log types.Log) (*MixinProcessMixinTransaction, error) {
	event := new(MixinProcessMixinTransaction)
	if err := _MixinProcess.contract.UnpackLog(event, "MixinTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
