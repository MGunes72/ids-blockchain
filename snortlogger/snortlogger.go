// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package snortlogger

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

// SnortloggerMetaData contains all meta data concerning the Snortlogger contract.
var SnortloggerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"AlertLogged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"alerts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getAlert\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"hash\",\"type\":\"string\"}],\"name\":\"logAlert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SnortloggerABI is the input ABI used to generate the binding from.
// Deprecated: Use SnortloggerMetaData.ABI instead.
var SnortloggerABI = SnortloggerMetaData.ABI

// Snortlogger is an auto generated Go binding around an Ethereum contract.
type Snortlogger struct {
	SnortloggerCaller     // Read-only binding to the contract
	SnortloggerTransactor // Write-only binding to the contract
	SnortloggerFilterer   // Log filterer for contract events
}

// SnortloggerCaller is an auto generated read-only Go binding around an Ethereum contract.
type SnortloggerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnortloggerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SnortloggerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnortloggerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SnortloggerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnortloggerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SnortloggerSession struct {
	Contract     *Snortlogger      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnortloggerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SnortloggerCallerSession struct {
	Contract *SnortloggerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SnortloggerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SnortloggerTransactorSession struct {
	Contract     *SnortloggerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SnortloggerRaw is an auto generated low-level Go binding around an Ethereum contract.
type SnortloggerRaw struct {
	Contract *Snortlogger // Generic contract binding to access the raw methods on
}

// SnortloggerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SnortloggerCallerRaw struct {
	Contract *SnortloggerCaller // Generic read-only contract binding to access the raw methods on
}

// SnortloggerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SnortloggerTransactorRaw struct {
	Contract *SnortloggerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSnortlogger creates a new instance of Snortlogger, bound to a specific deployed contract.
func NewSnortlogger(address common.Address, backend bind.ContractBackend) (*Snortlogger, error) {
	contract, err := bindSnortlogger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Snortlogger{SnortloggerCaller: SnortloggerCaller{contract: contract}, SnortloggerTransactor: SnortloggerTransactor{contract: contract}, SnortloggerFilterer: SnortloggerFilterer{contract: contract}}, nil
}

// NewSnortloggerCaller creates a new read-only instance of Snortlogger, bound to a specific deployed contract.
func NewSnortloggerCaller(address common.Address, caller bind.ContractCaller) (*SnortloggerCaller, error) {
	contract, err := bindSnortlogger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SnortloggerCaller{contract: contract}, nil
}

// NewSnortloggerTransactor creates a new write-only instance of Snortlogger, bound to a specific deployed contract.
func NewSnortloggerTransactor(address common.Address, transactor bind.ContractTransactor) (*SnortloggerTransactor, error) {
	contract, err := bindSnortlogger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SnortloggerTransactor{contract: contract}, nil
}

// NewSnortloggerFilterer creates a new log filterer instance of Snortlogger, bound to a specific deployed contract.
func NewSnortloggerFilterer(address common.Address, filterer bind.ContractFilterer) (*SnortloggerFilterer, error) {
	contract, err := bindSnortlogger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SnortloggerFilterer{contract: contract}, nil
}

// bindSnortlogger binds a generic wrapper to an already deployed contract.
func bindSnortlogger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SnortloggerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Snortlogger *SnortloggerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Snortlogger.Contract.SnortloggerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Snortlogger *SnortloggerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snortlogger.Contract.SnortloggerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Snortlogger *SnortloggerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Snortlogger.Contract.SnortloggerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Snortlogger *SnortloggerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Snortlogger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Snortlogger *SnortloggerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Snortlogger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Snortlogger *SnortloggerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Snortlogger.Contract.contract.Transact(opts, method, params...)
}

// Alerts is a free data retrieval call binding the contract method 0x171d073a.
//
// Solidity: function alerts(uint256 ) view returns(address sender, string hash)
func (_Snortlogger *SnortloggerCaller) Alerts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Sender common.Address
	Hash   string
}, error) {
	var out []interface{}
	err := _Snortlogger.contract.Call(opts, &out, "alerts", arg0)

	outstruct := new(struct {
		Sender common.Address
		Hash   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Sender = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Hash = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Alerts is a free data retrieval call binding the contract method 0x171d073a.
//
// Solidity: function alerts(uint256 ) view returns(address sender, string hash)
func (_Snortlogger *SnortloggerSession) Alerts(arg0 *big.Int) (struct {
	Sender common.Address
	Hash   string
}, error) {
	return _Snortlogger.Contract.Alerts(&_Snortlogger.CallOpts, arg0)
}

// Alerts is a free data retrieval call binding the contract method 0x171d073a.
//
// Solidity: function alerts(uint256 ) view returns(address sender, string hash)
func (_Snortlogger *SnortloggerCallerSession) Alerts(arg0 *big.Int) (struct {
	Sender common.Address
	Hash   string
}, error) {
	return _Snortlogger.Contract.Alerts(&_Snortlogger.CallOpts, arg0)
}

// GetAlert is a free data retrieval call binding the contract method 0x85081886.
//
// Solidity: function getAlert(uint256 id) view returns(address, string)
func (_Snortlogger *SnortloggerCaller) GetAlert(opts *bind.CallOpts, id *big.Int) (common.Address, string, error) {
	var out []interface{}
	err := _Snortlogger.contract.Call(opts, &out, "getAlert", id)

	if err != nil {
		return *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetAlert is a free data retrieval call binding the contract method 0x85081886.
//
// Solidity: function getAlert(uint256 id) view returns(address, string)
func (_Snortlogger *SnortloggerSession) GetAlert(id *big.Int) (common.Address, string, error) {
	return _Snortlogger.Contract.GetAlert(&_Snortlogger.CallOpts, id)
}

// GetAlert is a free data retrieval call binding the contract method 0x85081886.
//
// Solidity: function getAlert(uint256 id) view returns(address, string)
func (_Snortlogger *SnortloggerCallerSession) GetAlert(id *big.Int) (common.Address, string, error) {
	return _Snortlogger.Contract.GetAlert(&_Snortlogger.CallOpts, id)
}

// GetAllIds is a free data retrieval call binding the contract method 0xaaa44e5c.
//
// Solidity: function getAllIds() view returns(uint256[])
func (_Snortlogger *SnortloggerCaller) GetAllIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Snortlogger.contract.Call(opts, &out, "getAllIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllIds is a free data retrieval call binding the contract method 0xaaa44e5c.
//
// Solidity: function getAllIds() view returns(uint256[])
func (_Snortlogger *SnortloggerSession) GetAllIds() ([]*big.Int, error) {
	return _Snortlogger.Contract.GetAllIds(&_Snortlogger.CallOpts)
}

// GetAllIds is a free data retrieval call binding the contract method 0xaaa44e5c.
//
// Solidity: function getAllIds() view returns(uint256[])
func (_Snortlogger *SnortloggerCallerSession) GetAllIds() ([]*big.Int, error) {
	return _Snortlogger.Contract.GetAllIds(&_Snortlogger.CallOpts)
}

// Ids is a free data retrieval call binding the contract method 0xfac333ac.
//
// Solidity: function ids(uint256 ) view returns(uint256)
func (_Snortlogger *SnortloggerCaller) Ids(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Snortlogger.contract.Call(opts, &out, "ids", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Ids is a free data retrieval call binding the contract method 0xfac333ac.
//
// Solidity: function ids(uint256 ) view returns(uint256)
func (_Snortlogger *SnortloggerSession) Ids(arg0 *big.Int) (*big.Int, error) {
	return _Snortlogger.Contract.Ids(&_Snortlogger.CallOpts, arg0)
}

// Ids is a free data retrieval call binding the contract method 0xfac333ac.
//
// Solidity: function ids(uint256 ) view returns(uint256)
func (_Snortlogger *SnortloggerCallerSession) Ids(arg0 *big.Int) (*big.Int, error) {
	return _Snortlogger.Contract.Ids(&_Snortlogger.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Snortlogger *SnortloggerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Snortlogger.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Snortlogger *SnortloggerSession) Owner() (common.Address, error) {
	return _Snortlogger.Contract.Owner(&_Snortlogger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Snortlogger *SnortloggerCallerSession) Owner() (common.Address, error) {
	return _Snortlogger.Contract.Owner(&_Snortlogger.CallOpts)
}

// LogAlert is a paid mutator transaction binding the contract method 0x75725fc6.
//
// Solidity: function logAlert(uint256 id, string hash) returns()
func (_Snortlogger *SnortloggerTransactor) LogAlert(opts *bind.TransactOpts, id *big.Int, hash string) (*types.Transaction, error) {
	return _Snortlogger.contract.Transact(opts, "logAlert", id, hash)
}

// LogAlert is a paid mutator transaction binding the contract method 0x75725fc6.
//
// Solidity: function logAlert(uint256 id, string hash) returns()
func (_Snortlogger *SnortloggerSession) LogAlert(id *big.Int, hash string) (*types.Transaction, error) {
	return _Snortlogger.Contract.LogAlert(&_Snortlogger.TransactOpts, id, hash)
}

// LogAlert is a paid mutator transaction binding the contract method 0x75725fc6.
//
// Solidity: function logAlert(uint256 id, string hash) returns()
func (_Snortlogger *SnortloggerTransactorSession) LogAlert(id *big.Int, hash string) (*types.Transaction, error) {
	return _Snortlogger.Contract.LogAlert(&_Snortlogger.TransactOpts, id, hash)
}

// SnortloggerAlertLoggedIterator is returned from FilterAlertLogged and is used to iterate over the raw logs and unpacked data for AlertLogged events raised by the Snortlogger contract.
type SnortloggerAlertLoggedIterator struct {
	Event *SnortloggerAlertLogged // Event containing the contract specifics and raw log

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
func (it *SnortloggerAlertLoggedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SnortloggerAlertLogged)
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
		it.Event = new(SnortloggerAlertLogged)
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
func (it *SnortloggerAlertLoggedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SnortloggerAlertLoggedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SnortloggerAlertLogged represents a AlertLogged event raised by the Snortlogger contract.
type SnortloggerAlertLogged struct {
	Id     *big.Int
	Sender common.Address
	Hash   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAlertLogged is a free log retrieval operation binding the contract event 0x925bda36ee4169eebc888b3b593fb5f32760175c6c8e32040ff4cbd4a109a212.
//
// Solidity: event AlertLogged(uint256 id, address sender, string hash)
func (_Snortlogger *SnortloggerFilterer) FilterAlertLogged(opts *bind.FilterOpts) (*SnortloggerAlertLoggedIterator, error) {

	logs, sub, err := _Snortlogger.contract.FilterLogs(opts, "AlertLogged")
	if err != nil {
		return nil, err
	}
	return &SnortloggerAlertLoggedIterator{contract: _Snortlogger.contract, event: "AlertLogged", logs: logs, sub: sub}, nil
}

// WatchAlertLogged is a free log subscription operation binding the contract event 0x925bda36ee4169eebc888b3b593fb5f32760175c6c8e32040ff4cbd4a109a212.
//
// Solidity: event AlertLogged(uint256 id, address sender, string hash)
func (_Snortlogger *SnortloggerFilterer) WatchAlertLogged(opts *bind.WatchOpts, sink chan<- *SnortloggerAlertLogged) (event.Subscription, error) {

	logs, sub, err := _Snortlogger.contract.WatchLogs(opts, "AlertLogged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SnortloggerAlertLogged)
				if err := _Snortlogger.contract.UnpackLog(event, "AlertLogged", log); err != nil {
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

// ParseAlertLogged is a log parse operation binding the contract event 0x925bda36ee4169eebc888b3b593fb5f32760175c6c8e32040ff4cbd4a109a212.
//
// Solidity: event AlertLogged(uint256 id, address sender, string hash)
func (_Snortlogger *SnortloggerFilterer) ParseAlertLogged(log types.Log) (*SnortloggerAlertLogged, error) {
	event := new(SnortloggerAlertLogged)
	if err := _Snortlogger.contract.UnpackLog(event, "AlertLogged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
