// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// CharityVaultMetaData contains all meta data concerning the CharityVault contract.
var CharityVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Donated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"donations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"donor\",\"type\":\"address\"}],\"name\":\"getDonationOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDonated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052348015600f57600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506109478061005f6000396000f3fe6080604052600436106100585760003560e01c8062b37044146100c05780632e1a7d4d146100eb5780638da5cb5b14610114578063cc6cb19a1461013f578063dc5b2ee41461017c578063ed88c68e146101b9576100bb565b366100bb573073ffffffffffffffffffffffffffffffffffffffff1663ed88c68e6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156100a557600080fd5b505af11580156100b9573d6000803e3d6000fd5b005b600080fd5b3480156100cc57600080fd5b506100d56101c3565b6040516100e29190610579565b60405180910390f35b3480156100f757600080fd5b50610112600480360381019061010d91906105c5565b6101c9565b005b34801561012057600080fd5b506101296103d9565b6040516101369190610633565b60405180910390f35b34801561014b57600080fd5b506101666004803603810190610161919061067a565b6103fd565b6040516101739190610579565b60405180910390f35b34801561018857600080fd5b506101a3600480360381019061019e919061067a565b610415565b6040516101b09190610579565b60405180910390f35b6101c161045e565b005b60015481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610257576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024e90610704565b60405180910390fd5b4781111561029a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161029190610770565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16826040516102e1906107c1565b60006040518083038185875af1925050503d806000811461031e576040519150601f19603f3d011682016040523d82523d6000602084013e610323565b606091505b5050905080610367576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035e90610822565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364836040516103cd9190610579565b60405180910390a25050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60026020528060005260406000206000915090505481565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600034116104a1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104989061088e565b60405180910390fd5b34600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104f091906108dd565b92505081905550346001600082825461050991906108dd565b925050819055503373ffffffffffffffffffffffffffffffffffffffff167f2a01595cddf097c90216094025db714da3f4e5bd8877b56ba86a24ecead8e543346040516105569190610579565b60405180910390a2565b6000819050919050565b61057381610560565b82525050565b600060208201905061058e600083018461056a565b92915050565b600080fd5b6105a281610560565b81146105ad57600080fd5b50565b6000813590506105bf81610599565b92915050565b6000602082840312156105db576105da610594565b5b60006105e9848285016105b0565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061061d826105f2565b9050919050565b61062d81610612565b82525050565b60006020820190506106486000830184610624565b92915050565b61065781610612565b811461066257600080fd5b50565b6000813590506106748161064e565b92915050565b6000602082840312156106905761068f610594565b5b600061069e84828501610665565b91505092915050565b600082825260208201905092915050565b7f4e6f74206f776e65720000000000000000000000000000000000000000000000600082015250565b60006106ee6009836106a7565b91506106f9826106b8565b602082019050919050565b6000602082019050818103600083015261071d816106e1565b9050919050565b7f496e73756666696369656e742062616c616e6365000000000000000000000000600082015250565b600061075a6014836106a7565b915061076582610724565b602082019050919050565b600060208201905081810360008301526107898161074d565b9050919050565b600081905092915050565b50565b60006107ab600083610790565b91506107b68261079b565b600082019050919050565b60006107cc8261079e565b9150819050919050565b7f5769746864726177206661696c65640000000000000000000000000000000000600082015250565b600061080c600f836106a7565b9150610817826107d6565b602082019050919050565b6000602082019050818103600083015261083b816107ff565b9050919050565b7f4d7573742073656e642045544800000000000000000000000000000000000000600082015250565b6000610878600d836106a7565b915061088382610842565b602082019050919050565b600060208201905081810360008301526108a78161086b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006108e882610560565b91506108f383610560565b925082820190508082111561090b5761090a6108ae565b5b9291505056fea26469706673582212205ef522f7b967c61e3c92be82e120f7fdfd4aeea9d91b1a16fea0543605e1e61864736f6c634300081c0033",
}

// CharityVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use CharityVaultMetaData.ABI instead.
var CharityVaultABI = CharityVaultMetaData.ABI

// CharityVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CharityVaultMetaData.Bin instead.
var CharityVaultBin = CharityVaultMetaData.Bin

// DeployCharityVault deploys a new Ethereum contract, binding an instance of CharityVault to it.
func DeployCharityVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CharityVault, error) {
	parsed, err := CharityVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CharityVaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CharityVault{CharityVaultCaller: CharityVaultCaller{contract: contract}, CharityVaultTransactor: CharityVaultTransactor{contract: contract}, CharityVaultFilterer: CharityVaultFilterer{contract: contract}}, nil
}

// CharityVault is an auto generated Go binding around an Ethereum contract.
type CharityVault struct {
	CharityVaultCaller     // Read-only binding to the contract
	CharityVaultTransactor // Write-only binding to the contract
	CharityVaultFilterer   // Log filterer for contract events
}

// CharityVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type CharityVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharityVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CharityVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharityVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CharityVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharityVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CharityVaultSession struct {
	Contract     *CharityVault     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CharityVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CharityVaultCallerSession struct {
	Contract *CharityVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CharityVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CharityVaultTransactorSession struct {
	Contract     *CharityVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CharityVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type CharityVaultRaw struct {
	Contract *CharityVault // Generic contract binding to access the raw methods on
}

// CharityVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CharityVaultCallerRaw struct {
	Contract *CharityVaultCaller // Generic read-only contract binding to access the raw methods on
}

// CharityVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CharityVaultTransactorRaw struct {
	Contract *CharityVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCharityVault creates a new instance of CharityVault, bound to a specific deployed contract.
func NewCharityVault(address common.Address, backend bind.ContractBackend) (*CharityVault, error) {
	contract, err := bindCharityVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CharityVault{CharityVaultCaller: CharityVaultCaller{contract: contract}, CharityVaultTransactor: CharityVaultTransactor{contract: contract}, CharityVaultFilterer: CharityVaultFilterer{contract: contract}}, nil
}

// NewCharityVaultCaller creates a new read-only instance of CharityVault, bound to a specific deployed contract.
func NewCharityVaultCaller(address common.Address, caller bind.ContractCaller) (*CharityVaultCaller, error) {
	contract, err := bindCharityVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CharityVaultCaller{contract: contract}, nil
}

// NewCharityVaultTransactor creates a new write-only instance of CharityVault, bound to a specific deployed contract.
func NewCharityVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*CharityVaultTransactor, error) {
	contract, err := bindCharityVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CharityVaultTransactor{contract: contract}, nil
}

// NewCharityVaultFilterer creates a new log filterer instance of CharityVault, bound to a specific deployed contract.
func NewCharityVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*CharityVaultFilterer, error) {
	contract, err := bindCharityVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CharityVaultFilterer{contract: contract}, nil
}

