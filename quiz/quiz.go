// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package quiz

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// QuizABI is the input ABI used to generate the binding from.
const QuizABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_qn\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_ans\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkBoard\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"question\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_ans\",\"type\":\"bytes32\"}],\"name\":\"sendAnswer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// QuizBin is the compiled bytecode used for deploying new contracts.
var QuizBin = `60806040523480156100115760006000fd5b506040516104f83803806104f8833981810160405260408110156100355760006000fd5b81019080805160405193929190846401000000008211156100565760006000fd5b8382019150602082018581111561006d5760006000fd5b825186600182028301116401000000008211171561008b5760006000fd5b8083526020830192505050908051906020019080838360005b838110156100c05780820151818401525b6020810190506100a4565b50505050905090810190601f1680156100ed5780820380516001836020036101000a031916815260200191505b50604052602001805190602001909291905050505b816000600050908051906020019061011b929190610134565b50806001600050819090600019169055505b50506101e4565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061017557805160ff19168380011785556101a8565b828001600101855582156101a8579182015b828111156101a75782518260005090905591602001919060010190610187565b5b5090506101b591906101b9565b5090565b6101e191906101c3565b808211156101dd57600081815060009055506001016101c3565b5090565b90565b610305806101f36000396000f3fe60806040523480156100115760006000fd5b50600436106100465760003560e01c806317d1653c1461004c5780633fad9ae01461009757806377f46bff1461011b57610046565b60006000fd5b61007d600480360360208110156100635760006000fd5b81019080803560001916906020019092919050505061013d565b604051808215151515815260200191505060405180910390f35b61009f610169565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100e05780820151818401525b6020810190506100c4565b50505050905090810190601f16801561010d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61012361020a565b604051808215151515815260200191505060405180910390f35b600061015d8260001916600160005054600019161461026663ffffffff16565b9050610164565b919050565b60006000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102025780601f106101d757610100808354040283529160200191610202565b820191906000526020600020905b8154815290600101906020018083116101e557829003601f168201915b505050505081565b6000600260005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050610263565b90565b600081600260005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550600190506102cb565b91905056fea265627a7a72315820410c7a8a888967afb4f14cf2008dd091d77d6d875129e02e7d1f8e1ac7fb3dd064736f6c634300050c0032`

