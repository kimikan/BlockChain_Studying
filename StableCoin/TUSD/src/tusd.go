// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tusd

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// AddressListABI is the input ABI used to generate the binding from.
const AddressListABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"onList\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_onList\",\"type\":\"bool\"}],\"name\":\"changeList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"nullValue\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"onList\",\"type\":\"bool\"}],\"name\":\"ChangeWhiteList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// AddressListBin is the compiled bytecode used for deploying new contracts.
const AddressListBin = `0x608060405234801561001057600080fd5b506040516105d23803806105d283398101604052805160208083015160008054600160a060020a031916331790559190920180519092610055916002918501906100d3565b5060036020527f3617319a054d772f909f7c479a2cebe5066e836a939412e32403c99029b92eff805491151560ff1992831617905573e14c47861b9f20a6bab730f10e8bb5d4ab420cd46000527fec2630a5e3d96a144a0cd3fb290c1336e4b6dee2d1d3140b683408e9317b34b0805490911660011790555061016e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061011457805160ff1916838001178555610141565b82800160010185558215610141579182015b82811115610141578251825591602001919060010190610126565b5061014d929150610151565b5090565b61016b91905b8082111561014d5760008155600101610157565b90565b6104558061017d6000396000f3006080604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100875780634e71e0c8146101115780636f626eb3146101285780638da5cb5b1461015d578063e30c39781461018e578063efbd9ec2146101a3578063f2fde38b146101c9575b600080fd5b34801561009357600080fd5b5061009c6101ea565b6040805160208082528351818301528351919283929083019185019080838360005b838110156100d65781810151838201526020016100be565b50505050905090810190601f1680156101035780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561011d57600080fd5b50610126610275565b005b34801561013457600080fd5b50610149600160a060020a03600435166102fd565b604080519115158252519081900360200190f35b34801561016957600080fd5b50610172610312565b60408051600160a060020a039092168252519081900360200190f35b34801561019a57600080fd5b50610172610321565b3480156101af57600080fd5b50610126600160a060020a03600435166024351515610330565b3480156101d557600080fd5b50610126600160a060020a03600435166103e3565b6002805460408051602060018416156101000260001901909316849004601f8101849004840282018401909252818152929183018282801561026d5780601f106102425761010080835404028352916020019161026d565b820191906000526020600020905b81548152906001019060200180831161025057829003601f168201915b505050505081565b600154600160a060020a0316331461028c57600080fd5b60015460008054604051600160a060020a0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b60036020526000908152604090205460ff1681565b600054600160a060020a031681565b600154600160a060020a031681565b600054600160a060020a0316331461034757600080fd5b600160a060020a038216151561035c57600080fd5b600160a060020a03821660009081526003602052604090205460ff161515811515146103df57600160a060020a038216600081815260036020908152604091829020805460ff1916851515908117909155825190815291517febf7001045f8f4b6a63ef446c723437b124accd3ba9cc203f2f369952621be599281900390910190a25b5050565b600054600160a060020a031633146103fa57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820789c448674eb488058ef8c153518a717350b7710cef5b2bbe67077c95f3921e00029`

// DeployAddressList deploys a new Ethereum contract, binding an instance of AddressList to it.
func DeployAddressList(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, nullValue bool) (common.Address, *types.Transaction, *AddressList, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressListABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressListBin), backend, _name, nullValue)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressList{AddressListCaller: AddressListCaller{contract: contract}, AddressListTransactor: AddressListTransactor{contract: contract}, AddressListFilterer: AddressListFilterer{contract: contract}}, nil
}

// AddressList is an auto generated Go binding around an Ethereum contract.
type AddressList struct {
	AddressListCaller     // Read-only binding to the contract
	AddressListTransactor // Write-only binding to the contract
	AddressListFilterer   // Log filterer for contract events
}

// AddressListCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressListSession struct {
	Contract     *AddressList      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressListCallerSession struct {
	Contract *AddressListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AddressListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressListTransactorSession struct {
	Contract     *AddressListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AddressListRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressListRaw struct {
	Contract *AddressList // Generic contract binding to access the raw methods on
}

// AddressListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressListCallerRaw struct {
	Contract *AddressListCaller // Generic read-only contract binding to access the raw methods on
}

// AddressListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressListTransactorRaw struct {
	Contract *AddressListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressList creates a new instance of AddressList, bound to a specific deployed contract.
func NewAddressList(address common.Address, backend bind.ContractBackend) (*AddressList, error) {
	contract, err := bindAddressList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressList{AddressListCaller: AddressListCaller{contract: contract}, AddressListTransactor: AddressListTransactor{contract: contract}, AddressListFilterer: AddressListFilterer{contract: contract}}, nil
}

// NewAddressListCaller creates a new read-only instance of AddressList, bound to a specific deployed contract.
func NewAddressListCaller(address common.Address, caller bind.ContractCaller) (*AddressListCaller, error) {
	contract, err := bindAddressList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressListCaller{contract: contract}, nil
}

// NewAddressListTransactor creates a new write-only instance of AddressList, bound to a specific deployed contract.
func NewAddressListTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressListTransactor, error) {
	contract, err := bindAddressList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressListTransactor{contract: contract}, nil
}

// NewAddressListFilterer creates a new log filterer instance of AddressList, bound to a specific deployed contract.
func NewAddressListFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressListFilterer, error) {
	contract, err := bindAddressList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressListFilterer{contract: contract}, nil
}

// bindAddressList binds a generic wrapper to an already deployed contract.
func bindAddressList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressList *AddressListRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressList.Contract.AddressListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressList *AddressListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressList.Contract.AddressListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressList *AddressListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressList.Contract.AddressListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressList *AddressListCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AddressList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressList *AddressListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressList *AddressListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressList.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_AddressList *AddressListCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AddressList.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_AddressList *AddressListSession) Name() (string, error) {
	return _AddressList.Contract.Name(&_AddressList.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_AddressList *AddressListCallerSession) Name() (string, error) {
	return _AddressList.Contract.Name(&_AddressList.CallOpts)
}

// OnList is a free data retrieval call binding the contract method 0x6f626eb3.
//
// Solidity: function onList( address) constant returns(bool)
func (_AddressList *AddressListCaller) OnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _AddressList.contract.Call(opts, out, "onList", arg0)
	return *ret0, err
}

// OnList is a free data retrieval call binding the contract method 0x6f626eb3.
//
// Solidity: function onList( address) constant returns(bool)
func (_AddressList *AddressListSession) OnList(arg0 common.Address) (bool, error) {
	return _AddressList.Contract.OnList(&_AddressList.CallOpts, arg0)
}