// bindCharityVault binds a generic wrapper to an already deployed contract.
func bindCharityVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CharityVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CharityVault *CharityVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CharityVault.Contract.CharityVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CharityVault *CharityVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CharityVault.Contract.CharityVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CharityVault *CharityVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CharityVault.Contract.CharityVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CharityVault *CharityVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CharityVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CharityVault *CharityVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CharityVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CharityVault *CharityVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CharityVault.Contract.contract.Transact(opts, method, params...)
}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_CharityVault *CharityVaultCaller) Donations(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CharityVault.contract.Call(opts, &out, "donations", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_CharityVault *CharityVaultSession) Donations(arg0 common.Address) (*big.Int, error) {
	return _CharityVault.Contract.Donations(&_CharityVault.CallOpts, arg0)
}

// Donations is a free data retrieval call binding the contract method 0xcc6cb19a.
//
// Solidity: function donations(address ) view returns(uint256)
func (_CharityVault *CharityVaultCallerSession) Donations(arg0 common.Address) (*big.Int, error) {
	return _CharityVault.Contract.Donations(&_CharityVault.CallOpts, arg0)
}

// GetDonationOf is a free data retrieval call binding the contract method 0xdc5b2ee4.
//
// Solidity: function getDonationOf(address donor) view returns(uint256)
func (_CharityVault *CharityVaultCaller) GetDonationOf(opts *bind.CallOpts, donor common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CharityVault.contract.Call(opts, &out, "getDonationOf", donor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDonationOf is a free data retrieval call binding the contract method 0xdc5b2ee4.
//
// Solidity: function getDonationOf(address donor) view returns(uint256)
func (_CharityVault *CharityVaultSession) GetDonationOf(donor common.Address) (*big.Int, error) {
	return _CharityVault.Contract.GetDonationOf(&_CharityVault.CallOpts, donor)
}

// GetDonationOf is a free data retrieval call binding the contract method 0xdc5b2ee4.
//
// Solidity: function getDonationOf(address donor) view returns(uint256)
func (_CharityVault *CharityVaultCallerSession) GetDonationOf(donor common.Address) (*big.Int, error) {
	return _CharityVault.Contract.GetDonationOf(&_CharityVault.CallOpts, donor)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CharityVault *CharityVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CharityVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CharityVault *CharityVaultSession) Owner() (common.Address, error) {
	return _CharityVault.Contract.Owner(&_CharityVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CharityVault *CharityVaultCallerSession) Owner() (common.Address, error) {
	return _CharityVault.Contract.Owner(&_CharityVault.CallOpts)
}

// TotalDonated is a free data retrieval call binding the contract method 0x00b37044.
//
// Solidity: function totalDonated() view returns(uint256)
func (_CharityVault *CharityVaultCaller) TotalDonated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CharityVault.contract.Call(opts, &out, "totalDonated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDonated is a free data retrieval call binding the contract method 0x00b37044.
//
// Solidity: function totalDonated() view returns(uint256)
func (_CharityVault *CharityVaultSession) TotalDonated() (*big.Int, error) {
	return _CharityVault.Contract.TotalDonated(&_CharityVault.CallOpts)
}

// TotalDonated is a free data retrieval call binding the contract method 0x00b37044.
//
// Solidity: function totalDonated() view returns(uint256)
func (_CharityVault *CharityVaultCallerSession) TotalDonated() (*big.Int, error) {
	return _CharityVault.Contract.TotalDonated(&_CharityVault.CallOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_CharityVault *CharityVaultTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CharityVault.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_CharityVault *CharityVaultSession) Donate() (*types.Transaction, error) {
	return _CharityVault.Contract.Donate(&_CharityVault.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_CharityVault *CharityVaultTransactorSession) Donate() (*types.Transaction, error) {
	return _CharityVault.Contract.Donate(&_CharityVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_CharityVault *CharityVaultTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CharityVault.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_CharityVault *CharityVaultSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _CharityVault.Contract.Withdraw(&_CharityVault.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_CharityVault *CharityVaultTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _CharityVault.Contract.Withdraw(&_CharityVault.TransactOpts, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CharityVault *CharityVaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CharityVault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CharityVault *CharityVaultSession) Receive() (*types.Transaction, error) {
	return _CharityVault.Contract.Receive(&_CharityVault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CharityVault *CharityVaultTransactorSession) Receive() (*types.Transaction, error) {
	return _CharityVault.Contract.Receive(&_CharityVault.TransactOpts)
}

// CharityVaultDonatedIterator is returned from FilterDonated and is used to iterate over the raw logs and unpacked data for Donated events raised by the CharityVault contract.
type CharityVaultDonatedIterator struct {
	Event *CharityVaultDonated // Event containing the contract specifics and raw log

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
func (it *CharityVaultDonatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharityVaultDonated)
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
		it.Event = new(CharityVaultDonated)
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
func (it *CharityVaultDonatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharityVaultDonatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharityVaultDonated represents a Donated event raised by the CharityVault contract.
type CharityVaultDonated struct {
	Donor  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDonated is a free log retrieval operation binding the contract event 0x2a01595cddf097c90216094025db714da3f4e5bd8877b56ba86a24ecead8e543.
//
// Solidity: event Donated(address indexed donor, uint256 amount)
func (_CharityVault *CharityVaultFilterer) FilterDonated(opts *bind.FilterOpts, donor []common.Address) (*CharityVaultDonatedIterator, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _CharityVault.contract.FilterLogs(opts, "Donated", donorRule)
	if err != nil {
		return nil, err
	}
	return &CharityVaultDonatedIterator{contract: _CharityVault.contract, event: "Donated", logs: logs, sub: sub}, nil
}

// WatchDonated is a free log subscription operation binding the contract event 0x2a01595cddf097c90216094025db714da3f4e5bd8877b56ba86a24ecead8e543.
//
// Solidity: event Donated(address indexed donor, uint256 amount)
func (_CharityVault *CharityVaultFilterer) WatchDonated(opts *bind.WatchOpts, sink chan<- *CharityVaultDonated, donor []common.Address) (event.Subscription, error) {

	var donorRule []interface{}
	for _, donorItem := range donor {
		donorRule = append(donorRule, donorItem)
	}

	logs, sub, err := _CharityVault.contract.WatchLogs(opts, "Donated", donorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharityVaultDonated)
				if err := _CharityVault.contract.UnpackLog(event, "Donated", log); err != nil {
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

// ParseDonated is a log parse operation binding the contract event 0x2a01595cddf097c90216094025db714da3f4e5bd8877b56ba86a24ecead8e543.
//
// Solidity: event Donated(address indexed donor, uint256 amount)
func (_CharityVault *CharityVaultFilterer) ParseDonated(log types.Log) (*CharityVaultDonated, error) {
	event := new(CharityVaultDonated)
	if err := _CharityVault.contract.UnpackLog(event, "Donated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CharityVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the CharityVault contract.
type CharityVaultWithdrawIterator struct {
	Event *CharityVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *CharityVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CharityVaultWithdraw)
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
		it.Event = new(CharityVaultWithdraw)
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
func (it *CharityVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CharityVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CharityVaultWithdraw represents a Withdraw event raised by the CharityVault contract.
type CharityVaultWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_CharityVault *CharityVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*CharityVaultWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CharityVault.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &CharityVaultWithdrawIterator{contract: _CharityVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_CharityVault *CharityVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *CharityVaultWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CharityVault.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CharityVaultWithdraw)
				if err := _CharityVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_CharityVault *CharityVaultFilterer) ParseWithdraw(log types.Log) (*CharityVaultWithdraw, error) {
	event := new(CharityVaultWithdraw)
	if err := _CharityVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