// DeployQuiz deploys a new Ethereum contract, binding an instance of Quiz to it.
func DeployQuiz(auth *bind.TransactOpts, backend bind.ContractBackend, _qn string, _ans [32]byte) (common.Address, *types.Transaction, *Quiz, error) {
	parsed, err := abi.JSON(strings.NewReader(QuizABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(QuizBin), backend, _qn, _ans)
	if err != nil {

		return common.Address{}, nil, nil, err
	}
	return address, tx, &Quiz{QuizCaller: QuizCaller{contract: contract}, QuizTransactor: QuizTransactor{contract: contract}, QuizFilterer: QuizFilterer{contract: contract}}, nil
}

// Quiz is an auto generated Go binding around an Ethereum contract.
type Quiz struct {
	QuizCaller     // Read-only binding to the contract
	QuizTransactor // Write-only binding to the contract
	QuizFilterer   // Log filterer for contract events
}

// QuizCaller is an auto generated read-only Go binding around an Ethereum contract.
type QuizCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuizTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QuizTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuizFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QuizFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QuizSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QuizSession struct {
	Contract     *Quiz             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuizCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QuizCallerSession struct {
	Contract *QuizCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// QuizTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QuizTransactorSession struct {
	Contract     *QuizTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QuizRaw is an auto generated low-level Go binding around an Ethereum contract.
type QuizRaw struct {
	Contract *Quiz // Generic contract binding to access the raw methods on
}

// QuizCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QuizCallerRaw struct {
	Contract *QuizCaller // Generic read-only contract binding to access the raw methods on
}

// QuizTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QuizTransactorRaw struct {
	Contract *QuizTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQuiz creates a new instance of Quiz, bound to a specific deployed contract.
func NewQuiz(address common.Address, backend bind.ContractBackend) (*Quiz, error) {
	contract, err := bindQuiz(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Quiz{QuizCaller: QuizCaller{contract: contract}, QuizTransactor: QuizTransactor{contract: contract}, QuizFilterer: QuizFilterer{contract: contract}}, nil
}

// NewQuizCaller creates a new read-only instance of Quiz, bound to a specific deployed contract.
func NewQuizCaller(address common.Address, caller bind.ContractCaller) (*QuizCaller, error) {
	contract, err := bindQuiz(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QuizCaller{contract: contract}, nil
}

// NewQuizTransactor creates a new write-only instance of Quiz, bound to a specific deployed contract.
func NewQuizTransactor(address common.Address, transactor bind.ContractTransactor) (*QuizTransactor, error) {
	contract, err := bindQuiz(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QuizTransactor{contract: contract}, nil
}

// NewQuizFilterer creates a new log filterer instance of Quiz, bound to a specific deployed contract.
func NewQuizFilterer(address common.Address, filterer bind.ContractFilterer) (*QuizFilterer, error) {
	contract, err := bindQuiz(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QuizFilterer{contract: contract}, nil
}

// bindQuiz binds a generic wrapper to an already deployed contract.
func bindQuiz(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QuizABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quiz *QuizRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Quiz.Contract.QuizCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quiz *QuizRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quiz.Contract.QuizTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quiz *QuizRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quiz.Contract.QuizTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Quiz *QuizCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Quiz.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Quiz *QuizTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Quiz.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Quiz *QuizTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Quiz.Contract.contract.Transact(opts, method, params...)
}

// CheckBoard is a free data retrieval call binding the contract method 0x77f46bff.
//
// Solidity: function checkBoard() constant returns(bool)
func (_Quiz *QuizCaller) CheckBoard(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Quiz.contract.Call(opts, out, "checkBoard")
	return *ret0, err
}

// CheckBoard is a free data retrieval call binding the contract method 0x77f46bff.
//
// Solidity: function checkBoard() constant returns(bool)
func (_Quiz *QuizSession) CheckBoard() (bool, error) {
	return _Quiz.Contract.CheckBoard(&_Quiz.CallOpts)
}

// CheckBoard is a free data retrieval call binding the contract method 0x77f46bff.
//
// Solidity: function checkBoard() constant returns(bool)
func (_Quiz *QuizCallerSession) CheckBoard() (bool, error) {
	return _Quiz.Contract.CheckBoard(&_Quiz.CallOpts)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() constant returns(string)
func (_Quiz *QuizCaller) Question(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Quiz.contract.Call(opts, out, "question")
	return *ret0, err
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() constant returns(string)
func (_Quiz *QuizSession) Question() (string, error) {
	return _Quiz.Contract.Question(&_Quiz.CallOpts)
}

// Question is a free data retrieval call binding the contract method 0x3fad9ae0.
//
// Solidity: function question() constant returns(string)
func (_Quiz *QuizCallerSession) Question() (string, error) {
	return _Quiz.Contract.Question(&_Quiz.CallOpts)
}

// SendAnswer is a paid mutator transaction binding the contract method 0x17d1653c.
//
// Solidity: function sendAnswer(bytes32 _ans) returns(bool)
func (_Quiz *QuizTransactor) SendAnswer(opts *bind.TransactOpts, _ans [32]byte) (*types.Transaction, error) {
	return _Quiz.contract.Transact(opts, "sendAnswer", _ans)
}

// SendAnswer is a paid mutator transaction binding the contract method 0x17d1653c.
//
// Solidity: function sendAnswer(bytes32 _ans) returns(bool)
func (_Quiz *QuizSession) SendAnswer(_ans [32]byte) (*types.Transaction, error) {
	return _Quiz.Contract.SendAnswer(&_Quiz.TransactOpts, _ans)
}

// SendAnswer is a paid mutator transaction binding the contract method 0x17d1653c.
//
// Solidity: function sendAnswer(bytes32 _ans) returns(bool)
func (_Quiz *QuizTransactorSession) SendAnswer(_ans [32]byte) (*types.Transaction, error) {
	return _Quiz.Contract.SendAnswer(&_Quiz.TransactOpts, _ans)
}