// OnList is a free data retrieval call binding the contract method 0x6f626eb3.
//
// Solidity: function onList( address) constant returns(bool)
func (_AddressList *AddressListCallerSession) OnList(arg0 common.Address) (bool, error) {
	return _AddressList.Contract.OnList(&_AddressList.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AddressList *AddressListCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AddressList.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AddressList *AddressListSession) Owner() (common.Address, error) {
	return _AddressList.Contract.Owner(&_AddressList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AddressList *AddressListCallerSession) Owner() (common.Address, error) {
	return _AddressList.Contract.Owner(&_AddressList.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_AddressList *AddressListCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AddressList.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_AddressList *AddressListSession) PendingOwner() (common.Address, error) {
	return _AddressList.Contract.PendingOwner(&_AddressList.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_AddressList *AddressListCallerSession) PendingOwner() (common.Address, error) {
	return _AddressList.Contract.PendingOwner(&_AddressList.CallOpts)
}

// ChangeList is a paid mutator transaction binding the contract method 0xefbd9ec2.
//
// Solidity: function changeList(_to address, _onList bool) returns()
func (_AddressList *AddressListTransactor) ChangeList(opts *bind.TransactOpts, _to common.Address, _onList bool) (*types.Transaction, error) {
	return _AddressList.contract.Transact(opts, "changeList", _to, _onList)
}

// ChangeList is a paid mutator transaction binding the contract method 0xefbd9ec2.
//
// Solidity: function changeList(_to address, _onList bool) returns()
func (_AddressList *AddressListSession) ChangeList(_to common.Address, _onList bool) (*types.Transaction, error) {
	return _AddressList.Contract.ChangeList(&_AddressList.TransactOpts, _to, _onList)
}

// ChangeList is a paid mutator transaction binding the contract method 0xefbd9ec2.
//
// Solidity: function changeList(_to address, _onList bool) returns()
func (_AddressList *AddressListTransactorSession) ChangeList(_to common.Address, _onList bool) (*types.Transaction, error) {
	return _AddressList.Contract.ChangeList(&_AddressList.TransactOpts, _to, _onList)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_AddressList *AddressListTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressList.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_AddressList *AddressListSession) ClaimOwnership() (*types.Transaction, error) {
	return _AddressList.Contract.ClaimOwnership(&_AddressList.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_AddressList *AddressListTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _AddressList.Contract.ClaimOwnership(&_AddressList.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AddressList *AddressListTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AddressList.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AddressList *AddressListSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressList.Contract.TransferOwnership(&_AddressList.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AddressList *AddressListTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AddressList.Contract.TransferOwnership(&_AddressList.TransactOpts, newOwner)
}

// AddressListChangeWhiteListIterator is returned from FilterChangeWhiteList and is used to iterate over the raw logs and unpacked data for ChangeWhiteList events raised by the AddressList contract.
type AddressListChangeWhiteListIterator struct {
	Event *AddressListChangeWhiteList // Event containing the contract specifics and raw log

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
func (it *AddressListChangeWhiteListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressListChangeWhiteList)
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
		it.Event = new(AddressListChangeWhiteList)
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
func (it *AddressListChangeWhiteListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressListChangeWhiteListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressListChangeWhiteList represents a ChangeWhiteList event raised by the AddressList contract.
type AddressListChangeWhiteList struct {
	To     common.Address
	OnList bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChangeWhiteList is a free log retrieval operation binding the contract event 0xebf7001045f8f4b6a63ef446c723437b124accd3ba9cc203f2f369952621be59.
//
// Solidity: e ChangeWhiteList(to indexed address, onList bool)
func (_AddressList *AddressListFilterer) FilterChangeWhiteList(opts *bind.FilterOpts, to []common.Address) (*AddressListChangeWhiteListIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AddressList.contract.FilterLogs(opts, "ChangeWhiteList", toRule)
	if err != nil {
		return nil, err
	}
	return &AddressListChangeWhiteListIterator{contract: _AddressList.contract, event: "ChangeWhiteList", logs: logs, sub: sub}, nil
}

// WatchChangeWhiteList is a free log subscription operation binding the contract event 0xebf7001045f8f4b6a63ef446c723437b124accd3ba9cc203f2f369952621be59.
//
// Solidity: e ChangeWhiteList(to indexed address, onList bool)
func (_AddressList *AddressListFilterer) WatchChangeWhiteList(opts *bind.WatchOpts, sink chan<- *AddressListChangeWhiteList, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AddressList.contract.WatchLogs(opts, "ChangeWhiteList", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressListChangeWhiteList)
				if err := _AddressList.contract.UnpackLog(event, "ChangeWhiteList", log); err != nil {
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

// AddressListOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AddressList contract.
type AddressListOwnershipTransferredIterator struct {
	Event *AddressListOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AddressListOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AddressListOwnershipTransferred)
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
		it.Event = new(AddressListOwnershipTransferred)
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
func (it *AddressListOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AddressListOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AddressListOwnershipTransferred represents a OwnershipTransferred event raised by the AddressList contract.
type AddressListOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AddressList *AddressListFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AddressListOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressList.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AddressListOwnershipTransferredIterator{contract: _AddressList.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AddressList *AddressListFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AddressListOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AddressList.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AddressListOwnershipTransferred)
				if err := _AddressList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// AllowanceSheetABI is the input ABI used to generate the binding from.
const AllowanceSheetABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenHolder\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"addAllowance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenHolder\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"subAllowance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenHolder\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setAllowance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// AllowanceSheetBin is the compiled bytecode used for deploying new contracts.
const AllowanceSheetBin = `0x608060405260008054600160a060020a03191633179055610432806100256000396000f30060806040526004361061008d5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631a46ec8281146100925780634e71e0c8146100cb5780635fd72d16146100e25780638da5cb5b1461010c57806397d88cd21461013d578063da46098c14610167578063e30c397814610191578063f2fde38b146101a6575b600080fd5b34801561009e57600080fd5b506100b9600160a060020a03600435811690602435166101c7565b60408051918252519081900360200190f35b3480156100d757600080fd5b506100e06101e4565b005b3480156100ee57600080fd5b506100e0600160a060020a036004358116906024351660443561026c565b34801561011857600080fd5b506101216102ea565b60408051600160a060020a039092168252519081900360200190f35b34801561014957600080fd5b506100e0600160a060020a03600435811690602435166044356102f9565b34801561017357600080fd5b506100e0600160a060020a0360043581169060243516604435610346565b34801561019d57600080fd5b50610121610389565b3480156101b257600080fd5b506100e0600160a060020a0360043516610398565b600260209081526000928352604080842090915290825290205481565b600154600160a060020a031633146101fb57600080fd5b60015460008054604051600160a060020a0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a0316331461028357600080fd5b600160a060020a038084166000908152600260209081526040808320938616835292905220546102b9908263ffffffff6103de16565b600160a060020a03938416600090815260026020908152604080832095909616825293909352929091209190915550565b600054600160a060020a031681565b600054600160a060020a0316331461031057600080fd5b600160a060020a038084166000908152600260209081526040808320938616835292905220546102b9908263ffffffff6103f416565b600054600160a060020a0316331461035d57600080fd5b600160a060020a0392831660009081526002602090815260408083209490951682529290925291902055565b600154600160a060020a031681565b600054600160a060020a031633146103af57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000828201838110156103ed57fe5b9392505050565b60008282111561040057fe5b509003905600a165627a7a72305820def44982447a3768f7b21b75ab84f16c2c28bf4de12ff3adeb6e0d0802c7d9e30029`

// DeployAllowanceSheet deploys a new Ethereum contract, binding an instance of AllowanceSheet to it.
func DeployAllowanceSheet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AllowanceSheet, error) {
	parsed, err := abi.JSON(strings.NewReader(AllowanceSheetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AllowanceSheetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AllowanceSheet{AllowanceSheetCaller: AllowanceSheetCaller{contract: contract}, AllowanceSheetTransactor: AllowanceSheetTransactor{contract: contract}, AllowanceSheetFilterer: AllowanceSheetFilterer{contract: contract}}, nil
}

// AllowanceSheet is an auto generated Go binding around an Ethereum contract.
type AllowanceSheet struct {
	AllowanceSheetCaller     // Read-only binding to the contract
	AllowanceSheetTransactor // Write-only binding to the contract
	AllowanceSheetFilterer   // Log filterer for contract events
}

// AllowanceSheetCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllowanceSheetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowanceSheetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllowanceSheetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowanceSheetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllowanceSheetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllowanceSheetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllowanceSheetSession struct {
	Contract     *AllowanceSheet   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AllowanceSheetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllowanceSheetCallerSession struct {
	Contract *AllowanceSheetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AllowanceSheetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllowanceSheetTransactorSession struct {
	Contract     *AllowanceSheetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AllowanceSheetRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllowanceSheetRaw struct {
	Contract *AllowanceSheet // Generic contract binding to access the raw methods on
}

// AllowanceSheetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllowanceSheetCallerRaw struct {
	Contract *AllowanceSheetCaller // Generic read-only contract binding to access the raw methods on
}

// AllowanceSheetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllowanceSheetTransactorRaw struct {
	Contract *AllowanceSheetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllowanceSheet creates a new instance of AllowanceSheet, bound to a specific deployed contract.
func NewAllowanceSheet(address common.Address, backend bind.ContractBackend) (*AllowanceSheet, error) {
	contract, err := bindAllowanceSheet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AllowanceSheet{AllowanceSheetCaller: AllowanceSheetCaller{contract: contract}, AllowanceSheetTransactor: AllowanceSheetTransactor{contract: contract}, AllowanceSheetFilterer: AllowanceSheetFilterer{contract: contract}}, nil
}

// NewAllowanceSheetCaller creates a new read-only instance of AllowanceSheet, bound to a specific deployed contract.
func NewAllowanceSheetCaller(address common.Address, caller bind.ContractCaller) (*AllowanceSheetCaller, error) {
	contract, err := bindAllowanceSheet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AllowanceSheetCaller{contract: contract}, nil
}

// NewAllowanceSheetTransactor creates a new write-only instance of AllowanceSheet, bound to a specific deployed contract.
func NewAllowanceSheetTransactor(address common.Address, transactor bind.ContractTransactor) (*AllowanceSheetTransactor, error) {
	contract, err := bindAllowanceSheet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AllowanceSheetTransactor{contract: contract}, nil
}

// NewAllowanceSheetFilterer creates a new log filterer instance of AllowanceSheet, bound to a specific deployed contract.
func NewAllowanceSheetFilterer(address common.Address, filterer bind.ContractFilterer) (*AllowanceSheetFilterer, error) {
	contract, err := bindAllowanceSheet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AllowanceSheetFilterer{contract: contract}, nil
}

// bindAllowanceSheet binds a generic wrapper to an already deployed contract.
func bindAllowanceSheet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AllowanceSheetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllowanceSheet *AllowanceSheetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AllowanceSheet.Contract.AllowanceSheetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllowanceSheet *AllowanceSheetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllowanceSheet.Contract.AllowanceSheetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllowanceSheet *AllowanceSheetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllowanceSheet.Contract.AllowanceSheetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllowanceSheet *AllowanceSheetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AllowanceSheet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllowanceSheet *AllowanceSheetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllowanceSheet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllowanceSheet *AllowanceSheetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllowanceSheet.Contract.contract.Transact(opts, method, params...)
}

// AllowanceOf is a free data retrieval call binding the contract method 0x1a46ec82.
//
// Solidity: function allowanceOf( address,  address) constant returns(uint256)
func (_AllowanceSheet *AllowanceSheetCaller) AllowanceOf(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AllowanceSheet.contract.Call(opts, out, "allowanceOf", arg0, arg1)
	return *ret0, err
}

// AllowanceOf is a free data retrieval call binding the contract method 0x1a46ec82.
//
// Solidity: function allowanceOf( address,  address) constant returns(uint256)
func (_AllowanceSheet *AllowanceSheetSession) AllowanceOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _AllowanceSheet.Contract.AllowanceOf(&_AllowanceSheet.CallOpts, arg0, arg1)
}

// AllowanceOf is a free data retrieval call binding the contract method 0x1a46ec82.
//
// Solidity: function allowanceOf( address,  address) constant returns(uint256)
func (_AllowanceSheet *AllowanceSheetCallerSession) AllowanceOf(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _AllowanceSheet.Contract.AllowanceOf(&_AllowanceSheet.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AllowanceSheet *AllowanceSheetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AllowanceSheet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AllowanceSheet *AllowanceSheetSession) Owner() (common.Address, error) {
	return _AllowanceSheet.Contract.Owner(&_AllowanceSheet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_AllowanceSheet *AllowanceSheetCallerSession) Owner() (common.Address, error) {
	return _AllowanceSheet.Contract.Owner(&_AllowanceSheet.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_AllowanceSheet *AllowanceSheetCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AllowanceSheet.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_AllowanceSheet *AllowanceSheetSession) PendingOwner() (common.Address, error) {
	return _AllowanceSheet.Contract.PendingOwner(&_AllowanceSheet.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_AllowanceSheet *AllowanceSheetCallerSession) PendingOwner() (common.Address, error) {
	return _AllowanceSheet.Contract.PendingOwner(&_AllowanceSheet.CallOpts)
}

// AddAllowance is a paid mutator transaction binding the contract method 0x5fd72d16.
//
// Solidity: function addAllowance(tokenHolder address, spender address, value uint256) returns()
func (_AllowanceSheet *AllowanceSheetTransactor) AddAllowance(opts *bind.TransactOpts, tokenHolder common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllowanceSheet.contract.Transact(opts, "addAllowance", tokenHolder, spender, value)
}

// AddAllowance is a paid mutator transaction binding the contract method 0x5fd72d16.
//
// Solidity: function addAllowance(tokenHolder address, spender address, value uint256) returns()
func (_AllowanceSheet *AllowanceSheetSession) AddAllowance(tokenHolder common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllowanceSheet.Contract.AddAllowance(&_AllowanceSheet.TransactOpts, tokenHolder, spender, value)
}

// AddAllowance is a paid mutator transaction binding the contract method 0x5fd72d16.
//
// Solidity: function addAllowance(tokenHolder address, spender address, value uint256) returns()
func (_AllowanceSheet *AllowanceSheetTransactorSession) AddAllowance(tokenHolder common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllowanceSheet.Contract.AddAllowance(&_AllowanceSheet.TransactOpts, tokenHolder, spender, value)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_AllowanceSheet *AllowanceSheetTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllowanceSheet.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_AllowanceSheet *AllowanceSheetSession) ClaimOwnership() (*types.Transaction, error) {
	return _AllowanceSheet.Contract.ClaimOwnership(&_AllowanceSheet.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_AllowanceSheet *AllowanceSheetTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _AllowanceSheet.Contract.ClaimOwnership(&_AllowanceSheet.TransactOpts)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(tokenHolder address, spender address, value uint256) returns()
func (_AllowanceSheet *AllowanceSheetTransactor) SetAllowance(opts *bind.TransactOpts, tokenHolder common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllowanceSheet.contract.Transact(opts, "setAllowance", tokenHolder, spender, value)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(tokenHolder address, spender address, value uint256) returns()
func (_AllowanceSheet *AllowanceSheetSession) SetAllowance(tokenHolder common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllowanceSheet.Contract.SetAllowance(&_AllowanceSheet.TransactOpts, tokenHolder, spender, value)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(tokenHolder address, spender address, value uint256) returns()
func (_AllowanceSheet *AllowanceSheetTransactorSession) SetAllowance(tokenHolder common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllowanceSheet.Contract.SetAllowance(&_AllowanceSheet.TransactOpts, tokenHolder, spender, value)
}

// SubAllowance is a paid mutator transaction binding the contract method 0x97d88cd2.
//
// Solidity: function subAllowance(tokenHolder address, spender address, value uint256) returns()
func (_AllowanceSheet *AllowanceSheetTransactor) SubAllowance(opts *bind.TransactOpts, tokenHolder common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllowanceSheet.contract.Transact(opts, "subAllowance", tokenHolder, spender, value)
}

// SubAllowance is a paid mutator transaction binding the contract method 0x97d88cd2.
//
// Solidity: function subAllowance(tokenHolder address, spender address, value uint256) returns()
func (_AllowanceSheet *AllowanceSheetSession) SubAllowance(tokenHolder common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllowanceSheet.Contract.SubAllowance(&_AllowanceSheet.TransactOpts, tokenHolder, spender, value)
}

// SubAllowance is a paid mutator transaction binding the contract method 0x97d88cd2.
//
// Solidity: function subAllowance(tokenHolder address, spender address, value uint256) returns()
func (_AllowanceSheet *AllowanceSheetTransactorSession) SubAllowance(tokenHolder common.Address, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _AllowanceSheet.Contract.SubAllowance(&_AllowanceSheet.TransactOpts, tokenHolder, spender, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AllowanceSheet *AllowanceSheetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AllowanceSheet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AllowanceSheet *AllowanceSheetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AllowanceSheet.Contract.TransferOwnership(&_AllowanceSheet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_AllowanceSheet *AllowanceSheetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AllowanceSheet.Contract.TransferOwnership(&_AllowanceSheet.TransactOpts, newOwner)
}

// AllowanceSheetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AllowanceSheet contract.
type AllowanceSheetOwnershipTransferredIterator struct {
	Event *AllowanceSheetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AllowanceSheetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AllowanceSheetOwnershipTransferred)
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
		it.Event = new(AllowanceSheetOwnershipTransferred)
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
func (it *AllowanceSheetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AllowanceSheetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AllowanceSheetOwnershipTransferred represents a OwnershipTransferred event raised by the AllowanceSheet contract.
type AllowanceSheetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AllowanceSheet *AllowanceSheetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AllowanceSheetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AllowanceSheet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AllowanceSheetOwnershipTransferredIterator{contract: _AllowanceSheet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_AllowanceSheet *AllowanceSheetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AllowanceSheetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AllowanceSheet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AllowanceSheetOwnershipTransferred)
				if err := _AllowanceSheet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BalanceSheetABI is the input ABI used to generate the binding from.
const BalanceSheetABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"addBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"subBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BalanceSheetBin is the compiled bytecode used for deploying new contracts.
const BalanceSheetBin = `0x608060405260008054600160a060020a031916331790556103d4806100256000396000f30060806040526004361061008d5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166321e5383a81146100925780634e71e0c8146100b857806370a08231146100cd5780638da5cb5b14610100578063cf8eeb7e14610131578063e30443bc14610155578063e30c397814610179578063f2fde38b1461018e575b600080fd5b34801561009e57600080fd5b506100b6600160a060020a03600435166024356101af565b005b3480156100c457600080fd5b506100b661020f565b3480156100d957600080fd5b506100ee600160a060020a0360043516610297565b60408051918252519081900360200190f35b34801561010c57600080fd5b506101156102a9565b60408051600160a060020a039092168252519081900360200190f35b34801561013d57600080fd5b506100b6600160a060020a03600435166024356102b8565b34801561016157600080fd5b506100b6600160a060020a03600435166024356102f8565b34801561018557600080fd5b5061011561032b565b34801561019a57600080fd5b506100b6600160a060020a036004351661033a565b600054600160a060020a031633146101c657600080fd5b600160a060020a0382166000908152600260205260409020546101ef908263ffffffff61038016565b600160a060020a0390921660009081526002602052604090209190915550565b600154600160a060020a0316331461022657600080fd5b60015460008054604051600160a060020a0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b60026020526000908152604090205481565b600054600160a060020a031681565b600054600160a060020a031633146102cf57600080fd5b600160a060020a0382166000908152600260205260409020546101ef908263ffffffff61039616565b600054600160a060020a0316331461030f57600080fd5b600160a060020a03909116600090815260026020526040902055565b600154600160a060020a031681565b600054600160a060020a0316331461035157600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008282018381101561038f57fe5b9392505050565b6000828211156103a257fe5b509003905600a165627a7a72305820cfb8197a6e4c2973c218ae25797a179bb90e3017c4392f331885cdbfbd5a59d60029`

// DeployBalanceSheet deploys a new Ethereum contract, binding an instance of BalanceSheet to it.
func DeployBalanceSheet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BalanceSheet, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceSheetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BalanceSheetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BalanceSheet{BalanceSheetCaller: BalanceSheetCaller{contract: contract}, BalanceSheetTransactor: BalanceSheetTransactor{contract: contract}, BalanceSheetFilterer: BalanceSheetFilterer{contract: contract}}, nil
}

// BalanceSheet is an auto generated Go binding around an Ethereum contract.
type BalanceSheet struct {
	BalanceSheetCaller     // Read-only binding to the contract
	BalanceSheetTransactor // Write-only binding to the contract
	BalanceSheetFilterer   // Log filterer for contract events
}

// BalanceSheetCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceSheetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceSheetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceSheetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceSheetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceSheetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceSheetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceSheetSession struct {
	Contract     *BalanceSheet     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceSheetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceSheetCallerSession struct {
	Contract *BalanceSheetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BalanceSheetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceSheetTransactorSession struct {
	Contract     *BalanceSheetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BalanceSheetRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceSheetRaw struct {
	Contract *BalanceSheet // Generic contract binding to access the raw methods on
}

// BalanceSheetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceSheetCallerRaw struct {
	Contract *BalanceSheetCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceSheetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceSheetTransactorRaw struct {
	Contract *BalanceSheetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalanceSheet creates a new instance of BalanceSheet, bound to a specific deployed contract.
func NewBalanceSheet(address common.Address, backend bind.ContractBackend) (*BalanceSheet, error) {
	contract, err := bindBalanceSheet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceSheet{BalanceSheetCaller: BalanceSheetCaller{contract: contract}, BalanceSheetTransactor: BalanceSheetTransactor{contract: contract}, BalanceSheetFilterer: BalanceSheetFilterer{contract: contract}}, nil
}

// NewBalanceSheetCaller creates a new read-only instance of BalanceSheet, bound to a specific deployed contract.
func NewBalanceSheetCaller(address common.Address, caller bind.ContractCaller) (*BalanceSheetCaller, error) {
	contract, err := bindBalanceSheet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceSheetCaller{contract: contract}, nil
}

// NewBalanceSheetTransactor creates a new write-only instance of BalanceSheet, bound to a specific deployed contract.
func NewBalanceSheetTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceSheetTransactor, error) {
	contract, err := bindBalanceSheet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceSheetTransactor{contract: contract}, nil
}

// NewBalanceSheetFilterer creates a new log filterer instance of BalanceSheet, bound to a specific deployed contract.
func NewBalanceSheetFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceSheetFilterer, error) {
	contract, err := bindBalanceSheet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceSheetFilterer{contract: contract}, nil
}

// bindBalanceSheet binds a generic wrapper to an already deployed contract.
func bindBalanceSheet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceSheetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceSheet *BalanceSheetRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalanceSheet.Contract.BalanceSheetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceSheet *BalanceSheetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceSheet.Contract.BalanceSheetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceSheet *BalanceSheetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceSheet.Contract.BalanceSheetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceSheet *BalanceSheetCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalanceSheet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceSheet *BalanceSheetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceSheet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceSheet *BalanceSheetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceSheet.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_BalanceSheet *BalanceSheetCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalanceSheet.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_BalanceSheet *BalanceSheetSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BalanceSheet.Contract.BalanceOf(&_BalanceSheet.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_BalanceSheet *BalanceSheetCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _BalanceSheet.Contract.BalanceOf(&_BalanceSheet.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BalanceSheet *BalanceSheetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BalanceSheet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BalanceSheet *BalanceSheetSession) Owner() (common.Address, error) {
	return _BalanceSheet.Contract.Owner(&_BalanceSheet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BalanceSheet *BalanceSheetCallerSession) Owner() (common.Address, error) {
	return _BalanceSheet.Contract.Owner(&_BalanceSheet.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_BalanceSheet *BalanceSheetCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BalanceSheet.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_BalanceSheet *BalanceSheetSession) PendingOwner() (common.Address, error) {
	return _BalanceSheet.Contract.PendingOwner(&_BalanceSheet.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_BalanceSheet *BalanceSheetCallerSession) PendingOwner() (common.Address, error) {
	return _BalanceSheet.Contract.PendingOwner(&_BalanceSheet.CallOpts)
}

// AddBalance is a paid mutator transaction binding the contract method 0x21e5383a.
//
// Solidity: function addBalance(addr address, value uint256) returns()
func (_BalanceSheet *BalanceSheetTransactor) AddBalance(opts *bind.TransactOpts, addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _BalanceSheet.contract.Transact(opts, "addBalance", addr, value)
}

// AddBalance is a paid mutator transaction binding the contract method 0x21e5383a.
//
// Solidity: function addBalance(addr address, value uint256) returns()
func (_BalanceSheet *BalanceSheetSession) AddBalance(addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _BalanceSheet.Contract.AddBalance(&_BalanceSheet.TransactOpts, addr, value)
}

// AddBalance is a paid mutator transaction binding the contract method 0x21e5383a.
//
// Solidity: function addBalance(addr address, value uint256) returns()
func (_BalanceSheet *BalanceSheetTransactorSession) AddBalance(addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _BalanceSheet.Contract.AddBalance(&_BalanceSheet.TransactOpts, addr, value)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BalanceSheet *BalanceSheetTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceSheet.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BalanceSheet *BalanceSheetSession) ClaimOwnership() (*types.Transaction, error) {
	return _BalanceSheet.Contract.ClaimOwnership(&_BalanceSheet.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BalanceSheet *BalanceSheetTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _BalanceSheet.Contract.ClaimOwnership(&_BalanceSheet.TransactOpts)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(addr address, value uint256) returns()
func (_BalanceSheet *BalanceSheetTransactor) SetBalance(opts *bind.TransactOpts, addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _BalanceSheet.contract.Transact(opts, "setBalance", addr, value)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(addr address, value uint256) returns()
func (_BalanceSheet *BalanceSheetSession) SetBalance(addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _BalanceSheet.Contract.SetBalance(&_BalanceSheet.TransactOpts, addr, value)
}

// SetBalance is a paid mutator transaction binding the contract method 0xe30443bc.
//
// Solidity: function setBalance(addr address, value uint256) returns()
func (_BalanceSheet *BalanceSheetTransactorSession) SetBalance(addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _BalanceSheet.Contract.SetBalance(&_BalanceSheet.TransactOpts, addr, value)
}

// SubBalance is a paid mutator transaction binding the contract method 0xcf8eeb7e.
//
// Solidity: function subBalance(addr address, value uint256) returns()
func (_BalanceSheet *BalanceSheetTransactor) SubBalance(opts *bind.TransactOpts, addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _BalanceSheet.contract.Transact(opts, "subBalance", addr, value)
}

// SubBalance is a paid mutator transaction binding the contract method 0xcf8eeb7e.
//
// Solidity: function subBalance(addr address, value uint256) returns()
func (_BalanceSheet *BalanceSheetSession) SubBalance(addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _BalanceSheet.Contract.SubBalance(&_BalanceSheet.TransactOpts, addr, value)
}

// SubBalance is a paid mutator transaction binding the contract method 0xcf8eeb7e.
//
// Solidity: function subBalance(addr address, value uint256) returns()
func (_BalanceSheet *BalanceSheetTransactorSession) SubBalance(addr common.Address, value *big.Int) (*types.Transaction, error) {
	return _BalanceSheet.Contract.SubBalance(&_BalanceSheet.TransactOpts, addr, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BalanceSheet *BalanceSheetTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BalanceSheet.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BalanceSheet *BalanceSheetSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BalanceSheet.Contract.TransferOwnership(&_BalanceSheet.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BalanceSheet *BalanceSheetTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BalanceSheet.Contract.TransferOwnership(&_BalanceSheet.TransactOpts, newOwner)
}

// BalanceSheetOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BalanceSheet contract.
type BalanceSheetOwnershipTransferredIterator struct {
	Event *BalanceSheetOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BalanceSheetOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceSheetOwnershipTransferred)
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
		it.Event = new(BalanceSheetOwnershipTransferred)
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
func (it *BalanceSheetOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceSheetOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceSheetOwnershipTransferred represents a OwnershipTransferred event raised by the BalanceSheet contract.
type BalanceSheetOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BalanceSheet *BalanceSheetFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BalanceSheetOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BalanceSheet.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BalanceSheetOwnershipTransferredIterator{contract: _BalanceSheet.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BalanceSheet *BalanceSheetFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BalanceSheetOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BalanceSheet.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceSheetOwnershipTransferred)
				if err := _BalanceSheet.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BasicTokenABI is the input ABI used to generate the binding from.
const BasicTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setBalanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BasicTokenBin is the compiled bytecode used for deploying new contracts.
const BasicTokenBin = `0x608060405260008054600160a060020a03191633179055610683806100256000396000f3006080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd811461009d5780634e71e0c8146100c457806354f78dad146100db57806370a08231146100fc5780637bb98a681461011d5780638da5cb5b1461014e578063a9059cbb14610163578063e30c39781461019b578063f2fde38b146101b0575b600080fd5b3480156100a957600080fd5b506100b26101d1565b60408051918252519081900360200190f35b3480156100d057600080fd5b506100d96101d7565b005b3480156100e757600080fd5b506100d9600160a060020a036004351661025f565b34801561010857600080fd5b506100b2600160a060020a0360043516610314565b34801561012957600080fd5b506101326103b1565b60408051600160a060020a039092168252519081900360200190f35b34801561015a57600080fd5b506101326103c0565b34801561016f57600080fd5b50610187600160a060020a03600435166024356103cf565b604080519115158252519081900360200190f35b3480156101a757600080fd5b506101326103e5565b3480156101bc57600080fd5b506100d9600160a060020a03600435166103f4565b60035490565b600154600160a060020a031633146101ee57600080fd5b60015460008054604051600160a060020a0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a0316331461027657600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055604080517f4e71e0c800000000000000000000000000000000000000000000000000000000815290519290911691634e71e0c89160048082019260009290919082900301818387803b1580156102f957600080fd5b505af115801561030d573d6000803e3d6000fd5b5050505050565b600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915160009392909216916370a082319160248082019260209290919082900301818787803b15801561037f57600080fd5b505af1158015610393573d6000803e3d6000fd5b505050506040513d60208110156103a957600080fd5b505192915050565b600254600160a060020a031681565b600054600160a060020a031681565b60006103dc33848461043a565b50600192915050565b600154600160a060020a031681565b600054600160a060020a0316331461040b57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600160a060020a038216151561044f57600080fd5b600160a060020a038316151561046457600080fd5b600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152915191909216916370a082319160248083019260209291908290030181600087803b1580156104cc57600080fd5b505af11580156104e0573d6000803e3d6000fd5b505050506040513d60208110156104f657600080fd5b505181111561050457600080fd5b600254604080517fcf8eeb7e000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152602482018590529151919092169163cf8eeb7e91604480830192600092919082900301818387803b15801561057257600080fd5b505af1158015610586573d6000803e3d6000fd5b5050600254604080517f21e5383a000000000000000000000000000000000000000000000000000000008152600160a060020a0387811660048301526024820187905291519190921693506321e5383a9250604480830192600092919082900301818387803b1580156105f857600080fd5b505af115801561060c573d6000803e3d6000fd5b5050604080518481529051600160a060020a038087169450871692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050505600a165627a7a7230582065e917d95bc37ca13bebf2969c9f1bfa73ea3dd660bbbd3839fb9a937a8036570029`

// DeployBasicToken deploys a new Ethereum contract, binding an instance of BasicToken to it.
func DeployBasicToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BasicToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BasicTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// BasicToken is an auto generated Go binding around an Ethereum contract.
type BasicToken struct {
	BasicTokenCaller     // Read-only binding to the contract
	BasicTokenTransactor // Write-only binding to the contract
	BasicTokenFilterer   // Log filterer for contract events
}

// BasicTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BasicTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BasicTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BasicTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BasicTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BasicTokenSession struct {
	Contract     *BasicToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BasicTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BasicTokenCallerSession struct {
	Contract *BasicTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BasicTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BasicTokenTransactorSession struct {
	Contract     *BasicTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BasicTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BasicTokenRaw struct {
	Contract *BasicToken // Generic contract binding to access the raw methods on
}

// BasicTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BasicTokenCallerRaw struct {
	Contract *BasicTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BasicTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BasicTokenTransactorRaw struct {
	Contract *BasicTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBasicToken creates a new instance of BasicToken, bound to a specific deployed contract.
func NewBasicToken(address common.Address, backend bind.ContractBackend) (*BasicToken, error) {
	contract, err := bindBasicToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BasicToken{BasicTokenCaller: BasicTokenCaller{contract: contract}, BasicTokenTransactor: BasicTokenTransactor{contract: contract}, BasicTokenFilterer: BasicTokenFilterer{contract: contract}}, nil
}

// NewBasicTokenCaller creates a new read-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenCaller(address common.Address, caller bind.ContractCaller) (*BasicTokenCaller, error) {
	contract, err := bindBasicToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenCaller{contract: contract}, nil
}

// NewBasicTokenTransactor creates a new write-only instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BasicTokenTransactor, error) {
	contract, err := bindBasicToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransactor{contract: contract}, nil
}

// NewBasicTokenFilterer creates a new log filterer instance of BasicToken, bound to a specific deployed contract.
func NewBasicTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BasicTokenFilterer, error) {
	contract, err := bindBasicToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BasicTokenFilterer{contract: contract}, nil
}

// bindBasicToken binds a generic wrapper to an already deployed contract.
func bindBasicToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BasicTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.BasicTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.BasicTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BasicToken *BasicTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BasicToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BasicToken *BasicTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BasicToken *BasicTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BasicToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BasicToken *BasicTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BasicToken.Contract.BalanceOf(&_BasicToken.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_BasicToken *BasicTokenCaller) Balances(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "balances")
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_BasicToken *BasicTokenSession) Balances() (common.Address, error) {
	return _BasicToken.Contract.Balances(&_BasicToken.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_BasicToken *BasicTokenCallerSession) Balances() (common.Address, error) {
	return _BasicToken.Contract.Balances(&_BasicToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BasicToken *BasicTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BasicToken *BasicTokenSession) Owner() (common.Address, error) {
	return _BasicToken.Contract.Owner(&_BasicToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BasicToken *BasicTokenCallerSession) Owner() (common.Address, error) {
	return _BasicToken.Contract.Owner(&_BasicToken.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_BasicToken *BasicTokenCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_BasicToken *BasicTokenSession) PendingOwner() (common.Address, error) {
	return _BasicToken.Contract.PendingOwner(&_BasicToken.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_BasicToken *BasicTokenCallerSession) PendingOwner() (common.Address, error) {
	return _BasicToken.Contract.PendingOwner(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BasicToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BasicToken *BasicTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BasicToken.Contract.TotalSupply(&_BasicToken.CallOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BasicToken *BasicTokenTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BasicToken *BasicTokenSession) ClaimOwnership() (*types.Transaction, error) {
	return _BasicToken.Contract.ClaimOwnership(&_BasicToken.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BasicToken *BasicTokenTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _BasicToken.Contract.ClaimOwnership(&_BasicToken.TransactOpts)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_BasicToken *BasicTokenTransactor) SetBalanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "setBalanceSheet", sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_BasicToken *BasicTokenSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _BasicToken.Contract.SetBalanceSheet(&_BasicToken.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_BasicToken *BasicTokenTransactorSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _BasicToken.Contract.SetBalanceSheet(&_BasicToken.TransactOpts, sheet)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BasicToken *BasicTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BasicToken.Contract.Transfer(&_BasicToken.TransactOpts, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BasicToken *BasicTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BasicToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BasicToken *BasicTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasicToken.Contract.TransferOwnership(&_BasicToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BasicToken *BasicTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BasicToken.Contract.TransferOwnership(&_BasicToken.TransactOpts, newOwner)
}

// BasicTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BasicToken contract.
type BasicTokenOwnershipTransferredIterator struct {
	Event *BasicTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BasicTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenOwnershipTransferred)
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
		it.Event = new(BasicTokenOwnershipTransferred)
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
func (it *BasicTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenOwnershipTransferred represents a OwnershipTransferred event raised by the BasicToken contract.
type BasicTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BasicToken *BasicTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BasicTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenOwnershipTransferredIterator{contract: _BasicToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BasicToken *BasicTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BasicTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenOwnershipTransferred)
				if err := _BasicToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BasicTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BasicToken contract.
type BasicTokenTransferIterator struct {
	Event *BasicTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BasicTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BasicTokenTransfer)
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
		it.Event = new(BasicTokenTransfer)
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
func (it *BasicTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BasicTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BasicTokenTransfer represents a Transfer event raised by the BasicToken contract.
type BasicTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BasicTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BasicTokenTransferIterator{contract: _BasicToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BasicToken *BasicTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BasicTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BasicToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BasicTokenTransfer)
				if err := _BasicToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// BurnableTokenABI is the input ABI used to generate the binding from.
const BurnableTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setBalanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// BurnableTokenBin is the compiled bytecode used for deploying new contracts.
const BurnableTokenBin = `0x608060405260008054600160a060020a03191633179055610877806100256000396000f3006080604052600436106100a35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166318160ddd81146100a857806342966c68146100cf5780634e71e0c8146100e957806354f78dad146100fe57806370a082311461011f5780637bb98a68146101405780638da5cb5b14610171578063a9059cbb14610186578063e30c3978146101be578063f2fde38b146101d3575b600080fd5b3480156100b457600080fd5b506100bd6101f4565b60408051918252519081900360200190f35b3480156100db57600080fd5b506100e76004356101fa565b005b3480156100f557600080fd5b506100e76103b9565b34801561010a57600080fd5b506100e7600160a060020a0360043516610441565b34801561012b57600080fd5b506100bd600160a060020a03600435166104f6565b34801561014c57600080fd5b50610155610593565b60408051600160a060020a039092168252519081900360200190f35b34801561017d57600080fd5b506101556105a2565b34801561019257600080fd5b506101aa600160a060020a03600435166024356105b1565b604080519115158252519081900360200190f35b3480156101ca57600080fd5b506101556105c7565b3480156101df57600080fd5b506100e7600160a060020a03600435166105d6565b60035490565b600254604080517f70a082310000000000000000000000000000000000000000000000000000000081523360048201529051600092600160a060020a0316916370a0823191602480830192602092919082900301818787803b15801561025f57600080fd5b505af1158015610273573d6000803e3d6000fd5b505050506040513d602081101561028957600080fd5b505182111561029757600080fd5b50600254604080517fcf8eeb7e00000000000000000000000000000000000000000000000000000000815233600482018190526024820185905291519192600160a060020a03169163cf8eeb7e9160448082019260009290919082900301818387803b15801561030657600080fd5b505af115801561031a573d6000803e3d6000fd5b5050600354610332925090508363ffffffff61061c16565b600355604080518381529051600160a060020a038316917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2604080518381529051600091600160a060020a038416917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b600154600160a060020a031633146103d057600080fd5b60015460008054604051600160a060020a0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a0316331461045857600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055604080517f4e71e0c800000000000000000000000000000000000000000000000000000000815290519290911691634e71e0c89160048082019260009290919082900301818387803b1580156104db57600080fd5b505af11580156104ef573d6000803e3d6000fd5b5050505050565b600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915160009392909216916370a082319160248082019260209290919082900301818787803b15801561056157600080fd5b505af1158015610575573d6000803e3d6000fd5b505050506040513d602081101561058b57600080fd5b505192915050565b600254600160a060020a031681565b600054600160a060020a031681565b60006105be33848461062e565b50600192915050565b600154600160a060020a031681565b600054600160a060020a031633146105ed57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60008282111561062857fe5b50900390565b600160a060020a038216151561064357600080fd5b600160a060020a038316151561065857600080fd5b600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152915191909216916370a082319160248083019260209291908290030181600087803b1580156106c057600080fd5b505af11580156106d4573d6000803e3d6000fd5b505050506040513d60208110156106ea57600080fd5b50518111156106f857600080fd5b600254604080517fcf8eeb7e000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152602482018590529151919092169163cf8eeb7e91604480830192600092919082900301818387803b15801561076657600080fd5b505af115801561077a573d6000803e3d6000fd5b5050600254604080517f21e5383a000000000000000000000000000000000000000000000000000000008152600160a060020a0387811660048301526024820187905291519190921693506321e5383a9250604480830192600092919082900301818387803b1580156107ec57600080fd5b505af1158015610800573d6000803e3d6000fd5b5050604080518481529051600160a060020a038087169450871692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050505600a165627a7a72305820721ea3deb162940fb857391d16763b726717cfa47a4b00f66ce44736646606300029`

// DeployBurnableToken deploys a new Ethereum contract, binding an instance of BurnableToken to it.
func DeployBurnableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BurnableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnableToken{BurnableTokenCaller: BurnableTokenCaller{contract: contract}, BurnableTokenTransactor: BurnableTokenTransactor{contract: contract}, BurnableTokenFilterer: BurnableTokenFilterer{contract: contract}}, nil
}

// BurnableToken is an auto generated Go binding around an Ethereum contract.
type BurnableToken struct {
	BurnableTokenCaller     // Read-only binding to the contract
	BurnableTokenTransactor // Write-only binding to the contract
	BurnableTokenFilterer   // Log filterer for contract events
}

// BurnableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnableTokenSession struct {
	Contract     *BurnableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnableTokenCallerSession struct {
	Contract *BurnableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BurnableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnableTokenTransactorSession struct {
	Contract     *BurnableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BurnableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnableTokenRaw struct {
	Contract *BurnableToken // Generic contract binding to access the raw methods on
}

// BurnableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnableTokenCallerRaw struct {
	Contract *BurnableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BurnableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnableTokenTransactorRaw struct {
	Contract *BurnableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnableToken creates a new instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableToken(address common.Address, backend bind.ContractBackend) (*BurnableToken, error) {
	contract, err := bindBurnableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnableToken{BurnableTokenCaller: BurnableTokenCaller{contract: contract}, BurnableTokenTransactor: BurnableTokenTransactor{contract: contract}, BurnableTokenFilterer: BurnableTokenFilterer{contract: contract}}, nil
}

// NewBurnableTokenCaller creates a new read-only instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenCaller(address common.Address, caller bind.ContractCaller) (*BurnableTokenCaller, error) {
	contract, err := bindBurnableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenCaller{contract: contract}, nil
}

// NewBurnableTokenTransactor creates a new write-only instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnableTokenTransactor, error) {
	contract, err := bindBurnableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenTransactor{contract: contract}, nil
}

// NewBurnableTokenFilterer creates a new log filterer instance of BurnableToken, bound to a specific deployed contract.
func NewBurnableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnableTokenFilterer, error) {
	contract, err := bindBurnableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenFilterer{contract: contract}, nil
}

// bindBurnableToken binds a generic wrapper to an already deployed contract.
func bindBurnableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BurnableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableToken *BurnableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnableToken.Contract.BurnableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableToken *BurnableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableToken.Contract.BurnableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableToken *BurnableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableToken.Contract.BurnableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnableToken *BurnableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BurnableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnableToken *BurnableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnableToken *BurnableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnableToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BurnableToken *BurnableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BurnableToken *BurnableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.BalanceOf(&_BurnableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_BurnableToken *BurnableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _BurnableToken.Contract.BalanceOf(&_BurnableToken.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_BurnableToken *BurnableTokenCaller) Balances(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "balances")
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_BurnableToken *BurnableTokenSession) Balances() (common.Address, error) {
	return _BurnableToken.Contract.Balances(&_BurnableToken.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_BurnableToken *BurnableTokenCallerSession) Balances() (common.Address, error) {
	return _BurnableToken.Contract.Balances(&_BurnableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BurnableToken *BurnableTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BurnableToken *BurnableTokenSession) Owner() (common.Address, error) {
	return _BurnableToken.Contract.Owner(&_BurnableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BurnableToken *BurnableTokenCallerSession) Owner() (common.Address, error) {
	return _BurnableToken.Contract.Owner(&_BurnableToken.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_BurnableToken *BurnableTokenCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_BurnableToken *BurnableTokenSession) PendingOwner() (common.Address, error) {
	return _BurnableToken.Contract.PendingOwner(&_BurnableToken.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_BurnableToken *BurnableTokenCallerSession) PendingOwner() (common.Address, error) {
	return _BurnableToken.Contract.PendingOwner(&_BurnableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableToken *BurnableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BurnableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableToken *BurnableTokenSession) TotalSupply() (*big.Int, error) {
	return _BurnableToken.Contract.TotalSupply(&_BurnableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BurnableToken *BurnableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BurnableToken.Contract.TotalSupply(&_BurnableToken.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableToken *BurnableTokenTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableToken *BurnableTokenSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Burn(&_BurnableToken.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_BurnableToken *BurnableTokenTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Burn(&_BurnableToken.TransactOpts, _value)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BurnableToken *BurnableTokenTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BurnableToken *BurnableTokenSession) ClaimOwnership() (*types.Transaction, error) {
	return _BurnableToken.Contract.ClaimOwnership(&_BurnableToken.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_BurnableToken *BurnableTokenTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _BurnableToken.Contract.ClaimOwnership(&_BurnableToken.TransactOpts)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_BurnableToken *BurnableTokenTransactor) SetBalanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "setBalanceSheet", sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_BurnableToken *BurnableTokenSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _BurnableToken.Contract.SetBalanceSheet(&_BurnableToken.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_BurnableToken *BurnableTokenTransactorSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _BurnableToken.Contract.SetBalanceSheet(&_BurnableToken.TransactOpts, sheet)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Transfer(&_BurnableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_BurnableToken *BurnableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BurnableToken.Contract.Transfer(&_BurnableToken.TransactOpts, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BurnableToken *BurnableTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BurnableToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BurnableToken *BurnableTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BurnableToken.Contract.TransferOwnership(&_BurnableToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_BurnableToken *BurnableTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BurnableToken.Contract.TransferOwnership(&_BurnableToken.TransactOpts, newOwner)
}

// BurnableTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the BurnableToken contract.
type BurnableTokenBurnIterator struct {
	Event *BurnableTokenBurn // Event containing the contract specifics and raw log

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
func (it *BurnableTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenBurn)
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
		it.Event = new(BurnableTokenBurn)
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
func (it *BurnableTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenBurn represents a Burn event raised by the BurnableToken contract.
type BurnableTokenBurn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*BurnableTokenBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenBurnIterator{contract: _BurnableToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *BurnableTokenBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenBurn)
				if err := _BurnableToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// BurnableTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BurnableToken contract.
type BurnableTokenOwnershipTransferredIterator struct {
	Event *BurnableTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BurnableTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenOwnershipTransferred)
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
		it.Event = new(BurnableTokenOwnershipTransferred)
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
func (it *BurnableTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenOwnershipTransferred represents a OwnershipTransferred event raised by the BurnableToken contract.
type BurnableTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BurnableToken *BurnableTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BurnableTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenOwnershipTransferredIterator{contract: _BurnableToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_BurnableToken *BurnableTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BurnableTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenOwnershipTransferred)
				if err := _BurnableToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BurnableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BurnableToken contract.
type BurnableTokenTransferIterator struct {
	Event *BurnableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BurnableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BurnableTokenTransfer)
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
		it.Event = new(BurnableTokenTransfer)
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
func (it *BurnableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BurnableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BurnableTokenTransfer represents a Transfer event raised by the BurnableToken contract.
type BurnableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BurnableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BurnableTokenTransferIterator{contract: _BurnableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_BurnableToken *BurnableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BurnableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BurnableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BurnableTokenTransfer)
				if err := _BurnableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// CanDelegateABI is the input ABI used to generate the binding from.
const CanDelegateABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"delegateToNewContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setBalanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setAllowanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"DelegatedTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// CanDelegateBin is the compiled bytecode used for deploying new contracts.
const CanDelegateBin = `0x608060405260008054600160a060020a03191633179055611505806100256000396000f3006080604052600436106100fb5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461010057806318160ddd146101385780631d2d84001461015f57806323b872dd146101825780633ed10b92146101ac5780634e71e0c8146101dd57806354f78dad146101f2578063661884631461021357806370a08231146102375780637bb98a68146102585780638da5cb5b1461026d578063a9059cbb14610282578063c89e4361146102a6578063d73dd623146102bb578063dd62ed3e146102df578063e30c397814610306578063edc1e4f91461031b578063f2fde38b1461033c575b600080fd5b34801561010c57600080fd5b50610124600160a060020a036004351660243561035d565b604080519115158252519081900360200190f35b34801561014457600080fd5b5061014d61042d565b60408051918252519081900360200190f35b34801561016b57600080fd5b50610180600160a060020a03600435166104ef565b005b34801561018e57600080fd5b50610124600160a060020a0360043581169060243516604435610563565b3480156101b857600080fd5b506101c161063d565b60408051600160a060020a039092168252519081900360200190f35b3480156101e957600080fd5b5061018061064c565b3480156101fe57600080fd5b50610180600160a060020a03600435166106d4565b34801561021f57600080fd5b50610124600160a060020a0360043516602435610789565b34801561024357600080fd5b5061014d600160a060020a036004351661081d565b34801561026457600080fd5b506101c16108de565b34801561027957600080fd5b506101c16108ed565b34801561028e57600080fd5b50610124600160a060020a03600435166024356108fc565b3480156102b257600080fd5b506101c1610990565b3480156102c757600080fd5b50610124600160a060020a036004351660243561099f565b3480156102eb57600080fd5b5061014d600160a060020a0360043581169060243516610a33565b34801561031257600080fd5b506101c1610ac2565b34801561032757600080fd5b50610180600160a060020a0360043516610ad1565b34801561034857600080fd5b50610180600160a060020a0360043516610b67565b600554600090600160a060020a031615156103835761037c8383610bad565b9050610427565b600554604080517f296f4000000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152602482018690523360448301529151919092169163296f40009160648083019260209291908290030181600087803b1580156103f857600080fd5b505af115801561040c573d6000803e3d6000fd5b505050506040513d602081101561042257600080fd5b505190505b92915050565b600554600090600160a060020a031615156104515761044a610bc3565b90506104ec565b600560009054906101000a9004600160a060020a0316600160a060020a03166376e71dd86040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b1580156104bd57600080fd5b505af11580156104d1573d6000803e3d6000fd5b505050506040513d60208110156104e757600080fd5b505190505b90565b600054600160a060020a0316331461050657600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383811691909117918290556040519116907feef3c91406f155f6bf1d8754e73003590b8bfa5cfa5472ee9ea936761864ea3090600090a250565b600554600090600160a060020a0316151561058a57610583848484610bc9565b9050610636565b600554604080517f4df6b45d000000000000000000000000000000000000000000000000000000008152600160a060020a03878116600483015286811660248301526044820186905233606483015291519190921691634df6b45d9160848083019260209291908290030181600087803b15801561060757600080fd5b505af115801561061b573d6000803e3d6000fd5b505050506040513d602081101561063157600080fd5b505190505b9392505050565b600454600160a060020a031681565b600154600160a060020a0316331461066357600080fd5b60015460008054604051600160a060020a0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031633146106eb57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055604080517f4e71e0c800000000000000000000000000000000000000000000000000000000815290519290911691634e71e0c89160048082019260009290919082900301818387803b15801561076e57600080fd5b505af1158015610782573d6000803e3d6000fd5b5050505050565b600554600090600160a060020a031615156107a85761037c8383610be1565b600554604080517f93d3173a000000000000000000000000000000000000000000000000000000008152600160a060020a03868116600483015260248201869052336044830152915191909216916393d3173a9160648083019260209291908290030181600087803b1580156103f857600080fd5b600554600090600160a060020a031615156108425761083b82610bee565b90506108d9565b600554604080517f43a468c8000000000000000000000000000000000000000000000000000000008152600160a060020a038581166004830152915191909216916343a468c89160248083019260209291908290030181600087803b1580156108aa57600080fd5b505af11580156108be573d6000803e3d6000fd5b505050506040513d60208110156108d457600080fd5b505190505b919050565b600254600160a060020a031681565b600054600160a060020a031681565b600554600090600160a060020a0316151561091b5761037c8383610c59565b600554604080517f9cd1a121000000000000000000000000000000000000000000000000000000008152600160a060020a0386811660048301526024820186905233604483015291519190921691639cd1a1219160648083019260209291908290030181600087803b1580156103f857600080fd5b600554600160a060020a031681565b600554600090600160a060020a031615156109be5761037c8383610c66565b600554604080517f554249b3000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152602482018690523360448301529151919092169163554249b39160648083019260209291908290030181600087803b1580156103f857600080fd5b600554600090600160a060020a03161515610a525761037c8383610c73565b600554604080517f09ab8bba000000000000000000000000000000000000000000000000000000008152600160a060020a0386811660048301528581166024830152915191909216916309ab8bba9160448083019260209291908290030181600087803b1580156103f857600080fd5b600154600160a060020a031681565b600054600160a060020a03163314610ae857600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383811691909117808355604080517f4e71e0c800000000000000000000000000000000000000000000000000000000815290519190921692634e71e0c89280820192600092909182900301818387803b15801561076e57600080fd5b600054600160a060020a03163314610b7e57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000610bba838333610cd2565b50600192915050565b60035490565b6000610bd784848433610dab565b5060019392505050565b6000610bba838333610ee4565b600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915160009392909216916370a082319160248082019260209290919082900301818787803b1580156108aa57600080fd5b6000610bba338484611165565b6000610bba838333611382565b600480546040805160e160020a630d237641028152600160a060020a03868116948201949094528484166024820152905160009390921691631a46ec829160448082019260209290919082900301818787803b1580156103f857600080fd5b60048054604080517fda46098c000000000000000000000000000000000000000000000000000000008152600160a060020a03858116948201949094528684166024820152604481018690529051929091169163da46098c9160648082019260009290919082900301818387803b158015610d4c57600080fd5b505af1158015610d60573d6000803e3d6000fd5b5050604080518581529051600160a060020a038088169450851692507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360200190a3505050565b600480546040805160e160020a630d237641028152600160a060020a0388811694820194909452848416602482015290519290911691631a46ec82916044808201926020929091908290030181600087803b158015610e0957600080fd5b505af1158015610e1d573d6000803e3d6000fd5b505050506040513d6020811015610e3357600080fd5b5051821115610e4157600080fd5b60048054604080517f97d88cd2000000000000000000000000000000000000000000000000000000008152600160a060020a0388811694820194909452848416602482015260448101869052905192909116916397d88cd29160648082019260009290919082900301818387803b158015610ebb57600080fd5b505af1158015610ecf573d6000803e3d6000fd5b50505050610ede848484611165565b50505050565b600480546040805160e160020a630d237641028152600160a060020a03858116948201949094528684166024820152905160009390921691631a46ec829160448082019260209290919082900301818787803b158015610f4357600080fd5b505af1158015610f57573d6000803e3d6000fd5b505050506040513d6020811015610f6d57600080fd5b505190508083111561100e5760048054604080517fda46098c000000000000000000000000000000000000000000000000000000008152600160a060020a038681169482019490945287841660248201526000604482018190529151939092169263da46098c92606480820193929182900301818387803b158015610ff157600080fd5b505af1158015611005573d6000803e3d6000fd5b505050506110a1565b60048054604080517f97d88cd2000000000000000000000000000000000000000000000000000000008152600160a060020a0386811694820194909452878416602482015260448101879052905192909116916397d88cd29160648082019260009290919082900301818387803b15801561108857600080fd5b505af115801561109c573d6000803e3d6000fd5b505050505b600480546040805160e160020a630d237641028152600160a060020a038681169482018590528089166024830181905292519294937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92593911691631a46ec829160448083019260209291908290030181600087803b15801561112257600080fd5b505af1158015611136573d6000803e3d6000fd5b505050506040513d602081101561114c57600080fd5b505160408051918252519081900360200190a350505050565b600160a060020a038216151561117a57600080fd5b600160a060020a038316151561118f57600080fd5b600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152915191909216916370a082319160248083019260209291908290030181600087803b1580156111f757600080fd5b505af115801561120b573d6000803e3d6000fd5b505050506040513d602081101561122157600080fd5b505181111561122f57600080fd5b600254604080517fcf8eeb7e000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152602482018590529151919092169163cf8eeb7e91604480830192600092919082900301818387803b15801561129d57600080fd5b505af11580156112b1573d6000803e3d6000fd5b5050600254604080517f21e5383a000000000000000000000000000000000000000000000000000000008152600160a060020a0387811660048301526024820187905291519190921693506321e5383a9250604480830192600092919082900301818387803b15801561132357600080fd5b505af1158015611337573d6000803e3d6000fd5b5050604080518481529051600160a060020a038087169450871692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3505050565b60048054604080517f5fd72d16000000000000000000000000000000000000000000000000000000008152600160a060020a038581169482019490945286841660248201526044810186905290519290911691635fd72d169160648082019260009290919082900301818387803b1580156113fc57600080fd5b505af1158015611410573d6000803e3d6000fd5b5050600480546040805160e160020a630d237641028152600160a060020a038781169482018590528981166024830181905292519296509394507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259390921691631a46ec829160448083019260209291908290030181600087803b15801561149757600080fd5b505af11580156114ab573d6000803e3d6000fd5b505050506040513d60208110156114c157600080fd5b505160408051918252519081900360200190a35050505600a165627a7a72305820ee21c7c5024cf40bd6dd0cc0535d73adab8c854ed70d7710885792253e0c479f0029`

// DeployCanDelegate deploys a new Ethereum contract, binding an instance of CanDelegate to it.
func DeployCanDelegate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CanDelegate, error) {
	parsed, err := abi.JSON(strings.NewReader(CanDelegateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CanDelegateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CanDelegate{CanDelegateCaller: CanDelegateCaller{contract: contract}, CanDelegateTransactor: CanDelegateTransactor{contract: contract}, CanDelegateFilterer: CanDelegateFilterer{contract: contract}}, nil
}

// CanDelegate is an auto generated Go binding around an Ethereum contract.
type CanDelegate struct {
	CanDelegateCaller     // Read-only binding to the contract
	CanDelegateTransactor // Write-only binding to the contract
	CanDelegateFilterer   // Log filterer for contract events
}

// CanDelegateCaller is an auto generated read-only Go binding around an Ethereum contract.
type CanDelegateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanDelegateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CanDelegateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanDelegateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CanDelegateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanDelegateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CanDelegateSession struct {
	Contract     *CanDelegate      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CanDelegateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CanDelegateCallerSession struct {
	Contract *CanDelegateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CanDelegateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CanDelegateTransactorSession struct {
	Contract     *CanDelegateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CanDelegateRaw is an auto generated low-level Go binding around an Ethereum contract.
type CanDelegateRaw struct {
	Contract *CanDelegate // Generic contract binding to access the raw methods on
}

// CanDelegateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CanDelegateCallerRaw struct {
	Contract *CanDelegateCaller // Generic read-only contract binding to access the raw methods on
}

// CanDelegateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CanDelegateTransactorRaw struct {
	Contract *CanDelegateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCanDelegate creates a new instance of CanDelegate, bound to a specific deployed contract.
func NewCanDelegate(address common.Address, backend bind.ContractBackend) (*CanDelegate, error) {
	contract, err := bindCanDelegate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CanDelegate{CanDelegateCaller: CanDelegateCaller{contract: contract}, CanDelegateTransactor: CanDelegateTransactor{contract: contract}, CanDelegateFilterer: CanDelegateFilterer{contract: contract}}, nil
}

// NewCanDelegateCaller creates a new read-only instance of CanDelegate, bound to a specific deployed contract.
func NewCanDelegateCaller(address common.Address, caller bind.ContractCaller) (*CanDelegateCaller, error) {
	contract, err := bindCanDelegate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CanDelegateCaller{contract: contract}, nil
}

// NewCanDelegateTransactor creates a new write-only instance of CanDelegate, bound to a specific deployed contract.
func NewCanDelegateTransactor(address common.Address, transactor bind.ContractTransactor) (*CanDelegateTransactor, error) {
	contract, err := bindCanDelegate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CanDelegateTransactor{contract: contract}, nil
}

// NewCanDelegateFilterer creates a new log filterer instance of CanDelegate, bound to a specific deployed contract.
func NewCanDelegateFilterer(address common.Address, filterer bind.ContractFilterer) (*CanDelegateFilterer, error) {
	contract, err := bindCanDelegate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CanDelegateFilterer{contract: contract}, nil
}

// bindCanDelegate binds a generic wrapper to an already deployed contract.
func bindCanDelegate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CanDelegateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanDelegate *CanDelegateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CanDelegate.Contract.CanDelegateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanDelegate *CanDelegateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanDelegate.Contract.CanDelegateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanDelegate *CanDelegateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanDelegate.Contract.CanDelegateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanDelegate *CanDelegateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CanDelegate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanDelegate *CanDelegateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanDelegate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanDelegate *CanDelegateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanDelegate.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, spender address) constant returns(uint256)
func (_CanDelegate *CanDelegateCaller) Allowance(opts *bind.CallOpts, _owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CanDelegate.contract.Call(opts, out, "allowance", _owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, spender address) constant returns(uint256)
func (_CanDelegate *CanDelegateSession) Allowance(_owner common.Address, spender common.Address) (*big.Int, error) {
	return _CanDelegate.Contract.Allowance(&_CanDelegate.CallOpts, _owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, spender address) constant returns(uint256)
func (_CanDelegate *CanDelegateCallerSession) Allowance(_owner common.Address, spender common.Address) (*big.Int, error) {
	return _CanDelegate.Contract.Allowance(&_CanDelegate.CallOpts, _owner, spender)
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_CanDelegate *CanDelegateCaller) Allowances(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CanDelegate.contract.Call(opts, out, "allowances")
	return *ret0, err
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_CanDelegate *CanDelegateSession) Allowances() (common.Address, error) {
	return _CanDelegate.Contract.Allowances(&_CanDelegate.CallOpts)
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_CanDelegate *CanDelegateCallerSession) Allowances() (common.Address, error) {
	return _CanDelegate.Contract.Allowances(&_CanDelegate.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_CanDelegate *CanDelegateCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CanDelegate.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_CanDelegate *CanDelegateSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _CanDelegate.Contract.BalanceOf(&_CanDelegate.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_CanDelegate *CanDelegateCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _CanDelegate.Contract.BalanceOf(&_CanDelegate.CallOpts, who)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_CanDelegate *CanDelegateCaller) Balances(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CanDelegate.contract.Call(opts, out, "balances")
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_CanDelegate *CanDelegateSession) Balances() (common.Address, error) {
	return _CanDelegate.Contract.Balances(&_CanDelegate.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_CanDelegate *CanDelegateCallerSession) Balances() (common.Address, error) {
	return _CanDelegate.Contract.Balances(&_CanDelegate.CallOpts)
}

// Delegate is a free data retrieval call binding the contract method 0xc89e4361.
//
// Solidity: function delegate() constant returns(address)
func (_CanDelegate *CanDelegateCaller) Delegate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CanDelegate.contract.Call(opts, out, "delegate")
	return *ret0, err
}

// Delegate is a free data retrieval call binding the contract method 0xc89e4361.
//
// Solidity: function delegate() constant returns(address)
func (_CanDelegate *CanDelegateSession) Delegate() (common.Address, error) {
	return _CanDelegate.Contract.Delegate(&_CanDelegate.CallOpts)
}

// Delegate is a free data retrieval call binding the contract method 0xc89e4361.
//
// Solidity: function delegate() constant returns(address)
func (_CanDelegate *CanDelegateCallerSession) Delegate() (common.Address, error) {
	return _CanDelegate.Contract.Delegate(&_CanDelegate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CanDelegate *CanDelegateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CanDelegate.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CanDelegate *CanDelegateSession) Owner() (common.Address, error) {
	return _CanDelegate.Contract.Owner(&_CanDelegate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CanDelegate *CanDelegateCallerSession) Owner() (common.Address, error) {
	return _CanDelegate.Contract.Owner(&_CanDelegate.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_CanDelegate *CanDelegateCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CanDelegate.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_CanDelegate *CanDelegateSession) PendingOwner() (common.Address, error) {
	return _CanDelegate.Contract.PendingOwner(&_CanDelegate.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_CanDelegate *CanDelegateCallerSession) PendingOwner() (common.Address, error) {
	return _CanDelegate.Contract.PendingOwner(&_CanDelegate.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CanDelegate *CanDelegateCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _CanDelegate.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CanDelegate *CanDelegateSession) TotalSupply() (*big.Int, error) {
	return _CanDelegate.Contract.TotalSupply(&_CanDelegate.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_CanDelegate *CanDelegateCallerSession) TotalSupply() (*big.Int, error) {
	return _CanDelegate.Contract.TotalSupply(&_CanDelegate.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_CanDelegate *CanDelegateTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CanDelegate.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_CanDelegate *CanDelegateSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CanDelegate.Contract.Approve(&_CanDelegate.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_CanDelegate *CanDelegateTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _CanDelegate.Contract.Approve(&_CanDelegate.TransactOpts, spender, value)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_CanDelegate *CanDelegateTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanDelegate.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_CanDelegate *CanDelegateSession) ClaimOwnership() (*types.Transaction, error) {
	return _CanDelegate.Contract.ClaimOwnership(&_CanDelegate.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_CanDelegate *CanDelegateTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _CanDelegate.Contract.ClaimOwnership(&_CanDelegate.TransactOpts)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, subtractedValue uint256) returns(bool)
func (_CanDelegate *CanDelegateTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CanDelegate.contract.Transact(opts, "decreaseApproval", spender, subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, subtractedValue uint256) returns(bool)
func (_CanDelegate *CanDelegateSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CanDelegate.Contract.DecreaseApproval(&_CanDelegate.TransactOpts, spender, subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, subtractedValue uint256) returns(bool)
func (_CanDelegate *CanDelegateTransactorSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _CanDelegate.Contract.DecreaseApproval(&_CanDelegate.TransactOpts, spender, subtractedValue)
}

// DelegateToNewContract is a paid mutator transaction binding the contract method 0x1d2d8400.
//
// Solidity: function delegateToNewContract(newContract address) returns()
func (_CanDelegate *CanDelegateTransactor) DelegateToNewContract(opts *bind.TransactOpts, newContract common.Address) (*types.Transaction, error) {
	return _CanDelegate.contract.Transact(opts, "delegateToNewContract", newContract)
}

// DelegateToNewContract is a paid mutator transaction binding the contract method 0x1d2d8400.
//
// Solidity: function delegateToNewContract(newContract address) returns()
func (_CanDelegate *CanDelegateSession) DelegateToNewContract(newContract common.Address) (*types.Transaction, error) {
	return _CanDelegate.Contract.DelegateToNewContract(&_CanDelegate.TransactOpts, newContract)
}

// DelegateToNewContract is a paid mutator transaction binding the contract method 0x1d2d8400.
//
// Solidity: function delegateToNewContract(newContract address) returns()
func (_CanDelegate *CanDelegateTransactorSession) DelegateToNewContract(newContract common.Address) (*types.Transaction, error) {
	return _CanDelegate.Contract.DelegateToNewContract(&_CanDelegate.TransactOpts, newContract)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, addedValue uint256) returns(bool)
func (_CanDelegate *CanDelegateTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CanDelegate.contract.Transact(opts, "increaseApproval", spender, addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, addedValue uint256) returns(bool)
func (_CanDelegate *CanDelegateSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CanDelegate.Contract.IncreaseApproval(&_CanDelegate.TransactOpts, spender, addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, addedValue uint256) returns(bool)
func (_CanDelegate *CanDelegateTransactorSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _CanDelegate.Contract.IncreaseApproval(&_CanDelegate.TransactOpts, spender, addedValue)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_CanDelegate *CanDelegateTransactor) SetAllowanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _CanDelegate.contract.Transact(opts, "setAllowanceSheet", sheet)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_CanDelegate *CanDelegateSession) SetAllowanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _CanDelegate.Contract.SetAllowanceSheet(&_CanDelegate.TransactOpts, sheet)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_CanDelegate *CanDelegateTransactorSession) SetAllowanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _CanDelegate.Contract.SetAllowanceSheet(&_CanDelegate.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_CanDelegate *CanDelegateTransactor) SetBalanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _CanDelegate.contract.Transact(opts, "setBalanceSheet", sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_CanDelegate *CanDelegateSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _CanDelegate.Contract.SetBalanceSheet(&_CanDelegate.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_CanDelegate *CanDelegateTransactorSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _CanDelegate.Contract.SetBalanceSheet(&_CanDelegate.TransactOpts, sheet)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_CanDelegate *CanDelegateTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CanDelegate.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_CanDelegate *CanDelegateSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CanDelegate.Contract.Transfer(&_CanDelegate.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_CanDelegate *CanDelegateTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CanDelegate.Contract.Transfer(&_CanDelegate.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_CanDelegate *CanDelegateTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CanDelegate.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_CanDelegate *CanDelegateSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CanDelegate.Contract.TransferFrom(&_CanDelegate.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_CanDelegate *CanDelegateTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _CanDelegate.Contract.TransferFrom(&_CanDelegate.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CanDelegate *CanDelegateTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CanDelegate.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CanDelegate *CanDelegateSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CanDelegate.Contract.TransferOwnership(&_CanDelegate.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CanDelegate *CanDelegateTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CanDelegate.Contract.TransferOwnership(&_CanDelegate.TransactOpts, newOwner)
}

// CanDelegateApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CanDelegate contract.
type CanDelegateApprovalIterator struct {
	Event *CanDelegateApproval // Event containing the contract specifics and raw log

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
func (it *CanDelegateApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanDelegateApproval)
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
		it.Event = new(CanDelegateApproval)
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
func (it *CanDelegateApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanDelegateApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanDelegateApproval represents a Approval event raised by the CanDelegate contract.
type CanDelegateApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_CanDelegate *CanDelegateFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CanDelegateApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CanDelegate.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CanDelegateApprovalIterator{contract: _CanDelegate.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_CanDelegate *CanDelegateFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CanDelegateApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CanDelegate.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanDelegateApproval)
				if err := _CanDelegate.contract.UnpackLog(event, "Approval", log); err != nil {
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

// CanDelegateDelegatedToIterator is returned from FilterDelegatedTo and is used to iterate over the raw logs and unpacked data for DelegatedTo events raised by the CanDelegate contract.
type CanDelegateDelegatedToIterator struct {
	Event *CanDelegateDelegatedTo // Event containing the contract specifics and raw log

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
func (it *CanDelegateDelegatedToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanDelegateDelegatedTo)
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
		it.Event = new(CanDelegateDelegatedTo)
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
func (it *CanDelegateDelegatedToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanDelegateDelegatedToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanDelegateDelegatedTo represents a DelegatedTo event raised by the CanDelegate contract.
type CanDelegateDelegatedTo struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDelegatedTo is a free log retrieval operation binding the contract event 0xeef3c91406f155f6bf1d8754e73003590b8bfa5cfa5472ee9ea936761864ea30.
//
// Solidity: e DelegatedTo(newContract indexed address)
func (_CanDelegate *CanDelegateFilterer) FilterDelegatedTo(opts *bind.FilterOpts, newContract []common.Address) (*CanDelegateDelegatedToIterator, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _CanDelegate.contract.FilterLogs(opts, "DelegatedTo", newContractRule)
	if err != nil {
		return nil, err
	}
	return &CanDelegateDelegatedToIterator{contract: _CanDelegate.contract, event: "DelegatedTo", logs: logs, sub: sub}, nil
}

// WatchDelegatedTo is a free log subscription operation binding the contract event 0xeef3c91406f155f6bf1d8754e73003590b8bfa5cfa5472ee9ea936761864ea30.
//
// Solidity: e DelegatedTo(newContract indexed address)
func (_CanDelegate *CanDelegateFilterer) WatchDelegatedTo(opts *bind.WatchOpts, sink chan<- *CanDelegateDelegatedTo, newContract []common.Address) (event.Subscription, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _CanDelegate.contract.WatchLogs(opts, "DelegatedTo", newContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanDelegateDelegatedTo)
				if err := _CanDelegate.contract.UnpackLog(event, "DelegatedTo", log); err != nil {
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

// CanDelegateOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CanDelegate contract.
type CanDelegateOwnershipTransferredIterator struct {
	Event *CanDelegateOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CanDelegateOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanDelegateOwnershipTransferred)
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
		it.Event = new(CanDelegateOwnershipTransferred)
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
func (it *CanDelegateOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanDelegateOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanDelegateOwnershipTransferred represents a OwnershipTransferred event raised by the CanDelegate contract.
type CanDelegateOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CanDelegate *CanDelegateFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CanDelegateOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CanDelegate.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CanDelegateOwnershipTransferredIterator{contract: _CanDelegate.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CanDelegate *CanDelegateFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CanDelegateOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CanDelegate.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanDelegateOwnershipTransferred)
				if err := _CanDelegate.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// CanDelegateTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CanDelegate contract.
type CanDelegateTransferIterator struct {
	Event *CanDelegateTransfer // Event containing the contract specifics and raw log

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
func (it *CanDelegateTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanDelegateTransfer)
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
		it.Event = new(CanDelegateTransfer)
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
func (it *CanDelegateTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanDelegateTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanDelegateTransfer represents a Transfer event raised by the CanDelegate contract.
type CanDelegateTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_CanDelegate *CanDelegateFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CanDelegateTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CanDelegate.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CanDelegateTransferIterator{contract: _CanDelegate.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_CanDelegate *CanDelegateFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CanDelegateTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CanDelegate.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanDelegateTransfer)
				if err := _CanDelegate.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// CanReclaimTokenABI is the input ABI used to generate the binding from.
const CanReclaimTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// CanReclaimTokenBin is the compiled bytecode used for deploying new contracts.
const CanReclaimTokenBin = `0x608060405260008054600160a060020a03191633179055610320806100256000396000f3006080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317ffc320811461005b5780638da5cb5b1461007e578063f2fde38b146100af575b600080fd5b34801561006757600080fd5b5061007c600160a060020a03600435166100d0565b005b34801561008a57600080fd5b5061009361019c565b60408051600160a060020a039092168252519081900360200190f35b3480156100bb57600080fd5b5061007c600160a060020a03600435166101ab565b60008054600160a060020a031633146100e857600080fd5b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038416916370a082319160248083019260209291908290030181600087803b15801561014957600080fd5b505af115801561015d573d6000803e3d6000fd5b505050506040513d602081101561017357600080fd5b505160005490915061019890600160a060020a0384811691168363ffffffff61023f16565b5050565b600054600160a060020a031681565b600054600160a060020a031633146101c257600080fd5b600160a060020a03811615156101d757600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156102bb57600080fd5b505af11580156102cf573d6000803e3d6000fd5b505050506040513d60208110156102e557600080fd5b505115156102ef57fe5b5050505600a165627a7a72305820715f8a7f3681e45a390e8c63fd2d089d705b98aa06b71d29f9e7ae7c0dfecac60029`

// DeployCanReclaimToken deploys a new Ethereum contract, binding an instance of CanReclaimToken to it.
func DeployCanReclaimToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CanReclaimToken, error) {
	parsed, err := abi.JSON(strings.NewReader(CanReclaimTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CanReclaimTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CanReclaimToken{CanReclaimTokenCaller: CanReclaimTokenCaller{contract: contract}, CanReclaimTokenTransactor: CanReclaimTokenTransactor{contract: contract}, CanReclaimTokenFilterer: CanReclaimTokenFilterer{contract: contract}}, nil
}

// CanReclaimToken is an auto generated Go binding around an Ethereum contract.
type CanReclaimToken struct {
	CanReclaimTokenCaller     // Read-only binding to the contract
	CanReclaimTokenTransactor // Write-only binding to the contract
	CanReclaimTokenFilterer   // Log filterer for contract events
}

// CanReclaimTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CanReclaimTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanReclaimTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CanReclaimTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanReclaimTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CanReclaimTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CanReclaimTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CanReclaimTokenSession struct {
	Contract     *CanReclaimToken  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CanReclaimTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CanReclaimTokenCallerSession struct {
	Contract *CanReclaimTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CanReclaimTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CanReclaimTokenTransactorSession struct {
	Contract     *CanReclaimTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CanReclaimTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CanReclaimTokenRaw struct {
	Contract *CanReclaimToken // Generic contract binding to access the raw methods on
}

// CanReclaimTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CanReclaimTokenCallerRaw struct {
	Contract *CanReclaimTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CanReclaimTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CanReclaimTokenTransactorRaw struct {
	Contract *CanReclaimTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCanReclaimToken creates a new instance of CanReclaimToken, bound to a specific deployed contract.
func NewCanReclaimToken(address common.Address, backend bind.ContractBackend) (*CanReclaimToken, error) {
	contract, err := bindCanReclaimToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CanReclaimToken{CanReclaimTokenCaller: CanReclaimTokenCaller{contract: contract}, CanReclaimTokenTransactor: CanReclaimTokenTransactor{contract: contract}, CanReclaimTokenFilterer: CanReclaimTokenFilterer{contract: contract}}, nil
}

// NewCanReclaimTokenCaller creates a new read-only instance of CanReclaimToken, bound to a specific deployed contract.
func NewCanReclaimTokenCaller(address common.Address, caller bind.ContractCaller) (*CanReclaimTokenCaller, error) {
	contract, err := bindCanReclaimToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CanReclaimTokenCaller{contract: contract}, nil
}

// NewCanReclaimTokenTransactor creates a new write-only instance of CanReclaimToken, bound to a specific deployed contract.
func NewCanReclaimTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CanReclaimTokenTransactor, error) {
	contract, err := bindCanReclaimToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CanReclaimTokenTransactor{contract: contract}, nil
}

// NewCanReclaimTokenFilterer creates a new log filterer instance of CanReclaimToken, bound to a specific deployed contract.
func NewCanReclaimTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CanReclaimTokenFilterer, error) {
	contract, err := bindCanReclaimToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CanReclaimTokenFilterer{contract: contract}, nil
}

// bindCanReclaimToken binds a generic wrapper to an already deployed contract.
func bindCanReclaimToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CanReclaimTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanReclaimToken *CanReclaimTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CanReclaimToken.Contract.CanReclaimTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanReclaimToken *CanReclaimTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.CanReclaimTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanReclaimToken *CanReclaimTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.CanReclaimTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CanReclaimToken *CanReclaimTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CanReclaimToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CanReclaimToken *CanReclaimTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CanReclaimToken *CanReclaimTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CanReclaimToken *CanReclaimTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CanReclaimToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CanReclaimToken *CanReclaimTokenSession) Owner() (common.Address, error) {
	return _CanReclaimToken.Contract.Owner(&_CanReclaimToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_CanReclaimToken *CanReclaimTokenCallerSession) Owner() (common.Address, error) {
	return _CanReclaimToken.Contract.Owner(&_CanReclaimToken.CallOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_CanReclaimToken *CanReclaimTokenTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _CanReclaimToken.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_CanReclaimToken *CanReclaimTokenSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.ReclaimToken(&_CanReclaimToken.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_CanReclaimToken *CanReclaimTokenTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.ReclaimToken(&_CanReclaimToken.TransactOpts, token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CanReclaimToken *CanReclaimTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CanReclaimToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CanReclaimToken *CanReclaimTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.TransferOwnership(&_CanReclaimToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_CanReclaimToken *CanReclaimTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CanReclaimToken.Contract.TransferOwnership(&_CanReclaimToken.TransactOpts, newOwner)
}

// CanReclaimTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CanReclaimToken contract.
type CanReclaimTokenOwnershipTransferredIterator struct {
	Event *CanReclaimTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CanReclaimTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CanReclaimTokenOwnershipTransferred)
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
		it.Event = new(CanReclaimTokenOwnershipTransferred)
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
func (it *CanReclaimTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CanReclaimTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CanReclaimTokenOwnershipTransferred represents a OwnershipTransferred event raised by the CanReclaimToken contract.
type CanReclaimTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CanReclaimToken *CanReclaimTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CanReclaimTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CanReclaimToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CanReclaimTokenOwnershipTransferredIterator{contract: _CanReclaimToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_CanReclaimToken *CanReclaimTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CanReclaimTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CanReclaimToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CanReclaimTokenOwnershipTransferred)
				if err := _CanReclaimToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ClaimableABI is the input ABI used to generate the binding from.
const ClaimableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// ClaimableBin is the compiled bytecode used for deploying new contracts.
const ClaimableBin = `0x608060405260008054600160a060020a031916331790556101fc806100256000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416634e71e0c881146100665780638da5cb5b1461007d578063e30c3978146100ae578063f2fde38b146100c3575b600080fd5b34801561007257600080fd5b5061007b6100e4565b005b34801561008957600080fd5b5061009261016c565b60408051600160a060020a039092168252519081900360200190f35b3480156100ba57600080fd5b5061009261017b565b3480156100cf57600080fd5b5061007b600160a060020a036004351661018a565b600154600160a060020a031633146100fb57600080fd5b60015460008054604051600160a060020a0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031681565b600154600160a060020a031681565b600054600160a060020a031633146101a157600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a7230582054b71ec10c3b1e9354d680bb69ac23e7af3d1479aeea7017b708b8003683e97e0029`

// DeployClaimable deploys a new Ethereum contract, binding an instance of Claimable to it.
func DeployClaimable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Claimable, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClaimableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Claimable{ClaimableCaller: ClaimableCaller{contract: contract}, ClaimableTransactor: ClaimableTransactor{contract: contract}, ClaimableFilterer: ClaimableFilterer{contract: contract}}, nil
}

// Claimable is an auto generated Go binding around an Ethereum contract.
type Claimable struct {
	ClaimableCaller     // Read-only binding to the contract
	ClaimableTransactor // Write-only binding to the contract
	ClaimableFilterer   // Log filterer for contract events
}

// ClaimableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimableSession struct {
	Contract     *Claimable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimableCallerSession struct {
	Contract *ClaimableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ClaimableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimableTransactorSession struct {
	Contract     *ClaimableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ClaimableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimableRaw struct {
	Contract *Claimable // Generic contract binding to access the raw methods on
}

// ClaimableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimableCallerRaw struct {
	Contract *ClaimableCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimableTransactorRaw struct {
	Contract *ClaimableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimable creates a new instance of Claimable, bound to a specific deployed contract.
func NewClaimable(address common.Address, backend bind.ContractBackend) (*Claimable, error) {
	contract, err := bindClaimable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Claimable{ClaimableCaller: ClaimableCaller{contract: contract}, ClaimableTransactor: ClaimableTransactor{contract: contract}, ClaimableFilterer: ClaimableFilterer{contract: contract}}, nil
}

// NewClaimableCaller creates a new read-only instance of Claimable, bound to a specific deployed contract.
func NewClaimableCaller(address common.Address, caller bind.ContractCaller) (*ClaimableCaller, error) {
	contract, err := bindClaimable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimableCaller{contract: contract}, nil
}

// NewClaimableTransactor creates a new write-only instance of Claimable, bound to a specific deployed contract.
func NewClaimableTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimableTransactor, error) {
	contract, err := bindClaimable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimableTransactor{contract: contract}, nil
}

// NewClaimableFilterer creates a new log filterer instance of Claimable, bound to a specific deployed contract.
func NewClaimableFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimableFilterer, error) {
	contract, err := bindClaimable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimableFilterer{contract: contract}, nil
}

// bindClaimable binds a generic wrapper to an already deployed contract.
func bindClaimable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claimable *ClaimableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Claimable.Contract.ClaimableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claimable *ClaimableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claimable.Contract.ClaimableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claimable *ClaimableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claimable.Contract.ClaimableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claimable *ClaimableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Claimable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claimable *ClaimableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claimable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claimable *ClaimableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claimable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Claimable *ClaimableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Claimable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Claimable *ClaimableSession) Owner() (common.Address, error) {
	return _Claimable.Contract.Owner(&_Claimable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Claimable *ClaimableCallerSession) Owner() (common.Address, error) {
	return _Claimable.Contract.Owner(&_Claimable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_Claimable *ClaimableCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Claimable.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_Claimable *ClaimableSession) PendingOwner() (common.Address, error) {
	return _Claimable.Contract.PendingOwner(&_Claimable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_Claimable *ClaimableCallerSession) PendingOwner() (common.Address, error) {
	return _Claimable.Contract.PendingOwner(&_Claimable.CallOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Claimable *ClaimableTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claimable.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Claimable *ClaimableSession) ClaimOwnership() (*types.Transaction, error) {
	return _Claimable.Contract.ClaimOwnership(&_Claimable.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Claimable *ClaimableTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _Claimable.Contract.ClaimOwnership(&_Claimable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Claimable *ClaimableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Claimable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Claimable *ClaimableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Claimable.Contract.TransferOwnership(&_Claimable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Claimable *ClaimableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Claimable.Contract.TransferOwnership(&_Claimable.TransactOpts, newOwner)
}

// ClaimableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Claimable contract.
type ClaimableOwnershipTransferredIterator struct {
	Event *ClaimableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ClaimableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimableOwnershipTransferred)
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
		it.Event = new(ClaimableOwnershipTransferred)
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
func (it *ClaimableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimableOwnershipTransferred represents a OwnershipTransferred event raised by the Claimable contract.
type ClaimableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Claimable *ClaimableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ClaimableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Claimable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ClaimableOwnershipTransferredIterator{contract: _Claimable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Claimable *ClaimableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ClaimableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Claimable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimableOwnershipTransferred)
				if err := _Claimable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// DelegateERC20ABI is the input ABI used to generate the binding from.
const DelegateERC20ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"delegateAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateApprove\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"delegateBalanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateTransferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateIncreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegateTotalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateDecreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DelegateERC20Bin is the compiled bytecode used for deploying new contracts.
const DelegateERC20Bin = `0x`

// DeployDelegateERC20 deploys a new Ethereum contract, binding an instance of DelegateERC20 to it.
func DeployDelegateERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DelegateERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegateERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DelegateERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DelegateERC20{DelegateERC20Caller: DelegateERC20Caller{contract: contract}, DelegateERC20Transactor: DelegateERC20Transactor{contract: contract}, DelegateERC20Filterer: DelegateERC20Filterer{contract: contract}}, nil
}

// DelegateERC20 is an auto generated Go binding around an Ethereum contract.
type DelegateERC20 struct {
	DelegateERC20Caller     // Read-only binding to the contract
	DelegateERC20Transactor // Write-only binding to the contract
	DelegateERC20Filterer   // Log filterer for contract events
}

// DelegateERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type DelegateERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type DelegateERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DelegateERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DelegateERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DelegateERC20Session struct {
	Contract     *DelegateERC20    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DelegateERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DelegateERC20CallerSession struct {
	Contract *DelegateERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// DelegateERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DelegateERC20TransactorSession struct {
	Contract     *DelegateERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// DelegateERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type DelegateERC20Raw struct {
	Contract *DelegateERC20 // Generic contract binding to access the raw methods on
}

// DelegateERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DelegateERC20CallerRaw struct {
	Contract *DelegateERC20Caller // Generic read-only contract binding to access the raw methods on
}

// DelegateERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DelegateERC20TransactorRaw struct {
	Contract *DelegateERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewDelegateERC20 creates a new instance of DelegateERC20, bound to a specific deployed contract.
func NewDelegateERC20(address common.Address, backend bind.ContractBackend) (*DelegateERC20, error) {
	contract, err := bindDelegateERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DelegateERC20{DelegateERC20Caller: DelegateERC20Caller{contract: contract}, DelegateERC20Transactor: DelegateERC20Transactor{contract: contract}, DelegateERC20Filterer: DelegateERC20Filterer{contract: contract}}, nil
}

// NewDelegateERC20Caller creates a new read-only instance of DelegateERC20, bound to a specific deployed contract.
func NewDelegateERC20Caller(address common.Address, caller bind.ContractCaller) (*DelegateERC20Caller, error) {
	contract, err := bindDelegateERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DelegateERC20Caller{contract: contract}, nil
}

// NewDelegateERC20Transactor creates a new write-only instance of DelegateERC20, bound to a specific deployed contract.
func NewDelegateERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*DelegateERC20Transactor, error) {
	contract, err := bindDelegateERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DelegateERC20Transactor{contract: contract}, nil
}

// NewDelegateERC20Filterer creates a new log filterer instance of DelegateERC20, bound to a specific deployed contract.
func NewDelegateERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*DelegateERC20Filterer, error) {
	contract, err := bindDelegateERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DelegateERC20Filterer{contract: contract}, nil
}

// bindDelegateERC20 binds a generic wrapper to an already deployed contract.
func bindDelegateERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DelegateERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegateERC20 *DelegateERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DelegateERC20.Contract.DelegateERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegateERC20 *DelegateERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegateERC20.Contract.DelegateERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegateERC20 *DelegateERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegateERC20.Contract.DelegateERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DelegateERC20 *DelegateERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DelegateERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DelegateERC20 *DelegateERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DelegateERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DelegateERC20 *DelegateERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DelegateERC20.Contract.contract.Transact(opts, method, params...)
}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(owner address, spender address) constant returns(uint256)
func (_DelegateERC20 *DelegateERC20Caller) DelegateAllowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DelegateERC20.contract.Call(opts, out, "delegateAllowance", owner, spender)
	return *ret0, err
}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(owner address, spender address) constant returns(uint256)
func (_DelegateERC20 *DelegateERC20Session) DelegateAllowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DelegateERC20.Contract.DelegateAllowance(&_DelegateERC20.CallOpts, owner, spender)
}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(owner address, spender address) constant returns(uint256)
func (_DelegateERC20 *DelegateERC20CallerSession) DelegateAllowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _DelegateERC20.Contract.DelegateAllowance(&_DelegateERC20.CallOpts, owner, spender)
}

// DelegateBalanceOf is a free data retrieval call binding the contract method 0x43a468c8.
//
// Solidity: function delegateBalanceOf(who address) constant returns(uint256)
func (_DelegateERC20 *DelegateERC20Caller) DelegateBalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DelegateERC20.contract.Call(opts, out, "delegateBalanceOf", who)
	return *ret0, err
}

// DelegateBalanceOf is a free data retrieval call binding the contract method 0x43a468c8.
//
// Solidity: function delegateBalanceOf(who address) constant returns(uint256)
func (_DelegateERC20 *DelegateERC20Session) DelegateBalanceOf(who common.Address) (*big.Int, error) {
	return _DelegateERC20.Contract.DelegateBalanceOf(&_DelegateERC20.CallOpts, who)
}

// DelegateBalanceOf is a free data retrieval call binding the contract method 0x43a468c8.
//
// Solidity: function delegateBalanceOf(who address) constant returns(uint256)
func (_DelegateERC20 *DelegateERC20CallerSession) DelegateBalanceOf(who common.Address) (*big.Int, error) {
	return _DelegateERC20.Contract.DelegateBalanceOf(&_DelegateERC20.CallOpts, who)
}

// DelegateTotalSupply is a free data retrieval call binding the contract method 0x76e71dd8.
//
// Solidity: function delegateTotalSupply() constant returns(uint256)
func (_DelegateERC20 *DelegateERC20Caller) DelegateTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DelegateERC20.contract.Call(opts, out, "delegateTotalSupply")
	return *ret0, err
}

// DelegateTotalSupply is a free data retrieval call binding the contract method 0x76e71dd8.
//
// Solidity: function delegateTotalSupply() constant returns(uint256)
func (_DelegateERC20 *DelegateERC20Session) DelegateTotalSupply() (*big.Int, error) {
	return _DelegateERC20.Contract.DelegateTotalSupply(&_DelegateERC20.CallOpts)
}

// DelegateTotalSupply is a free data retrieval call binding the contract method 0x76e71dd8.
//
// Solidity: function delegateTotalSupply() constant returns(uint256)
func (_DelegateERC20 *DelegateERC20CallerSession) DelegateTotalSupply() (*big.Int, error) {
	return _DelegateERC20.Contract.DelegateTotalSupply(&_DelegateERC20.CallOpts)
}

// DelegateApprove is a paid mutator transaction binding the contract method 0x296f4000.
//
// Solidity: function delegateApprove(spender address, value uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20Transactor) DelegateApprove(opts *bind.TransactOpts, spender common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.contract.Transact(opts, "delegateApprove", spender, value, origSender)
}

// DelegateApprove is a paid mutator transaction binding the contract method 0x296f4000.
//
// Solidity: function delegateApprove(spender address, value uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20Session) DelegateApprove(spender common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.Contract.DelegateApprove(&_DelegateERC20.TransactOpts, spender, value, origSender)
}

// DelegateApprove is a paid mutator transaction binding the contract method 0x296f4000.
//
// Solidity: function delegateApprove(spender address, value uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20TransactorSession) DelegateApprove(spender common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.Contract.DelegateApprove(&_DelegateERC20.TransactOpts, spender, value, origSender)
}

// DelegateDecreaseApproval is a paid mutator transaction binding the contract method 0x93d3173a.
//
// Solidity: function delegateDecreaseApproval(spender address, subtractedValue uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20Transactor) DelegateDecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.contract.Transact(opts, "delegateDecreaseApproval", spender, subtractedValue, origSender)
}

// DelegateDecreaseApproval is a paid mutator transaction binding the contract method 0x93d3173a.
//
// Solidity: function delegateDecreaseApproval(spender address, subtractedValue uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20Session) DelegateDecreaseApproval(spender common.Address, subtractedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.Contract.DelegateDecreaseApproval(&_DelegateERC20.TransactOpts, spender, subtractedValue, origSender)
}

// DelegateDecreaseApproval is a paid mutator transaction binding the contract method 0x93d3173a.
//
// Solidity: function delegateDecreaseApproval(spender address, subtractedValue uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20TransactorSession) DelegateDecreaseApproval(spender common.Address, subtractedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.Contract.DelegateDecreaseApproval(&_DelegateERC20.TransactOpts, spender, subtractedValue, origSender)
}

// DelegateIncreaseApproval is a paid mutator transaction binding the contract method 0x554249b3.
//
// Solidity: function delegateIncreaseApproval(spender address, addedValue uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20Transactor) DelegateIncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.contract.Transact(opts, "delegateIncreaseApproval", spender, addedValue, origSender)
}

// DelegateIncreaseApproval is a paid mutator transaction binding the contract method 0x554249b3.
//
// Solidity: function delegateIncreaseApproval(spender address, addedValue uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20Session) DelegateIncreaseApproval(spender common.Address, addedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.Contract.DelegateIncreaseApproval(&_DelegateERC20.TransactOpts, spender, addedValue, origSender)
}

// DelegateIncreaseApproval is a paid mutator transaction binding the contract method 0x554249b3.
//
// Solidity: function delegateIncreaseApproval(spender address, addedValue uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20TransactorSession) DelegateIncreaseApproval(spender common.Address, addedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.Contract.DelegateIncreaseApproval(&_DelegateERC20.TransactOpts, spender, addedValue, origSender)
}

// DelegateTransfer is a paid mutator transaction binding the contract method 0x9cd1a121.
//
// Solidity: function delegateTransfer(to address, value uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20Transactor) DelegateTransfer(opts *bind.TransactOpts, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.contract.Transact(opts, "delegateTransfer", to, value, origSender)
}

// DelegateTransfer is a paid mutator transaction binding the contract method 0x9cd1a121.
//
// Solidity: function delegateTransfer(to address, value uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20Session) DelegateTransfer(to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.Contract.DelegateTransfer(&_DelegateERC20.TransactOpts, to, value, origSender)
}

// DelegateTransfer is a paid mutator transaction binding the contract method 0x9cd1a121.
//
// Solidity: function delegateTransfer(to address, value uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20TransactorSession) DelegateTransfer(to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.Contract.DelegateTransfer(&_DelegateERC20.TransactOpts, to, value, origSender)
}

// DelegateTransferFrom is a paid mutator transaction binding the contract method 0x4df6b45d.
//
// Solidity: function delegateTransferFrom(from address, to address, value uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20Transactor) DelegateTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.contract.Transact(opts, "delegateTransferFrom", from, to, value, origSender)
}

// DelegateTransferFrom is a paid mutator transaction binding the contract method 0x4df6b45d.
//
// Solidity: function delegateTransferFrom(from address, to address, value uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20Session) DelegateTransferFrom(from common.Address, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.Contract.DelegateTransferFrom(&_DelegateERC20.TransactOpts, from, to, value, origSender)
}

// DelegateTransferFrom is a paid mutator transaction binding the contract method 0x4df6b45d.
//
// Solidity: function delegateTransferFrom(from address, to address, value uint256, origSender address) returns(bool)
func (_DelegateERC20 *DelegateERC20TransactorSession) DelegateTransferFrom(from common.Address, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _DelegateERC20.Contract.DelegateTransferFrom(&_DelegateERC20.TransactOpts, from, to, value, origSender)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20Bin is the compiled bytecode used for deploying new contracts.
const ERC20Bin = `0x`

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, from, to, value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// ERC20BasicBin is the compiled bytecode used for deploying new contracts.
const ERC20BasicBin = `0x`

// DeployERC20Basic deploys a new Ethereum contract, binding an instance of ERC20Basic to it.
func DeployERC20Basic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Basic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20BasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, to, value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// HasNoContractsABI is the input ABI used to generate the binding from.
const HasNoContractsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"reclaimContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// HasNoContractsBin is the compiled bytecode used for deploying new contracts.
const HasNoContractsBin = `0x608060405260008054600160a060020a03191633179055610239806100256000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416632aed7f3f811461005b5780638da5cb5b1461007e578063f2fde38b146100af575b600080fd5b34801561006757600080fd5b5061007c600160a060020a03600435166100d0565b005b34801561008a57600080fd5b5061009361016a565b60408051600160a060020a039092168252519081900360200190f35b3480156100bb57600080fd5b5061007c600160a060020a0360043516610179565b60008054600160a060020a031633146100e857600080fd5b5060008054604080517ff2fde38b000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820152905184939284169263f2fde38b926024808201939182900301818387803b15801561014e57600080fd5b505af1158015610162573d6000803e3d6000fd5b505050505050565b600054600160a060020a031681565b600054600160a060020a0316331461019057600080fd5b600160a060020a03811615156101a557600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820cd2ae2c386e5a09dcc0789ab8cbcecf6cff2474c29d4e5f4e3c06943711852670029`

// DeployHasNoContracts deploys a new Ethereum contract, binding an instance of HasNoContracts to it.
func DeployHasNoContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HasNoContracts, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoContractsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HasNoContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HasNoContracts{HasNoContractsCaller: HasNoContractsCaller{contract: contract}, HasNoContractsTransactor: HasNoContractsTransactor{contract: contract}, HasNoContractsFilterer: HasNoContractsFilterer{contract: contract}}, nil
}

// HasNoContracts is an auto generated Go binding around an Ethereum contract.
type HasNoContracts struct {
	HasNoContractsCaller     // Read-only binding to the contract
	HasNoContractsTransactor // Write-only binding to the contract
	HasNoContractsFilterer   // Log filterer for contract events
}

// HasNoContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type HasNoContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HasNoContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HasNoContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HasNoContractsSession struct {
	Contract     *HasNoContracts   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HasNoContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HasNoContractsCallerSession struct {
	Contract *HasNoContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// HasNoContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HasNoContractsTransactorSession struct {
	Contract     *HasNoContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// HasNoContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type HasNoContractsRaw struct {
	Contract *HasNoContracts // Generic contract binding to access the raw methods on
}

// HasNoContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HasNoContractsCallerRaw struct {
	Contract *HasNoContractsCaller // Generic read-only contract binding to access the raw methods on
}

// HasNoContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HasNoContractsTransactorRaw struct {
	Contract *HasNoContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHasNoContracts creates a new instance of HasNoContracts, bound to a specific deployed contract.
func NewHasNoContracts(address common.Address, backend bind.ContractBackend) (*HasNoContracts, error) {
	contract, err := bindHasNoContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HasNoContracts{HasNoContractsCaller: HasNoContractsCaller{contract: contract}, HasNoContractsTransactor: HasNoContractsTransactor{contract: contract}, HasNoContractsFilterer: HasNoContractsFilterer{contract: contract}}, nil
}

// NewHasNoContractsCaller creates a new read-only instance of HasNoContracts, bound to a specific deployed contract.
func NewHasNoContractsCaller(address common.Address, caller bind.ContractCaller) (*HasNoContractsCaller, error) {
	contract, err := bindHasNoContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoContractsCaller{contract: contract}, nil
}

// NewHasNoContractsTransactor creates a new write-only instance of HasNoContracts, bound to a specific deployed contract.
func NewHasNoContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*HasNoContractsTransactor, error) {
	contract, err := bindHasNoContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoContractsTransactor{contract: contract}, nil
}

// NewHasNoContractsFilterer creates a new log filterer instance of HasNoContracts, bound to a specific deployed contract.
func NewHasNoContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*HasNoContractsFilterer, error) {
	contract, err := bindHasNoContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HasNoContractsFilterer{contract: contract}, nil
}

// bindHasNoContracts binds a generic wrapper to an already deployed contract.
func bindHasNoContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoContracts *HasNoContractsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoContracts.Contract.HasNoContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoContracts *HasNoContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoContracts.Contract.HasNoContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoContracts *HasNoContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoContracts.Contract.HasNoContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoContracts *HasNoContractsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoContracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoContracts *HasNoContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoContracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoContracts *HasNoContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoContracts.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoContracts *HasNoContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HasNoContracts.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoContracts *HasNoContractsSession) Owner() (common.Address, error) {
	return _HasNoContracts.Contract.Owner(&_HasNoContracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoContracts *HasNoContractsCallerSession) Owner() (common.Address, error) {
	return _HasNoContracts.Contract.Owner(&_HasNoContracts.CallOpts)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_HasNoContracts *HasNoContractsTransactor) ReclaimContract(opts *bind.TransactOpts, contractAddr common.Address) (*types.Transaction, error) {
	return _HasNoContracts.contract.Transact(opts, "reclaimContract", contractAddr)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_HasNoContracts *HasNoContractsSession) ReclaimContract(contractAddr common.Address) (*types.Transaction, error) {
	return _HasNoContracts.Contract.ReclaimContract(&_HasNoContracts.TransactOpts, contractAddr)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_HasNoContracts *HasNoContractsTransactorSession) ReclaimContract(contractAddr common.Address) (*types.Transaction, error) {
	return _HasNoContracts.Contract.ReclaimContract(&_HasNoContracts.TransactOpts, contractAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoContracts *HasNoContractsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HasNoContracts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoContracts *HasNoContractsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoContracts.Contract.TransferOwnership(&_HasNoContracts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoContracts *HasNoContractsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoContracts.Contract.TransferOwnership(&_HasNoContracts.TransactOpts, newOwner)
}

// HasNoContractsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HasNoContracts contract.
type HasNoContractsOwnershipTransferredIterator struct {
	Event *HasNoContractsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HasNoContractsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HasNoContractsOwnershipTransferred)
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
		it.Event = new(HasNoContractsOwnershipTransferred)
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
func (it *HasNoContractsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HasNoContractsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HasNoContractsOwnershipTransferred represents a OwnershipTransferred event raised by the HasNoContracts contract.
type HasNoContractsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HasNoContracts *HasNoContractsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HasNoContractsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HasNoContracts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HasNoContractsOwnershipTransferredIterator{contract: _HasNoContracts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HasNoContracts *HasNoContractsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HasNoContractsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HasNoContracts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HasNoContractsOwnershipTransferred)
				if err := _HasNoContracts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// HasNoEtherABI is the input ABI used to generate the binding from.
const HasNoEtherABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// HasNoEtherBin is the compiled bytecode used for deploying new contracts.
const HasNoEtherBin = `0x608060405260008054600160a060020a03191633179055341561002157600080fd5b6101e8806100306000396000f3006080604052600436106100565763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b81146100655780639f727c2714610096578063f2fde38b146100ad575b34801561006257600080fd5b50005b34801561007157600080fd5b5061007a6100ce565b60408051600160a060020a039092168252519081900360200190f35b3480156100a257600080fd5b506100ab6100dd565b005b3480156100b957600080fd5b506100ab600160a060020a0360043516610128565b600054600160a060020a031681565b600054600160a060020a031633146100f457600080fd5b60008054604051600160a060020a0390911691303180156108fc02929091818181858888f19350505050151561012657fe5b565b600054600160a060020a0316331461013f57600080fd5b600160a060020a038116151561015457600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820836bf284c51b8037ce053f335aee2b14a11b87565b34802444ec5d4f8aa385e60029`

// DeployHasNoEther deploys a new Ethereum contract, binding an instance of HasNoEther to it.
func DeployHasNoEther(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HasNoEther, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoEtherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HasNoEtherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HasNoEther{HasNoEtherCaller: HasNoEtherCaller{contract: contract}, HasNoEtherTransactor: HasNoEtherTransactor{contract: contract}, HasNoEtherFilterer: HasNoEtherFilterer{contract: contract}}, nil
}

// HasNoEther is an auto generated Go binding around an Ethereum contract.
type HasNoEther struct {
	HasNoEtherCaller     // Read-only binding to the contract
	HasNoEtherTransactor // Write-only binding to the contract
	HasNoEtherFilterer   // Log filterer for contract events
}

// HasNoEtherCaller is an auto generated read-only Go binding around an Ethereum contract.
type HasNoEtherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoEtherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HasNoEtherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoEtherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HasNoEtherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoEtherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HasNoEtherSession struct {
	Contract     *HasNoEther       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HasNoEtherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HasNoEtherCallerSession struct {
	Contract *HasNoEtherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// HasNoEtherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HasNoEtherTransactorSession struct {
	Contract     *HasNoEtherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// HasNoEtherRaw is an auto generated low-level Go binding around an Ethereum contract.
type HasNoEtherRaw struct {
	Contract *HasNoEther // Generic contract binding to access the raw methods on
}

// HasNoEtherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HasNoEtherCallerRaw struct {
	Contract *HasNoEtherCaller // Generic read-only contract binding to access the raw methods on
}

// HasNoEtherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HasNoEtherTransactorRaw struct {
	Contract *HasNoEtherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHasNoEther creates a new instance of HasNoEther, bound to a specific deployed contract.
func NewHasNoEther(address common.Address, backend bind.ContractBackend) (*HasNoEther, error) {
	contract, err := bindHasNoEther(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HasNoEther{HasNoEtherCaller: HasNoEtherCaller{contract: contract}, HasNoEtherTransactor: HasNoEtherTransactor{contract: contract}, HasNoEtherFilterer: HasNoEtherFilterer{contract: contract}}, nil
}

// NewHasNoEtherCaller creates a new read-only instance of HasNoEther, bound to a specific deployed contract.
func NewHasNoEtherCaller(address common.Address, caller bind.ContractCaller) (*HasNoEtherCaller, error) {
	contract, err := bindHasNoEther(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoEtherCaller{contract: contract}, nil
}

// NewHasNoEtherTransactor creates a new write-only instance of HasNoEther, bound to a specific deployed contract.
func NewHasNoEtherTransactor(address common.Address, transactor bind.ContractTransactor) (*HasNoEtherTransactor, error) {
	contract, err := bindHasNoEther(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoEtherTransactor{contract: contract}, nil
}

// NewHasNoEtherFilterer creates a new log filterer instance of HasNoEther, bound to a specific deployed contract.
func NewHasNoEtherFilterer(address common.Address, filterer bind.ContractFilterer) (*HasNoEtherFilterer, error) {
	contract, err := bindHasNoEther(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HasNoEtherFilterer{contract: contract}, nil
}

// bindHasNoEther binds a generic wrapper to an already deployed contract.
func bindHasNoEther(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoEtherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoEther *HasNoEtherRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoEther.Contract.HasNoEtherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoEther *HasNoEtherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoEther.Contract.HasNoEtherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoEther *HasNoEtherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoEther.Contract.HasNoEtherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoEther *HasNoEtherCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoEther.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoEther *HasNoEtherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoEther.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoEther *HasNoEtherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoEther.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoEther *HasNoEtherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HasNoEther.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoEther *HasNoEtherSession) Owner() (common.Address, error) {
	return _HasNoEther.Contract.Owner(&_HasNoEther.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoEther *HasNoEtherCallerSession) Owner() (common.Address, error) {
	return _HasNoEther.Contract.Owner(&_HasNoEther.CallOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_HasNoEther *HasNoEtherTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoEther.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_HasNoEther *HasNoEtherSession) ReclaimEther() (*types.Transaction, error) {
	return _HasNoEther.Contract.ReclaimEther(&_HasNoEther.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_HasNoEther *HasNoEtherTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _HasNoEther.Contract.ReclaimEther(&_HasNoEther.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoEther *HasNoEtherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HasNoEther.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoEther *HasNoEtherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoEther.Contract.TransferOwnership(&_HasNoEther.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoEther *HasNoEtherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoEther.Contract.TransferOwnership(&_HasNoEther.TransactOpts, newOwner)
}

// HasNoEtherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HasNoEther contract.
type HasNoEtherOwnershipTransferredIterator struct {
	Event *HasNoEtherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HasNoEtherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HasNoEtherOwnershipTransferred)
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
		it.Event = new(HasNoEtherOwnershipTransferred)
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
func (it *HasNoEtherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HasNoEtherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HasNoEtherOwnershipTransferred represents a OwnershipTransferred event raised by the HasNoEther contract.
type HasNoEtherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HasNoEther *HasNoEtherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HasNoEtherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HasNoEther.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HasNoEtherOwnershipTransferredIterator{contract: _HasNoEther.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HasNoEther *HasNoEtherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HasNoEtherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HasNoEther.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HasNoEtherOwnershipTransferred)
				if err := _HasNoEther.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// HasNoTokensABI is the input ABI used to generate the binding from.
const HasNoTokensABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// HasNoTokensBin is the compiled bytecode used for deploying new contracts.
const HasNoTokensBin = `0x608060405260008054600160a060020a0319163317905561035c806100256000396000f3006080604052600436106100615763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317ffc32081146100665780638da5cb5b14610089578063c0ee0b8a146100ba578063f2fde38b146100eb575b600080fd5b34801561007257600080fd5b50610087600160a060020a036004351661010c565b005b34801561009557600080fd5b5061009e6101d8565b60408051600160a060020a039092168252519081900360200190f35b3480156100c657600080fd5b5061008760048035600160a060020a0316906024803591604435918201910135610061565b3480156100f757600080fd5b50610087600160a060020a03600435166101e7565b60008054600160a060020a0316331461012457600080fd5b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038416916370a082319160248083019260209291908290030181600087803b15801561018557600080fd5b505af1158015610199573d6000803e3d6000fd5b505050506040513d60208110156101af57600080fd5b50516000549091506101d490600160a060020a0384811691168363ffffffff61027b16565b5050565b600054600160a060020a031681565b600054600160a060020a031633146101fe57600080fd5b600160a060020a038116151561021357600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b1580156102f757600080fd5b505af115801561030b573d6000803e3d6000fd5b505050506040513d602081101561032157600080fd5b5051151561032b57fe5b5050505600a165627a7a72305820ede4e6cc7a26ff2c0d6fc9b6898b43b4cc08d20e670c08ca8d0330808fd110350029`

// DeployHasNoTokens deploys a new Ethereum contract, binding an instance of HasNoTokens to it.
func DeployHasNoTokens(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HasNoTokens, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoTokensABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HasNoTokensBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HasNoTokens{HasNoTokensCaller: HasNoTokensCaller{contract: contract}, HasNoTokensTransactor: HasNoTokensTransactor{contract: contract}, HasNoTokensFilterer: HasNoTokensFilterer{contract: contract}}, nil
}

// HasNoTokens is an auto generated Go binding around an Ethereum contract.
type HasNoTokens struct {
	HasNoTokensCaller     // Read-only binding to the contract
	HasNoTokensTransactor // Write-only binding to the contract
	HasNoTokensFilterer   // Log filterer for contract events
}

// HasNoTokensCaller is an auto generated read-only Go binding around an Ethereum contract.
type HasNoTokensCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoTokensTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HasNoTokensTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoTokensFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HasNoTokensFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HasNoTokensSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HasNoTokensSession struct {
	Contract     *HasNoTokens      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HasNoTokensCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HasNoTokensCallerSession struct {
	Contract *HasNoTokensCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// HasNoTokensTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HasNoTokensTransactorSession struct {
	Contract     *HasNoTokensTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HasNoTokensRaw is an auto generated low-level Go binding around an Ethereum contract.
type HasNoTokensRaw struct {
	Contract *HasNoTokens // Generic contract binding to access the raw methods on
}

// HasNoTokensCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HasNoTokensCallerRaw struct {
	Contract *HasNoTokensCaller // Generic read-only contract binding to access the raw methods on
}

// HasNoTokensTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HasNoTokensTransactorRaw struct {
	Contract *HasNoTokensTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHasNoTokens creates a new instance of HasNoTokens, bound to a specific deployed contract.
func NewHasNoTokens(address common.Address, backend bind.ContractBackend) (*HasNoTokens, error) {
	contract, err := bindHasNoTokens(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HasNoTokens{HasNoTokensCaller: HasNoTokensCaller{contract: contract}, HasNoTokensTransactor: HasNoTokensTransactor{contract: contract}, HasNoTokensFilterer: HasNoTokensFilterer{contract: contract}}, nil
}

// NewHasNoTokensCaller creates a new read-only instance of HasNoTokens, bound to a specific deployed contract.
func NewHasNoTokensCaller(address common.Address, caller bind.ContractCaller) (*HasNoTokensCaller, error) {
	contract, err := bindHasNoTokens(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoTokensCaller{contract: contract}, nil
}

// NewHasNoTokensTransactor creates a new write-only instance of HasNoTokens, bound to a specific deployed contract.
func NewHasNoTokensTransactor(address common.Address, transactor bind.ContractTransactor) (*HasNoTokensTransactor, error) {
	contract, err := bindHasNoTokens(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HasNoTokensTransactor{contract: contract}, nil
}

// NewHasNoTokensFilterer creates a new log filterer instance of HasNoTokens, bound to a specific deployed contract.
func NewHasNoTokensFilterer(address common.Address, filterer bind.ContractFilterer) (*HasNoTokensFilterer, error) {
	contract, err := bindHasNoTokens(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HasNoTokensFilterer{contract: contract}, nil
}

// bindHasNoTokens binds a generic wrapper to an already deployed contract.
func bindHasNoTokens(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HasNoTokensABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoTokens *HasNoTokensRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoTokens.Contract.HasNoTokensCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoTokens *HasNoTokensRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoTokens.Contract.HasNoTokensTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoTokens *HasNoTokensRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoTokens.Contract.HasNoTokensTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HasNoTokens *HasNoTokensCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HasNoTokens.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HasNoTokens *HasNoTokensTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HasNoTokens.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HasNoTokens *HasNoTokensTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HasNoTokens.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoTokens *HasNoTokensCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _HasNoTokens.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoTokens *HasNoTokensSession) Owner() (common.Address, error) {
	return _HasNoTokens.Contract.Owner(&_HasNoTokens.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_HasNoTokens *HasNoTokensCallerSession) Owner() (common.Address, error) {
	return _HasNoTokens.Contract.Owner(&_HasNoTokens.CallOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_HasNoTokens *HasNoTokensTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _HasNoTokens.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_HasNoTokens *HasNoTokensSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _HasNoTokens.Contract.ReclaimToken(&_HasNoTokens.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_HasNoTokens *HasNoTokensTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _HasNoTokens.Contract.ReclaimToken(&_HasNoTokens.TransactOpts, token)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_HasNoTokens *HasNoTokensTransactor) TokenFallback(opts *bind.TransactOpts, from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _HasNoTokens.contract.Transact(opts, "tokenFallback", from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_HasNoTokens *HasNoTokensSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _HasNoTokens.Contract.TokenFallback(&_HasNoTokens.TransactOpts, from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_HasNoTokens *HasNoTokensTransactorSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _HasNoTokens.Contract.TokenFallback(&_HasNoTokens.TransactOpts, from_, value_, data_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoTokens *HasNoTokensTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HasNoTokens.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoTokens *HasNoTokensSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoTokens.Contract.TransferOwnership(&_HasNoTokens.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_HasNoTokens *HasNoTokensTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HasNoTokens.Contract.TransferOwnership(&_HasNoTokens.TransactOpts, newOwner)
}

// HasNoTokensOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HasNoTokens contract.
type HasNoTokensOwnershipTransferredIterator struct {
	Event *HasNoTokensOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HasNoTokensOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HasNoTokensOwnershipTransferred)
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
		it.Event = new(HasNoTokensOwnershipTransferred)
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
func (it *HasNoTokensOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HasNoTokensOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HasNoTokensOwnershipTransferred represents a OwnershipTransferred event raised by the HasNoTokens contract.
type HasNoTokensOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HasNoTokens *HasNoTokensFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HasNoTokensOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HasNoTokens.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HasNoTokensOwnershipTransferredIterator{contract: _HasNoTokens.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_HasNoTokens *HasNoTokensFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HasNoTokensOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HasNoTokens.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HasNoTokensOwnershipTransferred)
				if err := _HasNoTokens.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// NoOwnerABI is the input ABI used to generate the binding from.
const NoOwnerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"reclaimContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// NoOwnerBin is the compiled bytecode used for deploying new contracts.
const NoOwnerBin = `0x608060405260008054600160a060020a03191633179055341561002157600080fd5b61049c806100306000396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166317ffc32081146100865780632aed7f3f146100a95780638da5cb5b146100ca5780639f727c27146100fb578063c0ee0b8a14610110578063f2fde38b14610141575b34801561008357600080fd5b50005b34801561009257600080fd5b506100a7600160a060020a0360043516610162565b005b3480156100b557600080fd5b506100a7600160a060020a036004351661022e565b3480156100d657600080fd5b506100df6102c8565b60408051600160a060020a039092168252519081900360200190f35b34801561010757600080fd5b506100a76102d7565b34801561011c57600080fd5b506100a760048035600160a060020a0316906024803591604435918201910135610322565b34801561014d57600080fd5b506100a7600160a060020a0360043516610327565b60008054600160a060020a0316331461017a57600080fd5b604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a038416916370a082319160248083019260209291908290030181600087803b1580156101db57600080fd5b505af11580156101ef573d6000803e3d6000fd5b505050506040513d602081101561020557600080fd5b505160005490915061022a90600160a060020a0384811691168363ffffffff6103bb16565b5050565b60008054600160a060020a0316331461024657600080fd5b5060008054604080517ff2fde38b000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820152905184939284169263f2fde38b926024808201939182900301818387803b1580156102ac57600080fd5b505af11580156102c0573d6000803e3d6000fd5b505050505050565b600054600160a060020a031681565b600054600160a060020a031633146102ee57600080fd5b60008054604051600160a060020a0390911691303180156108fc02929091818181858888f19350505050151561032057fe5b565b600080fd5b600054600160a060020a0316331461033e57600080fd5b600160a060020a038116151561035357600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561043757600080fd5b505af115801561044b573d6000803e3d6000fd5b505050506040513d602081101561046157600080fd5b5051151561046b57fe5b5050505600a165627a7a72305820110192117fd3ede75d17a8ca54c3176e7d3260e7911a49a996ef4e721867807a0029`

// DeployNoOwner deploys a new Ethereum contract, binding an instance of NoOwner to it.
func DeployNoOwner(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NoOwner, error) {
	parsed, err := abi.JSON(strings.NewReader(NoOwnerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NoOwnerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NoOwner{NoOwnerCaller: NoOwnerCaller{contract: contract}, NoOwnerTransactor: NoOwnerTransactor{contract: contract}, NoOwnerFilterer: NoOwnerFilterer{contract: contract}}, nil
}

// NoOwner is an auto generated Go binding around an Ethereum contract.
type NoOwner struct {
	NoOwnerCaller     // Read-only binding to the contract
	NoOwnerTransactor // Write-only binding to the contract
	NoOwnerFilterer   // Log filterer for contract events
}

// NoOwnerCaller is an auto generated read-only Go binding around an Ethereum contract.
type NoOwnerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOwnerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NoOwnerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOwnerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NoOwnerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoOwnerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NoOwnerSession struct {
	Contract     *NoOwner          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NoOwnerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NoOwnerCallerSession struct {
	Contract *NoOwnerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// NoOwnerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NoOwnerTransactorSession struct {
	Contract     *NoOwnerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NoOwnerRaw is an auto generated low-level Go binding around an Ethereum contract.
type NoOwnerRaw struct {
	Contract *NoOwner // Generic contract binding to access the raw methods on
}

// NoOwnerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NoOwnerCallerRaw struct {
	Contract *NoOwnerCaller // Generic read-only contract binding to access the raw methods on
}

// NoOwnerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NoOwnerTransactorRaw struct {
	Contract *NoOwnerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNoOwner creates a new instance of NoOwner, bound to a specific deployed contract.
func NewNoOwner(address common.Address, backend bind.ContractBackend) (*NoOwner, error) {
	contract, err := bindNoOwner(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NoOwner{NoOwnerCaller: NoOwnerCaller{contract: contract}, NoOwnerTransactor: NoOwnerTransactor{contract: contract}, NoOwnerFilterer: NoOwnerFilterer{contract: contract}}, nil
}

// NewNoOwnerCaller creates a new read-only instance of NoOwner, bound to a specific deployed contract.
func NewNoOwnerCaller(address common.Address, caller bind.ContractCaller) (*NoOwnerCaller, error) {
	contract, err := bindNoOwner(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NoOwnerCaller{contract: contract}, nil
}

// NewNoOwnerTransactor creates a new write-only instance of NoOwner, bound to a specific deployed contract.
func NewNoOwnerTransactor(address common.Address, transactor bind.ContractTransactor) (*NoOwnerTransactor, error) {
	contract, err := bindNoOwner(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NoOwnerTransactor{contract: contract}, nil
}

// NewNoOwnerFilterer creates a new log filterer instance of NoOwner, bound to a specific deployed contract.
func NewNoOwnerFilterer(address common.Address, filterer bind.ContractFilterer) (*NoOwnerFilterer, error) {
	contract, err := bindNoOwner(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NoOwnerFilterer{contract: contract}, nil
}

// bindNoOwner binds a generic wrapper to an already deployed contract.
func bindNoOwner(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NoOwnerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoOwner *NoOwnerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NoOwner.Contract.NoOwnerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoOwner *NoOwnerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOwner.Contract.NoOwnerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoOwner *NoOwnerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOwner.Contract.NoOwnerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NoOwner *NoOwnerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NoOwner.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NoOwner *NoOwnerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOwner.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NoOwner *NoOwnerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NoOwner.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NoOwner *NoOwnerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _NoOwner.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NoOwner *NoOwnerSession) Owner() (common.Address, error) {
	return _NoOwner.Contract.Owner(&_NoOwner.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_NoOwner *NoOwnerCallerSession) Owner() (common.Address, error) {
	return _NoOwner.Contract.Owner(&_NoOwner.CallOpts)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_NoOwner *NoOwnerTransactor) ReclaimContract(opts *bind.TransactOpts, contractAddr common.Address) (*types.Transaction, error) {
	return _NoOwner.contract.Transact(opts, "reclaimContract", contractAddr)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_NoOwner *NoOwnerSession) ReclaimContract(contractAddr common.Address) (*types.Transaction, error) {
	return _NoOwner.Contract.ReclaimContract(&_NoOwner.TransactOpts, contractAddr)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_NoOwner *NoOwnerTransactorSession) ReclaimContract(contractAddr common.Address) (*types.Transaction, error) {
	return _NoOwner.Contract.ReclaimContract(&_NoOwner.TransactOpts, contractAddr)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_NoOwner *NoOwnerTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NoOwner.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_NoOwner *NoOwnerSession) ReclaimEther() (*types.Transaction, error) {
	return _NoOwner.Contract.ReclaimEther(&_NoOwner.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_NoOwner *NoOwnerTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _NoOwner.Contract.ReclaimEther(&_NoOwner.TransactOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_NoOwner *NoOwnerTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _NoOwner.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_NoOwner *NoOwnerSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _NoOwner.Contract.ReclaimToken(&_NoOwner.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_NoOwner *NoOwnerTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _NoOwner.Contract.ReclaimToken(&_NoOwner.TransactOpts, token)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_NoOwner *NoOwnerTransactor) TokenFallback(opts *bind.TransactOpts, from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _NoOwner.contract.Transact(opts, "tokenFallback", from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_NoOwner *NoOwnerSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _NoOwner.Contract.TokenFallback(&_NoOwner.TransactOpts, from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_NoOwner *NoOwnerTransactorSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _NoOwner.Contract.TokenFallback(&_NoOwner.TransactOpts, from_, value_, data_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_NoOwner *NoOwnerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NoOwner.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_NoOwner *NoOwnerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NoOwner.Contract.TransferOwnership(&_NoOwner.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_NoOwner *NoOwnerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NoOwner.Contract.TransferOwnership(&_NoOwner.TransactOpts, newOwner)
}

// NoOwnerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NoOwner contract.
type NoOwnerOwnershipTransferredIterator struct {
	Event *NoOwnerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NoOwnerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NoOwnerOwnershipTransferred)
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
		it.Event = new(NoOwnerOwnershipTransferred)
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
func (it *NoOwnerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NoOwnerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NoOwnerOwnershipTransferred represents a OwnershipTransferred event raised by the NoOwner contract.
type NoOwnerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_NoOwner *NoOwnerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NoOwnerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NoOwner.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NoOwnerOwnershipTransferredIterator{contract: _NoOwner.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_NoOwner *NoOwnerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NoOwnerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NoOwner.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NoOwnerOwnershipTransferred)
				if err := _NoOwner.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a03191633179055610173806100326000396000f30060806040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b14610081575b600080fd5b34801561005c57600080fd5b506100656100a4565b60408051600160a060020a039092168252519081900360200190f35b34801561008d57600080fd5b506100a2600160a060020a03600435166100b3565b005b600054600160a060020a031681565b600054600160a060020a031633146100ca57600080fd5b600160a060020a03811615156100df57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a72305820a21cc510b62115642e274e42bfc9519a69db371633e9b14dd35a97b0485afe6a0029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// PausableBin is the compiled bytecode used for deploying new contracts.
const PausableBin = `0x608060405260008054600160a860020a0319163317905561032c806100256000396000f30060806040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f4ba83a81146100715780635c975abb146100885780638456cb59146100b15780638da5cb5b146100c6578063f2fde38b146100f7575b600080fd5b34801561007d57600080fd5b50610086610118565b005b34801561009457600080fd5b5061009d61019f565b604080519115158252519081900360200190f35b3480156100bd57600080fd5b506100866101c0565b3480156100d257600080fd5b506100db61025d565b60408051600160a060020a039092168252519081900360200190f35b34801561010357600080fd5b50610086600160a060020a036004351661026c565b600054600160a060020a0316331461012f57600080fd5b60005474010000000000000000000000000000000000000000900460ff16151561015857600080fd5b6000805474ff0000000000000000000000000000000000000000191681556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b339190a1565b60005474010000000000000000000000000000000000000000900460ff1681565b600054600160a060020a031633146101d757600080fd5b60005474010000000000000000000000000000000000000000900460ff16156101ff57600080fd5b6000805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001781556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff6259190a1565b600054600160a060020a031681565b600054600160a060020a0316331461028357600080fd5b600160a060020a038116151561029857600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555600a165627a7a723058202b0f55ef161d48fa9801106ff6505a692460c74b95feb61a72dfc4fa65152f800029`

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Pausable *PausableCallerSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Pausable.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Pausable *PausableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// PausableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pausable contract.
type PausableOwnershipTransferredIterator struct {
	Event *PausableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipTransferred)
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
		it.Event = new(PausableOwnershipTransferred)
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
func (it *PausableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipTransferred represents a OwnershipTransferred event raised by the Pausable contract.
type PausableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Pausable *PausableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipTransferredIterator{contract: _Pausable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Pausable *PausableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PausableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipTransferred)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PausablePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Pausable contract.
type PausablePauseIterator struct {
	Event *PausablePause // Event containing the contract specifics and raw log

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
func (it *PausablePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePause)
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
		it.Event = new(PausablePause)
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
func (it *PausablePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePause represents a Pause event raised by the Pausable contract.
type PausablePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) FilterPause(opts *bind.FilterOpts) (*PausablePauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausablePauseIterator{contract: _Pausable.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_Pausable *PausableFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausablePause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePause)
				if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PausableUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Pausable contract.
type PausableUnpauseIterator struct {
	Event *PausableUnpause // Event containing the contract specifics and raw log

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
func (it *PausableUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpause)
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
		it.Event = new(PausableUnpause)
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
func (it *PausableUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpause represents a Unpause event raised by the Pausable contract.
type PausableUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableUnpauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableUnpauseIterator{contract: _Pausable.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_Pausable *PausableFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableUnpause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpause)
				if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// PausableTokenABI is the input ABI used to generate the binding from.
const PausableTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setBalanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setAllowanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// PausableTokenBin is the compiled bytecode used for deploying new contracts.
const PausableTokenBin = `0x60806040526004805460a060020a60ff021916905560008054600160a060020a0319163317905561114c806100356000396000f3006080604052600436106101065763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461010b57806318160ddd1461014357806323b872dd1461016a5780633ed10b92146101945780633f4ba83a146101c55780634e71e0c8146101dc57806354f78dad146101f15780635c975abb14610212578063661884631461022757806370a082311461024b5780637bb98a681461026c5780638456cb59146102815780638da5cb5b14610296578063a9059cbb146102ab578063d73dd623146102cf578063dd62ed3e146102f3578063e30c39781461031a578063edc1e4f91461032f578063f2fde38b14610350575b600080fd5b34801561011757600080fd5b5061012f600160a060020a0360043516602435610371565b604080519115158252519081900360200190f35b34801561014f57600080fd5b5061015861039c565b60408051918252519081900360200190f35b34801561017657600080fd5b5061012f600160a060020a03600435811690602435166044356103a2565b3480156101a057600080fd5b506101a96103cf565b60408051600160a060020a039092168252519081900360200190f35b3480156101d157600080fd5b506101da6103de565b005b3480156101e857600080fd5b506101da610456565b3480156101fd57600080fd5b506101da600160a060020a03600435166104de565b34801561021e57600080fd5b5061012f610593565b34801561023357600080fd5b5061012f600160a060020a03600435166024356105a3565b34801561025757600080fd5b50610158600160a060020a03600435166105c7565b34801561027857600080fd5b506101a9610664565b34801561028d57600080fd5b506101da610673565b3480156102a257600080fd5b506101a96106f0565b3480156102b757600080fd5b5061012f600160a060020a03600435166024356106ff565b3480156102db57600080fd5b5061012f600160a060020a0360043516602435610723565b3480156102ff57600080fd5b50610158600160a060020a0360043581169060243516610747565b34801561032657600080fd5b506101a96107d9565b34801561033b57600080fd5b506101da600160a060020a03600435166107e8565b34801561035c57600080fd5b506101da600160a060020a036004351661087e565b60045460009060a060020a900460ff161561038b57600080fd5b61039583836108c4565b9392505050565b60035490565b60045460009060a060020a900460ff16156103bc57600080fd5b6103c78484846108da565b949350505050565b600454600160a060020a031681565b600054600160a060020a031633146103f557600080fd5b60045460a060020a900460ff16151561040d57600080fd5b6004805474ff0000000000000000000000000000000000000000191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b600154600160a060020a0316331461046d57600080fd5b60015460008054604051600160a060020a0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031633146104f557600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055604080517f4e71e0c800000000000000000000000000000000000000000000000000000000815290519290911691634e71e0c89160048082019260009290919082900301818387803b15801561057857600080fd5b505af115801561058c573d6000803e3d6000fd5b5050505050565b60045460a060020a900460ff1681565b60045460009060a060020a900460ff16156105bd57600080fd5b61039583836108f2565b600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915160009392909216916370a082319160248082019260209290919082900301818787803b15801561063257600080fd5b505af1158015610646573d6000803e3d6000fd5b505050506040513d602081101561065c57600080fd5b505192915050565b600254600160a060020a031681565b600054600160a060020a0316331461068a57600080fd5b60045460a060020a900460ff16156106a157600080fd5b6004805474ff0000000000000000000000000000000000000000191660a060020a1790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b600054600160a060020a031681565b60045460009060a060020a900460ff161561071957600080fd5b61039583836108ff565b60045460009060a060020a900460ff161561073d57600080fd5b610395838361090c565b600480546040805160e160020a630d237641028152600160a060020a03868116948201949094528484166024820152905160009390921691631a46ec829160448082019260209290919082900301818787803b1580156107a657600080fd5b505af11580156107ba573d6000803e3d6000fd5b505050506040513d60208110156107d057600080fd5b50519392505050565b600154600160a060020a031681565b600054600160a060020a031633146107ff57600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383811691909117808355604080517f4e71e0c800000000000000000000000000000000000000000000000000000000815290519190921692634e71e0c89280820192600092909182900301818387803b15801561057857600080fd5b600054600160a060020a0316331461089557600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60006108d1838333610919565b50600192915050565b60006108e8848484336109f2565b5060019392505050565b60006108d1838333610b2b565b60006108d1338484610dac565b60006108d1838333610fc9565b60048054604080517fda46098c000000000000000000000000000000000000000000000000000000008152600160a060020a03858116948201949094528684166024820152604481018690529051929091169163da46098c9160648082019260009290919082900301818387803b15801561099357600080fd5b505af11580156109a7573d6000803e3d6000fd5b5050604080518581529051600160a060020a038088169450851692507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360200190a3505050565b600480546040805160e160020a630d237641028152600160a060020a0388811694820194909452848416602482015290519290911691631a46ec82916044808201926020929091908290030181600087803b158015610a5057600080fd5b505af1158015610a64573d6000803e3d6000fd5b505050506040513d6020811015610a7a57600080fd5b5051821115610a8857600080fd5b60048054604080517f97d88cd2000000000000000000000000000000000000000000000000000000008152600160a060020a0388811694820194909452848416602482015260448101869052905192909116916397d88cd29160648082019260009290919082900301818387803b158015610b0257600080fd5b505af1158015610b16573d6000803e3d6000fd5b50505050610b25848484610dac565b50505050565b600480546040805160e160020a630d237641028152600160a060020a03858116948201949094528684166024820152905160009390921691631a46ec829160448082019260209290919082900301818787803b158015610b8a57600080fd5b505af1158015610b9e573d6000803e3d6000fd5b505050506040513d6020811015610bb457600080fd5b5051905080831115610c555760048054604080517fda46098c000000000000000000000000000000000000000000000000000000008152600160a060020a038681169482019490945287841660248201526000604482018190529151939092169263da46098c92606480820193929182900301818387803b158015610c3857600080fd5b505af1158015610c4c573d6000803e3d6000fd5b50505050610ce8565b60048054604080517f97d88cd2000000000000000000000000000000000000000000000000000000008152600160a060020a0386811694820194909452878416602482015260448101879052905192909116916397d88cd29160648082019260009290919082900301818387803b158015610ccf57600080fd5b505af1158015610ce3573d6000803e3d6000fd5b505050505b600480546040805160e160020a630d237641028152600160a060020a038681169482018590528089166024830181905292519294937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92593911691631a46ec829160448083019260209291908290030181600087803b158015610d6957600080fd5b505af1158015610d7d573d6000803e3d6000fd5b505050506040513d6020811015610d9357600080fd5b505160408051918252519081900360200190a350505050565b600160a060020a0382161515610dc157600080fd5b600160a060020a0383161515610dd657600080fd5b600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152915191909216916370a082319160248083019260209291908290030181600087803b158015610e3e57600080fd5b505af1158015610e52573d6000803e3d6000fd5b505050506040513d6020811015610e6857600080fd5b5051811115610e7657600080fd5b600254604080517fcf8eeb7e000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152602482018590529151919092169163cf8eeb7e91604480830192600092919082900301818387803b158015610ee457600080fd5b505af1158015610ef8573d6000803e3d6000fd5b5050600254604080517f21e5383a000000000000000000000000000000000000000000000000000000008152600160a060020a0387811660048301526024820187905291519190921693506321e5383a9250604480830192600092919082900301818387803b158015610f6a57600080fd5b505af1158015610f7e573d6000803e3d6000fd5b5050604080518481529051600160a060020a038087169450871692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3505050565b60048054604080517f5fd72d16000000000000000000000000000000000000000000000000000000008152600160a060020a038581169482019490945286841660248201526044810186905290519290911691635fd72d169160648082019260009290919082900301818387803b15801561104357600080fd5b505af1158015611057573d6000803e3d6000fd5b5050600480546040805160e160020a630d237641028152600160a060020a038781169482018590528981166024830181905292519296509394507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259390921691631a46ec829160448083019260209291908290030181600087803b1580156110de57600080fd5b505af11580156110f2573d6000803e3d6000fd5b505050506040513d602081101561110857600080fd5b505160408051918252519081900360200190a35050505600a165627a7a723058201d939cb3ed6f709ef780f281a20a1f68bd1da46db77de7af36c702922bc8e4730029`

// DeployPausableToken deploys a new Ethereum contract, binding an instance of PausableToken to it.
func DeployPausableToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PausableToken, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PausableToken{PausableTokenCaller: PausableTokenCaller{contract: contract}, PausableTokenTransactor: PausableTokenTransactor{contract: contract}, PausableTokenFilterer: PausableTokenFilterer{contract: contract}}, nil
}

// PausableToken is an auto generated Go binding around an Ethereum contract.
type PausableToken struct {
	PausableTokenCaller     // Read-only binding to the contract
	PausableTokenTransactor // Write-only binding to the contract
	PausableTokenFilterer   // Log filterer for contract events
}

// PausableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableTokenSession struct {
	Contract     *PausableToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableTokenCallerSession struct {
	Contract *PausableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PausableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTokenTransactorSession struct {
	Contract     *PausableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PausableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableTokenRaw struct {
	Contract *PausableToken // Generic contract binding to access the raw methods on
}

// PausableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableTokenCallerRaw struct {
	Contract *PausableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTokenTransactorRaw struct {
	Contract *PausableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausableToken creates a new instance of PausableToken, bound to a specific deployed contract.
func NewPausableToken(address common.Address, backend bind.ContractBackend) (*PausableToken, error) {
	contract, err := bindPausableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PausableToken{PausableTokenCaller: PausableTokenCaller{contract: contract}, PausableTokenTransactor: PausableTokenTransactor{contract: contract}, PausableTokenFilterer: PausableTokenFilterer{contract: contract}}, nil
}

// NewPausableTokenCaller creates a new read-only instance of PausableToken, bound to a specific deployed contract.
func NewPausableTokenCaller(address common.Address, caller bind.ContractCaller) (*PausableTokenCaller, error) {
	contract, err := bindPausableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTokenCaller{contract: contract}, nil
}

// NewPausableTokenTransactor creates a new write-only instance of PausableToken, bound to a specific deployed contract.
func NewPausableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTokenTransactor, error) {
	contract, err := bindPausableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTokenTransactor{contract: contract}, nil
}

// NewPausableTokenFilterer creates a new log filterer instance of PausableToken, bound to a specific deployed contract.
func NewPausableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableTokenFilterer, error) {
	contract, err := bindPausableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableTokenFilterer{contract: contract}, nil
}

// bindPausableToken binds a generic wrapper to an already deployed contract.
func bindPausableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableToken *PausableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PausableToken.Contract.PausableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableToken *PausableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.Contract.PausableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableToken *PausableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableToken.Contract.PausableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableToken *PausableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PausableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableToken *PausableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableToken *PausableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableToken *PausableTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableToken *PausableTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PausableToken.Contract.Allowance(&_PausableToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PausableToken *PausableTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PausableToken.Contract.Allowance(&_PausableToken.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_PausableToken *PausableTokenCaller) Allowances(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "allowances")
	return *ret0, err
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_PausableToken *PausableTokenSession) Allowances() (common.Address, error) {
	return _PausableToken.Contract.Allowances(&_PausableToken.CallOpts)
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_PausableToken *PausableTokenCallerSession) Allowances() (common.Address, error) {
	return _PausableToken.Contract.Allowances(&_PausableToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_PausableToken *PausableTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_PausableToken *PausableTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PausableToken.Contract.BalanceOf(&_PausableToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_PausableToken *PausableTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PausableToken.Contract.BalanceOf(&_PausableToken.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_PausableToken *PausableTokenCaller) Balances(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "balances")
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_PausableToken *PausableTokenSession) Balances() (common.Address, error) {
	return _PausableToken.Contract.Balances(&_PausableToken.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_PausableToken *PausableTokenCallerSession) Balances() (common.Address, error) {
	return _PausableToken.Contract.Balances(&_PausableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableToken *PausableTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableToken *PausableTokenSession) Owner() (common.Address, error) {
	return _PausableToken.Contract.Owner(&_PausableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PausableToken *PausableTokenCallerSession) Owner() (common.Address, error) {
	return _PausableToken.Contract.Owner(&_PausableToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableToken *PausableTokenCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableToken *PausableTokenSession) Paused() (bool, error) {
	return _PausableToken.Contract.Paused(&_PausableToken.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PausableToken *PausableTokenCallerSession) Paused() (bool, error) {
	return _PausableToken.Contract.Paused(&_PausableToken.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_PausableToken *PausableTokenCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_PausableToken *PausableTokenSession) PendingOwner() (common.Address, error) {
	return _PausableToken.Contract.PendingOwner(&_PausableToken.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_PausableToken *PausableTokenCallerSession) PendingOwner() (common.Address, error) {
	return _PausableToken.Contract.PendingOwner(&_PausableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableToken *PausableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PausableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableToken *PausableTokenSession) TotalSupply() (*big.Int, error) {
	return _PausableToken.Contract.TotalSupply(&_PausableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PausableToken *PausableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _PausableToken.Contract.TotalSupply(&_PausableToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Approve(&_PausableToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Approve(&_PausableToken.TransactOpts, _spender, _value)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PausableToken *PausableTokenTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PausableToken *PausableTokenSession) ClaimOwnership() (*types.Transaction, error) {
	return _PausableToken.Contract.ClaimOwnership(&_PausableToken.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_PausableToken *PausableTokenTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _PausableToken.Contract.ClaimOwnership(&_PausableToken.TransactOpts)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.DecreaseApproval(&_PausableToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.DecreaseApproval(&_PausableToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.IncreaseApproval(&_PausableToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PausableToken *PausableTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.IncreaseApproval(&_PausableToken.TransactOpts, _spender, _addedValue)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableToken *PausableTokenTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableToken *PausableTokenSession) Pause() (*types.Transaction, error) {
	return _PausableToken.Contract.Pause(&_PausableToken.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PausableToken *PausableTokenTransactorSession) Pause() (*types.Transaction, error) {
	return _PausableToken.Contract.Pause(&_PausableToken.TransactOpts)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_PausableToken *PausableTokenTransactor) SetAllowanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "setAllowanceSheet", sheet)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_PausableToken *PausableTokenSession) SetAllowanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _PausableToken.Contract.SetAllowanceSheet(&_PausableToken.TransactOpts, sheet)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_PausableToken *PausableTokenTransactorSession) SetAllowanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _PausableToken.Contract.SetAllowanceSheet(&_PausableToken.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_PausableToken *PausableTokenTransactor) SetBalanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "setBalanceSheet", sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_PausableToken *PausableTokenSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _PausableToken.Contract.SetBalanceSheet(&_PausableToken.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_PausableToken *PausableTokenTransactorSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _PausableToken.Contract.SetBalanceSheet(&_PausableToken.TransactOpts, sheet)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Transfer(&_PausableToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.Transfer(&_PausableToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferFrom(&_PausableToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PausableToken *PausableTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferFrom(&_PausableToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PausableToken *PausableTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PausableToken *PausableTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferOwnership(&_PausableToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PausableToken *PausableTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PausableToken.Contract.TransferOwnership(&_PausableToken.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableToken *PausableTokenTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableToken.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableToken *PausableTokenSession) Unpause() (*types.Transaction, error) {
	return _PausableToken.Contract.Unpause(&_PausableToken.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PausableToken *PausableTokenTransactorSession) Unpause() (*types.Transaction, error) {
	return _PausableToken.Contract.Unpause(&_PausableToken.TransactOpts)
}

// PausableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PausableToken contract.
type PausableTokenApprovalIterator struct {
	Event *PausableTokenApproval // Event containing the contract specifics and raw log

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
func (it *PausableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenApproval)
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
		it.Event = new(PausableTokenApproval)
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
func (it *PausableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenApproval represents a Approval event raised by the PausableToken contract.
type PausableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PausableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenApprovalIterator{contract: _PausableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PausableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenApproval)
				if err := _PausableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PausableTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PausableToken contract.
type PausableTokenOwnershipTransferredIterator struct {
	Event *PausableTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PausableTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenOwnershipTransferred)
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
		it.Event = new(PausableTokenOwnershipTransferred)
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
func (it *PausableTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenOwnershipTransferred represents a OwnershipTransferred event raised by the PausableToken contract.
type PausableTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PausableToken *PausableTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenOwnershipTransferredIterator{contract: _PausableToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PausableToken *PausableTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PausableTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenOwnershipTransferred)
				if err := _PausableToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PausableTokenPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PausableToken contract.
type PausableTokenPauseIterator struct {
	Event *PausableTokenPause // Event containing the contract specifics and raw log

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
func (it *PausableTokenPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenPause)
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
		it.Event = new(PausableTokenPause)
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
func (it *PausableTokenPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenPause represents a Pause event raised by the PausableToken contract.
type PausableTokenPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PausableToken *PausableTokenFilterer) FilterPause(opts *bind.FilterOpts) (*PausableTokenPauseIterator, error) {

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausableTokenPauseIterator{contract: _PausableToken.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PausableToken *PausableTokenFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausableTokenPause) (event.Subscription, error) {

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenPause)
				if err := _PausableToken.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PausableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PausableToken contract.
type PausableTokenTransferIterator struct {
	Event *PausableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *PausableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenTransfer)
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
		it.Event = new(PausableTokenTransfer)
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
func (it *PausableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenTransfer represents a Transfer event raised by the PausableToken contract.
type PausableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PausableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PausableTokenTransferIterator{contract: _PausableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_PausableToken *PausableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PausableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenTransfer)
				if err := _PausableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// PausableTokenUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the PausableToken contract.
type PausableTokenUnpauseIterator struct {
	Event *PausableTokenUnpause // Event containing the contract specifics and raw log

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
func (it *PausableTokenUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableTokenUnpause)
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
		it.Event = new(PausableTokenUnpause)
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
func (it *PausableTokenUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableTokenUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableTokenUnpause represents a Unpause event raised by the PausableToken contract.
type PausableTokenUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PausableToken *PausableTokenFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableTokenUnpauseIterator, error) {

	logs, sub, err := _PausableToken.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableTokenUnpauseIterator{contract: _PausableToken.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PausableToken *PausableTokenFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableTokenUnpause) (event.Subscription, error) {

	logs, sub, err := _PausableToken.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableTokenUnpause)
				if err := _PausableToken.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
const SafeERC20Bin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a72305820dc1d94a92bfe6756a65ae2dd89ecd04eb12ad13a68622ea1767d6e9cf1ef7caa0029`

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
const SafeMathBin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058200e3b24b3bf2d8b2fce43659b35e4c27a5ba043a8ee41a0d89f6ec2fb55b2b1b60029`

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StandardDelegateABI is the input ABI used to generate the binding from.
const StandardDelegateABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"delegateAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegatedFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateApprove\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"delegateBalanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateTransferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setBalanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateIncreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setDelegatedFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegateTotalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateDecreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setAllowanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StandardDelegateBin is the compiled bytecode used for deploying new contracts.
const StandardDelegateBin = `0x608060405260008054600160a060020a0319163317905561125f806100256000396000f3006080604052600436106101535763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b3811461015857806309ab8bba1461019057806318160ddd146101c957806323b872dd146101de57806326fe995114610208578063296f4000146102395780633ed10b921461026457806343a468c8146102795780634df6b45d1461029a5780634e71e0c8146102cb57806354f78dad146102e2578063554249b31461030357806361927adb1461032e578063661884631461034f57806370a082311461037357806376e71dd8146103945780637bb98a68146103a95780638da5cb5b146103be57806393d3173a146103d35780639cd1a121146103fe578063a9059cbb14610429578063d73dd6231461044d578063dd62ed3e14610471578063e30c397814610498578063edc1e4f9146104ad578063f2fde38b146104ce575b600080fd5b34801561016457600080fd5b5061017c600160a060020a03600435166024356104ef565b604080519115158252519081900360200190f35b34801561019c57600080fd5b506101b7600160a060020a0360043581169060243516610505565b60408051918252519081900360200190f35b3480156101d557600080fd5b506101b7610518565b3480156101ea57600080fd5b5061017c600160a060020a036004358116906024351660443561051e565b34801561021457600080fd5b5061021d610536565b60408051600160a060020a039092168252519081900360200190f35b34801561024557600080fd5b5061017c600160a060020a036004358116906024359060443516610545565b34801561027057600080fd5b5061021d610576565b34801561028557600080fd5b506101b7600160a060020a0360043516610585565b3480156102a657600080fd5b5061017c600160a060020a036004358116906024358116906044359060643516610596565b3480156102d757600080fd5b506102e06105c9565b005b3480156102ee57600080fd5b506102e0600160a060020a0360043516610651565b34801561030f57600080fd5b5061017c600160a060020a036004358116906024359060443516610706565b34801561033a57600080fd5b506102e0600160a060020a036004351661072c565b34801561035b57600080fd5b5061017c600160a060020a0360043516602435610772565b34801561037f57600080fd5b506101b7600160a060020a036004351661077f565b3480156103a057600080fd5b506101b761081c565b3480156103b557600080fd5b5061021d61082b565b3480156103ca57600080fd5b5061021d61083a565b3480156103df57600080fd5b5061017c600160a060020a036004358116906024359060443516610849565b34801561040a57600080fd5b5061017c600160a060020a03600435811690602435906044351661086f565b34801561043557600080fd5b5061017c600160a060020a0360043516602435610895565b34801561045957600080fd5b5061017c600160a060020a03600435166024356108a2565b34801561047d57600080fd5b506101b7600160a060020a03600435811690602435166108af565b3480156104a457600080fd5b5061021d610941565b3480156104b957600080fd5b506102e0600160a060020a0360043516610950565b3480156104da57600080fd5b506102e0600160a060020a03600435166109e6565b60006104fc838333610a2c565b50600192915050565b600061051183836108af565b9392505050565b60035490565b600061052c84848433610b05565b5060019392505050565b600554600160a060020a031681565b600554600090600160a060020a031633811461056057600080fd5b61056b858585610a2c565b506001949350505050565b600454600160a060020a031681565b60006105908261077f565b92915050565b600554600090600160a060020a03163381146105b157600080fd5b6105bd86868686610b05565b50600195945050505050565b600154600160a060020a031633146105e057600080fd5b60015460008054604051600160a060020a0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a0316331461066857600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055604080517f4e71e0c800000000000000000000000000000000000000000000000000000000815290519290911691634e71e0c89160048082019260009290919082900301818387803b1580156106eb57600080fd5b505af11580156106ff573d6000803e3d6000fd5b5050505050565b600554600090600160a060020a031633811461072157600080fd5b61056b858585610c3e565b600054600160a060020a0316331461074357600080fd5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60006104fc838333610d95565b600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915160009392909216916370a082319160248082019260209290919082900301818787803b1580156107ea57600080fd5b505af11580156107fe573d6000803e3d6000fd5b505050506040513d602081101561081457600080fd5b505192915050565b6000610826610518565b905090565b600254600160a060020a031681565b600054600160a060020a031681565b600554600090600160a060020a031633811461086457600080fd5b61056b858585610d95565b600554600090600160a060020a031633811461088a57600080fd5b61056b838686611016565b60006104fc338484611016565b60006104fc838333610c3e565b600480546040805160e160020a630d237641028152600160a060020a03868116948201949094528484166024820152905160009390921691631a46ec829160448082019260209290919082900301818787803b15801561090e57600080fd5b505af1158015610922573d6000803e3d6000fd5b505050506040513d602081101561093857600080fd5b50519392505050565b600154600160a060020a031681565b600054600160a060020a0316331461096757600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383811691909117808355604080517f4e71e0c800000000000000000000000000000000000000000000000000000000815290519190921692634e71e0c89280820192600092909182900301818387803b1580156106eb57600080fd5b600054600160a060020a031633146109fd57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60048054604080517fda46098c000000000000000000000000000000000000000000000000000000008152600160a060020a03858116948201949094528684166024820152604481018690529051929091169163da46098c9160648082019260009290919082900301818387803b158015610aa657600080fd5b505af1158015610aba573d6000803e3d6000fd5b5050604080518581529051600160a060020a038088169450851692507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360200190a3505050565b600480546040805160e160020a630d237641028152600160a060020a0388811694820194909452848416602482015290519290911691631a46ec82916044808201926020929091908290030181600087803b158015610b6357600080fd5b505af1158015610b77573d6000803e3d6000fd5b505050506040513d6020811015610b8d57600080fd5b5051821115610b9b57600080fd5b60048054604080517f97d88cd2000000000000000000000000000000000000000000000000000000008152600160a060020a0388811694820194909452848416602482015260448101869052905192909116916397d88cd29160648082019260009290919082900301818387803b158015610c1557600080fd5b505af1158015610c29573d6000803e3d6000fd5b50505050610c38848484611016565b50505050565b60048054604080517f5fd72d16000000000000000000000000000000000000000000000000000000008152600160a060020a038581169482019490945286841660248201526044810186905290519290911691635fd72d169160648082019260009290919082900301818387803b158015610cb857600080fd5b505af1158015610ccc573d6000803e3d6000fd5b5050600480546040805160e160020a630d237641028152600160a060020a038781169482018590528981166024830181905292519296509394507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259390921691631a46ec829160448083019260209291908290030181600087803b158015610d5357600080fd5b505af1158015610d67573d6000803e3d6000fd5b505050506040513d6020811015610d7d57600080fd5b505160408051918252519081900360200190a3505050565b600480546040805160e160020a630d237641028152600160a060020a03858116948201949094528684166024820152905160009390921691631a46ec829160448082019260209290919082900301818787803b158015610df457600080fd5b505af1158015610e08573d6000803e3d6000fd5b505050506040513d6020811015610e1e57600080fd5b5051905080831115610ebf5760048054604080517fda46098c000000000000000000000000000000000000000000000000000000008152600160a060020a038681169482019490945287841660248201526000604482018190529151939092169263da46098c92606480820193929182900301818387803b158015610ea257600080fd5b505af1158015610eb6573d6000803e3d6000fd5b50505050610f52565b60048054604080517f97d88cd2000000000000000000000000000000000000000000000000000000008152600160a060020a0386811694820194909452878416602482015260448101879052905192909116916397d88cd29160648082019260009290919082900301818387803b158015610f3957600080fd5b505af1158015610f4d573d6000803e3d6000fd5b505050505b600480546040805160e160020a630d237641028152600160a060020a038681169482018590528089166024830181905292519294937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92593911691631a46ec829160448083019260209291908290030181600087803b158015610fd357600080fd5b505af1158015610fe7573d6000803e3d6000fd5b505050506040513d6020811015610ffd57600080fd5b505160408051918252519081900360200190a350505050565b600160a060020a038216151561102b57600080fd5b600160a060020a038316151561104057600080fd5b600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152915191909216916370a082319160248083019260209291908290030181600087803b1580156110a857600080fd5b505af11580156110bc573d6000803e3d6000fd5b505050506040513d60208110156110d257600080fd5b50518111156110e057600080fd5b600254604080517fcf8eeb7e000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152602482018590529151919092169163cf8eeb7e91604480830192600092919082900301818387803b15801561114e57600080fd5b505af1158015611162573d6000803e3d6000fd5b5050600254604080517f21e5383a000000000000000000000000000000000000000000000000000000008152600160a060020a0387811660048301526024820187905291519190921693506321e5383a9250604480830192600092919082900301818387803b1580156111d457600080fd5b505af11580156111e8573d6000803e3d6000fd5b5050604080518481529051600160a060020a038087169450871692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050505600a165627a7a72305820343f3c57c84f5dd40185e51a7a32cdc692b6908e78c36815f151ad7016fa6fcd0029`

// DeployStandardDelegate deploys a new Ethereum contract, binding an instance of StandardDelegate to it.
func DeployStandardDelegate(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardDelegate, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardDelegateABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardDelegateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardDelegate{StandardDelegateCaller: StandardDelegateCaller{contract: contract}, StandardDelegateTransactor: StandardDelegateTransactor{contract: contract}, StandardDelegateFilterer: StandardDelegateFilterer{contract: contract}}, nil
}

// StandardDelegate is an auto generated Go binding around an Ethereum contract.
type StandardDelegate struct {
	StandardDelegateCaller     // Read-only binding to the contract
	StandardDelegateTransactor // Write-only binding to the contract
	StandardDelegateFilterer   // Log filterer for contract events
}

// StandardDelegateCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardDelegateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardDelegateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardDelegateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardDelegateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardDelegateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardDelegateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardDelegateSession struct {
	Contract     *StandardDelegate // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardDelegateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardDelegateCallerSession struct {
	Contract *StandardDelegateCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// StandardDelegateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardDelegateTransactorSession struct {
	Contract     *StandardDelegateTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// StandardDelegateRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardDelegateRaw struct {
	Contract *StandardDelegate // Generic contract binding to access the raw methods on
}

// StandardDelegateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardDelegateCallerRaw struct {
	Contract *StandardDelegateCaller // Generic read-only contract binding to access the raw methods on
}

// StandardDelegateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardDelegateTransactorRaw struct {
	Contract *StandardDelegateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardDelegate creates a new instance of StandardDelegate, bound to a specific deployed contract.
func NewStandardDelegate(address common.Address, backend bind.ContractBackend) (*StandardDelegate, error) {
	contract, err := bindStandardDelegate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardDelegate{StandardDelegateCaller: StandardDelegateCaller{contract: contract}, StandardDelegateTransactor: StandardDelegateTransactor{contract: contract}, StandardDelegateFilterer: StandardDelegateFilterer{contract: contract}}, nil
}

// NewStandardDelegateCaller creates a new read-only instance of StandardDelegate, bound to a specific deployed contract.
func NewStandardDelegateCaller(address common.Address, caller bind.ContractCaller) (*StandardDelegateCaller, error) {
	contract, err := bindStandardDelegate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardDelegateCaller{contract: contract}, nil
}

// NewStandardDelegateTransactor creates a new write-only instance of StandardDelegate, bound to a specific deployed contract.
func NewStandardDelegateTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardDelegateTransactor, error) {
	contract, err := bindStandardDelegate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardDelegateTransactor{contract: contract}, nil
}

// NewStandardDelegateFilterer creates a new log filterer instance of StandardDelegate, bound to a specific deployed contract.
func NewStandardDelegateFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardDelegateFilterer, error) {
	contract, err := bindStandardDelegate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardDelegateFilterer{contract: contract}, nil
}

// bindStandardDelegate binds a generic wrapper to an already deployed contract.
func bindStandardDelegate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardDelegateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardDelegate *StandardDelegateRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardDelegate.Contract.StandardDelegateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardDelegate *StandardDelegateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardDelegate.Contract.StandardDelegateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardDelegate *StandardDelegateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardDelegate.Contract.StandardDelegateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardDelegate *StandardDelegateCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardDelegate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardDelegate *StandardDelegateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardDelegate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardDelegate *StandardDelegateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardDelegate.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardDelegate *StandardDelegateCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardDelegate.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardDelegate *StandardDelegateSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardDelegate.Contract.Allowance(&_StandardDelegate.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardDelegate *StandardDelegateCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardDelegate.Contract.Allowance(&_StandardDelegate.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_StandardDelegate *StandardDelegateCaller) Allowances(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StandardDelegate.contract.Call(opts, out, "allowances")
	return *ret0, err
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_StandardDelegate *StandardDelegateSession) Allowances() (common.Address, error) {
	return _StandardDelegate.Contract.Allowances(&_StandardDelegate.CallOpts)
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_StandardDelegate *StandardDelegateCallerSession) Allowances() (common.Address, error) {
	return _StandardDelegate.Contract.Allowances(&_StandardDelegate.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardDelegate *StandardDelegateCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardDelegate.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardDelegate *StandardDelegateSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardDelegate.Contract.BalanceOf(&_StandardDelegate.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardDelegate *StandardDelegateCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardDelegate.Contract.BalanceOf(&_StandardDelegate.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_StandardDelegate *StandardDelegateCaller) Balances(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StandardDelegate.contract.Call(opts, out, "balances")
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_StandardDelegate *StandardDelegateSession) Balances() (common.Address, error) {
	return _StandardDelegate.Contract.Balances(&_StandardDelegate.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_StandardDelegate *StandardDelegateCallerSession) Balances() (common.Address, error) {
	return _StandardDelegate.Contract.Balances(&_StandardDelegate.CallOpts)
}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(owner address, spender address) constant returns(uint256)
func (_StandardDelegate *StandardDelegateCaller) DelegateAllowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardDelegate.contract.Call(opts, out, "delegateAllowance", owner, spender)
	return *ret0, err
}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(owner address, spender address) constant returns(uint256)
func (_StandardDelegate *StandardDelegateSession) DelegateAllowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StandardDelegate.Contract.DelegateAllowance(&_StandardDelegate.CallOpts, owner, spender)
}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(owner address, spender address) constant returns(uint256)
func (_StandardDelegate *StandardDelegateCallerSession) DelegateAllowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StandardDelegate.Contract.DelegateAllowance(&_StandardDelegate.CallOpts, owner, spender)
}

// DelegateBalanceOf is a free data retrieval call binding the contract method 0x43a468c8.
//
// Solidity: function delegateBalanceOf(who address) constant returns(uint256)
func (_StandardDelegate *StandardDelegateCaller) DelegateBalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardDelegate.contract.Call(opts, out, "delegateBalanceOf", who)
	return *ret0, err
}

// DelegateBalanceOf is a free data retrieval call binding the contract method 0x43a468c8.
//
// Solidity: function delegateBalanceOf(who address) constant returns(uint256)
func (_StandardDelegate *StandardDelegateSession) DelegateBalanceOf(who common.Address) (*big.Int, error) {
	return _StandardDelegate.Contract.DelegateBalanceOf(&_StandardDelegate.CallOpts, who)
}

// DelegateBalanceOf is a free data retrieval call binding the contract method 0x43a468c8.
//
// Solidity: function delegateBalanceOf(who address) constant returns(uint256)
func (_StandardDelegate *StandardDelegateCallerSession) DelegateBalanceOf(who common.Address) (*big.Int, error) {
	return _StandardDelegate.Contract.DelegateBalanceOf(&_StandardDelegate.CallOpts, who)
}

// DelegateTotalSupply is a free data retrieval call binding the contract method 0x76e71dd8.
//
// Solidity: function delegateTotalSupply() constant returns(uint256)
func (_StandardDelegate *StandardDelegateCaller) DelegateTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardDelegate.contract.Call(opts, out, "delegateTotalSupply")
	return *ret0, err
}

// DelegateTotalSupply is a free data retrieval call binding the contract method 0x76e71dd8.
//
// Solidity: function delegateTotalSupply() constant returns(uint256)
func (_StandardDelegate *StandardDelegateSession) DelegateTotalSupply() (*big.Int, error) {
	return _StandardDelegate.Contract.DelegateTotalSupply(&_StandardDelegate.CallOpts)
}

// DelegateTotalSupply is a free data retrieval call binding the contract method 0x76e71dd8.
//
// Solidity: function delegateTotalSupply() constant returns(uint256)
func (_StandardDelegate *StandardDelegateCallerSession) DelegateTotalSupply() (*big.Int, error) {
	return _StandardDelegate.Contract.DelegateTotalSupply(&_StandardDelegate.CallOpts)
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x26fe9951.
//
// Solidity: function delegatedFrom() constant returns(address)
func (_StandardDelegate *StandardDelegateCaller) DelegatedFrom(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StandardDelegate.contract.Call(opts, out, "delegatedFrom")
	return *ret0, err
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x26fe9951.
//
// Solidity: function delegatedFrom() constant returns(address)
func (_StandardDelegate *StandardDelegateSession) DelegatedFrom() (common.Address, error) {
	return _StandardDelegate.Contract.DelegatedFrom(&_StandardDelegate.CallOpts)
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x26fe9951.
//
// Solidity: function delegatedFrom() constant returns(address)
func (_StandardDelegate *StandardDelegateCallerSession) DelegatedFrom() (common.Address, error) {
	return _StandardDelegate.Contract.DelegatedFrom(&_StandardDelegate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StandardDelegate *StandardDelegateCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StandardDelegate.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StandardDelegate *StandardDelegateSession) Owner() (common.Address, error) {
	return _StandardDelegate.Contract.Owner(&_StandardDelegate.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StandardDelegate *StandardDelegateCallerSession) Owner() (common.Address, error) {
	return _StandardDelegate.Contract.Owner(&_StandardDelegate.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_StandardDelegate *StandardDelegateCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StandardDelegate.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_StandardDelegate *StandardDelegateSession) PendingOwner() (common.Address, error) {
	return _StandardDelegate.Contract.PendingOwner(&_StandardDelegate.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_StandardDelegate *StandardDelegateCallerSession) PendingOwner() (common.Address, error) {
	return _StandardDelegate.Contract.PendingOwner(&_StandardDelegate.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardDelegate *StandardDelegateCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardDelegate.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardDelegate *StandardDelegateSession) TotalSupply() (*big.Int, error) {
	return _StandardDelegate.Contract.TotalSupply(&_StandardDelegate.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardDelegate *StandardDelegateCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardDelegate.Contract.TotalSupply(&_StandardDelegate.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardDelegate *StandardDelegateTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardDelegate *StandardDelegateSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.Contract.Approve(&_StandardDelegate.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardDelegate *StandardDelegateTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.Contract.Approve(&_StandardDelegate.TransactOpts, _spender, _value)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_StandardDelegate *StandardDelegateTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_StandardDelegate *StandardDelegateSession) ClaimOwnership() (*types.Transaction, error) {
	return _StandardDelegate.Contract.ClaimOwnership(&_StandardDelegate.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_StandardDelegate *StandardDelegateTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _StandardDelegate.Contract.ClaimOwnership(&_StandardDelegate.TransactOpts)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardDelegate *StandardDelegateTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardDelegate *StandardDelegateSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.Contract.DecreaseApproval(&_StandardDelegate.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardDelegate *StandardDelegateTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.Contract.DecreaseApproval(&_StandardDelegate.TransactOpts, _spender, _subtractedValue)
}

// DelegateApprove is a paid mutator transaction binding the contract method 0x296f4000.
//
// Solidity: function delegateApprove(spender address, value uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateTransactor) DelegateApprove(opts *bind.TransactOpts, spender common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "delegateApprove", spender, value, origSender)
}

// DelegateApprove is a paid mutator transaction binding the contract method 0x296f4000.
//
// Solidity: function delegateApprove(spender address, value uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateSession) DelegateApprove(spender common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.DelegateApprove(&_StandardDelegate.TransactOpts, spender, value, origSender)
}

// DelegateApprove is a paid mutator transaction binding the contract method 0x296f4000.
//
// Solidity: function delegateApprove(spender address, value uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateTransactorSession) DelegateApprove(spender common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.DelegateApprove(&_StandardDelegate.TransactOpts, spender, value, origSender)
}

// DelegateDecreaseApproval is a paid mutator transaction binding the contract method 0x93d3173a.
//
// Solidity: function delegateDecreaseApproval(spender address, subtractedValue uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateTransactor) DelegateDecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "delegateDecreaseApproval", spender, subtractedValue, origSender)
}

// DelegateDecreaseApproval is a paid mutator transaction binding the contract method 0x93d3173a.
//
// Solidity: function delegateDecreaseApproval(spender address, subtractedValue uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateSession) DelegateDecreaseApproval(spender common.Address, subtractedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.DelegateDecreaseApproval(&_StandardDelegate.TransactOpts, spender, subtractedValue, origSender)
}

// DelegateDecreaseApproval is a paid mutator transaction binding the contract method 0x93d3173a.
//
// Solidity: function delegateDecreaseApproval(spender address, subtractedValue uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateTransactorSession) DelegateDecreaseApproval(spender common.Address, subtractedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.DelegateDecreaseApproval(&_StandardDelegate.TransactOpts, spender, subtractedValue, origSender)
}

// DelegateIncreaseApproval is a paid mutator transaction binding the contract method 0x554249b3.
//
// Solidity: function delegateIncreaseApproval(spender address, addedValue uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateTransactor) DelegateIncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "delegateIncreaseApproval", spender, addedValue, origSender)
}

// DelegateIncreaseApproval is a paid mutator transaction binding the contract method 0x554249b3.
//
// Solidity: function delegateIncreaseApproval(spender address, addedValue uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateSession) DelegateIncreaseApproval(spender common.Address, addedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.DelegateIncreaseApproval(&_StandardDelegate.TransactOpts, spender, addedValue, origSender)
}

// DelegateIncreaseApproval is a paid mutator transaction binding the contract method 0x554249b3.
//
// Solidity: function delegateIncreaseApproval(spender address, addedValue uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateTransactorSession) DelegateIncreaseApproval(spender common.Address, addedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.DelegateIncreaseApproval(&_StandardDelegate.TransactOpts, spender, addedValue, origSender)
}

// DelegateTransfer is a paid mutator transaction binding the contract method 0x9cd1a121.
//
// Solidity: function delegateTransfer(to address, value uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateTransactor) DelegateTransfer(opts *bind.TransactOpts, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "delegateTransfer", to, value, origSender)
}

// DelegateTransfer is a paid mutator transaction binding the contract method 0x9cd1a121.
//
// Solidity: function delegateTransfer(to address, value uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateSession) DelegateTransfer(to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.DelegateTransfer(&_StandardDelegate.TransactOpts, to, value, origSender)
}

// DelegateTransfer is a paid mutator transaction binding the contract method 0x9cd1a121.
//
// Solidity: function delegateTransfer(to address, value uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateTransactorSession) DelegateTransfer(to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.DelegateTransfer(&_StandardDelegate.TransactOpts, to, value, origSender)
}

// DelegateTransferFrom is a paid mutator transaction binding the contract method 0x4df6b45d.
//
// Solidity: function delegateTransferFrom(from address, to address, value uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateTransactor) DelegateTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "delegateTransferFrom", from, to, value, origSender)
}

// DelegateTransferFrom is a paid mutator transaction binding the contract method 0x4df6b45d.
//
// Solidity: function delegateTransferFrom(from address, to address, value uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateSession) DelegateTransferFrom(from common.Address, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.DelegateTransferFrom(&_StandardDelegate.TransactOpts, from, to, value, origSender)
}

// DelegateTransferFrom is a paid mutator transaction binding the contract method 0x4df6b45d.
//
// Solidity: function delegateTransferFrom(from address, to address, value uint256, origSender address) returns(bool)
func (_StandardDelegate *StandardDelegateTransactorSession) DelegateTransferFrom(from common.Address, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.DelegateTransferFrom(&_StandardDelegate.TransactOpts, from, to, value, origSender)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardDelegate *StandardDelegateTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardDelegate *StandardDelegateSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.Contract.IncreaseApproval(&_StandardDelegate.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardDelegate *StandardDelegateTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.Contract.IncreaseApproval(&_StandardDelegate.TransactOpts, _spender, _addedValue)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_StandardDelegate *StandardDelegateTransactor) SetAllowanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "setAllowanceSheet", sheet)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_StandardDelegate *StandardDelegateSession) SetAllowanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.SetAllowanceSheet(&_StandardDelegate.TransactOpts, sheet)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_StandardDelegate *StandardDelegateTransactorSession) SetAllowanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.SetAllowanceSheet(&_StandardDelegate.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_StandardDelegate *StandardDelegateTransactor) SetBalanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "setBalanceSheet", sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_StandardDelegate *StandardDelegateSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.SetBalanceSheet(&_StandardDelegate.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_StandardDelegate *StandardDelegateTransactorSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.SetBalanceSheet(&_StandardDelegate.TransactOpts, sheet)
}

// SetDelegatedFrom is a paid mutator transaction binding the contract method 0x61927adb.
//
// Solidity: function setDelegatedFrom(addr address) returns()
func (_StandardDelegate *StandardDelegateTransactor) SetDelegatedFrom(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "setDelegatedFrom", addr)
}

// SetDelegatedFrom is a paid mutator transaction binding the contract method 0x61927adb.
//
// Solidity: function setDelegatedFrom(addr address) returns()
func (_StandardDelegate *StandardDelegateSession) SetDelegatedFrom(addr common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.SetDelegatedFrom(&_StandardDelegate.TransactOpts, addr)
}

// SetDelegatedFrom is a paid mutator transaction binding the contract method 0x61927adb.
//
// Solidity: function setDelegatedFrom(addr address) returns()
func (_StandardDelegate *StandardDelegateTransactorSession) SetDelegatedFrom(addr common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.SetDelegatedFrom(&_StandardDelegate.TransactOpts, addr)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardDelegate *StandardDelegateTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardDelegate *StandardDelegateSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.Contract.Transfer(&_StandardDelegate.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardDelegate *StandardDelegateTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.Contract.Transfer(&_StandardDelegate.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardDelegate *StandardDelegateTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardDelegate *StandardDelegateSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.Contract.TransferFrom(&_StandardDelegate.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardDelegate *StandardDelegateTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardDelegate.Contract.TransferFrom(&_StandardDelegate.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_StandardDelegate *StandardDelegateTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StandardDelegate.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_StandardDelegate *StandardDelegateSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.TransferOwnership(&_StandardDelegate.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_StandardDelegate *StandardDelegateTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StandardDelegate.Contract.TransferOwnership(&_StandardDelegate.TransactOpts, newOwner)
}

// StandardDelegateApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardDelegate contract.
type StandardDelegateApprovalIterator struct {
	Event *StandardDelegateApproval // Event containing the contract specifics and raw log

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
func (it *StandardDelegateApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardDelegateApproval)
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
		it.Event = new(StandardDelegateApproval)
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
func (it *StandardDelegateApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardDelegateApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardDelegateApproval represents a Approval event raised by the StandardDelegate contract.
type StandardDelegateApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardDelegate *StandardDelegateFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardDelegateApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardDelegate.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardDelegateApprovalIterator{contract: _StandardDelegate.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardDelegate *StandardDelegateFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardDelegateApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardDelegate.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardDelegateApproval)
				if err := _StandardDelegate.contract.UnpackLog(event, "Approval", log); err != nil {
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

// StandardDelegateOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StandardDelegate contract.
type StandardDelegateOwnershipTransferredIterator struct {
	Event *StandardDelegateOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StandardDelegateOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardDelegateOwnershipTransferred)
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
		it.Event = new(StandardDelegateOwnershipTransferred)
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
func (it *StandardDelegateOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardDelegateOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardDelegateOwnershipTransferred represents a OwnershipTransferred event raised by the StandardDelegate contract.
type StandardDelegateOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_StandardDelegate *StandardDelegateFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StandardDelegateOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StandardDelegate.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StandardDelegateOwnershipTransferredIterator{contract: _StandardDelegate.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_StandardDelegate *StandardDelegateFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StandardDelegateOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StandardDelegate.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardDelegateOwnershipTransferred)
				if err := _StandardDelegate.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// StandardDelegateTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardDelegate contract.
type StandardDelegateTransferIterator struct {
	Event *StandardDelegateTransfer // Event containing the contract specifics and raw log

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
func (it *StandardDelegateTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardDelegateTransfer)
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
		it.Event = new(StandardDelegateTransfer)
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
func (it *StandardDelegateTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardDelegateTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardDelegateTransfer represents a Transfer event raised by the StandardDelegate contract.
type StandardDelegateTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardDelegate *StandardDelegateFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardDelegateTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardDelegate.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardDelegateTransferIterator{contract: _StandardDelegate.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardDelegate *StandardDelegateFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardDelegateTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardDelegate.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardDelegateTransfer)
				if err := _StandardDelegate.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// StandardTokenABI is the input ABI used to generate the binding from.
const StandardTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setBalanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setAllowanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// StandardTokenBin is the compiled bytecode used for deploying new contracts.
const StandardTokenBin = `0x608060405260008054600160a060020a03191633179055610f23806100256000396000f3006080604052600436106100e55763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663095ea7b381146100ea57806318160ddd1461012257806323b872dd146101495780633ed10b92146101735780634e71e0c8146101a457806354f78dad146101bb57806366188463146101dc57806370a08231146102005780637bb98a68146102215780638da5cb5b14610236578063a9059cbb1461024b578063d73dd6231461026f578063dd62ed3e14610293578063e30c3978146102ba578063edc1e4f9146102cf578063f2fde38b146102f0575b600080fd5b3480156100f657600080fd5b5061010e600160a060020a0360043516602435610311565b604080519115158252519081900360200190f35b34801561012e57600080fd5b50610137610327565b60408051918252519081900360200190f35b34801561015557600080fd5b5061010e600160a060020a036004358116906024351660443561032d565b34801561017f57600080fd5b50610188610345565b60408051600160a060020a039092168252519081900360200190f35b3480156101b057600080fd5b506101b9610354565b005b3480156101c757600080fd5b506101b9600160a060020a03600435166103dc565b3480156101e857600080fd5b5061010e600160a060020a0360043516602435610491565b34801561020c57600080fd5b50610137600160a060020a036004351661049e565b34801561022d57600080fd5b5061018861053b565b34801561024257600080fd5b5061018861054a565b34801561025757600080fd5b5061010e600160a060020a0360043516602435610559565b34801561027b57600080fd5b5061010e600160a060020a0360043516602435610566565b34801561029f57600080fd5b50610137600160a060020a0360043581169060243516610573565b3480156102c657600080fd5b50610188610605565b3480156102db57600080fd5b506101b9600160a060020a0360043516610614565b3480156102fc57600080fd5b506101b9600160a060020a03600435166106aa565b600061031e8383336106f0565b50600192915050565b60035490565b600061033b848484336107c9565b5060019392505050565b600454600160a060020a031681565b600154600160a060020a0316331461036b57600080fd5b60015460008054604051600160a060020a0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600180546000805473ffffffffffffffffffffffffffffffffffffffff19908116600160a060020a03841617909155169055565b600054600160a060020a031633146103f357600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055604080517f4e71e0c800000000000000000000000000000000000000000000000000000000815290519290911691634e71e0c89160048082019260009290919082900301818387803b15801561047657600080fd5b505af115801561048a573d6000803e3d6000fd5b5050505050565b600061031e838333610902565b600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038481166004830152915160009392909216916370a082319160248082019260209290919082900301818787803b15801561050957600080fd5b505af115801561051d573d6000803e3d6000fd5b505050506040513d602081101561053357600080fd5b505192915050565b600254600160a060020a031681565b600054600160a060020a031681565b600061031e338484610b83565b600061031e838333610da0565b600480546040805160e160020a630d237641028152600160a060020a03868116948201949094528484166024820152905160009390921691631a46ec829160448082019260209290919082900301818787803b1580156105d257600080fd5b505af11580156105e6573d6000803e3d6000fd5b505050506040513d60208110156105fc57600080fd5b50519392505050565b600154600160a060020a031681565b600054600160a060020a0316331461062b57600080fd5b6004805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383811691909117808355604080517f4e71e0c800000000000000000000000000000000000000000000000000000000815290519190921692634e71e0c89280820192600092909182900301818387803b15801561047657600080fd5b600054600160a060020a031633146106c157600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60048054604080517fda46098c000000000000000000000000000000000000000000000000000000008152600160a060020a03858116948201949094528684166024820152604481018690529051929091169163da46098c9160648082019260009290919082900301818387803b15801561076a57600080fd5b505af115801561077e573d6000803e3d6000fd5b5050604080518581529051600160a060020a038088169450851692507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360200190a3505050565b600480546040805160e160020a630d237641028152600160a060020a0388811694820194909452848416602482015290519290911691631a46ec82916044808201926020929091908290030181600087803b15801561082757600080fd5b505af115801561083b573d6000803e3d6000fd5b505050506040513d602081101561085157600080fd5b505182111561085f57600080fd5b60048054604080517f97d88cd2000000000000000000000000000000000000000000000000000000008152600160a060020a0388811694820194909452848416602482015260448101869052905192909116916397d88cd29160648082019260009290919082900301818387803b1580156108d957600080fd5b505af11580156108ed573d6000803e3d6000fd5b505050506108fc848484610b83565b50505050565b600480546040805160e160020a630d237641028152600160a060020a03858116948201949094528684166024820152905160009390921691631a46ec829160448082019260209290919082900301818787803b15801561096157600080fd5b505af1158015610975573d6000803e3d6000fd5b505050506040513d602081101561098b57600080fd5b5051905080831115610a2c5760048054604080517fda46098c000000000000000000000000000000000000000000000000000000008152600160a060020a038681169482019490945287841660248201526000604482018190529151939092169263da46098c92606480820193929182900301818387803b158015610a0f57600080fd5b505af1158015610a23573d6000803e3d6000fd5b50505050610abf565b60048054604080517f97d88cd2000000000000000000000000000000000000000000000000000000008152600160a060020a0386811694820194909452878416602482015260448101879052905192909116916397d88cd29160648082019260009290919082900301818387803b158015610aa657600080fd5b505af1158015610aba573d6000803e3d6000fd5b505050505b600480546040805160e160020a630d237641028152600160a060020a038681169482018590528089166024830181905292519294937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92593911691631a46ec829160448083019260209291908290030181600087803b158015610b4057600080fd5b505af1158015610b54573d6000803e3d6000fd5b505050506040513d6020811015610b6a57600080fd5b505160408051918252519081900360200190a350505050565b600160a060020a0382161515610b9857600080fd5b600160a060020a0383161515610bad57600080fd5b600254604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152915191909216916370a082319160248083019260209291908290030181600087803b158015610c1557600080fd5b505af1158015610c29573d6000803e3d6000fd5b505050506040513d6020811015610c3f57600080fd5b5051811115610c4d57600080fd5b600254604080517fcf8eeb7e000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152602482018590529151919092169163cf8eeb7e91604480830192600092919082900301818387803b158015610cbb57600080fd5b505af1158015610ccf573d6000803e3d6000fd5b5050600254604080517f21e5383a000000000000000000000000000000000000000000000000000000008152600160a060020a0387811660048301526024820187905291519190921693506321e5383a9250604480830192600092919082900301818387803b158015610d4157600080fd5b505af1158015610d55573d6000803e3d6000fd5b5050604080518481529051600160a060020a038087169450871692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3505050565b60048054604080517f5fd72d16000000000000000000000000000000000000000000000000000000008152600160a060020a038581169482019490945286841660248201526044810186905290519290911691635fd72d169160648082019260009290919082900301818387803b158015610e1a57600080fd5b505af1158015610e2e573d6000803e3d6000fd5b5050600480546040805160e160020a630d237641028152600160a060020a038781169482018590528981166024830181905292519296509394507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259390921691631a46ec829160448083019260209291908290030181600087803b158015610eb557600080fd5b505af1158015610ec9573d6000803e3d6000fd5b505050506040513d6020811015610edf57600080fd5b505160408051918252519081900360200190a35050505600a165627a7a723058201d7fcff8931aade5685bcbe2cd04f8063b1a6032bd7be573cd4663f394bf77d90029`

// DeployStandardToken deploys a new Ethereum contract, binding an instance of StandardToken to it.
func DeployStandardToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StandardToken, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StandardTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// StandardToken is an auto generated Go binding around an Ethereum contract.
type StandardToken struct {
	StandardTokenCaller     // Read-only binding to the contract
	StandardTokenTransactor // Write-only binding to the contract
	StandardTokenFilterer   // Log filterer for contract events
}

// StandardTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandardTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandardTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandardTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandardTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandardTokenSession struct {
	Contract     *StandardToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandardTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandardTokenCallerSession struct {
	Contract *StandardTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandardTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandardTokenTransactorSession struct {
	Contract     *StandardTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandardTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandardTokenRaw struct {
	Contract *StandardToken // Generic contract binding to access the raw methods on
}

// StandardTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandardTokenCallerRaw struct {
	Contract *StandardTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandardTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandardTokenTransactorRaw struct {
	Contract *StandardTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandardToken creates a new instance of StandardToken, bound to a specific deployed contract.
func NewStandardToken(address common.Address, backend bind.ContractBackend) (*StandardToken, error) {
	contract, err := bindStandardToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandardToken{StandardTokenCaller: StandardTokenCaller{contract: contract}, StandardTokenTransactor: StandardTokenTransactor{contract: contract}, StandardTokenFilterer: StandardTokenFilterer{contract: contract}}, nil
}

// NewStandardTokenCaller creates a new read-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenCaller(address common.Address, caller bind.ContractCaller) (*StandardTokenCaller, error) {
	contract, err := bindStandardToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenCaller{contract: contract}, nil
}

// NewStandardTokenTransactor creates a new write-only instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandardTokenTransactor, error) {
	contract, err := bindStandardToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransactor{contract: contract}, nil
}

// NewStandardTokenFilterer creates a new log filterer instance of StandardToken, bound to a specific deployed contract.
func NewStandardTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandardTokenFilterer, error) {
	contract, err := bindStandardToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandardTokenFilterer{contract: contract}, nil
}

// bindStandardToken binds a generic wrapper to an already deployed contract.
func bindStandardToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandardTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.StandardTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.StandardTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandardToken *StandardTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StandardToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandardToken *StandardTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandardToken *StandardTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandardToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _StandardToken.Contract.Allowance(&_StandardToken.CallOpts, _owner, _spender)
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_StandardToken *StandardTokenCaller) Allowances(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "allowances")
	return *ret0, err
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_StandardToken *StandardTokenSession) Allowances() (common.Address, error) {
	return _StandardToken.Contract.Allowances(&_StandardToken.CallOpts)
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_StandardToken *StandardTokenCallerSession) Allowances() (common.Address, error) {
	return _StandardToken.Contract.Allowances(&_StandardToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_StandardToken *StandardTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _StandardToken.Contract.BalanceOf(&_StandardToken.CallOpts, _owner)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_StandardToken *StandardTokenCaller) Balances(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "balances")
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_StandardToken *StandardTokenSession) Balances() (common.Address, error) {
	return _StandardToken.Contract.Balances(&_StandardToken.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_StandardToken *StandardTokenCallerSession) Balances() (common.Address, error) {
	return _StandardToken.Contract.Balances(&_StandardToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StandardToken *StandardTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StandardToken *StandardTokenSession) Owner() (common.Address, error) {
	return _StandardToken.Contract.Owner(&_StandardToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StandardToken *StandardTokenCallerSession) Owner() (common.Address, error) {
	return _StandardToken.Contract.Owner(&_StandardToken.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_StandardToken *StandardTokenCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_StandardToken *StandardTokenSession) PendingOwner() (common.Address, error) {
	return _StandardToken.Contract.PendingOwner(&_StandardToken.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_StandardToken *StandardTokenCallerSession) PendingOwner() (common.Address, error) {
	return _StandardToken.Contract.PendingOwner(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StandardToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StandardToken *StandardTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandardToken.Contract.TotalSupply(&_StandardToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Approve(&_StandardToken.TransactOpts, _spender, _value)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_StandardToken *StandardTokenTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_StandardToken *StandardTokenSession) ClaimOwnership() (*types.Transaction, error) {
	return _StandardToken.Contract.ClaimOwnership(&_StandardToken.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_StandardToken *StandardTokenTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _StandardToken.Contract.ClaimOwnership(&_StandardToken.TransactOpts)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.DecreaseApproval(&_StandardToken.TransactOpts, _spender, _subtractedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.IncreaseApproval(&_StandardToken.TransactOpts, _spender, _addedValue)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_StandardToken *StandardTokenTransactor) SetAllowanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "setAllowanceSheet", sheet)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_StandardToken *StandardTokenSession) SetAllowanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.SetAllowanceSheet(&_StandardToken.TransactOpts, sheet)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_StandardToken *StandardTokenTransactorSession) SetAllowanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.SetAllowanceSheet(&_StandardToken.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_StandardToken *StandardTokenTransactor) SetBalanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "setBalanceSheet", sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_StandardToken *StandardTokenSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.SetBalanceSheet(&_StandardToken.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_StandardToken *StandardTokenTransactorSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.SetBalanceSheet(&_StandardToken.TransactOpts, sheet)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.Transfer(&_StandardToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_StandardToken *StandardTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferFrom(&_StandardToken.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_StandardToken *StandardTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StandardToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_StandardToken *StandardTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferOwnership(&_StandardToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_StandardToken *StandardTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StandardToken.Contract.TransferOwnership(&_StandardToken.TransactOpts, newOwner)
}

// StandardTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandardToken contract.
type StandardTokenApprovalIterator struct {
	Event *StandardTokenApproval // Event containing the contract specifics and raw log

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
func (it *StandardTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenApproval)
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
		it.Event = new(StandardTokenApproval)
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
func (it *StandardTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenApproval represents a Approval event raised by the StandardToken contract.
type StandardTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandardTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenApprovalIterator{contract: _StandardToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandardTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenApproval)
				if err := _StandardToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// StandardTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StandardToken contract.
type StandardTokenOwnershipTransferredIterator struct {
	Event *StandardTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StandardTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenOwnershipTransferred)
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
		it.Event = new(StandardTokenOwnershipTransferred)
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
func (it *StandardTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenOwnershipTransferred represents a OwnershipTransferred event raised by the StandardToken contract.
type StandardTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_StandardToken *StandardTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StandardTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenOwnershipTransferredIterator{contract: _StandardToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_StandardToken *StandardTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StandardTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenOwnershipTransferred)
				if err := _StandardToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// StandardTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandardToken contract.
type StandardTokenTransferIterator struct {
	Event *StandardTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StandardTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandardTokenTransfer)
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
		it.Event = new(StandardTokenTransfer)
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
func (it *StandardTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandardTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandardTokenTransfer represents a Transfer event raised by the StandardToken contract.
type StandardTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandardTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandardTokenTransferIterator{contract: _StandardToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_StandardToken *StandardTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandardTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandardToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandardTokenTransfer)
				if err := _StandardToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// TrueUSDABI is the input ABI used to generate the binding from.
const TrueUSDABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"burnMin\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"delegateAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnFeeFlat\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_canReceiveMintWhiteList\",\"type\":\"address\"},{\"name\":\"_canBurnWhiteList\",\"type\":\"address\"},{\"name\":\"_blackList\",\"type\":\"address\"},{\"name\":\"_noFeesList\",\"type\":\"address\"}],\"name\":\"setLists\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"reclaimToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"delegateToNewContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_transferFeeNumerator\",\"type\":\"uint80\"},{\"name\":\"_transferFeeDenominator\",\"type\":\"uint80\"},{\"name\":\"_mintFeeNumerator\",\"type\":\"uint80\"},{\"name\":\"_mintFeeDenominator\",\"type\":\"uint80\"},{\"name\":\"_mintFeeFlat\",\"type\":\"uint256\"},{\"name\":\"_burnFeeNumerator\",\"type\":\"uint80\"},{\"name\":\"_burnFeeDenominator\",\"type\":\"uint80\"},{\"name\":\"_burnFeeFlat\",\"type\":\"uint256\"}],\"name\":\"changeStakingFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canReceiveMintWhiteList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegatedFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateApprove\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"contractAddr\",\"type\":\"address\"}],\"name\":\"reclaimContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"delegateBalanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateTransferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setBalanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateIncreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnFeeNumerator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canBurnWhiteList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnMax\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintFeeDenominator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"staker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setDelegatedFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"noFeesList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newMin\",\"type\":\"uint256\"},{\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"changeBurnBounds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegateTotalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"changeName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintFeeNumerator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferFeeNumerator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateDecreaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"origSender\",\"type\":\"address\"}],\"name\":\"delegateTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"reclaimEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newStaker\",\"type\":\"address\"}],\"name\":\"changeStaker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"wipeBlacklistedAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from_\",\"type\":\"address\"},{\"name\":\"value_\",\"type\":\"uint256\"},{\"name\":\"data_\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnFeeDenominator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"delegate\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blackList\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferFeeDenominator\",\"outputs\":[{\"name\":\"\",\"type\":\"uint80\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"mintFeeFlat\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sheet\",\"type\":\"address\"}],\"name\":\"setAllowanceSheet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newMin\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"ChangeBurnBoundsEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"WipedAccount\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"DelegatedTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// TrueUSDBin is the compiled bytecode used for deploying new contracts.
const TrueUSDBin = `0x6005805460a060020a60ff021916905560c0604052600760808190527f547275655553440000000000000000000000000000000000000000000000000060a09081526200004e9190816200016b565b506040805180820190915260048082527f5455534400000000000000000000000000000000000000000000000000000000602090920191825262000095916008916200016b565b5069021e19e0c9bab2400000600d556a108b2a2c28029094000000600e55600f805460a060020a60f060020a03196b271000000000000000000000605060020a60a060020a03196001605060020a031993841660071716811791909116909255601080549091166127101790556000601181905560128054600160a060020a0319169092179091556013553480156200012d57600080fd5b5060008054600160a060020a0319163317905534156200014c57600080fd5b620f424060035560148054600160a060020a0319163317905562000210565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001ae57805160ff1916838001178555620001de565b82800160010185558215620001de579182015b82811115620001de578251825591602001919060010190620001c1565b50620001ec929150620001f0565b5090565b6200020d91905b80821115620001ec5760008155600101620001f7565b90565b61302580620002206000396000f3006080604052600436106102bb5763ffffffff60e060020a60003504166302d3fdc981146102ca57806306fdde03146102f1578063095ea7b31461037b57806309ab8bba146103b35780630b8e845a146103da5780630ce51179146103ef57806317ffc3201461042457806318160ddd146104455780631d2d84001461045a5780631db8cb3f1461047b5780631f7af1df146104c457806323b872dd146104f557806326fe99511461051f578063296f4000146105345780632aed7f3f1461055f578063313ce567146105805780633ed10b92146105ab5780633f4ba83a146105c057806340c10f19146105d557806342966c68146105f957806343a468c8146106115780634df6b45d146106325780634e71e0c81461066357806354f78dad14610678578063554249b31461069957806356e1c40d146106c45780635a444139146106f85780635c131d701461070d5780635c975abb146107225780635db07aee146107375780635ebaf1db1461074c57806361927adb1461076157806366188463146107825780636d4717fe146107a657806370a08231146107bb57806370df42e1146107dc57806376e71dd8146107f75780637bb98a681461080c5780638456cb591461082157806386575e40146108365780638d93eac2146108cd5780638da5cb5b146108e25780638f98ce8f146108f757806393d3173a1461090c57806395d89b41146109375780639cd1a1211461094c5780639f727c2714610977578063a9059cbb1461098c578063ab55979d146109b0578063bd7243f6146109d1578063c0ee0b8a146109f2578063c18f483114610a23578063c89e436114610a38578063cdab73b514610a4d578063d42cfc4114610a62578063d63a138914610a77578063d73dd62314610a8c578063dd62ed3e14610ab0578063e30c397814610ad7578063edc1e4f914610aec578063f2fde38b14610b0d575b3480156102c757600080fd5b50005b3480156102d657600080fd5b506102df610b2e565b60408051918252519081900360200190f35b3480156102fd57600080fd5b50610306610b34565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610340578181015183820152602001610328565b50505050905090810190601f16801561036d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561038757600080fd5b5061039f600160a060020a0360043516602435610bc2565b604080519115158252519081900360200190f35b3480156103bf57600080fd5b506102df600160a060020a0360043581169060243516610c92565b3480156103e657600080fd5b506102df610ca5565b3480156103fb57600080fd5b50610422600160a060020a0360043581169060243581169060443581169060643516610cab565b005b34801561043057600080fd5b50610422600160a060020a0360043516610d12565b34801561045157600080fd5b506102df610dc8565b34801561046657600080fd5b50610422600160a060020a0360043516610e71565b34801561048757600080fd5b5061042269ffffffffffffffffffff6004358116906024358116906044358116906064358116906084359060a43581169060c4351660e435610ed8565b3480156104d057600080fd5b506104d9610fef565b60408051600160a060020a039092168252519081900360200190f35b34801561050157600080fd5b5061039f600160a060020a0360043581169060243516604435610ffe565b34801561052b57600080fd5b506104d96110d6565b34801561054057600080fd5b5061039f600160a060020a0360043581169060243590604435166110e5565b34801561056b57600080fd5b50610422600160a060020a0360043516611116565b34801561058c57600080fd5b506105956111b0565b6040805160ff9092168252519081900360200190f35b3480156105b757600080fd5b506104d96111b5565b3480156105cc57600080fd5b506104226111c4565b3480156105e157600080fd5b50610422600160a060020a036004351660243561123c565b34801561060557600080fd5b5061042260043561142f565b34801561061d57600080fd5b506102df600160a060020a036004351661152a565b34801561063e57600080fd5b5061039f600160a060020a03600435811690602435811690604435906064351661153d565b34801561066f57600080fd5b50610422611570565b34801561068457600080fd5b50610422600160a060020a03600435166115eb565b3480156106a557600080fd5b5061039f600160a060020a036004358116906024359060443516611693565b3480156106d057600080fd5b506106d96116b9565b6040805169ffffffffffffffffffff9092168252519081900360200190f35b34801561070457600080fd5b506104d96116cb565b34801561071957600080fd5b506102df6116da565b34801561072e57600080fd5b5061039f6116e0565b34801561074357600080fd5b506106d96116f0565b34801561075857600080fd5b506104d9611702565b34801561076d57600080fd5b50610422600160a060020a0360043516611711565b34801561078e57600080fd5b5061039f600160a060020a036004351660243561174a565b3480156107b257600080fd5b506104d96117de565b3480156107c757600080fd5b506102df600160a060020a03600435166117ed565b3480156107e857600080fd5b506104226004356024356118ad565b34801561080357600080fd5b506102df61191a565b34801561081857600080fd5b506104d9611929565b34801561082d57600080fd5b50610422611938565b34801561084257600080fd5b506040805160206004803580820135601f810184900484028501840190955284845261042294369492936024939284019190819084018382808284375050604080516020601f89358b018035918201839004830284018301909452808352979a9998810197919650918201945092508291508401838280828437509497506119b59650505050505050565b3480156108d957600080fd5b506106d96119f3565b3480156108ee57600080fd5b506104d9611a0c565b34801561090357600080fd5b506106d9611a1b565b34801561091857600080fd5b5061039f600160a060020a036004358116906024359060443516611a2d565b34801561094357600080fd5b50610306611a53565b34801561095857600080fd5b5061039f600160a060020a036004358116906024359060443516611aae565b34801561098357600080fd5b50610422611ad4565b34801561099857600080fd5b5061039f600160a060020a0360043516602435611b1f565b3480156109bc57600080fd5b50610422600160a060020a0360043516611bb3565b3480156109dd57600080fd5b50610422600160a060020a0360043516611c01565b3480156109fe57600080fd5b5061042260048035600160a060020a0316906024803591604435918201910135611d8c565b348015610a2f57600080fd5b506106d9611d91565b348015610a4457600080fd5b506104d9611db1565b348015610a5957600080fd5b506104d9611dc0565b348015610a6e57600080fd5b506106d9611dcf565b348015610a8357600080fd5b506102df611def565b348015610a9857600080fd5b5061039f600160a060020a0360043516602435611df5565b348015610abc57600080fd5b506102df600160a060020a0360043581169060243516611e89565b348015610ae357600080fd5b506104d9611f18565b348015610af857600080fd5b50610422600160a060020a0360043516611f27565b348015610b1957600080fd5b50610422600160a060020a0360043516611fb0565b600d5481565b6007805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610bba5780601f10610b8f57610100808354040283529160200191610bba565b820191906000526020600020905b815481529060010190602001808311610b9d57829003601f168201915b505050505081565b600654600090600160a060020a03161515610be857610be18383611fe9565b9050610c8c565b600654604080517f296f4000000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152602482018690523360448301529151919092169163296f40009160648083019260209291908290030181600087803b158015610c5d57600080fd5b505af1158015610c71573d6000803e3d6000fd5b505050506040513d6020811015610c8757600080fd5b505190505b92915050565b6000610c9e8383611e89565b9392505050565b60135481565b600054600160a060020a03163314610cc257600080fd5b60098054600160a060020a03958616600160a060020a031991821617909155600a805494861694821694909417909355600b805492851692841692909217909155600c8054919093169116179055565b60008054600160a060020a03163314610d2a57600080fd5b6040805160e060020a6370a082310281523060048201529051600160a060020a038416916370a082319160248083019260209291908290030181600087803b158015610d7557600080fd5b505af1158015610d89573d6000803e3d6000fd5b505050506040513d6020811015610d9f57600080fd5b5051600054909150610dc490600160a060020a0384811691168363ffffffff61200d16565b5050565b600654600090600160a060020a03161515610dec57610de56120a4565b9050610e6e565b600660009054906101000a9004600160a060020a0316600160a060020a03166376e71dd86040518163ffffffff1660e060020a028152600401602060405180830381600087803b158015610e3f57600080fd5b505af1158015610e53573d6000803e3d6000fd5b505050506040513d6020811015610e6957600080fd5b505190505b90565b600054600160a060020a03163314610e8857600080fd5b60068054600160a060020a031916600160a060020a0383811691909117918290556040519116907feef3c91406f155f6bf1d8754e73003590b8bfa5cfa5472ee9ea936761864ea3090600090a250565b600054600160a060020a03163314610eef57600080fd5b69ffffffffffffffffffff87161515610f0757600080fd5b69ffffffffffffffffffff85161515610f1f57600080fd5b69ffffffffffffffffffff82161515610f3757600080fd5b600f805469ffffffffffffffffffff1990811669ffffffffffffffffffff9a8b161773ffffffffffffffffffff00000000000000000000199081166a01000000000000000000009a8c168b02177fffff00000000000000000000ffffffffffffffffffffffffffffffffffffffff1660a060020a998c169990990298909817909155601080548216968a1696909617909555601193909355601280549094169187169190911790931692909416909202179055601355565b600954600160a060020a031681565b600654600090600160a060020a031615156110255761101e8484846120aa565b9050610c9e565b600654604080517f4df6b45d000000000000000000000000000000000000000000000000000000008152600160a060020a03878116600483015286811660248301526044820186905233606483015291519190921691634df6b45d9160848083019260209291908290030181600087803b1580156110a257600080fd5b505af11580156110b6573d6000803e3d6000fd5b505050506040513d60208110156110cc57600080fd5b5051949350505050565b600554600160a060020a031681565b600554600090600160a060020a031633811461110057600080fd5b61110b8585856120d7565b506001949350505050565b60008054600160a060020a0316331461112e57600080fd5b5060008054604080517ff2fde38b000000000000000000000000000000000000000000000000000000008152600160a060020a039283166004820152905184939284169263f2fde38b926024808201939182900301818387803b15801561119457600080fd5b505af11580156111a8573d6000803e3d6000fd5b505050505050565b601281565b600454600160a060020a031681565b600054600160a060020a031633146111db57600080fd5b60055460a060020a900460ff1615156111f357600080fd5b6005805474ff0000000000000000000000000000000000000000191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b600054600160a060020a0316331461125357600080fd5b6009546040805160e060020a636f626eb3028152600160a060020a03858116600483015291519190921691636f626eb39160248083019260209291908290030181600087803b1580156112a557600080fd5b505af11580156112b9573d6000803e3d6000fd5b505050506040513d60208110156112cf57600080fd5b505115156112dc57600080fd5b6003546112ef908263ffffffff6121b016565b600355600254604080517f21e5383a000000000000000000000000000000000000000000000000000000008152600160a060020a03858116600483015260248201859052915191909216916321e5383a91604480830192600092919082900301818387803b15801561136057600080fd5b505af1158015611374573d6000803e3d6000fd5b5050604080518481529051600160a060020a03861693507f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688592509081900360200190a2604080518281529051600160a060020a038416916000917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3600f5460105460115461142a928592859260a060020a90920469ffffffffffffffffffff9081169291169060006121ca565b505050565b600a546040805160e060020a636f626eb302815233600482015290516000928392600160a060020a0390911691636f626eb39160248082019260209290919082900301818787803b15801561148357600080fd5b505af1158015611497573d6000803e3d6000fd5b505050506040513d60208110156114ad57600080fd5b505115156114ba57600080fd5b600d548310156114c957600080fd5b600e548311156114d857600080fd5b60125460135461150d913391869169ffffffffffffffffffff808216926a0100000000000000000000909204169060006121ca565b915061151f838363ffffffff61235916565b905061142a8161236b565b6000611535826117ed565b90505b919050565b600554600090600160a060020a031633811461155857600080fd5b61156486868686612514565b50600195945050505050565b600154600160a060020a0316331461158757600080fd5b60015460008054604051600160a060020a0393841693909116917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36001805460008054600160a060020a0319908116600160a060020a03841617909155169055565b600054600160a060020a0316331461160257600080fd5b60028054600160a060020a031916600160a060020a038381169190911791829055604080517f4e71e0c800000000000000000000000000000000000000000000000000000000815290519290911691634e71e0c89160048082019260009290919082900301818387803b15801561167857600080fd5b505af115801561168c573d6000803e3d6000fd5b5050505050565b600554600090600160a060020a03163381146116ae57600080fd5b61110b85858561264d565b60125469ffffffffffffffffffff1681565b600a54600160a060020a031681565b600e5481565b60055460a060020a900460ff1681565b60105469ffffffffffffffffffff1681565b601454600160a060020a031681565b600054600160a060020a0316331461172857600080fd5b60058054600160a060020a031916600160a060020a0392909216919091179055565b600654600090600160a060020a0316151561176957610be183836127a4565b600654604080517f93d3173a000000000000000000000000000000000000000000000000000000008152600160a060020a03868116600483015260248201869052336044830152915191909216916393d3173a9160648083019260209291908290030181600087803b158015610c5d57600080fd5b600c54600160a060020a031681565b600654600090600160a060020a031615156118125761180b826127c8565b9050611538565b600654604080517f43a468c8000000000000000000000000000000000000000000000000000000008152600160a060020a038581166004830152915191909216916343a468c89160248083019260209291908290030181600087803b15801561187a57600080fd5b505af115801561188e573d6000803e3d6000fd5b505050506040513d60208110156118a457600080fd5b50519050611538565b600054600160a060020a031633146118c457600080fd5b808211156118d157600080fd5b600d829055600e819055604080518381526020810183905281517ff8f7312d8aa9257dcfe43287f24cacc0f267875658809b6c7953b27756562522929181900390910190a15050565b6000611924610dc8565b905090565b600254600160a060020a031681565b600054600160a060020a0316331461194f57600080fd5b60055460a060020a900460ff161561196657600080fd5b6005805474ff0000000000000000000000000000000000000000191660a060020a1790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b600054600160a060020a031633146119cc57600080fd5b81516119df906007906020850190612f61565b50805161142a906008906020840190612f61565b600f5460a060020a900469ffffffffffffffffffff1681565b600054600160a060020a031681565b600f5469ffffffffffffffffffff1681565b600554600090600160a060020a0316338114611a4857600080fd5b61110b85858561284f565b6008805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610bba5780601f10610b8f57610100808354040283529160200191610bba565b600554600090600160a060020a0316338114611ac957600080fd5b61110b838686612ad0565b600054600160a060020a03163314611aeb57600080fd5b60008054604051600160a060020a0390911691303180156108fc02929091818181858888f193505050501515611b1d57fe5b565b600654600090600160a060020a03161515611b3e57610be18383612c1c565b600654604080517f9cd1a121000000000000000000000000000000000000000000000000000000008152600160a060020a0386811660048301526024820186905233604483015291519190921691639cd1a1219160648083019260209291908290030181600087803b158015610c5d57600080fd5b600054600160a060020a03163314611bca57600080fd5b600160a060020a0381161515611bdf57600080fd5b60148054600160a060020a031916600160a060020a0392909216919091179055565b60008054600160a060020a03163314611c1957600080fd5b600b546040805160e060020a636f626eb3028152600160a060020a03858116600483015291519190921691636f626eb39160248083019260209291908290030181600087803b158015611c6b57600080fd5b505af1158015611c7f573d6000803e3d6000fd5b505050506040513d6020811015611c9557600080fd5b50511515611ca257600080fd5b611cab826117ed565b600254604080517fe30443bc000000000000000000000000000000000000000000000000000000008152600160a060020a03868116600483015260006024830181905292519495509092169263e30443bc92604480820193929182900301818387803b158015611d1a57600080fd5b505af1158015611d2e573d6000803e3d6000fd5b5050600354611d46925090508263ffffffff61235916565b600355604080518281529051600160a060020a038416917fdf58d2368c06216a398f05a7a88c8edc64a25c33f33fd2bd8b56fbc8822c02d8919081900360200190a25050565b600080fd5b6012546a0100000000000000000000900469ffffffffffffffffffff1681565b600654600160a060020a031681565b600b54600160a060020a031681565b600f546a0100000000000000000000900469ffffffffffffffffffff1681565b60115481565b600654600090600160a060020a03161515611e1457610be18383612c40565b600654604080517f554249b3000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152602482018690523360448301529151919092169163554249b39160648083019260209291908290030181600087803b158015610c5d57600080fd5b600654600090600160a060020a03161515611ea857610be18383612c64565b600654604080517f09ab8bba000000000000000000000000000000000000000000000000000000008152600160a060020a0386811660048301528581166024830152915191909216916309ab8bba9160448083019260209291908290030181600087803b158015610c5d57600080fd5b600154600160a060020a031681565b600054600160a060020a03163314611f3e57600080fd5b60048054600160a060020a031916600160a060020a0383811691909117808355604080517f4e71e0c800000000000000000000000000000000000000000000000000000000815290519190921692634e71e0c89280820192600092909182900301818387803b15801561167857600080fd5b600054600160a060020a03163314611fc757600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b60055460009060a060020a900460ff161561200357600080fd5b610c9e8383612cc3565b82600160a060020a031663a9059cbb83836040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561207057600080fd5b505af1158015612084573d6000803e3d6000fd5b505050506040513d602081101561209a57600080fd5b5051151561142a57fe5b60035490565b60055460009060a060020a900460ff16156120c457600080fd5b6120cf848484612cd9565b949350505050565b60048054604080517fda46098c000000000000000000000000000000000000000000000000000000008152600160a060020a03858116948201949094528684166024820152604481018690529051929091169163da46098c9160648082019260009290919082900301818387803b15801561215157600080fd5b505af1158015612165573d6000803e3d6000fd5b5050604080518581529051600160a060020a038088169450851692507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360200190a3505050565b6000828201838110156121bf57fe5b8091505b5092915050565b600c546040805160e060020a636f626eb3028152600160a060020a038981166004830152915160009384931691636f626eb391602480830192602092919082900301818787803b15801561221d57600080fd5b505af1158015612231573d6000803e3d6000fd5b505050506040513d602081101561224757600080fd5b5051806122ce5750600c546040805160e060020a636f626eb3028152600160a060020a03868116600483015291519190921691636f626eb39160248083019260209291908290030181600087803b1580156122a157600080fd5b505af11580156122b5573d6000803e3d6000fd5b505050506040513d60208110156122cb57600080fd5b50515b156122dc576000915061234e565b6123278461231b8769ffffffffffffffffffff1661230f8a69ffffffffffffffffffff168c612cf190919063ffffffff16565b9063ffffffff612d1c16565b9063ffffffff6121b016565b9050600081111561234a5760145461234a908990600160a060020a031683612d33565b8091505b509695505050505050565b60008282111561236557fe5b50900390565b6002546040805160e060020a6370a082310281523360048201529051600092600160a060020a0316916370a0823191602480830192602092919082900301818787803b1580156123ba57600080fd5b505af11580156123ce573d6000803e3d6000fd5b505050506040513d60208110156123e457600080fd5b50518211156123f257600080fd5b50600254604080517fcf8eeb7e00000000000000000000000000000000000000000000000000000000815233600482018190526024820185905291519192600160a060020a03169163cf8eeb7e9160448082019260009290919082900301818387803b15801561246157600080fd5b505af1158015612475573d6000803e3d6000fd5b505060035461248d925090508363ffffffff61235916565b600355604080518381529051600160a060020a038316917fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5919081900360200190a2604080518381529051600091600160a060020a038416917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b600480546040805160e160020a630d237641028152600160a060020a0388811694820194909452848416602482015290519290911691631a46ec82916044808201926020929091908290030181600087803b15801561257257600080fd5b505af1158015612586573d6000803e3d6000fd5b505050506040513d602081101561259c57600080fd5b50518211156125aa57600080fd5b60048054604080517f97d88cd2000000000000000000000000000000000000000000000000000000008152600160a060020a0388811694820194909452848416602482015260448101869052905192909116916397d88cd29160648082019260009290919082900301818387803b15801561262457600080fd5b505af1158015612638573d6000803e3d6000fd5b50505050612647848484612ad0565b50505050565b60048054604080517f5fd72d16000000000000000000000000000000000000000000000000000000008152600160a060020a038581169482019490945286841660248201526044810186905290519290911691635fd72d169160648082019260009290919082900301818387803b1580156126c757600080fd5b505af11580156126db573d6000803e3d6000fd5b5050600480546040805160e160020a630d237641028152600160a060020a038781169482018590528981166024830181905292519296509394507f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259390921691631a46ec829160448083019260209291908290030181600087803b15801561276257600080fd5b505af1158015612776573d6000803e3d6000fd5b505050506040513d602081101561278c57600080fd5b505160408051918252519081900360200190a3505050565b60055460009060a060020a900460ff16156127be57600080fd5b610c9e8383612f3a565b6002546040805160e060020a6370a08231028152600160a060020a038481166004830152915160009392909216916370a082319160248082019260209290919082900301818787803b15801561281d57600080fd5b505af1158015612831573d6000803e3d6000fd5b505050506040513d602081101561284757600080fd5b505192915050565b600480546040805160e160020a630d237641028152600160a060020a03858116948201949094528684166024820152905160009390921691631a46ec829160448082019260209290919082900301818787803b1580156128ae57600080fd5b505af11580156128c2573d6000803e3d6000fd5b505050506040513d60208110156128d857600080fd5b50519050808311156129795760048054604080517fda46098c000000000000000000000000000000000000000000000000000000008152600160a060020a038681169482019490945287841660248201526000604482018190529151939092169263da46098c92606480820193929182900301818387803b15801561295c57600080fd5b505af1158015612970573d6000803e3d6000fd5b50505050612a0c565b60048054604080517f97d88cd2000000000000000000000000000000000000000000000000000000008152600160a060020a0386811694820194909452878416602482015260448101879052905192909116916397d88cd29160648082019260009290919082900301818387803b1580156129f357600080fd5b505af1158015612a07573d6000803e3d6000fd5b505050505b600480546040805160e160020a630d237641028152600160a060020a038681169482018590528089166024830181905292519294937f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92593911691631a46ec829160448083019260209291908290030181600087803b158015612a8d57600080fd5b505af1158015612aa1573d6000803e3d6000fd5b505050506040513d6020811015612ab757600080fd5b505160408051918252519081900360200190a350505050565b600b546040805160e060020a636f626eb3028152600160a060020a03868116600483015291519190921691636f626eb39160248083019260209291908290030181600087803b158015612b2257600080fd5b505af1158015612b36573d6000803e3d6000fd5b505050506040513d6020811015612b4c57600080fd5b505115612b5857600080fd5b600b546040805160e060020a636f626eb3028152600160a060020a03858116600483015291519190921691636f626eb39160248083019260209291908290030181600087803b158015612baa57600080fd5b505af1158015612bbe573d6000803e3d6000fd5b505050506040513d6020811015612bd457600080fd5b505115612be057600080fd5b612beb838383612d33565b600f54612647908390839069ffffffffffffffffffff808216916a01000000000000000000009004166000886121ca565b60055460009060a060020a900460ff1615612c3657600080fd5b610c9e8383612f47565b60055460009060a060020a900460ff1615612c5a57600080fd5b610c9e8383612f54565b600480546040805160e160020a630d237641028152600160a060020a03868116948201949094528484166024820152905160009390921691631a46ec829160448082019260209290919082900301818787803b158015610c5d57600080fd5b6000612cd08383336120d7565b50600192915050565b6000612ce784848433612514565b5060019392505050565b600080831515612d0457600091506121c3565b50828202828482811515612d1457fe5b04146121bf57fe5b6000808284811515612d2a57fe5b04949350505050565b600160a060020a0382161515612d4857600080fd5b600160a060020a0383161515612d5d57600080fd5b6002546040805160e060020a6370a08231028152600160a060020a038681166004830152915191909216916370a082319160248083019260209291908290030181600087803b158015612daf57600080fd5b505af1158015612dc3573d6000803e3d6000fd5b505050506040513d6020811015612dd957600080fd5b5051811115612de757600080fd5b600254604080517fcf8eeb7e000000000000000000000000000000000000000000000000000000008152600160a060020a038681166004830152602482018590529151919092169163cf8eeb7e91604480830192600092919082900301818387803b158015612e5557600080fd5b505af1158015612e69573d6000803e3d6000fd5b5050600254604080517f21e5383a000000000000000000000000000000000000000000000000000000008152600160a060020a0387811660048301526024820187905291519190921693506321e5383a9250604480830192600092919082900301818387803b158015612edb57600080fd5b505af1158015612eef573d6000803e3d6000fd5b5050604080518481529051600160a060020a038087169450871692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a3505050565b6000612cd083833361284f565b6000612cd0338484612ad0565b6000612cd083833361264d565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10612fa257805160ff1916838001178555612fcf565b82800160010185558215612fcf579182015b82811115612fcf578251825591602001919060010190612fb4565b50612fdb929150612fdf565b5090565b610e6e91905b80821115612fdb5760008155600101612fe55600a165627a7a723058208095922ec04fc21d4eb484d00e3aa382b44b1e73ca812e8689d3db89b0c250bc0029`

// DeployTrueUSD deploys a new Ethereum contract, binding an instance of TrueUSD to it.
func DeployTrueUSD(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TrueUSD, error) {
	parsed, err := abi.JSON(strings.NewReader(TrueUSDABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TrueUSDBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TrueUSD{TrueUSDCaller: TrueUSDCaller{contract: contract}, TrueUSDTransactor: TrueUSDTransactor{contract: contract}, TrueUSDFilterer: TrueUSDFilterer{contract: contract}}, nil
}

// TrueUSD is an auto generated Go binding around an Ethereum contract.
type TrueUSD struct {
	TrueUSDCaller     // Read-only binding to the contract
	TrueUSDTransactor // Write-only binding to the contract
	TrueUSDFilterer   // Log filterer for contract events
}

// TrueUSDCaller is an auto generated read-only Go binding around an Ethereum contract.
type TrueUSDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrueUSDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TrueUSDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrueUSDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TrueUSDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TrueUSDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TrueUSDSession struct {
	Contract     *TrueUSD          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TrueUSDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TrueUSDCallerSession struct {
	Contract *TrueUSDCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TrueUSDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TrueUSDTransactorSession struct {
	Contract     *TrueUSDTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TrueUSDRaw is an auto generated low-level Go binding around an Ethereum contract.
type TrueUSDRaw struct {
	Contract *TrueUSD // Generic contract binding to access the raw methods on
}

// TrueUSDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TrueUSDCallerRaw struct {
	Contract *TrueUSDCaller // Generic read-only contract binding to access the raw methods on
}

// TrueUSDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TrueUSDTransactorRaw struct {
	Contract *TrueUSDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrueUSD creates a new instance of TrueUSD, bound to a specific deployed contract.
func NewTrueUSD(address common.Address, backend bind.ContractBackend) (*TrueUSD, error) {
	contract, err := bindTrueUSD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TrueUSD{TrueUSDCaller: TrueUSDCaller{contract: contract}, TrueUSDTransactor: TrueUSDTransactor{contract: contract}, TrueUSDFilterer: TrueUSDFilterer{contract: contract}}, nil
}

// NewTrueUSDCaller creates a new read-only instance of TrueUSD, bound to a specific deployed contract.
func NewTrueUSDCaller(address common.Address, caller bind.ContractCaller) (*TrueUSDCaller, error) {
	contract, err := bindTrueUSD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TrueUSDCaller{contract: contract}, nil
}

// NewTrueUSDTransactor creates a new write-only instance of TrueUSD, bound to a specific deployed contract.
func NewTrueUSDTransactor(address common.Address, transactor bind.ContractTransactor) (*TrueUSDTransactor, error) {
	contract, err := bindTrueUSD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TrueUSDTransactor{contract: contract}, nil
}

// NewTrueUSDFilterer creates a new log filterer instance of TrueUSD, bound to a specific deployed contract.
func NewTrueUSDFilterer(address common.Address, filterer bind.ContractFilterer) (*TrueUSDFilterer, error) {
	contract, err := bindTrueUSD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TrueUSDFilterer{contract: contract}, nil
}

// bindTrueUSD binds a generic wrapper to an already deployed contract.
func bindTrueUSD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TrueUSDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrueUSD *TrueUSDRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TrueUSD.Contract.TrueUSDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrueUSD *TrueUSDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrueUSD.Contract.TrueUSDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrueUSD *TrueUSDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrueUSD.Contract.TrueUSDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TrueUSD *TrueUSDCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TrueUSD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TrueUSD *TrueUSDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrueUSD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TrueUSD *TrueUSDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TrueUSD.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, spender address) constant returns(uint256)
func (_TrueUSD *TrueUSDCaller) Allowance(opts *bind.CallOpts, _owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "allowance", _owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, spender address) constant returns(uint256)
func (_TrueUSD *TrueUSDSession) Allowance(_owner common.Address, spender common.Address) (*big.Int, error) {
	return _TrueUSD.Contract.Allowance(&_TrueUSD.CallOpts, _owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, spender address) constant returns(uint256)
func (_TrueUSD *TrueUSDCallerSession) Allowance(_owner common.Address, spender common.Address) (*big.Int, error) {
	return _TrueUSD.Contract.Allowance(&_TrueUSD.CallOpts, _owner, spender)
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_TrueUSD *TrueUSDCaller) Allowances(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "allowances")
	return *ret0, err
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_TrueUSD *TrueUSDSession) Allowances() (common.Address, error) {
	return _TrueUSD.Contract.Allowances(&_TrueUSD.CallOpts)
}

// Allowances is a free data retrieval call binding the contract method 0x3ed10b92.
//
// Solidity: function allowances() constant returns(address)
func (_TrueUSD *TrueUSDCallerSession) Allowances() (common.Address, error) {
	return _TrueUSD.Contract.Allowances(&_TrueUSD.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_TrueUSD *TrueUSDCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_TrueUSD *TrueUSDSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _TrueUSD.Contract.BalanceOf(&_TrueUSD.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_TrueUSD *TrueUSDCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _TrueUSD.Contract.BalanceOf(&_TrueUSD.CallOpts, who)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_TrueUSD *TrueUSDCaller) Balances(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "balances")
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_TrueUSD *TrueUSDSession) Balances() (common.Address, error) {
	return _TrueUSD.Contract.Balances(&_TrueUSD.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address)
func (_TrueUSD *TrueUSDCallerSession) Balances() (common.Address, error) {
	return _TrueUSD.Contract.Balances(&_TrueUSD.CallOpts)
}

// BlackList is a free data retrieval call binding the contract method 0xcdab73b5.
//
// Solidity: function blackList() constant returns(address)
func (_TrueUSD *TrueUSDCaller) BlackList(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "blackList")
	return *ret0, err
}

// BlackList is a free data retrieval call binding the contract method 0xcdab73b5.
//
// Solidity: function blackList() constant returns(address)
func (_TrueUSD *TrueUSDSession) BlackList() (common.Address, error) {
	return _TrueUSD.Contract.BlackList(&_TrueUSD.CallOpts)
}

// BlackList is a free data retrieval call binding the contract method 0xcdab73b5.
//
// Solidity: function blackList() constant returns(address)
func (_TrueUSD *TrueUSDCallerSession) BlackList() (common.Address, error) {
	return _TrueUSD.Contract.BlackList(&_TrueUSD.CallOpts)
}

// BurnFeeDenominator is a free data retrieval call binding the contract method 0xc18f4831.
//
// Solidity: function burnFeeDenominator() constant returns(uint80)
func (_TrueUSD *TrueUSDCaller) BurnFeeDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "burnFeeDenominator")
	return *ret0, err
}

// BurnFeeDenominator is a free data retrieval call binding the contract method 0xc18f4831.
//
// Solidity: function burnFeeDenominator() constant returns(uint80)
func (_TrueUSD *TrueUSDSession) BurnFeeDenominator() (*big.Int, error) {
	return _TrueUSD.Contract.BurnFeeDenominator(&_TrueUSD.CallOpts)
}

// BurnFeeDenominator is a free data retrieval call binding the contract method 0xc18f4831.
//
// Solidity: function burnFeeDenominator() constant returns(uint80)
func (_TrueUSD *TrueUSDCallerSession) BurnFeeDenominator() (*big.Int, error) {
	return _TrueUSD.Contract.BurnFeeDenominator(&_TrueUSD.CallOpts)
}

// BurnFeeFlat is a free data retrieval call binding the contract method 0x0b8e845a.
//
// Solidity: function burnFeeFlat() constant returns(uint256)
func (_TrueUSD *TrueUSDCaller) BurnFeeFlat(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "burnFeeFlat")
	return *ret0, err
}

// BurnFeeFlat is a free data retrieval call binding the contract method 0x0b8e845a.
//
// Solidity: function burnFeeFlat() constant returns(uint256)
func (_TrueUSD *TrueUSDSession) BurnFeeFlat() (*big.Int, error) {
	return _TrueUSD.Contract.BurnFeeFlat(&_TrueUSD.CallOpts)
}

// BurnFeeFlat is a free data retrieval call binding the contract method 0x0b8e845a.
//
// Solidity: function burnFeeFlat() constant returns(uint256)
func (_TrueUSD *TrueUSDCallerSession) BurnFeeFlat() (*big.Int, error) {
	return _TrueUSD.Contract.BurnFeeFlat(&_TrueUSD.CallOpts)
}

// BurnFeeNumerator is a free data retrieval call binding the contract method 0x56e1c40d.
//
// Solidity: function burnFeeNumerator() constant returns(uint80)
func (_TrueUSD *TrueUSDCaller) BurnFeeNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "burnFeeNumerator")
	return *ret0, err
}

// BurnFeeNumerator is a free data retrieval call binding the contract method 0x56e1c40d.
//
// Solidity: function burnFeeNumerator() constant returns(uint80)
func (_TrueUSD *TrueUSDSession) BurnFeeNumerator() (*big.Int, error) {
	return _TrueUSD.Contract.BurnFeeNumerator(&_TrueUSD.CallOpts)
}

// BurnFeeNumerator is a free data retrieval call binding the contract method 0x56e1c40d.
//
// Solidity: function burnFeeNumerator() constant returns(uint80)
func (_TrueUSD *TrueUSDCallerSession) BurnFeeNumerator() (*big.Int, error) {
	return _TrueUSD.Contract.BurnFeeNumerator(&_TrueUSD.CallOpts)
}

// BurnMax is a free data retrieval call binding the contract method 0x5c131d70.
//
// Solidity: function burnMax() constant returns(uint256)
func (_TrueUSD *TrueUSDCaller) BurnMax(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "burnMax")
	return *ret0, err
}

// BurnMax is a free data retrieval call binding the contract method 0x5c131d70.
//
// Solidity: function burnMax() constant returns(uint256)
func (_TrueUSD *TrueUSDSession) BurnMax() (*big.Int, error) {
	return _TrueUSD.Contract.BurnMax(&_TrueUSD.CallOpts)
}

// BurnMax is a free data retrieval call binding the contract method 0x5c131d70.
//
// Solidity: function burnMax() constant returns(uint256)
func (_TrueUSD *TrueUSDCallerSession) BurnMax() (*big.Int, error) {
	return _TrueUSD.Contract.BurnMax(&_TrueUSD.CallOpts)
}

// BurnMin is a free data retrieval call binding the contract method 0x02d3fdc9.
//
// Solidity: function burnMin() constant returns(uint256)
func (_TrueUSD *TrueUSDCaller) BurnMin(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "burnMin")
	return *ret0, err
}

// BurnMin is a free data retrieval call binding the contract method 0x02d3fdc9.
//
// Solidity: function burnMin() constant returns(uint256)
func (_TrueUSD *TrueUSDSession) BurnMin() (*big.Int, error) {
	return _TrueUSD.Contract.BurnMin(&_TrueUSD.CallOpts)
}

// BurnMin is a free data retrieval call binding the contract method 0x02d3fdc9.
//
// Solidity: function burnMin() constant returns(uint256)
func (_TrueUSD *TrueUSDCallerSession) BurnMin() (*big.Int, error) {
	return _TrueUSD.Contract.BurnMin(&_TrueUSD.CallOpts)
}

// CanBurnWhiteList is a free data retrieval call binding the contract method 0x5a444139.
//
// Solidity: function canBurnWhiteList() constant returns(address)
func (_TrueUSD *TrueUSDCaller) CanBurnWhiteList(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "canBurnWhiteList")
	return *ret0, err
}

// CanBurnWhiteList is a free data retrieval call binding the contract method 0x5a444139.
//
// Solidity: function canBurnWhiteList() constant returns(address)
func (_TrueUSD *TrueUSDSession) CanBurnWhiteList() (common.Address, error) {
	return _TrueUSD.Contract.CanBurnWhiteList(&_TrueUSD.CallOpts)
}

// CanBurnWhiteList is a free data retrieval call binding the contract method 0x5a444139.
//
// Solidity: function canBurnWhiteList() constant returns(address)
func (_TrueUSD *TrueUSDCallerSession) CanBurnWhiteList() (common.Address, error) {
	return _TrueUSD.Contract.CanBurnWhiteList(&_TrueUSD.CallOpts)
}

// CanReceiveMintWhiteList is a free data retrieval call binding the contract method 0x1f7af1df.
//
// Solidity: function canReceiveMintWhiteList() constant returns(address)
func (_TrueUSD *TrueUSDCaller) CanReceiveMintWhiteList(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "canReceiveMintWhiteList")
	return *ret0, err
}

// CanReceiveMintWhiteList is a free data retrieval call binding the contract method 0x1f7af1df.
//
// Solidity: function canReceiveMintWhiteList() constant returns(address)
func (_TrueUSD *TrueUSDSession) CanReceiveMintWhiteList() (common.Address, error) {
	return _TrueUSD.Contract.CanReceiveMintWhiteList(&_TrueUSD.CallOpts)
}

// CanReceiveMintWhiteList is a free data retrieval call binding the contract method 0x1f7af1df.
//
// Solidity: function canReceiveMintWhiteList() constant returns(address)
func (_TrueUSD *TrueUSDCallerSession) CanReceiveMintWhiteList() (common.Address, error) {
	return _TrueUSD.Contract.CanReceiveMintWhiteList(&_TrueUSD.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TrueUSD *TrueUSDCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TrueUSD *TrueUSDSession) Decimals() (uint8, error) {
	return _TrueUSD.Contract.Decimals(&_TrueUSD.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_TrueUSD *TrueUSDCallerSession) Decimals() (uint8, error) {
	return _TrueUSD.Contract.Decimals(&_TrueUSD.CallOpts)
}

// Delegate is a free data retrieval call binding the contract method 0xc89e4361.
//
// Solidity: function delegate() constant returns(address)
func (_TrueUSD *TrueUSDCaller) Delegate(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "delegate")
	return *ret0, err
}

// Delegate is a free data retrieval call binding the contract method 0xc89e4361.
//
// Solidity: function delegate() constant returns(address)
func (_TrueUSD *TrueUSDSession) Delegate() (common.Address, error) {
	return _TrueUSD.Contract.Delegate(&_TrueUSD.CallOpts)
}

// Delegate is a free data retrieval call binding the contract method 0xc89e4361.
//
// Solidity: function delegate() constant returns(address)
func (_TrueUSD *TrueUSDCallerSession) Delegate() (common.Address, error) {
	return _TrueUSD.Contract.Delegate(&_TrueUSD.CallOpts)
}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(owner address, spender address) constant returns(uint256)
func (_TrueUSD *TrueUSDCaller) DelegateAllowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "delegateAllowance", owner, spender)
	return *ret0, err
}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(owner address, spender address) constant returns(uint256)
func (_TrueUSD *TrueUSDSession) DelegateAllowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TrueUSD.Contract.DelegateAllowance(&_TrueUSD.CallOpts, owner, spender)
}

// DelegateAllowance is a free data retrieval call binding the contract method 0x09ab8bba.
//
// Solidity: function delegateAllowance(owner address, spender address) constant returns(uint256)
func (_TrueUSD *TrueUSDCallerSession) DelegateAllowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TrueUSD.Contract.DelegateAllowance(&_TrueUSD.CallOpts, owner, spender)
}

// DelegateBalanceOf is a free data retrieval call binding the contract method 0x43a468c8.
//
// Solidity: function delegateBalanceOf(who address) constant returns(uint256)
func (_TrueUSD *TrueUSDCaller) DelegateBalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "delegateBalanceOf", who)
	return *ret0, err
}

// DelegateBalanceOf is a free data retrieval call binding the contract method 0x43a468c8.
//
// Solidity: function delegateBalanceOf(who address) constant returns(uint256)
func (_TrueUSD *TrueUSDSession) DelegateBalanceOf(who common.Address) (*big.Int, error) {
	return _TrueUSD.Contract.DelegateBalanceOf(&_TrueUSD.CallOpts, who)
}

// DelegateBalanceOf is a free data retrieval call binding the contract method 0x43a468c8.
//
// Solidity: function delegateBalanceOf(who address) constant returns(uint256)
func (_TrueUSD *TrueUSDCallerSession) DelegateBalanceOf(who common.Address) (*big.Int, error) {
	return _TrueUSD.Contract.DelegateBalanceOf(&_TrueUSD.CallOpts, who)
}

// DelegateTotalSupply is a free data retrieval call binding the contract method 0x76e71dd8.
//
// Solidity: function delegateTotalSupply() constant returns(uint256)
func (_TrueUSD *TrueUSDCaller) DelegateTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "delegateTotalSupply")
	return *ret0, err
}

// DelegateTotalSupply is a free data retrieval call binding the contract method 0x76e71dd8.
//
// Solidity: function delegateTotalSupply() constant returns(uint256)
func (_TrueUSD *TrueUSDSession) DelegateTotalSupply() (*big.Int, error) {
	return _TrueUSD.Contract.DelegateTotalSupply(&_TrueUSD.CallOpts)
}

// DelegateTotalSupply is a free data retrieval call binding the contract method 0x76e71dd8.
//
// Solidity: function delegateTotalSupply() constant returns(uint256)
func (_TrueUSD *TrueUSDCallerSession) DelegateTotalSupply() (*big.Int, error) {
	return _TrueUSD.Contract.DelegateTotalSupply(&_TrueUSD.CallOpts)
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x26fe9951.
//
// Solidity: function delegatedFrom() constant returns(address)
func (_TrueUSD *TrueUSDCaller) DelegatedFrom(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "delegatedFrom")
	return *ret0, err
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x26fe9951.
//
// Solidity: function delegatedFrom() constant returns(address)
func (_TrueUSD *TrueUSDSession) DelegatedFrom() (common.Address, error) {
	return _TrueUSD.Contract.DelegatedFrom(&_TrueUSD.CallOpts)
}

// DelegatedFrom is a free data retrieval call binding the contract method 0x26fe9951.
//
// Solidity: function delegatedFrom() constant returns(address)
func (_TrueUSD *TrueUSDCallerSession) DelegatedFrom() (common.Address, error) {
	return _TrueUSD.Contract.DelegatedFrom(&_TrueUSD.CallOpts)
}

// MintFeeDenominator is a free data retrieval call binding the contract method 0x5db07aee.
//
// Solidity: function mintFeeDenominator() constant returns(uint80)
func (_TrueUSD *TrueUSDCaller) MintFeeDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "mintFeeDenominator")
	return *ret0, err
}

// MintFeeDenominator is a free data retrieval call binding the contract method 0x5db07aee.
//
// Solidity: function mintFeeDenominator() constant returns(uint80)
func (_TrueUSD *TrueUSDSession) MintFeeDenominator() (*big.Int, error) {
	return _TrueUSD.Contract.MintFeeDenominator(&_TrueUSD.CallOpts)
}

// MintFeeDenominator is a free data retrieval call binding the contract method 0x5db07aee.
//
// Solidity: function mintFeeDenominator() constant returns(uint80)
func (_TrueUSD *TrueUSDCallerSession) MintFeeDenominator() (*big.Int, error) {
	return _TrueUSD.Contract.MintFeeDenominator(&_TrueUSD.CallOpts)
}

// MintFeeFlat is a free data retrieval call binding the contract method 0xd63a1389.
//
// Solidity: function mintFeeFlat() constant returns(uint256)
func (_TrueUSD *TrueUSDCaller) MintFeeFlat(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "mintFeeFlat")
	return *ret0, err
}

// MintFeeFlat is a free data retrieval call binding the contract method 0xd63a1389.
//
// Solidity: function mintFeeFlat() constant returns(uint256)
func (_TrueUSD *TrueUSDSession) MintFeeFlat() (*big.Int, error) {
	return _TrueUSD.Contract.MintFeeFlat(&_TrueUSD.CallOpts)
}

// MintFeeFlat is a free data retrieval call binding the contract method 0xd63a1389.
//
// Solidity: function mintFeeFlat() constant returns(uint256)
func (_TrueUSD *TrueUSDCallerSession) MintFeeFlat() (*big.Int, error) {
	return _TrueUSD.Contract.MintFeeFlat(&_TrueUSD.CallOpts)
}

// MintFeeNumerator is a free data retrieval call binding the contract method 0x8d93eac2.
//
// Solidity: function mintFeeNumerator() constant returns(uint80)
func (_TrueUSD *TrueUSDCaller) MintFeeNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "mintFeeNumerator")
	return *ret0, err
}

// MintFeeNumerator is a free data retrieval call binding the contract method 0x8d93eac2.
//
// Solidity: function mintFeeNumerator() constant returns(uint80)
func (_TrueUSD *TrueUSDSession) MintFeeNumerator() (*big.Int, error) {
	return _TrueUSD.Contract.MintFeeNumerator(&_TrueUSD.CallOpts)
}

// MintFeeNumerator is a free data retrieval call binding the contract method 0x8d93eac2.
//
// Solidity: function mintFeeNumerator() constant returns(uint80)
func (_TrueUSD *TrueUSDCallerSession) MintFeeNumerator() (*big.Int, error) {
	return _TrueUSD.Contract.MintFeeNumerator(&_TrueUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TrueUSD *TrueUSDCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TrueUSD *TrueUSDSession) Name() (string, error) {
	return _TrueUSD.Contract.Name(&_TrueUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_TrueUSD *TrueUSDCallerSession) Name() (string, error) {
	return _TrueUSD.Contract.Name(&_TrueUSD.CallOpts)
}

// NoFeesList is a free data retrieval call binding the contract method 0x6d4717fe.
//
// Solidity: function noFeesList() constant returns(address)
func (_TrueUSD *TrueUSDCaller) NoFeesList(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "noFeesList")
	return *ret0, err
}

// NoFeesList is a free data retrieval call binding the contract method 0x6d4717fe.
//
// Solidity: function noFeesList() constant returns(address)
func (_TrueUSD *TrueUSDSession) NoFeesList() (common.Address, error) {
	return _TrueUSD.Contract.NoFeesList(&_TrueUSD.CallOpts)
}

// NoFeesList is a free data retrieval call binding the contract method 0x6d4717fe.
//
// Solidity: function noFeesList() constant returns(address)
func (_TrueUSD *TrueUSDCallerSession) NoFeesList() (common.Address, error) {
	return _TrueUSD.Contract.NoFeesList(&_TrueUSD.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TrueUSD *TrueUSDCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TrueUSD *TrueUSDSession) Owner() (common.Address, error) {
	return _TrueUSD.Contract.Owner(&_TrueUSD.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TrueUSD *TrueUSDCallerSession) Owner() (common.Address, error) {
	return _TrueUSD.Contract.Owner(&_TrueUSD.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_TrueUSD *TrueUSDCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_TrueUSD *TrueUSDSession) Paused() (bool, error) {
	return _TrueUSD.Contract.Paused(&_TrueUSD.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_TrueUSD *TrueUSDCallerSession) Paused() (bool, error) {
	return _TrueUSD.Contract.Paused(&_TrueUSD.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_TrueUSD *TrueUSDCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "pendingOwner")
	return *ret0, err
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_TrueUSD *TrueUSDSession) PendingOwner() (common.Address, error) {
	return _TrueUSD.Contract.PendingOwner(&_TrueUSD.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() constant returns(address)
func (_TrueUSD *TrueUSDCallerSession) PendingOwner() (common.Address, error) {
	return _TrueUSD.Contract.PendingOwner(&_TrueUSD.CallOpts)
}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() constant returns(address)
func (_TrueUSD *TrueUSDCaller) Staker(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "staker")
	return *ret0, err
}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() constant returns(address)
func (_TrueUSD *TrueUSDSession) Staker() (common.Address, error) {
	return _TrueUSD.Contract.Staker(&_TrueUSD.CallOpts)
}

// Staker is a free data retrieval call binding the contract method 0x5ebaf1db.
//
// Solidity: function staker() constant returns(address)
func (_TrueUSD *TrueUSDCallerSession) Staker() (common.Address, error) {
	return _TrueUSD.Contract.Staker(&_TrueUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TrueUSD *TrueUSDCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TrueUSD *TrueUSDSession) Symbol() (string, error) {
	return _TrueUSD.Contract.Symbol(&_TrueUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_TrueUSD *TrueUSDCallerSession) Symbol() (string, error) {
	return _TrueUSD.Contract.Symbol(&_TrueUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TrueUSD *TrueUSDCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TrueUSD *TrueUSDSession) TotalSupply() (*big.Int, error) {
	return _TrueUSD.Contract.TotalSupply(&_TrueUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TrueUSD *TrueUSDCallerSession) TotalSupply() (*big.Int, error) {
	return _TrueUSD.Contract.TotalSupply(&_TrueUSD.CallOpts)
}

// TransferFeeDenominator is a free data retrieval call binding the contract method 0xd42cfc41.
//
// Solidity: function transferFeeDenominator() constant returns(uint80)
func (_TrueUSD *TrueUSDCaller) TransferFeeDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "transferFeeDenominator")
	return *ret0, err
}

// TransferFeeDenominator is a free data retrieval call binding the contract method 0xd42cfc41.
//
// Solidity: function transferFeeDenominator() constant returns(uint80)
func (_TrueUSD *TrueUSDSession) TransferFeeDenominator() (*big.Int, error) {
	return _TrueUSD.Contract.TransferFeeDenominator(&_TrueUSD.CallOpts)
}

// TransferFeeDenominator is a free data retrieval call binding the contract method 0xd42cfc41.
//
// Solidity: function transferFeeDenominator() constant returns(uint80)
func (_TrueUSD *TrueUSDCallerSession) TransferFeeDenominator() (*big.Int, error) {
	return _TrueUSD.Contract.TransferFeeDenominator(&_TrueUSD.CallOpts)
}

// TransferFeeNumerator is a free data retrieval call binding the contract method 0x8f98ce8f.
//
// Solidity: function transferFeeNumerator() constant returns(uint80)
func (_TrueUSD *TrueUSDCaller) TransferFeeNumerator(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TrueUSD.contract.Call(opts, out, "transferFeeNumerator")
	return *ret0, err
}

// TransferFeeNumerator is a free data retrieval call binding the contract method 0x8f98ce8f.
//
// Solidity: function transferFeeNumerator() constant returns(uint80)
func (_TrueUSD *TrueUSDSession) TransferFeeNumerator() (*big.Int, error) {
	return _TrueUSD.Contract.TransferFeeNumerator(&_TrueUSD.CallOpts)
}

// TransferFeeNumerator is a free data retrieval call binding the contract method 0x8f98ce8f.
//
// Solidity: function transferFeeNumerator() constant returns(uint80)
func (_TrueUSD *TrueUSDCallerSession) TransferFeeNumerator() (*big.Int, error) {
	return _TrueUSD.Contract.TransferFeeNumerator(&_TrueUSD.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_TrueUSD *TrueUSDTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_TrueUSD *TrueUSDSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.Approve(&_TrueUSD.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_TrueUSD *TrueUSDTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.Approve(&_TrueUSD.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_TrueUSD *TrueUSDTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_TrueUSD *TrueUSDSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.Burn(&_TrueUSD.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(_value uint256) returns()
func (_TrueUSD *TrueUSDTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.Burn(&_TrueUSD.TransactOpts, _value)
}

// ChangeBurnBounds is a paid mutator transaction binding the contract method 0x70df42e1.
//
// Solidity: function changeBurnBounds(newMin uint256, newMax uint256) returns()
func (_TrueUSD *TrueUSDTransactor) ChangeBurnBounds(opts *bind.TransactOpts, newMin *big.Int, newMax *big.Int) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "changeBurnBounds", newMin, newMax)
}

// ChangeBurnBounds is a paid mutator transaction binding the contract method 0x70df42e1.
//
// Solidity: function changeBurnBounds(newMin uint256, newMax uint256) returns()
func (_TrueUSD *TrueUSDSession) ChangeBurnBounds(newMin *big.Int, newMax *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.ChangeBurnBounds(&_TrueUSD.TransactOpts, newMin, newMax)
}

// ChangeBurnBounds is a paid mutator transaction binding the contract method 0x70df42e1.
//
// Solidity: function changeBurnBounds(newMin uint256, newMax uint256) returns()
func (_TrueUSD *TrueUSDTransactorSession) ChangeBurnBounds(newMin *big.Int, newMax *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.ChangeBurnBounds(&_TrueUSD.TransactOpts, newMin, newMax)
}

// ChangeName is a paid mutator transaction binding the contract method 0x86575e40.
//
// Solidity: function changeName(_name string, _symbol string) returns()
func (_TrueUSD *TrueUSDTransactor) ChangeName(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "changeName", _name, _symbol)
}

// ChangeName is a paid mutator transaction binding the contract method 0x86575e40.
//
// Solidity: function changeName(_name string, _symbol string) returns()
func (_TrueUSD *TrueUSDSession) ChangeName(_name string, _symbol string) (*types.Transaction, error) {
	return _TrueUSD.Contract.ChangeName(&_TrueUSD.TransactOpts, _name, _symbol)
}

// ChangeName is a paid mutator transaction binding the contract method 0x86575e40.
//
// Solidity: function changeName(_name string, _symbol string) returns()
func (_TrueUSD *TrueUSDTransactorSession) ChangeName(_name string, _symbol string) (*types.Transaction, error) {
	return _TrueUSD.Contract.ChangeName(&_TrueUSD.TransactOpts, _name, _symbol)
}

// ChangeStaker is a paid mutator transaction binding the contract method 0xab55979d.
//
// Solidity: function changeStaker(newStaker address) returns()
func (_TrueUSD *TrueUSDTransactor) ChangeStaker(opts *bind.TransactOpts, newStaker common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "changeStaker", newStaker)
}

// ChangeStaker is a paid mutator transaction binding the contract method 0xab55979d.
//
// Solidity: function changeStaker(newStaker address) returns()
func (_TrueUSD *TrueUSDSession) ChangeStaker(newStaker common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.ChangeStaker(&_TrueUSD.TransactOpts, newStaker)
}

// ChangeStaker is a paid mutator transaction binding the contract method 0xab55979d.
//
// Solidity: function changeStaker(newStaker address) returns()
func (_TrueUSD *TrueUSDTransactorSession) ChangeStaker(newStaker common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.ChangeStaker(&_TrueUSD.TransactOpts, newStaker)
}

// ChangeStakingFees is a paid mutator transaction binding the contract method 0x1db8cb3f.
//
// Solidity: function changeStakingFees(_transferFeeNumerator uint80, _transferFeeDenominator uint80, _mintFeeNumerator uint80, _mintFeeDenominator uint80, _mintFeeFlat uint256, _burnFeeNumerator uint80, _burnFeeDenominator uint80, _burnFeeFlat uint256) returns()
func (_TrueUSD *TrueUSDTransactor) ChangeStakingFees(opts *bind.TransactOpts, _transferFeeNumerator *big.Int, _transferFeeDenominator *big.Int, _mintFeeNumerator *big.Int, _mintFeeDenominator *big.Int, _mintFeeFlat *big.Int, _burnFeeNumerator *big.Int, _burnFeeDenominator *big.Int, _burnFeeFlat *big.Int) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "changeStakingFees", _transferFeeNumerator, _transferFeeDenominator, _mintFeeNumerator, _mintFeeDenominator, _mintFeeFlat, _burnFeeNumerator, _burnFeeDenominator, _burnFeeFlat)
}

// ChangeStakingFees is a paid mutator transaction binding the contract method 0x1db8cb3f.
//
// Solidity: function changeStakingFees(_transferFeeNumerator uint80, _transferFeeDenominator uint80, _mintFeeNumerator uint80, _mintFeeDenominator uint80, _mintFeeFlat uint256, _burnFeeNumerator uint80, _burnFeeDenominator uint80, _burnFeeFlat uint256) returns()
func (_TrueUSD *TrueUSDSession) ChangeStakingFees(_transferFeeNumerator *big.Int, _transferFeeDenominator *big.Int, _mintFeeNumerator *big.Int, _mintFeeDenominator *big.Int, _mintFeeFlat *big.Int, _burnFeeNumerator *big.Int, _burnFeeDenominator *big.Int, _burnFeeFlat *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.ChangeStakingFees(&_TrueUSD.TransactOpts, _transferFeeNumerator, _transferFeeDenominator, _mintFeeNumerator, _mintFeeDenominator, _mintFeeFlat, _burnFeeNumerator, _burnFeeDenominator, _burnFeeFlat)
}

// ChangeStakingFees is a paid mutator transaction binding the contract method 0x1db8cb3f.
//
// Solidity: function changeStakingFees(_transferFeeNumerator uint80, _transferFeeDenominator uint80, _mintFeeNumerator uint80, _mintFeeDenominator uint80, _mintFeeFlat uint256, _burnFeeNumerator uint80, _burnFeeDenominator uint80, _burnFeeFlat uint256) returns()
func (_TrueUSD *TrueUSDTransactorSession) ChangeStakingFees(_transferFeeNumerator *big.Int, _transferFeeDenominator *big.Int, _mintFeeNumerator *big.Int, _mintFeeDenominator *big.Int, _mintFeeFlat *big.Int, _burnFeeNumerator *big.Int, _burnFeeDenominator *big.Int, _burnFeeFlat *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.ChangeStakingFees(&_TrueUSD.TransactOpts, _transferFeeNumerator, _transferFeeDenominator, _mintFeeNumerator, _mintFeeDenominator, _mintFeeFlat, _burnFeeNumerator, _burnFeeDenominator, _burnFeeFlat)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_TrueUSD *TrueUSDTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_TrueUSD *TrueUSDSession) ClaimOwnership() (*types.Transaction, error) {
	return _TrueUSD.Contract.ClaimOwnership(&_TrueUSD.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_TrueUSD *TrueUSDTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _TrueUSD.Contract.ClaimOwnership(&_TrueUSD.TransactOpts)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, subtractedValue uint256) returns(bool)
func (_TrueUSD *TrueUSDTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "decreaseApproval", spender, subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, subtractedValue uint256) returns(bool)
func (_TrueUSD *TrueUSDSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.DecreaseApproval(&_TrueUSD.TransactOpts, spender, subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(spender address, subtractedValue uint256) returns(bool)
func (_TrueUSD *TrueUSDTransactorSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.DecreaseApproval(&_TrueUSD.TransactOpts, spender, subtractedValue)
}

// DelegateApprove is a paid mutator transaction binding the contract method 0x296f4000.
//
// Solidity: function delegateApprove(spender address, value uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDTransactor) DelegateApprove(opts *bind.TransactOpts, spender common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "delegateApprove", spender, value, origSender)
}

// DelegateApprove is a paid mutator transaction binding the contract method 0x296f4000.
//
// Solidity: function delegateApprove(spender address, value uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDSession) DelegateApprove(spender common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.DelegateApprove(&_TrueUSD.TransactOpts, spender, value, origSender)
}

// DelegateApprove is a paid mutator transaction binding the contract method 0x296f4000.
//
// Solidity: function delegateApprove(spender address, value uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDTransactorSession) DelegateApprove(spender common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.DelegateApprove(&_TrueUSD.TransactOpts, spender, value, origSender)
}

// DelegateDecreaseApproval is a paid mutator transaction binding the contract method 0x93d3173a.
//
// Solidity: function delegateDecreaseApproval(spender address, subtractedValue uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDTransactor) DelegateDecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "delegateDecreaseApproval", spender, subtractedValue, origSender)
}

// DelegateDecreaseApproval is a paid mutator transaction binding the contract method 0x93d3173a.
//
// Solidity: function delegateDecreaseApproval(spender address, subtractedValue uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDSession) DelegateDecreaseApproval(spender common.Address, subtractedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.DelegateDecreaseApproval(&_TrueUSD.TransactOpts, spender, subtractedValue, origSender)
}

// DelegateDecreaseApproval is a paid mutator transaction binding the contract method 0x93d3173a.
//
// Solidity: function delegateDecreaseApproval(spender address, subtractedValue uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDTransactorSession) DelegateDecreaseApproval(spender common.Address, subtractedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.DelegateDecreaseApproval(&_TrueUSD.TransactOpts, spender, subtractedValue, origSender)
}

// DelegateIncreaseApproval is a paid mutator transaction binding the contract method 0x554249b3.
//
// Solidity: function delegateIncreaseApproval(spender address, addedValue uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDTransactor) DelegateIncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "delegateIncreaseApproval", spender, addedValue, origSender)
}

// DelegateIncreaseApproval is a paid mutator transaction binding the contract method 0x554249b3.
//
// Solidity: function delegateIncreaseApproval(spender address, addedValue uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDSession) DelegateIncreaseApproval(spender common.Address, addedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.DelegateIncreaseApproval(&_TrueUSD.TransactOpts, spender, addedValue, origSender)
}

// DelegateIncreaseApproval is a paid mutator transaction binding the contract method 0x554249b3.
//
// Solidity: function delegateIncreaseApproval(spender address, addedValue uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDTransactorSession) DelegateIncreaseApproval(spender common.Address, addedValue *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.DelegateIncreaseApproval(&_TrueUSD.TransactOpts, spender, addedValue, origSender)
}

// DelegateToNewContract is a paid mutator transaction binding the contract method 0x1d2d8400.
//
// Solidity: function delegateToNewContract(newContract address) returns()
func (_TrueUSD *TrueUSDTransactor) DelegateToNewContract(opts *bind.TransactOpts, newContract common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "delegateToNewContract", newContract)
}

// DelegateToNewContract is a paid mutator transaction binding the contract method 0x1d2d8400.
//
// Solidity: function delegateToNewContract(newContract address) returns()
func (_TrueUSD *TrueUSDSession) DelegateToNewContract(newContract common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.DelegateToNewContract(&_TrueUSD.TransactOpts, newContract)
}

// DelegateToNewContract is a paid mutator transaction binding the contract method 0x1d2d8400.
//
// Solidity: function delegateToNewContract(newContract address) returns()
func (_TrueUSD *TrueUSDTransactorSession) DelegateToNewContract(newContract common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.DelegateToNewContract(&_TrueUSD.TransactOpts, newContract)
}

// DelegateTransfer is a paid mutator transaction binding the contract method 0x9cd1a121.
//
// Solidity: function delegateTransfer(to address, value uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDTransactor) DelegateTransfer(opts *bind.TransactOpts, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "delegateTransfer", to, value, origSender)
}

// DelegateTransfer is a paid mutator transaction binding the contract method 0x9cd1a121.
//
// Solidity: function delegateTransfer(to address, value uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDSession) DelegateTransfer(to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.DelegateTransfer(&_TrueUSD.TransactOpts, to, value, origSender)
}

// DelegateTransfer is a paid mutator transaction binding the contract method 0x9cd1a121.
//
// Solidity: function delegateTransfer(to address, value uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDTransactorSession) DelegateTransfer(to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.DelegateTransfer(&_TrueUSD.TransactOpts, to, value, origSender)
}

// DelegateTransferFrom is a paid mutator transaction binding the contract method 0x4df6b45d.
//
// Solidity: function delegateTransferFrom(from address, to address, value uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDTransactor) DelegateTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "delegateTransferFrom", from, to, value, origSender)
}

// DelegateTransferFrom is a paid mutator transaction binding the contract method 0x4df6b45d.
//
// Solidity: function delegateTransferFrom(from address, to address, value uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDSession) DelegateTransferFrom(from common.Address, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.DelegateTransferFrom(&_TrueUSD.TransactOpts, from, to, value, origSender)
}

// DelegateTransferFrom is a paid mutator transaction binding the contract method 0x4df6b45d.
//
// Solidity: function delegateTransferFrom(from address, to address, value uint256, origSender address) returns(bool)
func (_TrueUSD *TrueUSDTransactorSession) DelegateTransferFrom(from common.Address, to common.Address, value *big.Int, origSender common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.DelegateTransferFrom(&_TrueUSD.TransactOpts, from, to, value, origSender)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, addedValue uint256) returns(bool)
func (_TrueUSD *TrueUSDTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "increaseApproval", spender, addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, addedValue uint256) returns(bool)
func (_TrueUSD *TrueUSDSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.IncreaseApproval(&_TrueUSD.TransactOpts, spender, addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(spender address, addedValue uint256) returns(bool)
func (_TrueUSD *TrueUSDTransactorSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.IncreaseApproval(&_TrueUSD.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns()
func (_TrueUSD *TrueUSDTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns()
func (_TrueUSD *TrueUSDSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.Mint(&_TrueUSD.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns()
func (_TrueUSD *TrueUSDTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.Mint(&_TrueUSD.TransactOpts, _to, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TrueUSD *TrueUSDTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TrueUSD *TrueUSDSession) Pause() (*types.Transaction, error) {
	return _TrueUSD.Contract.Pause(&_TrueUSD.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TrueUSD *TrueUSDTransactorSession) Pause() (*types.Transaction, error) {
	return _TrueUSD.Contract.Pause(&_TrueUSD.TransactOpts)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_TrueUSD *TrueUSDTransactor) ReclaimContract(opts *bind.TransactOpts, contractAddr common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "reclaimContract", contractAddr)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_TrueUSD *TrueUSDSession) ReclaimContract(contractAddr common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.ReclaimContract(&_TrueUSD.TransactOpts, contractAddr)
}

// ReclaimContract is a paid mutator transaction binding the contract method 0x2aed7f3f.
//
// Solidity: function reclaimContract(contractAddr address) returns()
func (_TrueUSD *TrueUSDTransactorSession) ReclaimContract(contractAddr common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.ReclaimContract(&_TrueUSD.TransactOpts, contractAddr)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_TrueUSD *TrueUSDTransactor) ReclaimEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "reclaimEther")
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_TrueUSD *TrueUSDSession) ReclaimEther() (*types.Transaction, error) {
	return _TrueUSD.Contract.ReclaimEther(&_TrueUSD.TransactOpts)
}

// ReclaimEther is a paid mutator transaction binding the contract method 0x9f727c27.
//
// Solidity: function reclaimEther() returns()
func (_TrueUSD *TrueUSDTransactorSession) ReclaimEther() (*types.Transaction, error) {
	return _TrueUSD.Contract.ReclaimEther(&_TrueUSD.TransactOpts)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_TrueUSD *TrueUSDTransactor) ReclaimToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "reclaimToken", token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_TrueUSD *TrueUSDSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.ReclaimToken(&_TrueUSD.TransactOpts, token)
}

// ReclaimToken is a paid mutator transaction binding the contract method 0x17ffc320.
//
// Solidity: function reclaimToken(token address) returns()
func (_TrueUSD *TrueUSDTransactorSession) ReclaimToken(token common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.ReclaimToken(&_TrueUSD.TransactOpts, token)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_TrueUSD *TrueUSDTransactor) SetAllowanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "setAllowanceSheet", sheet)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_TrueUSD *TrueUSDSession) SetAllowanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.SetAllowanceSheet(&_TrueUSD.TransactOpts, sheet)
}

// SetAllowanceSheet is a paid mutator transaction binding the contract method 0xedc1e4f9.
//
// Solidity: function setAllowanceSheet(sheet address) returns()
func (_TrueUSD *TrueUSDTransactorSession) SetAllowanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.SetAllowanceSheet(&_TrueUSD.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_TrueUSD *TrueUSDTransactor) SetBalanceSheet(opts *bind.TransactOpts, sheet common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "setBalanceSheet", sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_TrueUSD *TrueUSDSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.SetBalanceSheet(&_TrueUSD.TransactOpts, sheet)
}

// SetBalanceSheet is a paid mutator transaction binding the contract method 0x54f78dad.
//
// Solidity: function setBalanceSheet(sheet address) returns()
func (_TrueUSD *TrueUSDTransactorSession) SetBalanceSheet(sheet common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.SetBalanceSheet(&_TrueUSD.TransactOpts, sheet)
}

// SetDelegatedFrom is a paid mutator transaction binding the contract method 0x61927adb.
//
// Solidity: function setDelegatedFrom(addr address) returns()
func (_TrueUSD *TrueUSDTransactor) SetDelegatedFrom(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "setDelegatedFrom", addr)
}

// SetDelegatedFrom is a paid mutator transaction binding the contract method 0x61927adb.
//
// Solidity: function setDelegatedFrom(addr address) returns()
func (_TrueUSD *TrueUSDSession) SetDelegatedFrom(addr common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.SetDelegatedFrom(&_TrueUSD.TransactOpts, addr)
}

// SetDelegatedFrom is a paid mutator transaction binding the contract method 0x61927adb.
//
// Solidity: function setDelegatedFrom(addr address) returns()
func (_TrueUSD *TrueUSDTransactorSession) SetDelegatedFrom(addr common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.SetDelegatedFrom(&_TrueUSD.TransactOpts, addr)
}

// SetLists is a paid mutator transaction binding the contract method 0x0ce51179.
//
// Solidity: function setLists(_canReceiveMintWhiteList address, _canBurnWhiteList address, _blackList address, _noFeesList address) returns()
func (_TrueUSD *TrueUSDTransactor) SetLists(opts *bind.TransactOpts, _canReceiveMintWhiteList common.Address, _canBurnWhiteList common.Address, _blackList common.Address, _noFeesList common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "setLists", _canReceiveMintWhiteList, _canBurnWhiteList, _blackList, _noFeesList)
}

// SetLists is a paid mutator transaction binding the contract method 0x0ce51179.
//
// Solidity: function setLists(_canReceiveMintWhiteList address, _canBurnWhiteList address, _blackList address, _noFeesList address) returns()
func (_TrueUSD *TrueUSDSession) SetLists(_canReceiveMintWhiteList common.Address, _canBurnWhiteList common.Address, _blackList common.Address, _noFeesList common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.SetLists(&_TrueUSD.TransactOpts, _canReceiveMintWhiteList, _canBurnWhiteList, _blackList, _noFeesList)
}

// SetLists is a paid mutator transaction binding the contract method 0x0ce51179.
//
// Solidity: function setLists(_canReceiveMintWhiteList address, _canBurnWhiteList address, _blackList address, _noFeesList address) returns()
func (_TrueUSD *TrueUSDTransactorSession) SetLists(_canReceiveMintWhiteList common.Address, _canBurnWhiteList common.Address, _blackList common.Address, _noFeesList common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.SetLists(&_TrueUSD.TransactOpts, _canReceiveMintWhiteList, _canBurnWhiteList, _blackList, _noFeesList)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_TrueUSD *TrueUSDTransactor) TokenFallback(opts *bind.TransactOpts, from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "tokenFallback", from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_TrueUSD *TrueUSDSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _TrueUSD.Contract.TokenFallback(&_TrueUSD.TransactOpts, from_, value_, data_)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(from_ address, value_ uint256, data_ bytes) returns()
func (_TrueUSD *TrueUSDTransactorSession) TokenFallback(from_ common.Address, value_ *big.Int, data_ []byte) (*types.Transaction, error) {
	return _TrueUSD.Contract.TokenFallback(&_TrueUSD.TransactOpts, from_, value_, data_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_TrueUSD *TrueUSDTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_TrueUSD *TrueUSDSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.Transfer(&_TrueUSD.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_TrueUSD *TrueUSDTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.Transfer(&_TrueUSD.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_TrueUSD *TrueUSDTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_TrueUSD *TrueUSDSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.TransferFrom(&_TrueUSD.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_TrueUSD *TrueUSDTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _TrueUSD.Contract.TransferFrom(&_TrueUSD.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TrueUSD *TrueUSDTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TrueUSD *TrueUSDSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.TransferOwnership(&_TrueUSD.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TrueUSD *TrueUSDTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.TransferOwnership(&_TrueUSD.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TrueUSD *TrueUSDTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TrueUSD *TrueUSDSession) Unpause() (*types.Transaction, error) {
	return _TrueUSD.Contract.Unpause(&_TrueUSD.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TrueUSD *TrueUSDTransactorSession) Unpause() (*types.Transaction, error) {
	return _TrueUSD.Contract.Unpause(&_TrueUSD.TransactOpts)
}

// WipeBlacklistedAccount is a paid mutator transaction binding the contract method 0xbd7243f6.
//
// Solidity: function wipeBlacklistedAccount(account address) returns()
func (_TrueUSD *TrueUSDTransactor) WipeBlacklistedAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TrueUSD.contract.Transact(opts, "wipeBlacklistedAccount", account)
}

// WipeBlacklistedAccount is a paid mutator transaction binding the contract method 0xbd7243f6.
//
// Solidity: function wipeBlacklistedAccount(account address) returns()
func (_TrueUSD *TrueUSDSession) WipeBlacklistedAccount(account common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.WipeBlacklistedAccount(&_TrueUSD.TransactOpts, account)
}

// WipeBlacklistedAccount is a paid mutator transaction binding the contract method 0xbd7243f6.
//
// Solidity: function wipeBlacklistedAccount(account address) returns()
func (_TrueUSD *TrueUSDTransactorSession) WipeBlacklistedAccount(account common.Address) (*types.Transaction, error) {
	return _TrueUSD.Contract.WipeBlacklistedAccount(&_TrueUSD.TransactOpts, account)
}

// TrueUSDApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TrueUSD contract.
type TrueUSDApprovalIterator struct {
	Event *TrueUSDApproval // Event containing the contract specifics and raw log

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
func (it *TrueUSDApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrueUSDApproval)
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
		it.Event = new(TrueUSDApproval)
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
func (it *TrueUSDApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrueUSDApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrueUSDApproval represents a Approval event raised by the TrueUSD contract.
type TrueUSDApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_TrueUSD *TrueUSDFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TrueUSDApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TrueUSD.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TrueUSDApprovalIterator{contract: _TrueUSD.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_TrueUSD *TrueUSDFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TrueUSDApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TrueUSD.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrueUSDApproval)
				if err := _TrueUSD.contract.UnpackLog(event, "Approval", log); err != nil {
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

// TrueUSDBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the TrueUSD contract.
type TrueUSDBurnIterator struct {
	Event *TrueUSDBurn // Event containing the contract specifics and raw log

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
func (it *TrueUSDBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrueUSDBurn)
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
		it.Event = new(TrueUSDBurn)
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
func (it *TrueUSDBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrueUSDBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrueUSDBurn represents a Burn event raised by the TrueUSD contract.
type TrueUSDBurn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_TrueUSD *TrueUSDFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*TrueUSDBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _TrueUSD.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &TrueUSDBurnIterator{contract: _TrueUSD.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: e Burn(burner indexed address, value uint256)
func (_TrueUSD *TrueUSDFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *TrueUSDBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _TrueUSD.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrueUSDBurn)
				if err := _TrueUSD.contract.UnpackLog(event, "Burn", log); err != nil {
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

// TrueUSDChangeBurnBoundsEventIterator is returned from FilterChangeBurnBoundsEvent and is used to iterate over the raw logs and unpacked data for ChangeBurnBoundsEvent events raised by the TrueUSD contract.
type TrueUSDChangeBurnBoundsEventIterator struct {
	Event *TrueUSDChangeBurnBoundsEvent // Event containing the contract specifics and raw log

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
func (it *TrueUSDChangeBurnBoundsEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrueUSDChangeBurnBoundsEvent)
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
		it.Event = new(TrueUSDChangeBurnBoundsEvent)
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
func (it *TrueUSDChangeBurnBoundsEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrueUSDChangeBurnBoundsEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrueUSDChangeBurnBoundsEvent represents a ChangeBurnBoundsEvent event raised by the TrueUSD contract.
type TrueUSDChangeBurnBoundsEvent struct {
	NewMin *big.Int
	NewMax *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChangeBurnBoundsEvent is a free log retrieval operation binding the contract event 0xf8f7312d8aa9257dcfe43287f24cacc0f267875658809b6c7953b27756562522.
//
// Solidity: e ChangeBurnBoundsEvent(newMin uint256, newMax uint256)
func (_TrueUSD *TrueUSDFilterer) FilterChangeBurnBoundsEvent(opts *bind.FilterOpts) (*TrueUSDChangeBurnBoundsEventIterator, error) {

	logs, sub, err := _TrueUSD.contract.FilterLogs(opts, "ChangeBurnBoundsEvent")
	if err != nil {
		return nil, err
	}
	return &TrueUSDChangeBurnBoundsEventIterator{contract: _TrueUSD.contract, event: "ChangeBurnBoundsEvent", logs: logs, sub: sub}, nil
}

// WatchChangeBurnBoundsEvent is a free log subscription operation binding the contract event 0xf8f7312d8aa9257dcfe43287f24cacc0f267875658809b6c7953b27756562522.
//
// Solidity: e ChangeBurnBoundsEvent(newMin uint256, newMax uint256)
func (_TrueUSD *TrueUSDFilterer) WatchChangeBurnBoundsEvent(opts *bind.WatchOpts, sink chan<- *TrueUSDChangeBurnBoundsEvent) (event.Subscription, error) {

	logs, sub, err := _TrueUSD.contract.WatchLogs(opts, "ChangeBurnBoundsEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrueUSDChangeBurnBoundsEvent)
				if err := _TrueUSD.contract.UnpackLog(event, "ChangeBurnBoundsEvent", log); err != nil {
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

// TrueUSDDelegatedToIterator is returned from FilterDelegatedTo and is used to iterate over the raw logs and unpacked data for DelegatedTo events raised by the TrueUSD contract.
type TrueUSDDelegatedToIterator struct {
	Event *TrueUSDDelegatedTo // Event containing the contract specifics and raw log

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
func (it *TrueUSDDelegatedToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrueUSDDelegatedTo)
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
		it.Event = new(TrueUSDDelegatedTo)
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
func (it *TrueUSDDelegatedToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrueUSDDelegatedToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrueUSDDelegatedTo represents a DelegatedTo event raised by the TrueUSD contract.
type TrueUSDDelegatedTo struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDelegatedTo is a free log retrieval operation binding the contract event 0xeef3c91406f155f6bf1d8754e73003590b8bfa5cfa5472ee9ea936761864ea30.
//
// Solidity: e DelegatedTo(newContract indexed address)
func (_TrueUSD *TrueUSDFilterer) FilterDelegatedTo(opts *bind.FilterOpts, newContract []common.Address) (*TrueUSDDelegatedToIterator, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _TrueUSD.contract.FilterLogs(opts, "DelegatedTo", newContractRule)
	if err != nil {
		return nil, err
	}
	return &TrueUSDDelegatedToIterator{contract: _TrueUSD.contract, event: "DelegatedTo", logs: logs, sub: sub}, nil
}

// WatchDelegatedTo is a free log subscription operation binding the contract event 0xeef3c91406f155f6bf1d8754e73003590b8bfa5cfa5472ee9ea936761864ea30.
//
// Solidity: e DelegatedTo(newContract indexed address)
func (_TrueUSD *TrueUSDFilterer) WatchDelegatedTo(opts *bind.WatchOpts, sink chan<- *TrueUSDDelegatedTo, newContract []common.Address) (event.Subscription, error) {

	var newContractRule []interface{}
	for _, newContractItem := range newContract {
		newContractRule = append(newContractRule, newContractItem)
	}

	logs, sub, err := _TrueUSD.contract.WatchLogs(opts, "DelegatedTo", newContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrueUSDDelegatedTo)
				if err := _TrueUSD.contract.UnpackLog(event, "DelegatedTo", log); err != nil {
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

// TrueUSDMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the TrueUSD contract.
type TrueUSDMintIterator struct {
	Event *TrueUSDMint // Event containing the contract specifics and raw log

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
func (it *TrueUSDMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrueUSDMint)
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
		it.Event = new(TrueUSDMint)
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
func (it *TrueUSDMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrueUSDMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrueUSDMint represents a Mint event raised by the TrueUSD contract.
type TrueUSDMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_TrueUSD *TrueUSDFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*TrueUSDMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TrueUSD.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &TrueUSDMintIterator{contract: _TrueUSD.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_TrueUSD *TrueUSDFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *TrueUSDMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TrueUSD.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrueUSDMint)
				if err := _TrueUSD.contract.UnpackLog(event, "Mint", log); err != nil {
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

// TrueUSDOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TrueUSD contract.
type TrueUSDOwnershipTransferredIterator struct {
	Event *TrueUSDOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TrueUSDOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrueUSDOwnershipTransferred)
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
		it.Event = new(TrueUSDOwnershipTransferred)
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
func (it *TrueUSDOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrueUSDOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrueUSDOwnershipTransferred represents a OwnershipTransferred event raised by the TrueUSD contract.
type TrueUSDOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_TrueUSD *TrueUSDFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TrueUSDOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TrueUSD.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TrueUSDOwnershipTransferredIterator{contract: _TrueUSD.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_TrueUSD *TrueUSDFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TrueUSDOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TrueUSD.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrueUSDOwnershipTransferred)
				if err := _TrueUSD.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// TrueUSDPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the TrueUSD contract.
type TrueUSDPauseIterator struct {
	Event *TrueUSDPause // Event containing the contract specifics and raw log

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
func (it *TrueUSDPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrueUSDPause)
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
		it.Event = new(TrueUSDPause)
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
func (it *TrueUSDPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrueUSDPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrueUSDPause represents a Pause event raised by the TrueUSD contract.
type TrueUSDPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_TrueUSD *TrueUSDFilterer) FilterPause(opts *bind.FilterOpts) (*TrueUSDPauseIterator, error) {

	logs, sub, err := _TrueUSD.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &TrueUSDPauseIterator{contract: _TrueUSD.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_TrueUSD *TrueUSDFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *TrueUSDPause) (event.Subscription, error) {

	logs, sub, err := _TrueUSD.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrueUSDPause)
				if err := _TrueUSD.contract.UnpackLog(event, "Pause", log); err != nil {
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

// TrueUSDTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TrueUSD contract.
type TrueUSDTransferIterator struct {
	Event *TrueUSDTransfer // Event containing the contract specifics and raw log

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
func (it *TrueUSDTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrueUSDTransfer)
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
		it.Event = new(TrueUSDTransfer)
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
func (it *TrueUSDTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrueUSDTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrueUSDTransfer represents a Transfer event raised by the TrueUSD contract.
type TrueUSDTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_TrueUSD *TrueUSDFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TrueUSDTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TrueUSD.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TrueUSDTransferIterator{contract: _TrueUSD.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_TrueUSD *TrueUSDFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TrueUSDTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TrueUSD.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrueUSDTransfer)
				if err := _TrueUSD.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// TrueUSDUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the TrueUSD contract.
type TrueUSDUnpauseIterator struct {
	Event *TrueUSDUnpause // Event containing the contract specifics and raw log

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
func (it *TrueUSDUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrueUSDUnpause)
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
		it.Event = new(TrueUSDUnpause)
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
func (it *TrueUSDUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrueUSDUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrueUSDUnpause represents a Unpause event raised by the TrueUSD contract.
type TrueUSDUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_TrueUSD *TrueUSDFilterer) FilterUnpause(opts *bind.FilterOpts) (*TrueUSDUnpauseIterator, error) {

	logs, sub, err := _TrueUSD.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &TrueUSDUnpauseIterator{contract: _TrueUSD.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_TrueUSD *TrueUSDFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *TrueUSDUnpause) (event.Subscription, error) {

	logs, sub, err := _TrueUSD.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrueUSDUnpause)
				if err := _TrueUSD.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// TrueUSDWipedAccountIterator is returned from FilterWipedAccount and is used to iterate over the raw logs and unpacked data for WipedAccount events raised by the TrueUSD contract.
type TrueUSDWipedAccountIterator struct {
	Event *TrueUSDWipedAccount // Event containing the contract specifics and raw log

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
func (it *TrueUSDWipedAccountIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TrueUSDWipedAccount)
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
		it.Event = new(TrueUSDWipedAccount)
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
func (it *TrueUSDWipedAccountIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TrueUSDWipedAccountIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TrueUSDWipedAccount represents a WipedAccount event raised by the TrueUSD contract.
type TrueUSDWipedAccount struct {
	Account common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWipedAccount is a free log retrieval operation binding the contract event 0xdf58d2368c06216a398f05a7a88c8edc64a25c33f33fd2bd8b56fbc8822c02d8.
//
// Solidity: e WipedAccount(account indexed address, balance uint256)
func (_TrueUSD *TrueUSDFilterer) FilterWipedAccount(opts *bind.FilterOpts, account []common.Address) (*TrueUSDWipedAccountIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TrueUSD.contract.FilterLogs(opts, "WipedAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return &TrueUSDWipedAccountIterator{contract: _TrueUSD.contract, event: "WipedAccount", logs: logs, sub: sub}, nil
}

// WatchWipedAccount is a free log subscription operation binding the contract event 0xdf58d2368c06216a398f05a7a88c8edc64a25c33f33fd2bd8b56fbc8822c02d8.
//
// Solidity: e WipedAccount(account indexed address, balance uint256)
func (_TrueUSD *TrueUSDFilterer) WatchWipedAccount(opts *bind.WatchOpts, sink chan<- *TrueUSDWipedAccount, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _TrueUSD.contract.WatchLogs(opts, "WipedAccount", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TrueUSDWipedAccount)
				if err := _TrueUSD.contract.UnpackLog(event, "WipedAccount", log); err != nil {
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
