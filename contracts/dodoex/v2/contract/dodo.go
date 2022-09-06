// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receiveBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payQuote\",\"type\":\"uint256\"}],\"name\":\"BuyBaseToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maintainer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBaseToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ChargeMaintainerFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBaseToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ChargePenalty\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quoteTokenAmount\",\"type\":\"uint256\"}],\"name\":\"ClaimAssets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBaseToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenAmount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBaseToken\",\"type\":\"bool\"}],\"name\":\"Donate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payBase\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"receiveQuote\",\"type\":\"uint256\"}],\"name\":\"SellBaseToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldGasPriceLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGasPriceLimit\",\"type\":\"uint256\"}],\"name\":\"UpdateGasPriceLimit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldK\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newK\",\"type\":\"uint256\"}],\"name\":\"UpdateK\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidityProviderFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidityProviderFeeRate\",\"type\":\"uint256\"}],\"name\":\"UpdateLiquidityProviderFeeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldMaintainerFeeRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaintainerFeeRate\",\"type\":\"uint256\"}],\"name\":\"UpdateMaintainerFeeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBaseToken\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenAmount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"_BASE_BALANCE_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_CAPITAL_RECEIVE_QUOTE_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_CAPITAL_TOKEN_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_BASE_TOKEN_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_CLAIMED_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_CLOSED_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_DEPOSIT_BASE_ALLOWED_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_DEPOSIT_QUOTE_ALLOWED_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_GAS_PRICE_LIMIT_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_K_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_LP_FEE_RATE_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_MAINTAINER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_MT_FEE_RATE_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_NEW_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_ORACLE_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_OWNER_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_BALANCE_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_CAPITAL_RECEIVE_BASE_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_CAPITAL_TOKEN_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_QUOTE_TOKEN_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_R_STATUS_\",\"outputs\":[{\"internalType\":\"enumTypes.RStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_SUPERVISOR_\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_TARGET_BASE_TOKEN_AMOUNT_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_TARGET_QUOTE_TOKEN_AMOUNT_\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_TRADE_ALLOWED_\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPayQuote\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"buyBaseToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositBaseTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositQuoteTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableBaseDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableQuoteDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"disableTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"donateBaseToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"donateQuoteToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableBaseDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableQuoteDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"enableTrading\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalSettlement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"getBaseCapitalBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExpectedTarget\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseTarget\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quoteTarget\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"getLpBaseBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lpBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"getLpQuoteBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lpBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMidPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"midPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOraclePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"getQuoteCapitalBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalBaseCapital\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalQuoteCapital\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getWithdrawBasePenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getWithdrawQuotePenalty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"supervisor\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"maintainer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"baseToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"quoteToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lpFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mtFeeRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"k\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPriceLimit\",\"type\":\"uint256\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"queryBuyBaseToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"payQuote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"querySellBaseToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiveQuote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"retrieve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minReceiveQuote\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"sellBaseToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newGasPriceLimit\",\"type\":\"uint256\"}],\"name\":\"setGasPriceLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newK\",\"type\":\"uint256\"}],\"name\":\"setK\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLiquidityPorviderFeeRate\",\"type\":\"uint256\"}],\"name\":\"setLiquidityProviderFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMaintainer\",\"type\":\"address\"}],\"name\":\"setMaintainer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaintainerFeeRate\",\"type\":\"uint256\"}],\"name\":\"setMaintainerFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSupervisor\",\"type\":\"address\"}],\"name\":\"setSupervisor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawAllBaseTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAllQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawAllQuoteTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawBase\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawBaseTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawQuote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawQuoteTo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// BASEBALANCE is a free data retrieval call binding the contract method 0xeab5d20e.
//
// Solidity: function _BASE_BALANCE_() view returns(uint256)
func (_Contract *ContractCaller) BASEBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_BASE_BALANCE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASEBALANCE is a free data retrieval call binding the contract method 0xeab5d20e.
//
// Solidity: function _BASE_BALANCE_() view returns(uint256)
func (_Contract *ContractSession) BASEBALANCE() (*big.Int, error) {
	return _Contract.Contract.BASEBALANCE(&_Contract.CallOpts)
}

// BASEBALANCE is a free data retrieval call binding the contract method 0xeab5d20e.
//
// Solidity: function _BASE_BALANCE_() view returns(uint256)
func (_Contract *ContractCallerSession) BASEBALANCE() (*big.Int, error) {
	return _Contract.Contract.BASEBALANCE(&_Contract.CallOpts)
}

// BASECAPITALRECEIVEQUOTE is a free data retrieval call binding the contract method 0xc6b73cf9.
//
// Solidity: function _BASE_CAPITAL_RECEIVE_QUOTE_() view returns(uint256)
func (_Contract *ContractCaller) BASECAPITALRECEIVEQUOTE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_BASE_CAPITAL_RECEIVE_QUOTE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BASECAPITALRECEIVEQUOTE is a free data retrieval call binding the contract method 0xc6b73cf9.
//
// Solidity: function _BASE_CAPITAL_RECEIVE_QUOTE_() view returns(uint256)
func (_Contract *ContractSession) BASECAPITALRECEIVEQUOTE() (*big.Int, error) {
	return _Contract.Contract.BASECAPITALRECEIVEQUOTE(&_Contract.CallOpts)
}

// BASECAPITALRECEIVEQUOTE is a free data retrieval call binding the contract method 0xc6b73cf9.
//
// Solidity: function _BASE_CAPITAL_RECEIVE_QUOTE_() view returns(uint256)
func (_Contract *ContractCallerSession) BASECAPITALRECEIVEQUOTE() (*big.Int, error) {
	return _Contract.Contract.BASECAPITALRECEIVEQUOTE(&_Contract.CallOpts)
}

// BASECAPITALTOKEN is a free data retrieval call binding the contract method 0xd689107c.
//
// Solidity: function _BASE_CAPITAL_TOKEN_() view returns(address)
func (_Contract *ContractCaller) BASECAPITALTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_BASE_CAPITAL_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BASECAPITALTOKEN is a free data retrieval call binding the contract method 0xd689107c.
//
// Solidity: function _BASE_CAPITAL_TOKEN_() view returns(address)
func (_Contract *ContractSession) BASECAPITALTOKEN() (common.Address, error) {
	return _Contract.Contract.BASECAPITALTOKEN(&_Contract.CallOpts)
}

// BASECAPITALTOKEN is a free data retrieval call binding the contract method 0xd689107c.
//
// Solidity: function _BASE_CAPITAL_TOKEN_() view returns(address)
func (_Contract *ContractCallerSession) BASECAPITALTOKEN() (common.Address, error) {
	return _Contract.Contract.BASECAPITALTOKEN(&_Contract.CallOpts)
}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_Contract *ContractCaller) BASETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_BASE_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_Contract *ContractSession) BASETOKEN() (common.Address, error) {
	return _Contract.Contract.BASETOKEN(&_Contract.CallOpts)
}

// BASETOKEN is a free data retrieval call binding the contract method 0x4a248d2a.
//
// Solidity: function _BASE_TOKEN_() view returns(address)
func (_Contract *ContractCallerSession) BASETOKEN() (common.Address, error) {
	return _Contract.Contract.BASETOKEN(&_Contract.CallOpts)
}

// CLAIMED is a free data retrieval call binding the contract method 0x68be20ad.
//
// Solidity: function _CLAIMED_(address ) view returns(bool)
func (_Contract *ContractCaller) CLAIMED(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_CLAIMED_", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CLAIMED is a free data retrieval call binding the contract method 0x68be20ad.
//
// Solidity: function _CLAIMED_(address ) view returns(bool)
func (_Contract *ContractSession) CLAIMED(arg0 common.Address) (bool, error) {
	return _Contract.Contract.CLAIMED(&_Contract.CallOpts, arg0)
}

// CLAIMED is a free data retrieval call binding the contract method 0x68be20ad.
//
// Solidity: function _CLAIMED_(address ) view returns(bool)
func (_Contract *ContractCallerSession) CLAIMED(arg0 common.Address) (bool, error) {
	return _Contract.Contract.CLAIMED(&_Contract.CallOpts, arg0)
}

// CLOSED is a free data retrieval call binding the contract method 0x6ec6a58d.
//
// Solidity: function _CLOSED_() view returns(bool)
func (_Contract *ContractCaller) CLOSED(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_CLOSED_")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CLOSED is a free data retrieval call binding the contract method 0x6ec6a58d.
//
// Solidity: function _CLOSED_() view returns(bool)
func (_Contract *ContractSession) CLOSED() (bool, error) {
	return _Contract.Contract.CLOSED(&_Contract.CallOpts)
}

// CLOSED is a free data retrieval call binding the contract method 0x6ec6a58d.
//
// Solidity: function _CLOSED_() view returns(bool)
func (_Contract *ContractCallerSession) CLOSED() (bool, error) {
	return _Contract.Contract.CLOSED(&_Contract.CallOpts)
}

// DEPOSITBASEALLOWED is a free data retrieval call binding the contract method 0xa598aca7.
//
// Solidity: function _DEPOSIT_BASE_ALLOWED_() view returns(bool)
func (_Contract *ContractCaller) DEPOSITBASEALLOWED(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_DEPOSIT_BASE_ALLOWED_")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DEPOSITBASEALLOWED is a free data retrieval call binding the contract method 0xa598aca7.
//
// Solidity: function _DEPOSIT_BASE_ALLOWED_() view returns(bool)
func (_Contract *ContractSession) DEPOSITBASEALLOWED() (bool, error) {
	return _Contract.Contract.DEPOSITBASEALLOWED(&_Contract.CallOpts)
}

// DEPOSITBASEALLOWED is a free data retrieval call binding the contract method 0xa598aca7.
//
// Solidity: function _DEPOSIT_BASE_ALLOWED_() view returns(bool)
func (_Contract *ContractCallerSession) DEPOSITBASEALLOWED() (bool, error) {
	return _Contract.Contract.DEPOSITBASEALLOWED(&_Contract.CallOpts)
}

// DEPOSITQUOTEALLOWED is a free data retrieval call binding the contract method 0xc5bbffe8.
//
// Solidity: function _DEPOSIT_QUOTE_ALLOWED_() view returns(bool)
func (_Contract *ContractCaller) DEPOSITQUOTEALLOWED(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_DEPOSIT_QUOTE_ALLOWED_")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DEPOSITQUOTEALLOWED is a free data retrieval call binding the contract method 0xc5bbffe8.
//
// Solidity: function _DEPOSIT_QUOTE_ALLOWED_() view returns(bool)
func (_Contract *ContractSession) DEPOSITQUOTEALLOWED() (bool, error) {
	return _Contract.Contract.DEPOSITQUOTEALLOWED(&_Contract.CallOpts)
}

// DEPOSITQUOTEALLOWED is a free data retrieval call binding the contract method 0xc5bbffe8.
//
// Solidity: function _DEPOSIT_QUOTE_ALLOWED_() view returns(bool)
func (_Contract *ContractCallerSession) DEPOSITQUOTEALLOWED() (bool, error) {
	return _Contract.Contract.DEPOSITQUOTEALLOWED(&_Contract.CallOpts)
}

// GASPRICELIMIT is a free data retrieval call binding the contract method 0x4de4527e.
//
// Solidity: function _GAS_PRICE_LIMIT_() view returns(uint256)
func (_Contract *ContractCaller) GASPRICELIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_GAS_PRICE_LIMIT_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GASPRICELIMIT is a free data retrieval call binding the contract method 0x4de4527e.
//
// Solidity: function _GAS_PRICE_LIMIT_() view returns(uint256)
func (_Contract *ContractSession) GASPRICELIMIT() (*big.Int, error) {
	return _Contract.Contract.GASPRICELIMIT(&_Contract.CallOpts)
}

// GASPRICELIMIT is a free data retrieval call binding the contract method 0x4de4527e.
//
// Solidity: function _GAS_PRICE_LIMIT_() view returns(uint256)
func (_Contract *ContractCallerSession) GASPRICELIMIT() (*big.Int, error) {
	return _Contract.Contract.GASPRICELIMIT(&_Contract.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint256)
func (_Contract *ContractCaller) K(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_K_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint256)
func (_Contract *ContractSession) K() (*big.Int, error) {
	return _Contract.Contract.K(&_Contract.CallOpts)
}

// K is a free data retrieval call binding the contract method 0xec2fd46d.
//
// Solidity: function _K_() view returns(uint256)
func (_Contract *ContractCallerSession) K() (*big.Int, error) {
	return _Contract.Contract.K(&_Contract.CallOpts)
}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint256)
func (_Contract *ContractCaller) LPFEERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_LP_FEE_RATE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint256)
func (_Contract *ContractSession) LPFEERATE() (*big.Int, error) {
	return _Contract.Contract.LPFEERATE(&_Contract.CallOpts)
}

// LPFEERATE is a free data retrieval call binding the contract method 0xab44a7a3.
//
// Solidity: function _LP_FEE_RATE_() view returns(uint256)
func (_Contract *ContractCallerSession) LPFEERATE() (*big.Int, error) {
	return _Contract.Contract.LPFEERATE(&_Contract.CallOpts)
}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_Contract *ContractCaller) MAINTAINER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_MAINTAINER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_Contract *ContractSession) MAINTAINER() (common.Address, error) {
	return _Contract.Contract.MAINTAINER(&_Contract.CallOpts)
}

// MAINTAINER is a free data retrieval call binding the contract method 0x4322ec83.
//
// Solidity: function _MAINTAINER_() view returns(address)
func (_Contract *ContractCallerSession) MAINTAINER() (common.Address, error) {
	return _Contract.Contract.MAINTAINER(&_Contract.CallOpts)
}

// MTFEERATE is a free data retrieval call binding the contract method 0xc0ffa178.
//
// Solidity: function _MT_FEE_RATE_() view returns(uint256)
func (_Contract *ContractCaller) MTFEERATE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_MT_FEE_RATE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MTFEERATE is a free data retrieval call binding the contract method 0xc0ffa178.
//
// Solidity: function _MT_FEE_RATE_() view returns(uint256)
func (_Contract *ContractSession) MTFEERATE() (*big.Int, error) {
	return _Contract.Contract.MTFEERATE(&_Contract.CallOpts)
}

// MTFEERATE is a free data retrieval call binding the contract method 0xc0ffa178.
//
// Solidity: function _MT_FEE_RATE_() view returns(uint256)
func (_Contract *ContractCallerSession) MTFEERATE() (*big.Int, error) {
	return _Contract.Contract.MTFEERATE(&_Contract.CallOpts)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_Contract *ContractCaller) NEWOWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_NEW_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_Contract *ContractSession) NEWOWNER() (common.Address, error) {
	return _Contract.Contract.NEWOWNER(&_Contract.CallOpts)
}

// NEWOWNER is a free data retrieval call binding the contract method 0x8456db15.
//
// Solidity: function _NEW_OWNER_() view returns(address)
func (_Contract *ContractCallerSession) NEWOWNER() (common.Address, error) {
	return _Contract.Contract.NEWOWNER(&_Contract.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x73a2ab7c.
//
// Solidity: function _ORACLE_() view returns(address)
func (_Contract *ContractCaller) ORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_ORACLE_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ORACLE is a free data retrieval call binding the contract method 0x73a2ab7c.
//
// Solidity: function _ORACLE_() view returns(address)
func (_Contract *ContractSession) ORACLE() (common.Address, error) {
	return _Contract.Contract.ORACLE(&_Contract.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x73a2ab7c.
//
// Solidity: function _ORACLE_() view returns(address)
func (_Contract *ContractCallerSession) ORACLE() (common.Address, error) {
	return _Contract.Contract.ORACLE(&_Contract.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_Contract *ContractCaller) OWNER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_OWNER_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_Contract *ContractSession) OWNER() (common.Address, error) {
	return _Contract.Contract.OWNER(&_Contract.CallOpts)
}

// OWNER is a free data retrieval call binding the contract method 0x16048bc4.
//
// Solidity: function _OWNER_() view returns(address)
func (_Contract *ContractCallerSession) OWNER() (common.Address, error) {
	return _Contract.Contract.OWNER(&_Contract.CallOpts)
}

// QUOTEBALANCE is a free data retrieval call binding the contract method 0x7c9b8e89.
//
// Solidity: function _QUOTE_BALANCE_() view returns(uint256)
func (_Contract *ContractCaller) QUOTEBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_QUOTE_BALANCE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUOTEBALANCE is a free data retrieval call binding the contract method 0x7c9b8e89.
//
// Solidity: function _QUOTE_BALANCE_() view returns(uint256)
func (_Contract *ContractSession) QUOTEBALANCE() (*big.Int, error) {
	return _Contract.Contract.QUOTEBALANCE(&_Contract.CallOpts)
}

// QUOTEBALANCE is a free data retrieval call binding the contract method 0x7c9b8e89.
//
// Solidity: function _QUOTE_BALANCE_() view returns(uint256)
func (_Contract *ContractCallerSession) QUOTEBALANCE() (*big.Int, error) {
	return _Contract.Contract.QUOTEBALANCE(&_Contract.CallOpts)
}

// QUOTECAPITALRECEIVEBASE is a free data retrieval call binding the contract method 0x0e6518e9.
//
// Solidity: function _QUOTE_CAPITAL_RECEIVE_BASE_() view returns(uint256)
func (_Contract *ContractCaller) QUOTECAPITALRECEIVEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_QUOTE_CAPITAL_RECEIVE_BASE_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUOTECAPITALRECEIVEBASE is a free data retrieval call binding the contract method 0x0e6518e9.
//
// Solidity: function _QUOTE_CAPITAL_RECEIVE_BASE_() view returns(uint256)
func (_Contract *ContractSession) QUOTECAPITALRECEIVEBASE() (*big.Int, error) {
	return _Contract.Contract.QUOTECAPITALRECEIVEBASE(&_Contract.CallOpts)
}

// QUOTECAPITALRECEIVEBASE is a free data retrieval call binding the contract method 0x0e6518e9.
//
// Solidity: function _QUOTE_CAPITAL_RECEIVE_BASE_() view returns(uint256)
func (_Contract *ContractCallerSession) QUOTECAPITALRECEIVEBASE() (*big.Int, error) {
	return _Contract.Contract.QUOTECAPITALRECEIVEBASE(&_Contract.CallOpts)
}

// QUOTECAPITALTOKEN is a free data retrieval call binding the contract method 0xac1fbc98.
//
// Solidity: function _QUOTE_CAPITAL_TOKEN_() view returns(address)
func (_Contract *ContractCaller) QUOTECAPITALTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_QUOTE_CAPITAL_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTECAPITALTOKEN is a free data retrieval call binding the contract method 0xac1fbc98.
//
// Solidity: function _QUOTE_CAPITAL_TOKEN_() view returns(address)
func (_Contract *ContractSession) QUOTECAPITALTOKEN() (common.Address, error) {
	return _Contract.Contract.QUOTECAPITALTOKEN(&_Contract.CallOpts)
}

// QUOTECAPITALTOKEN is a free data retrieval call binding the contract method 0xac1fbc98.
//
// Solidity: function _QUOTE_CAPITAL_TOKEN_() view returns(address)
func (_Contract *ContractCallerSession) QUOTECAPITALTOKEN() (common.Address, error) {
	return _Contract.Contract.QUOTECAPITALTOKEN(&_Contract.CallOpts)
}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_Contract *ContractCaller) QUOTETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_QUOTE_TOKEN_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_Contract *ContractSession) QUOTETOKEN() (common.Address, error) {
	return _Contract.Contract.QUOTETOKEN(&_Contract.CallOpts)
}

// QUOTETOKEN is a free data retrieval call binding the contract method 0xd4b97046.
//
// Solidity: function _QUOTE_TOKEN_() view returns(address)
func (_Contract *ContractCallerSession) QUOTETOKEN() (common.Address, error) {
	return _Contract.Contract.QUOTETOKEN(&_Contract.CallOpts)
}

// RSTATUS is a free data retrieval call binding the contract method 0x17be952e.
//
// Solidity: function _R_STATUS_() view returns(uint8)
func (_Contract *ContractCaller) RSTATUS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_R_STATUS_")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RSTATUS is a free data retrieval call binding the contract method 0x17be952e.
//
// Solidity: function _R_STATUS_() view returns(uint8)
func (_Contract *ContractSession) RSTATUS() (uint8, error) {
	return _Contract.Contract.RSTATUS(&_Contract.CallOpts)
}

// RSTATUS is a free data retrieval call binding the contract method 0x17be952e.
//
// Solidity: function _R_STATUS_() view returns(uint8)
func (_Contract *ContractCallerSession) RSTATUS() (uint8, error) {
	return _Contract.Contract.RSTATUS(&_Contract.CallOpts)
}

// SUPERVISOR is a free data retrieval call binding the contract method 0x3960f142.
//
// Solidity: function _SUPERVISOR_() view returns(address)
func (_Contract *ContractCaller) SUPERVISOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_SUPERVISOR_")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SUPERVISOR is a free data retrieval call binding the contract method 0x3960f142.
//
// Solidity: function _SUPERVISOR_() view returns(address)
func (_Contract *ContractSession) SUPERVISOR() (common.Address, error) {
	return _Contract.Contract.SUPERVISOR(&_Contract.CallOpts)
}

// SUPERVISOR is a free data retrieval call binding the contract method 0x3960f142.
//
// Solidity: function _SUPERVISOR_() view returns(address)
func (_Contract *ContractCallerSession) SUPERVISOR() (common.Address, error) {
	return _Contract.Contract.SUPERVISOR(&_Contract.CallOpts)
}

// TARGETBASETOKENAMOUNT is a free data retrieval call binding the contract method 0xb2094fd3.
//
// Solidity: function _TARGET_BASE_TOKEN_AMOUNT_() view returns(uint256)
func (_Contract *ContractCaller) TARGETBASETOKENAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_TARGET_BASE_TOKEN_AMOUNT_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TARGETBASETOKENAMOUNT is a free data retrieval call binding the contract method 0xb2094fd3.
//
// Solidity: function _TARGET_BASE_TOKEN_AMOUNT_() view returns(uint256)
func (_Contract *ContractSession) TARGETBASETOKENAMOUNT() (*big.Int, error) {
	return _Contract.Contract.TARGETBASETOKENAMOUNT(&_Contract.CallOpts)
}

// TARGETBASETOKENAMOUNT is a free data retrieval call binding the contract method 0xb2094fd3.
//
// Solidity: function _TARGET_BASE_TOKEN_AMOUNT_() view returns(uint256)
func (_Contract *ContractCallerSession) TARGETBASETOKENAMOUNT() (*big.Int, error) {
	return _Contract.Contract.TARGETBASETOKENAMOUNT(&_Contract.CallOpts)
}

// TARGETQUOTETOKENAMOUNT is a free data retrieval call binding the contract method 0x245c9685.
//
// Solidity: function _TARGET_QUOTE_TOKEN_AMOUNT_() view returns(uint256)
func (_Contract *ContractCaller) TARGETQUOTETOKENAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_TARGET_QUOTE_TOKEN_AMOUNT_")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TARGETQUOTETOKENAMOUNT is a free data retrieval call binding the contract method 0x245c9685.
//
// Solidity: function _TARGET_QUOTE_TOKEN_AMOUNT_() view returns(uint256)
func (_Contract *ContractSession) TARGETQUOTETOKENAMOUNT() (*big.Int, error) {
	return _Contract.Contract.TARGETQUOTETOKENAMOUNT(&_Contract.CallOpts)
}

// TARGETQUOTETOKENAMOUNT is a free data retrieval call binding the contract method 0x245c9685.
//
// Solidity: function _TARGET_QUOTE_TOKEN_AMOUNT_() view returns(uint256)
func (_Contract *ContractCallerSession) TARGETQUOTETOKENAMOUNT() (*big.Int, error) {
	return _Contract.Contract.TARGETQUOTETOKENAMOUNT(&_Contract.CallOpts)
}

// TRADEALLOWED is a free data retrieval call binding the contract method 0xdd58b41c.
//
// Solidity: function _TRADE_ALLOWED_() view returns(bool)
func (_Contract *ContractCaller) TRADEALLOWED(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "_TRADE_ALLOWED_")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TRADEALLOWED is a free data retrieval call binding the contract method 0xdd58b41c.
//
// Solidity: function _TRADE_ALLOWED_() view returns(bool)
func (_Contract *ContractSession) TRADEALLOWED() (bool, error) {
	return _Contract.Contract.TRADEALLOWED(&_Contract.CallOpts)
}

// TRADEALLOWED is a free data retrieval call binding the contract method 0xdd58b41c.
//
// Solidity: function _TRADE_ALLOWED_() view returns(bool)
func (_Contract *ContractCallerSession) TRADEALLOWED() (bool, error) {
	return _Contract.Contract.TRADEALLOWED(&_Contract.CallOpts)
}

// GetBaseCapitalBalanceOf is a free data retrieval call binding the contract method 0x7aed942d.
//
// Solidity: function getBaseCapitalBalanceOf(address lp) view returns(uint256)
func (_Contract *ContractCaller) GetBaseCapitalBalanceOf(opts *bind.CallOpts, lp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getBaseCapitalBalanceOf", lp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseCapitalBalanceOf is a free data retrieval call binding the contract method 0x7aed942d.
//
// Solidity: function getBaseCapitalBalanceOf(address lp) view returns(uint256)
func (_Contract *ContractSession) GetBaseCapitalBalanceOf(lp common.Address) (*big.Int, error) {
	return _Contract.Contract.GetBaseCapitalBalanceOf(&_Contract.CallOpts, lp)
}

// GetBaseCapitalBalanceOf is a free data retrieval call binding the contract method 0x7aed942d.
//
// Solidity: function getBaseCapitalBalanceOf(address lp) view returns(uint256)
func (_Contract *ContractCallerSession) GetBaseCapitalBalanceOf(lp common.Address) (*big.Int, error) {
	return _Contract.Contract.GetBaseCapitalBalanceOf(&_Contract.CallOpts, lp)
}

// GetExpectedTarget is a free data retrieval call binding the contract method 0xffa64225.
//
// Solidity: function getExpectedTarget() view returns(uint256 baseTarget, uint256 quoteTarget)
func (_Contract *ContractCaller) GetExpectedTarget(opts *bind.CallOpts) (struct {
	BaseTarget  *big.Int
	QuoteTarget *big.Int
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getExpectedTarget")

	outstruct := new(struct {
		BaseTarget  *big.Int
		QuoteTarget *big.Int
	})

	outstruct.BaseTarget = out[0].(*big.Int)
	outstruct.QuoteTarget = out[1].(*big.Int)

	return *outstruct, err

}

// GetExpectedTarget is a free data retrieval call binding the contract method 0xffa64225.
//
// Solidity: function getExpectedTarget() view returns(uint256 baseTarget, uint256 quoteTarget)
func (_Contract *ContractSession) GetExpectedTarget() (struct {
	BaseTarget  *big.Int
	QuoteTarget *big.Int
}, error) {
	return _Contract.Contract.GetExpectedTarget(&_Contract.CallOpts)
}

// GetExpectedTarget is a free data retrieval call binding the contract method 0xffa64225.
//
// Solidity: function getExpectedTarget() view returns(uint256 baseTarget, uint256 quoteTarget)
func (_Contract *ContractCallerSession) GetExpectedTarget() (struct {
	BaseTarget  *big.Int
	QuoteTarget *big.Int
}, error) {
	return _Contract.Contract.GetExpectedTarget(&_Contract.CallOpts)
}

// GetLpBaseBalance is a free data retrieval call binding the contract method 0x95faa5f6.
//
// Solidity: function getLpBaseBalance(address lp) view returns(uint256 lpBalance)
func (_Contract *ContractCaller) GetLpBaseBalance(opts *bind.CallOpts, lp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLpBaseBalance", lp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLpBaseBalance is a free data retrieval call binding the contract method 0x95faa5f6.
//
// Solidity: function getLpBaseBalance(address lp) view returns(uint256 lpBalance)
func (_Contract *ContractSession) GetLpBaseBalance(lp common.Address) (*big.Int, error) {
	return _Contract.Contract.GetLpBaseBalance(&_Contract.CallOpts, lp)
}

// GetLpBaseBalance is a free data retrieval call binding the contract method 0x95faa5f6.
//
// Solidity: function getLpBaseBalance(address lp) view returns(uint256 lpBalance)
func (_Contract *ContractCallerSession) GetLpBaseBalance(lp common.Address) (*big.Int, error) {
	return _Contract.Contract.GetLpBaseBalance(&_Contract.CallOpts, lp)
}

// GetLpQuoteBalance is a free data retrieval call binding the contract method 0x36a53bbb.
//
// Solidity: function getLpQuoteBalance(address lp) view returns(uint256 lpBalance)
func (_Contract *ContractCaller) GetLpQuoteBalance(opts *bind.CallOpts, lp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getLpQuoteBalance", lp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLpQuoteBalance is a free data retrieval call binding the contract method 0x36a53bbb.
//
// Solidity: function getLpQuoteBalance(address lp) view returns(uint256 lpBalance)
func (_Contract *ContractSession) GetLpQuoteBalance(lp common.Address) (*big.Int, error) {
	return _Contract.Contract.GetLpQuoteBalance(&_Contract.CallOpts, lp)
}

// GetLpQuoteBalance is a free data retrieval call binding the contract method 0x36a53bbb.
//
// Solidity: function getLpQuoteBalance(address lp) view returns(uint256 lpBalance)
func (_Contract *ContractCallerSession) GetLpQuoteBalance(lp common.Address) (*big.Int, error) {
	return _Contract.Contract.GetLpQuoteBalance(&_Contract.CallOpts, lp)
}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_Contract *ContractCaller) GetMidPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getMidPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_Contract *ContractSession) GetMidPrice() (*big.Int, error) {
	return _Contract.Contract.GetMidPrice(&_Contract.CallOpts)
}

// GetMidPrice is a free data retrieval call binding the contract method 0xee27c689.
//
// Solidity: function getMidPrice() view returns(uint256 midPrice)
func (_Contract *ContractCallerSession) GetMidPrice() (*big.Int, error) {
	return _Contract.Contract.GetMidPrice(&_Contract.CallOpts)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x796da7af.
//
// Solidity: function getOraclePrice() view returns(uint256)
func (_Contract *ContractCaller) GetOraclePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOraclePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOraclePrice is a free data retrieval call binding the contract method 0x796da7af.
//
// Solidity: function getOraclePrice() view returns(uint256)
func (_Contract *ContractSession) GetOraclePrice() (*big.Int, error) {
	return _Contract.Contract.GetOraclePrice(&_Contract.CallOpts)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x796da7af.
//
// Solidity: function getOraclePrice() view returns(uint256)
func (_Contract *ContractCallerSession) GetOraclePrice() (*big.Int, error) {
	return _Contract.Contract.GetOraclePrice(&_Contract.CallOpts)
}

// GetQuoteCapitalBalanceOf is a free data retrieval call binding the contract method 0xf67ed448.
//
// Solidity: function getQuoteCapitalBalanceOf(address lp) view returns(uint256)
func (_Contract *ContractCaller) GetQuoteCapitalBalanceOf(opts *bind.CallOpts, lp common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getQuoteCapitalBalanceOf", lp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQuoteCapitalBalanceOf is a free data retrieval call binding the contract method 0xf67ed448.
//
// Solidity: function getQuoteCapitalBalanceOf(address lp) view returns(uint256)
func (_Contract *ContractSession) GetQuoteCapitalBalanceOf(lp common.Address) (*big.Int, error) {
	return _Contract.Contract.GetQuoteCapitalBalanceOf(&_Contract.CallOpts, lp)
}

// GetQuoteCapitalBalanceOf is a free data retrieval call binding the contract method 0xf67ed448.
//
// Solidity: function getQuoteCapitalBalanceOf(address lp) view returns(uint256)
func (_Contract *ContractCallerSession) GetQuoteCapitalBalanceOf(lp common.Address) (*big.Int, error) {
	return _Contract.Contract.GetQuoteCapitalBalanceOf(&_Contract.CallOpts, lp)
}

// GetTotalBaseCapital is a free data retrieval call binding the contract method 0x0cd1667d.
//
// Solidity: function getTotalBaseCapital() view returns(uint256)
func (_Contract *ContractCaller) GetTotalBaseCapital(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTotalBaseCapital")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalBaseCapital is a free data retrieval call binding the contract method 0x0cd1667d.
//
// Solidity: function getTotalBaseCapital() view returns(uint256)
func (_Contract *ContractSession) GetTotalBaseCapital() (*big.Int, error) {
	return _Contract.Contract.GetTotalBaseCapital(&_Contract.CallOpts)
}

// GetTotalBaseCapital is a free data retrieval call binding the contract method 0x0cd1667d.
//
// Solidity: function getTotalBaseCapital() view returns(uint256)
func (_Contract *ContractCallerSession) GetTotalBaseCapital() (*big.Int, error) {
	return _Contract.Contract.GetTotalBaseCapital(&_Contract.CallOpts)
}

// GetTotalQuoteCapital is a free data retrieval call binding the contract method 0x2aa82c65.
//
// Solidity: function getTotalQuoteCapital() view returns(uint256)
func (_Contract *ContractCaller) GetTotalQuoteCapital(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTotalQuoteCapital")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalQuoteCapital is a free data retrieval call binding the contract method 0x2aa82c65.
//
// Solidity: function getTotalQuoteCapital() view returns(uint256)
func (_Contract *ContractSession) GetTotalQuoteCapital() (*big.Int, error) {
	return _Contract.Contract.GetTotalQuoteCapital(&_Contract.CallOpts)
}

// GetTotalQuoteCapital is a free data retrieval call binding the contract method 0x2aa82c65.
//
// Solidity: function getTotalQuoteCapital() view returns(uint256)
func (_Contract *ContractCallerSession) GetTotalQuoteCapital() (*big.Int, error) {
	return _Contract.Contract.GetTotalQuoteCapital(&_Contract.CallOpts)
}

// GetWithdrawBasePenalty is a free data retrieval call binding the contract method 0xee5150b3.
//
// Solidity: function getWithdrawBasePenalty(uint256 amount) view returns(uint256 penalty)
func (_Contract *ContractCaller) GetWithdrawBasePenalty(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getWithdrawBasePenalty", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawBasePenalty is a free data retrieval call binding the contract method 0xee5150b3.
//
// Solidity: function getWithdrawBasePenalty(uint256 amount) view returns(uint256 penalty)
func (_Contract *ContractSession) GetWithdrawBasePenalty(amount *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetWithdrawBasePenalty(&_Contract.CallOpts, amount)
}

// GetWithdrawBasePenalty is a free data retrieval call binding the contract method 0xee5150b3.
//
// Solidity: function getWithdrawBasePenalty(uint256 amount) view returns(uint256 penalty)
func (_Contract *ContractCallerSession) GetWithdrawBasePenalty(amount *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetWithdrawBasePenalty(&_Contract.CallOpts, amount)
}

// GetWithdrawQuotePenalty is a free data retrieval call binding the contract method 0x0c9f7bd0.
//
// Solidity: function getWithdrawQuotePenalty(uint256 amount) view returns(uint256 penalty)
func (_Contract *ContractCaller) GetWithdrawQuotePenalty(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getWithdrawQuotePenalty", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWithdrawQuotePenalty is a free data retrieval call binding the contract method 0x0c9f7bd0.
//
// Solidity: function getWithdrawQuotePenalty(uint256 amount) view returns(uint256 penalty)
func (_Contract *ContractSession) GetWithdrawQuotePenalty(amount *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetWithdrawQuotePenalty(&_Contract.CallOpts, amount)
}

// GetWithdrawQuotePenalty is a free data retrieval call binding the contract method 0x0c9f7bd0.
//
// Solidity: function getWithdrawQuotePenalty(uint256 amount) view returns(uint256 penalty)
func (_Contract *ContractCallerSession) GetWithdrawQuotePenalty(amount *big.Int) (*big.Int, error) {
	return _Contract.Contract.GetWithdrawQuotePenalty(&_Contract.CallOpts, amount)
}

// QueryBuyBaseToken is a free data retrieval call binding the contract method 0x18c0bbe4.
//
// Solidity: function queryBuyBaseToken(uint256 amount) view returns(uint256 payQuote)
func (_Contract *ContractCaller) QueryBuyBaseToken(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "queryBuyBaseToken", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryBuyBaseToken is a free data retrieval call binding the contract method 0x18c0bbe4.
//
// Solidity: function queryBuyBaseToken(uint256 amount) view returns(uint256 payQuote)
func (_Contract *ContractSession) QueryBuyBaseToken(amount *big.Int) (*big.Int, error) {
	return _Contract.Contract.QueryBuyBaseToken(&_Contract.CallOpts, amount)
}

// QueryBuyBaseToken is a free data retrieval call binding the contract method 0x18c0bbe4.
//
// Solidity: function queryBuyBaseToken(uint256 amount) view returns(uint256 payQuote)
func (_Contract *ContractCallerSession) QueryBuyBaseToken(amount *big.Int) (*big.Int, error) {
	return _Contract.Contract.QueryBuyBaseToken(&_Contract.CallOpts, amount)
}

// QuerySellBaseToken is a free data retrieval call binding the contract method 0xa2801e16.
//
// Solidity: function querySellBaseToken(uint256 amount) view returns(uint256 receiveQuote)
func (_Contract *ContractCaller) QuerySellBaseToken(opts *bind.CallOpts, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "querySellBaseToken", amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuerySellBaseToken is a free data retrieval call binding the contract method 0xa2801e16.
//
// Solidity: function querySellBaseToken(uint256 amount) view returns(uint256 receiveQuote)
func (_Contract *ContractSession) QuerySellBaseToken(amount *big.Int) (*big.Int, error) {
	return _Contract.Contract.QuerySellBaseToken(&_Contract.CallOpts, amount)
}

// QuerySellBaseToken is a free data retrieval call binding the contract method 0xa2801e16.
//
// Solidity: function querySellBaseToken(uint256 amount) view returns(uint256 receiveQuote)
func (_Contract *ContractCallerSession) QuerySellBaseToken(amount *big.Int) (*big.Int, error) {
	return _Contract.Contract.QuerySellBaseToken(&_Contract.CallOpts, amount)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Contract *ContractCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Contract *ContractSession) Version() (*big.Int, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint256)
func (_Contract *ContractCallerSession) Version() (*big.Int, error) {
	return _Contract.Contract.Version(&_Contract.CallOpts)
}

// BuyBaseToken is a paid mutator transaction binding the contract method 0xe67ce706.
//
// Solidity: function buyBaseToken(uint256 amount, uint256 maxPayQuote, bytes data) returns(uint256)
func (_Contract *ContractTransactor) BuyBaseToken(opts *bind.TransactOpts, amount *big.Int, maxPayQuote *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "buyBaseToken", amount, maxPayQuote, data)
}

// BuyBaseToken is a paid mutator transaction binding the contract method 0xe67ce706.
//
// Solidity: function buyBaseToken(uint256 amount, uint256 maxPayQuote, bytes data) returns(uint256)
func (_Contract *ContractSession) BuyBaseToken(amount *big.Int, maxPayQuote *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.BuyBaseToken(&_Contract.TransactOpts, amount, maxPayQuote, data)
}

// BuyBaseToken is a paid mutator transaction binding the contract method 0xe67ce706.
//
// Solidity: function buyBaseToken(uint256 amount, uint256 maxPayQuote, bytes data) returns(uint256)
func (_Contract *ContractTransactorSession) BuyBaseToken(amount *big.Int, maxPayQuote *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.BuyBaseToken(&_Contract.TransactOpts, amount, maxPayQuote, data)
}

// ClaimAssets is a paid mutator transaction binding the contract method 0x1f3c156e.
//
// Solidity: function claimAssets() returns()
func (_Contract *ContractTransactor) ClaimAssets(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimAssets")
}

// ClaimAssets is a paid mutator transaction binding the contract method 0x1f3c156e.
//
// Solidity: function claimAssets() returns()
func (_Contract *ContractSession) ClaimAssets() (*types.Transaction, error) {
	return _Contract.Contract.ClaimAssets(&_Contract.TransactOpts)
}

// ClaimAssets is a paid mutator transaction binding the contract method 0x1f3c156e.
//
// Solidity: function claimAssets() returns()
func (_Contract *ContractTransactorSession) ClaimAssets() (*types.Transaction, error) {
	return _Contract.Contract.ClaimAssets(&_Contract.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Contract *ContractTransactor) ClaimOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "claimOwnership")
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Contract *ContractSession) ClaimOwnership() (*types.Transaction, error) {
	return _Contract.Contract.ClaimOwnership(&_Contract.TransactOpts)
}

// ClaimOwnership is a paid mutator transaction binding the contract method 0x4e71e0c8.
//
// Solidity: function claimOwnership() returns()
func (_Contract *ContractTransactorSession) ClaimOwnership() (*types.Transaction, error) {
	return _Contract.Contract.ClaimOwnership(&_Contract.TransactOpts)
}

// DepositBase is a paid mutator transaction binding the contract method 0x27bed8ee.
//
// Solidity: function depositBase(uint256 amount) returns(uint256)
func (_Contract *ContractTransactor) DepositBase(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "depositBase", amount)
}

// DepositBase is a paid mutator transaction binding the contract method 0x27bed8ee.
//
// Solidity: function depositBase(uint256 amount) returns(uint256)
func (_Contract *ContractSession) DepositBase(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DepositBase(&_Contract.TransactOpts, amount)
}

// DepositBase is a paid mutator transaction binding the contract method 0x27bed8ee.
//
// Solidity: function depositBase(uint256 amount) returns(uint256)
func (_Contract *ContractTransactorSession) DepositBase(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DepositBase(&_Contract.TransactOpts, amount)
}

// DepositBaseTo is a paid mutator transaction binding the contract method 0xaa06ce9b.
//
// Solidity: function depositBaseTo(address to, uint256 amount) returns(uint256)
func (_Contract *ContractTransactor) DepositBaseTo(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "depositBaseTo", to, amount)
}

// DepositBaseTo is a paid mutator transaction binding the contract method 0xaa06ce9b.
//
// Solidity: function depositBaseTo(address to, uint256 amount) returns(uint256)
func (_Contract *ContractSession) DepositBaseTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DepositBaseTo(&_Contract.TransactOpts, to, amount)
}

// DepositBaseTo is a paid mutator transaction binding the contract method 0xaa06ce9b.
//
// Solidity: function depositBaseTo(address to, uint256 amount) returns(uint256)
func (_Contract *ContractTransactorSession) DepositBaseTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DepositBaseTo(&_Contract.TransactOpts, to, amount)
}

// DepositQuote is a paid mutator transaction binding the contract method 0xf3ae6c5f.
//
// Solidity: function depositQuote(uint256 amount) returns(uint256)
func (_Contract *ContractTransactor) DepositQuote(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "depositQuote", amount)
}

// DepositQuote is a paid mutator transaction binding the contract method 0xf3ae6c5f.
//
// Solidity: function depositQuote(uint256 amount) returns(uint256)
func (_Contract *ContractSession) DepositQuote(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DepositQuote(&_Contract.TransactOpts, amount)
}

// DepositQuote is a paid mutator transaction binding the contract method 0xf3ae6c5f.
//
// Solidity: function depositQuote(uint256 amount) returns(uint256)
func (_Contract *ContractTransactorSession) DepositQuote(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DepositQuote(&_Contract.TransactOpts, amount)
}

// DepositQuoteTo is a paid mutator transaction binding the contract method 0x5f179f64.
//
// Solidity: function depositQuoteTo(address to, uint256 amount) returns(uint256)
func (_Contract *ContractTransactor) DepositQuoteTo(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "depositQuoteTo", to, amount)
}

// DepositQuoteTo is a paid mutator transaction binding the contract method 0x5f179f64.
//
// Solidity: function depositQuoteTo(address to, uint256 amount) returns(uint256)
func (_Contract *ContractSession) DepositQuoteTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DepositQuoteTo(&_Contract.TransactOpts, to, amount)
}

// DepositQuoteTo is a paid mutator transaction binding the contract method 0x5f179f64.
//
// Solidity: function depositQuoteTo(address to, uint256 amount) returns(uint256)
func (_Contract *ContractTransactorSession) DepositQuoteTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DepositQuoteTo(&_Contract.TransactOpts, to, amount)
}

// DisableBaseDeposit is a paid mutator transaction binding the contract method 0x13c57624.
//
// Solidity: function disableBaseDeposit() returns()
func (_Contract *ContractTransactor) DisableBaseDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "disableBaseDeposit")
}

// DisableBaseDeposit is a paid mutator transaction binding the contract method 0x13c57624.
//
// Solidity: function disableBaseDeposit() returns()
func (_Contract *ContractSession) DisableBaseDeposit() (*types.Transaction, error) {
	return _Contract.Contract.DisableBaseDeposit(&_Contract.TransactOpts)
}

// DisableBaseDeposit is a paid mutator transaction binding the contract method 0x13c57624.
//
// Solidity: function disableBaseDeposit() returns()
func (_Contract *ContractTransactorSession) DisableBaseDeposit() (*types.Transaction, error) {
	return _Contract.Contract.DisableBaseDeposit(&_Contract.TransactOpts)
}

// DisableQuoteDeposit is a paid mutator transaction binding the contract method 0xbc7d679d.
//
// Solidity: function disableQuoteDeposit() returns()
func (_Contract *ContractTransactor) DisableQuoteDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "disableQuoteDeposit")
}

// DisableQuoteDeposit is a paid mutator transaction binding the contract method 0xbc7d679d.
//
// Solidity: function disableQuoteDeposit() returns()
func (_Contract *ContractSession) DisableQuoteDeposit() (*types.Transaction, error) {
	return _Contract.Contract.DisableQuoteDeposit(&_Contract.TransactOpts)
}

// DisableQuoteDeposit is a paid mutator transaction binding the contract method 0xbc7d679d.
//
// Solidity: function disableQuoteDeposit() returns()
func (_Contract *ContractTransactorSession) DisableQuoteDeposit() (*types.Transaction, error) {
	return _Contract.Contract.DisableQuoteDeposit(&_Contract.TransactOpts)
}

// DisableTrading is a paid mutator transaction binding the contract method 0x17700f01.
//
// Solidity: function disableTrading() returns()
func (_Contract *ContractTransactor) DisableTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "disableTrading")
}

// DisableTrading is a paid mutator transaction binding the contract method 0x17700f01.
//
// Solidity: function disableTrading() returns()
func (_Contract *ContractSession) DisableTrading() (*types.Transaction, error) {
	return _Contract.Contract.DisableTrading(&_Contract.TransactOpts)
}

// DisableTrading is a paid mutator transaction binding the contract method 0x17700f01.
//
// Solidity: function disableTrading() returns()
func (_Contract *ContractTransactorSession) DisableTrading() (*types.Transaction, error) {
	return _Contract.Contract.DisableTrading(&_Contract.TransactOpts)
}

// DonateBaseToken is a paid mutator transaction binding the contract method 0xed0aa428.
//
// Solidity: function donateBaseToken(uint256 amount) returns()
func (_Contract *ContractTransactor) DonateBaseToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "donateBaseToken", amount)
}

// DonateBaseToken is a paid mutator transaction binding the contract method 0xed0aa428.
//
// Solidity: function donateBaseToken(uint256 amount) returns()
func (_Contract *ContractSession) DonateBaseToken(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DonateBaseToken(&_Contract.TransactOpts, amount)
}

// DonateBaseToken is a paid mutator transaction binding the contract method 0xed0aa428.
//
// Solidity: function donateBaseToken(uint256 amount) returns()
func (_Contract *ContractTransactorSession) DonateBaseToken(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DonateBaseToken(&_Contract.TransactOpts, amount)
}

// DonateQuoteToken is a paid mutator transaction binding the contract method 0x387b0c11.
//
// Solidity: function donateQuoteToken(uint256 amount) returns()
func (_Contract *ContractTransactor) DonateQuoteToken(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "donateQuoteToken", amount)
}

// DonateQuoteToken is a paid mutator transaction binding the contract method 0x387b0c11.
//
// Solidity: function donateQuoteToken(uint256 amount) returns()
func (_Contract *ContractSession) DonateQuoteToken(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DonateQuoteToken(&_Contract.TransactOpts, amount)
}

// DonateQuoteToken is a paid mutator transaction binding the contract method 0x387b0c11.
//
// Solidity: function donateQuoteToken(uint256 amount) returns()
func (_Contract *ContractTransactorSession) DonateQuoteToken(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.DonateQuoteToken(&_Contract.TransactOpts, amount)
}

// EnableBaseDeposit is a paid mutator transaction binding the contract method 0x1184d8be.
//
// Solidity: function enableBaseDeposit() returns()
func (_Contract *ContractTransactor) EnableBaseDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "enableBaseDeposit")
}

// EnableBaseDeposit is a paid mutator transaction binding the contract method 0x1184d8be.
//
// Solidity: function enableBaseDeposit() returns()
func (_Contract *ContractSession) EnableBaseDeposit() (*types.Transaction, error) {
	return _Contract.Contract.EnableBaseDeposit(&_Contract.TransactOpts)
}

// EnableBaseDeposit is a paid mutator transaction binding the contract method 0x1184d8be.
//
// Solidity: function enableBaseDeposit() returns()
func (_Contract *ContractTransactorSession) EnableBaseDeposit() (*types.Transaction, error) {
	return _Contract.Contract.EnableBaseDeposit(&_Contract.TransactOpts)
}

// EnableQuoteDeposit is a paid mutator transaction binding the contract method 0x36ac41a8.
//
// Solidity: function enableQuoteDeposit() returns()
func (_Contract *ContractTransactor) EnableQuoteDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "enableQuoteDeposit")
}

// EnableQuoteDeposit is a paid mutator transaction binding the contract method 0x36ac41a8.
//
// Solidity: function enableQuoteDeposit() returns()
func (_Contract *ContractSession) EnableQuoteDeposit() (*types.Transaction, error) {
	return _Contract.Contract.EnableQuoteDeposit(&_Contract.TransactOpts)
}

// EnableQuoteDeposit is a paid mutator transaction binding the contract method 0x36ac41a8.
//
// Solidity: function enableQuoteDeposit() returns()
func (_Contract *ContractTransactorSession) EnableQuoteDeposit() (*types.Transaction, error) {
	return _Contract.Contract.EnableQuoteDeposit(&_Contract.TransactOpts)
}

// EnableTrading is a paid mutator transaction binding the contract method 0x8a8c523c.
//
// Solidity: function enableTrading() returns()
func (_Contract *ContractTransactor) EnableTrading(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "enableTrading")
}

// EnableTrading is a paid mutator transaction binding the contract method 0x8a8c523c.
//
// Solidity: function enableTrading() returns()
func (_Contract *ContractSession) EnableTrading() (*types.Transaction, error) {
	return _Contract.Contract.EnableTrading(&_Contract.TransactOpts)
}

// EnableTrading is a paid mutator transaction binding the contract method 0x8a8c523c.
//
// Solidity: function enableTrading() returns()
func (_Contract *ContractTransactorSession) EnableTrading() (*types.Transaction, error) {
	return _Contract.Contract.EnableTrading(&_Contract.TransactOpts)
}

// FinalSettlement is a paid mutator transaction binding the contract method 0x648a4fac.
//
// Solidity: function finalSettlement() returns()
func (_Contract *ContractTransactor) FinalSettlement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "finalSettlement")
}

// FinalSettlement is a paid mutator transaction binding the contract method 0x648a4fac.
//
// Solidity: function finalSettlement() returns()
func (_Contract *ContractSession) FinalSettlement() (*types.Transaction, error) {
	return _Contract.Contract.FinalSettlement(&_Contract.TransactOpts)
}

// FinalSettlement is a paid mutator transaction binding the contract method 0x648a4fac.
//
// Solidity: function finalSettlement() returns()
func (_Contract *ContractTransactorSession) FinalSettlement() (*types.Transaction, error) {
	return _Contract.Contract.FinalSettlement(&_Contract.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xf00f9a09.
//
// Solidity: function init(address owner, address supervisor, address maintainer, address baseToken, address quoteToken, address oracle, uint256 lpFeeRate, uint256 mtFeeRate, uint256 k, uint256 gasPriceLimit) returns()
func (_Contract *ContractTransactor) Init(opts *bind.TransactOpts, owner common.Address, supervisor common.Address, maintainer common.Address, baseToken common.Address, quoteToken common.Address, oracle common.Address, lpFeeRate *big.Int, mtFeeRate *big.Int, k *big.Int, gasPriceLimit *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "init", owner, supervisor, maintainer, baseToken, quoteToken, oracle, lpFeeRate, mtFeeRate, k, gasPriceLimit)
}

// Init is a paid mutator transaction binding the contract method 0xf00f9a09.
//
// Solidity: function init(address owner, address supervisor, address maintainer, address baseToken, address quoteToken, address oracle, uint256 lpFeeRate, uint256 mtFeeRate, uint256 k, uint256 gasPriceLimit) returns()
func (_Contract *ContractSession) Init(owner common.Address, supervisor common.Address, maintainer common.Address, baseToken common.Address, quoteToken common.Address, oracle common.Address, lpFeeRate *big.Int, mtFeeRate *big.Int, k *big.Int, gasPriceLimit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Init(&_Contract.TransactOpts, owner, supervisor, maintainer, baseToken, quoteToken, oracle, lpFeeRate, mtFeeRate, k, gasPriceLimit)
}

// Init is a paid mutator transaction binding the contract method 0xf00f9a09.
//
// Solidity: function init(address owner, address supervisor, address maintainer, address baseToken, address quoteToken, address oracle, uint256 lpFeeRate, uint256 mtFeeRate, uint256 k, uint256 gasPriceLimit) returns()
func (_Contract *ContractTransactorSession) Init(owner common.Address, supervisor common.Address, maintainer common.Address, baseToken common.Address, quoteToken common.Address, oracle common.Address, lpFeeRate *big.Int, mtFeeRate *big.Int, k *big.Int, gasPriceLimit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Init(&_Contract.TransactOpts, owner, supervisor, maintainer, baseToken, quoteToken, oracle, lpFeeRate, mtFeeRate, k, gasPriceLimit)
}

// Retrieve is a paid mutator transaction binding the contract method 0xc3a2a665.
//
// Solidity: function retrieve(address token, uint256 amount) returns()
func (_Contract *ContractTransactor) Retrieve(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "retrieve", token, amount)
}

// Retrieve is a paid mutator transaction binding the contract method 0xc3a2a665.
//
// Solidity: function retrieve(address token, uint256 amount) returns()
func (_Contract *ContractSession) Retrieve(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Retrieve(&_Contract.TransactOpts, token, amount)
}

// Retrieve is a paid mutator transaction binding the contract method 0xc3a2a665.
//
// Solidity: function retrieve(address token, uint256 amount) returns()
func (_Contract *ContractTransactorSession) Retrieve(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Retrieve(&_Contract.TransactOpts, token, amount)
}

// SellBaseToken is a paid mutator transaction binding the contract method 0x8dae7333.
//
// Solidity: function sellBaseToken(uint256 amount, uint256 minReceiveQuote, bytes data) returns(uint256)
func (_Contract *ContractTransactor) SellBaseToken(opts *bind.TransactOpts, amount *big.Int, minReceiveQuote *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "sellBaseToken", amount, minReceiveQuote, data)
}

// SellBaseToken is a paid mutator transaction binding the contract method 0x8dae7333.
//
// Solidity: function sellBaseToken(uint256 amount, uint256 minReceiveQuote, bytes data) returns(uint256)
func (_Contract *ContractSession) SellBaseToken(amount *big.Int, minReceiveQuote *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SellBaseToken(&_Contract.TransactOpts, amount, minReceiveQuote, data)
}

// SellBaseToken is a paid mutator transaction binding the contract method 0x8dae7333.
//
// Solidity: function sellBaseToken(uint256 amount, uint256 minReceiveQuote, bytes data) returns(uint256)
func (_Contract *ContractTransactorSession) SellBaseToken(amount *big.Int, minReceiveQuote *big.Int, data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SellBaseToken(&_Contract.TransactOpts, amount, minReceiveQuote, data)
}

// SetGasPriceLimit is a paid mutator transaction binding the contract method 0x09231602.
//
// Solidity: function setGasPriceLimit(uint256 newGasPriceLimit) returns()
func (_Contract *ContractTransactor) SetGasPriceLimit(opts *bind.TransactOpts, newGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGasPriceLimit", newGasPriceLimit)
}

// SetGasPriceLimit is a paid mutator transaction binding the contract method 0x09231602.
//
// Solidity: function setGasPriceLimit(uint256 newGasPriceLimit) returns()
func (_Contract *ContractSession) SetGasPriceLimit(newGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGasPriceLimit(&_Contract.TransactOpts, newGasPriceLimit)
}

// SetGasPriceLimit is a paid mutator transaction binding the contract method 0x09231602.
//
// Solidity: function setGasPriceLimit(uint256 newGasPriceLimit) returns()
func (_Contract *ContractTransactorSession) SetGasPriceLimit(newGasPriceLimit *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetGasPriceLimit(&_Contract.TransactOpts, newGasPriceLimit)
}

// SetK is a paid mutator transaction binding the contract method 0x67de8be9.
//
// Solidity: function setK(uint256 newK) returns()
func (_Contract *ContractTransactor) SetK(opts *bind.TransactOpts, newK *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setK", newK)
}

// SetK is a paid mutator transaction binding the contract method 0x67de8be9.
//
// Solidity: function setK(uint256 newK) returns()
func (_Contract *ContractSession) SetK(newK *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetK(&_Contract.TransactOpts, newK)
}

// SetK is a paid mutator transaction binding the contract method 0x67de8be9.
//
// Solidity: function setK(uint256 newK) returns()
func (_Contract *ContractTransactorSession) SetK(newK *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetK(&_Contract.TransactOpts, newK)
}

// SetLiquidityProviderFeeRate is a paid mutator transaction binding the contract method 0x5bb7552a.
//
// Solidity: function setLiquidityProviderFeeRate(uint256 newLiquidityPorviderFeeRate) returns()
func (_Contract *ContractTransactor) SetLiquidityProviderFeeRate(opts *bind.TransactOpts, newLiquidityPorviderFeeRate *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLiquidityProviderFeeRate", newLiquidityPorviderFeeRate)
}

// SetLiquidityProviderFeeRate is a paid mutator transaction binding the contract method 0x5bb7552a.
//
// Solidity: function setLiquidityProviderFeeRate(uint256 newLiquidityPorviderFeeRate) returns()
func (_Contract *ContractSession) SetLiquidityProviderFeeRate(newLiquidityPorviderFeeRate *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetLiquidityProviderFeeRate(&_Contract.TransactOpts, newLiquidityPorviderFeeRate)
}

// SetLiquidityProviderFeeRate is a paid mutator transaction binding the contract method 0x5bb7552a.
//
// Solidity: function setLiquidityProviderFeeRate(uint256 newLiquidityPorviderFeeRate) returns()
func (_Contract *ContractTransactorSession) SetLiquidityProviderFeeRate(newLiquidityPorviderFeeRate *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetLiquidityProviderFeeRate(&_Contract.TransactOpts, newLiquidityPorviderFeeRate)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address newMaintainer) returns()
func (_Contract *ContractTransactor) SetMaintainer(opts *bind.TransactOpts, newMaintainer common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMaintainer", newMaintainer)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address newMaintainer) returns()
func (_Contract *ContractSession) SetMaintainer(newMaintainer common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetMaintainer(&_Contract.TransactOpts, newMaintainer)
}

// SetMaintainer is a paid mutator transaction binding the contract method 0x13ea5d29.
//
// Solidity: function setMaintainer(address newMaintainer) returns()
func (_Contract *ContractTransactorSession) SetMaintainer(newMaintainer common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetMaintainer(&_Contract.TransactOpts, newMaintainer)
}

// SetMaintainerFeeRate is a paid mutator transaction binding the contract method 0xf2220416.
//
// Solidity: function setMaintainerFeeRate(uint256 newMaintainerFeeRate) returns()
func (_Contract *ContractTransactor) SetMaintainerFeeRate(opts *bind.TransactOpts, newMaintainerFeeRate *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMaintainerFeeRate", newMaintainerFeeRate)
}

// SetMaintainerFeeRate is a paid mutator transaction binding the contract method 0xf2220416.
//
// Solidity: function setMaintainerFeeRate(uint256 newMaintainerFeeRate) returns()
func (_Contract *ContractSession) SetMaintainerFeeRate(newMaintainerFeeRate *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMaintainerFeeRate(&_Contract.TransactOpts, newMaintainerFeeRate)
}

// SetMaintainerFeeRate is a paid mutator transaction binding the contract method 0xf2220416.
//
// Solidity: function setMaintainerFeeRate(uint256 newMaintainerFeeRate) returns()
func (_Contract *ContractTransactorSession) SetMaintainerFeeRate(newMaintainerFeeRate *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMaintainerFeeRate(&_Contract.TransactOpts, newMaintainerFeeRate)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address newOracle) returns()
func (_Contract *ContractTransactor) SetOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setOracle", newOracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address newOracle) returns()
func (_Contract *ContractSession) SetOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOracle(&_Contract.TransactOpts, newOracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address newOracle) returns()
func (_Contract *ContractTransactorSession) SetOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOracle(&_Contract.TransactOpts, newOracle)
}

// SetSupervisor is a paid mutator transaction binding the contract method 0x9299eb30.
//
// Solidity: function setSupervisor(address newSupervisor) returns()
func (_Contract *ContractTransactor) SetSupervisor(opts *bind.TransactOpts, newSupervisor common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setSupervisor", newSupervisor)
}

// SetSupervisor is a paid mutator transaction binding the contract method 0x9299eb30.
//
// Solidity: function setSupervisor(address newSupervisor) returns()
func (_Contract *ContractSession) SetSupervisor(newSupervisor common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetSupervisor(&_Contract.TransactOpts, newSupervisor)
}

// SetSupervisor is a paid mutator transaction binding the contract method 0x9299eb30.
//
// Solidity: function setSupervisor(address newSupervisor) returns()
func (_Contract *ContractTransactorSession) SetSupervisor(newSupervisor common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetSupervisor(&_Contract.TransactOpts, newSupervisor)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Contract *ContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferOwnership(&_Contract.TransactOpts, newOwner)
}

// WithdrawAllBase is a paid mutator transaction binding the contract method 0xd47eaa37.
//
// Solidity: function withdrawAllBase() returns(uint256)
func (_Contract *ContractTransactor) WithdrawAllBase(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawAllBase")
}

// WithdrawAllBase is a paid mutator transaction binding the contract method 0xd47eaa37.
//
// Solidity: function withdrawAllBase() returns(uint256)
func (_Contract *ContractSession) WithdrawAllBase() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAllBase(&_Contract.TransactOpts)
}

// WithdrawAllBase is a paid mutator transaction binding the contract method 0xd47eaa37.
//
// Solidity: function withdrawAllBase() returns(uint256)
func (_Contract *ContractTransactorSession) WithdrawAllBase() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAllBase(&_Contract.TransactOpts)
}

// WithdrawAllBaseTo is a paid mutator transaction binding the contract method 0x1e34b9cc.
//
// Solidity: function withdrawAllBaseTo(address to) returns(uint256)
func (_Contract *ContractTransactor) WithdrawAllBaseTo(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawAllBaseTo", to)
}

// WithdrawAllBaseTo is a paid mutator transaction binding the contract method 0x1e34b9cc.
//
// Solidity: function withdrawAllBaseTo(address to) returns(uint256)
func (_Contract *ContractSession) WithdrawAllBaseTo(to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAllBaseTo(&_Contract.TransactOpts, to)
}

// WithdrawAllBaseTo is a paid mutator transaction binding the contract method 0x1e34b9cc.
//
// Solidity: function withdrawAllBaseTo(address to) returns(uint256)
func (_Contract *ContractTransactorSession) WithdrawAllBaseTo(to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAllBaseTo(&_Contract.TransactOpts, to)
}

// WithdrawAllQuote is a paid mutator transaction binding the contract method 0xc59203af.
//
// Solidity: function withdrawAllQuote() returns(uint256)
func (_Contract *ContractTransactor) WithdrawAllQuote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawAllQuote")
}

// WithdrawAllQuote is a paid mutator transaction binding the contract method 0xc59203af.
//
// Solidity: function withdrawAllQuote() returns(uint256)
func (_Contract *ContractSession) WithdrawAllQuote() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAllQuote(&_Contract.TransactOpts)
}

// WithdrawAllQuote is a paid mutator transaction binding the contract method 0xc59203af.
//
// Solidity: function withdrawAllQuote() returns(uint256)
func (_Contract *ContractTransactorSession) WithdrawAllQuote() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAllQuote(&_Contract.TransactOpts)
}

// WithdrawAllQuoteTo is a paid mutator transaction binding the contract method 0x04512dc4.
//
// Solidity: function withdrawAllQuoteTo(address to) returns(uint256)
func (_Contract *ContractTransactor) WithdrawAllQuoteTo(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawAllQuoteTo", to)
}

// WithdrawAllQuoteTo is a paid mutator transaction binding the contract method 0x04512dc4.
//
// Solidity: function withdrawAllQuoteTo(address to) returns(uint256)
func (_Contract *ContractSession) WithdrawAllQuoteTo(to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAllQuoteTo(&_Contract.TransactOpts, to)
}

// WithdrawAllQuoteTo is a paid mutator transaction binding the contract method 0x04512dc4.
//
// Solidity: function withdrawAllQuoteTo(address to) returns(uint256)
func (_Contract *ContractTransactorSession) WithdrawAllQuoteTo(to common.Address) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawAllQuoteTo(&_Contract.TransactOpts, to)
}

// WithdrawBase is a paid mutator transaction binding the contract method 0xf98bea15.
//
// Solidity: function withdrawBase(uint256 amount) returns(uint256)
func (_Contract *ContractTransactor) WithdrawBase(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawBase", amount)
}

// WithdrawBase is a paid mutator transaction binding the contract method 0xf98bea15.
//
// Solidity: function withdrawBase(uint256 amount) returns(uint256)
func (_Contract *ContractSession) WithdrawBase(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawBase(&_Contract.TransactOpts, amount)
}

// WithdrawBase is a paid mutator transaction binding the contract method 0xf98bea15.
//
// Solidity: function withdrawBase(uint256 amount) returns(uint256)
func (_Contract *ContractTransactorSession) WithdrawBase(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawBase(&_Contract.TransactOpts, amount)
}

// WithdrawBaseTo is a paid mutator transaction binding the contract method 0x947cf92b.
//
// Solidity: function withdrawBaseTo(address to, uint256 amount) returns(uint256)
func (_Contract *ContractTransactor) WithdrawBaseTo(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawBaseTo", to, amount)
}

// WithdrawBaseTo is a paid mutator transaction binding the contract method 0x947cf92b.
//
// Solidity: function withdrawBaseTo(address to, uint256 amount) returns(uint256)
func (_Contract *ContractSession) WithdrawBaseTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawBaseTo(&_Contract.TransactOpts, to, amount)
}

// WithdrawBaseTo is a paid mutator transaction binding the contract method 0x947cf92b.
//
// Solidity: function withdrawBaseTo(address to, uint256 amount) returns(uint256)
func (_Contract *ContractTransactorSession) WithdrawBaseTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawBaseTo(&_Contract.TransactOpts, to, amount)
}

// WithdrawQuote is a paid mutator transaction binding the contract method 0xc0a5f6ff.
//
// Solidity: function withdrawQuote(uint256 amount) returns(uint256)
func (_Contract *ContractTransactor) WithdrawQuote(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawQuote", amount)
}

// WithdrawQuote is a paid mutator transaction binding the contract method 0xc0a5f6ff.
//
// Solidity: function withdrawQuote(uint256 amount) returns(uint256)
func (_Contract *ContractSession) WithdrawQuote(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawQuote(&_Contract.TransactOpts, amount)
}

// WithdrawQuote is a paid mutator transaction binding the contract method 0xc0a5f6ff.
//
// Solidity: function withdrawQuote(uint256 amount) returns(uint256)
func (_Contract *ContractTransactorSession) WithdrawQuote(amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawQuote(&_Contract.TransactOpts, amount)
}

// WithdrawQuoteTo is a paid mutator transaction binding the contract method 0x108db744.
//
// Solidity: function withdrawQuoteTo(address to, uint256 amount) returns(uint256)
func (_Contract *ContractTransactor) WithdrawQuoteTo(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawQuoteTo", to, amount)
}

// WithdrawQuoteTo is a paid mutator transaction binding the contract method 0x108db744.
//
// Solidity: function withdrawQuoteTo(address to, uint256 amount) returns(uint256)
func (_Contract *ContractSession) WithdrawQuoteTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawQuoteTo(&_Contract.TransactOpts, to, amount)
}

// WithdrawQuoteTo is a paid mutator transaction binding the contract method 0x108db744.
//
// Solidity: function withdrawQuoteTo(address to, uint256 amount) returns(uint256)
func (_Contract *ContractTransactorSession) WithdrawQuoteTo(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.WithdrawQuoteTo(&_Contract.TransactOpts, to, amount)
}

// ContractBuyBaseTokenIterator is returned from FilterBuyBaseToken and is used to iterate over the raw logs and unpacked data for BuyBaseToken events raised by the Contract contract.
type ContractBuyBaseTokenIterator struct {
	Event *ContractBuyBaseToken // Event containing the contract specifics and raw log

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
func (it *ContractBuyBaseTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBuyBaseToken)
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
		it.Event = new(ContractBuyBaseToken)
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
func (it *ContractBuyBaseTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBuyBaseTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBuyBaseToken represents a BuyBaseToken event raised by the Contract contract.
type ContractBuyBaseToken struct {
	Buyer       common.Address
	ReceiveBase *big.Int
	PayQuote    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBuyBaseToken is a free log retrieval operation binding the contract event 0xe93ad76094f247c0dafc1c61adc2187de1ac2738f7a3b49cb20b2263420251a3.
//
// Solidity: event BuyBaseToken(address indexed buyer, uint256 receiveBase, uint256 payQuote)
func (_Contract *ContractFilterer) FilterBuyBaseToken(opts *bind.FilterOpts, buyer []common.Address) (*ContractBuyBaseTokenIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "BuyBaseToken", buyerRule)
	if err != nil {
		return nil, err
	}
	return &ContractBuyBaseTokenIterator{contract: _Contract.contract, event: "BuyBaseToken", logs: logs, sub: sub}, nil
}

// WatchBuyBaseToken is a free log subscription operation binding the contract event 0xe93ad76094f247c0dafc1c61adc2187de1ac2738f7a3b49cb20b2263420251a3.
//
// Solidity: event BuyBaseToken(address indexed buyer, uint256 receiveBase, uint256 payQuote)
func (_Contract *ContractFilterer) WatchBuyBaseToken(opts *bind.WatchOpts, sink chan<- *ContractBuyBaseToken, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "BuyBaseToken", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBuyBaseToken)
				if err := _Contract.contract.UnpackLog(event, "BuyBaseToken", log); err != nil {
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

// ParseBuyBaseToken is a log parse operation binding the contract event 0xe93ad76094f247c0dafc1c61adc2187de1ac2738f7a3b49cb20b2263420251a3.
//
// Solidity: event BuyBaseToken(address indexed buyer, uint256 receiveBase, uint256 payQuote)
func (_Contract *ContractFilterer) ParseBuyBaseToken(log types.Log) (*ContractBuyBaseToken, error) {
	event := new(ContractBuyBaseToken)
	if err := _Contract.contract.UnpackLog(event, "BuyBaseToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractChargeMaintainerFeeIterator is returned from FilterChargeMaintainerFee and is used to iterate over the raw logs and unpacked data for ChargeMaintainerFee events raised by the Contract contract.
type ContractChargeMaintainerFeeIterator struct {
	Event *ContractChargeMaintainerFee // Event containing the contract specifics and raw log

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
func (it *ContractChargeMaintainerFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractChargeMaintainerFee)
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
		it.Event = new(ContractChargeMaintainerFee)
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
func (it *ContractChargeMaintainerFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractChargeMaintainerFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractChargeMaintainerFee represents a ChargeMaintainerFee event raised by the Contract contract.
type ContractChargeMaintainerFee struct {
	Maintainer  common.Address
	IsBaseToken bool
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChargeMaintainerFee is a free log retrieval operation binding the contract event 0xe4fed5362e2669c70e5da5a18942d1e617d8917f6adc0164d9668bd3a6d0cebe.
//
// Solidity: event ChargeMaintainerFee(address indexed maintainer, bool isBaseToken, uint256 amount)
func (_Contract *ContractFilterer) FilterChargeMaintainerFee(opts *bind.FilterOpts, maintainer []common.Address) (*ContractChargeMaintainerFeeIterator, error) {

	var maintainerRule []interface{}
	for _, maintainerItem := range maintainer {
		maintainerRule = append(maintainerRule, maintainerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ChargeMaintainerFee", maintainerRule)
	if err != nil {
		return nil, err
	}
	return &ContractChargeMaintainerFeeIterator{contract: _Contract.contract, event: "ChargeMaintainerFee", logs: logs, sub: sub}, nil
}

// WatchChargeMaintainerFee is a free log subscription operation binding the contract event 0xe4fed5362e2669c70e5da5a18942d1e617d8917f6adc0164d9668bd3a6d0cebe.
//
// Solidity: event ChargeMaintainerFee(address indexed maintainer, bool isBaseToken, uint256 amount)
func (_Contract *ContractFilterer) WatchChargeMaintainerFee(opts *bind.WatchOpts, sink chan<- *ContractChargeMaintainerFee, maintainer []common.Address) (event.Subscription, error) {

	var maintainerRule []interface{}
	for _, maintainerItem := range maintainer {
		maintainerRule = append(maintainerRule, maintainerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ChargeMaintainerFee", maintainerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractChargeMaintainerFee)
				if err := _Contract.contract.UnpackLog(event, "ChargeMaintainerFee", log); err != nil {
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

// ParseChargeMaintainerFee is a log parse operation binding the contract event 0xe4fed5362e2669c70e5da5a18942d1e617d8917f6adc0164d9668bd3a6d0cebe.
//
// Solidity: event ChargeMaintainerFee(address indexed maintainer, bool isBaseToken, uint256 amount)
func (_Contract *ContractFilterer) ParseChargeMaintainerFee(log types.Log) (*ContractChargeMaintainerFee, error) {
	event := new(ContractChargeMaintainerFee)
	if err := _Contract.contract.UnpackLog(event, "ChargeMaintainerFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractChargePenaltyIterator is returned from FilterChargePenalty and is used to iterate over the raw logs and unpacked data for ChargePenalty events raised by the Contract contract.
type ContractChargePenaltyIterator struct {
	Event *ContractChargePenalty // Event containing the contract specifics and raw log

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
func (it *ContractChargePenaltyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractChargePenalty)
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
		it.Event = new(ContractChargePenalty)
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
func (it *ContractChargePenaltyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractChargePenaltyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractChargePenalty represents a ChargePenalty event raised by the Contract contract.
type ContractChargePenalty struct {
	Payer       common.Address
	IsBaseToken bool
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChargePenalty is a free log retrieval operation binding the contract event 0x581f351e2bdb9fa9021bb2a24def989f06ac236f8a92aac14bcbc618ddf3826a.
//
// Solidity: event ChargePenalty(address indexed payer, bool isBaseToken, uint256 amount)
func (_Contract *ContractFilterer) FilterChargePenalty(opts *bind.FilterOpts, payer []common.Address) (*ContractChargePenaltyIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ChargePenalty", payerRule)
	if err != nil {
		return nil, err
	}
	return &ContractChargePenaltyIterator{contract: _Contract.contract, event: "ChargePenalty", logs: logs, sub: sub}, nil
}

// WatchChargePenalty is a free log subscription operation binding the contract event 0x581f351e2bdb9fa9021bb2a24def989f06ac236f8a92aac14bcbc618ddf3826a.
//
// Solidity: event ChargePenalty(address indexed payer, bool isBaseToken, uint256 amount)
func (_Contract *ContractFilterer) WatchChargePenalty(opts *bind.WatchOpts, sink chan<- *ContractChargePenalty, payer []common.Address) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ChargePenalty", payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractChargePenalty)
				if err := _Contract.contract.UnpackLog(event, "ChargePenalty", log); err != nil {
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

// ParseChargePenalty is a log parse operation binding the contract event 0x581f351e2bdb9fa9021bb2a24def989f06ac236f8a92aac14bcbc618ddf3826a.
//
// Solidity: event ChargePenalty(address indexed payer, bool isBaseToken, uint256 amount)
func (_Contract *ContractFilterer) ParseChargePenalty(log types.Log) (*ContractChargePenalty, error) {
	event := new(ContractChargePenalty)
	if err := _Contract.contract.UnpackLog(event, "ChargePenalty", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractClaimAssetsIterator is returned from FilterClaimAssets and is used to iterate over the raw logs and unpacked data for ClaimAssets events raised by the Contract contract.
type ContractClaimAssetsIterator struct {
	Event *ContractClaimAssets // Event containing the contract specifics and raw log

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
func (it *ContractClaimAssetsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractClaimAssets)
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
		it.Event = new(ContractClaimAssets)
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
func (it *ContractClaimAssetsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractClaimAssetsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractClaimAssets represents a ClaimAssets event raised by the Contract contract.
type ContractClaimAssets struct {
	User             common.Address
	BaseTokenAmount  *big.Int
	QuoteTokenAmount *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClaimAssets is a free log retrieval operation binding the contract event 0xbe5f7fe66d16c6a87bb5b8b08a96634fe4f1c2bac9e5e413efe41a782d4d0c43.
//
// Solidity: event ClaimAssets(address indexed user, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Contract *ContractFilterer) FilterClaimAssets(opts *bind.FilterOpts, user []common.Address) (*ContractClaimAssetsIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ClaimAssets", userRule)
	if err != nil {
		return nil, err
	}
	return &ContractClaimAssetsIterator{contract: _Contract.contract, event: "ClaimAssets", logs: logs, sub: sub}, nil
}

// WatchClaimAssets is a free log subscription operation binding the contract event 0xbe5f7fe66d16c6a87bb5b8b08a96634fe4f1c2bac9e5e413efe41a782d4d0c43.
//
// Solidity: event ClaimAssets(address indexed user, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Contract *ContractFilterer) WatchClaimAssets(opts *bind.WatchOpts, sink chan<- *ContractClaimAssets, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ClaimAssets", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractClaimAssets)
				if err := _Contract.contract.UnpackLog(event, "ClaimAssets", log); err != nil {
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

// ParseClaimAssets is a log parse operation binding the contract event 0xbe5f7fe66d16c6a87bb5b8b08a96634fe4f1c2bac9e5e413efe41a782d4d0c43.
//
// Solidity: event ClaimAssets(address indexed user, uint256 baseTokenAmount, uint256 quoteTokenAmount)
func (_Contract *ContractFilterer) ParseClaimAssets(log types.Log) (*ContractClaimAssets, error) {
	event := new(ContractClaimAssets)
	if err := _Contract.contract.UnpackLog(event, "ClaimAssets", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Contract contract.
type ContractDepositIterator struct {
	Event *ContractDeposit // Event containing the contract specifics and raw log

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
func (it *ContractDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDeposit)
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
		it.Event = new(ContractDeposit)
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
func (it *ContractDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDeposit represents a Deposit event raised by the Contract contract.
type ContractDeposit struct {
	Payer         common.Address
	Receiver      common.Address
	IsBaseToken   bool
	Amount        *big.Int
	LpTokenAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x18081cde2fa64894914e1080b98cca17bb6d1acf633e57f6e26ebdb945ad830b.
//
// Solidity: event Deposit(address indexed payer, address indexed receiver, bool isBaseToken, uint256 amount, uint256 lpTokenAmount)
func (_Contract *ContractFilterer) FilterDeposit(opts *bind.FilterOpts, payer []common.Address, receiver []common.Address) (*ContractDepositIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Deposit", payerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ContractDepositIterator{contract: _Contract.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x18081cde2fa64894914e1080b98cca17bb6d1acf633e57f6e26ebdb945ad830b.
//
// Solidity: event Deposit(address indexed payer, address indexed receiver, bool isBaseToken, uint256 amount, uint256 lpTokenAmount)
func (_Contract *ContractFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ContractDeposit, payer []common.Address, receiver []common.Address) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Deposit", payerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDeposit)
				if err := _Contract.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x18081cde2fa64894914e1080b98cca17bb6d1acf633e57f6e26ebdb945ad830b.
//
// Solidity: event Deposit(address indexed payer, address indexed receiver, bool isBaseToken, uint256 amount, uint256 lpTokenAmount)
func (_Contract *ContractFilterer) ParseDeposit(log types.Log) (*ContractDeposit, error) {
	event := new(ContractDeposit)
	if err := _Contract.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractDonateIterator is returned from FilterDonate and is used to iterate over the raw logs and unpacked data for Donate events raised by the Contract contract.
type ContractDonateIterator struct {
	Event *ContractDonate // Event containing the contract specifics and raw log

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
func (it *ContractDonateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractDonate)
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
		it.Event = new(ContractDonate)
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
func (it *ContractDonateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractDonateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractDonate represents a Donate event raised by the Contract contract.
type ContractDonate struct {
	Amount      *big.Int
	IsBaseToken bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDonate is a free log retrieval operation binding the contract event 0xa259c93818139b6bc90fb80e8feb75122b42edaae49560f81392cf4e1946726e.
//
// Solidity: event Donate(uint256 amount, bool isBaseToken)
func (_Contract *ContractFilterer) FilterDonate(opts *bind.FilterOpts) (*ContractDonateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Donate")
	if err != nil {
		return nil, err
	}
	return &ContractDonateIterator{contract: _Contract.contract, event: "Donate", logs: logs, sub: sub}, nil
}

// WatchDonate is a free log subscription operation binding the contract event 0xa259c93818139b6bc90fb80e8feb75122b42edaae49560f81392cf4e1946726e.
//
// Solidity: event Donate(uint256 amount, bool isBaseToken)
func (_Contract *ContractFilterer) WatchDonate(opts *bind.WatchOpts, sink chan<- *ContractDonate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Donate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractDonate)
				if err := _Contract.contract.UnpackLog(event, "Donate", log); err != nil {
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

// ParseDonate is a log parse operation binding the contract event 0xa259c93818139b6bc90fb80e8feb75122b42edaae49560f81392cf4e1946726e.
//
// Solidity: event Donate(uint256 amount, bool isBaseToken)
func (_Contract *ContractFilterer) ParseDonate(log types.Log) (*ContractDonate, error) {
	event := new(ContractDonate)
	if err := _Contract.contract.UnpackLog(event, "Donate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferPreparedIterator is returned from FilterOwnershipTransferPrepared and is used to iterate over the raw logs and unpacked data for OwnershipTransferPrepared events raised by the Contract contract.
type ContractOwnershipTransferPreparedIterator struct {
	Event *ContractOwnershipTransferPrepared // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferPrepared)
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
		it.Event = new(ContractOwnershipTransferPrepared)
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
func (it *ContractOwnershipTransferPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferPrepared represents a OwnershipTransferPrepared event raised by the Contract contract.
type ContractOwnershipTransferPrepared struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferPrepared is a free log retrieval operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferPrepared(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferPreparedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferPreparedIterator{contract: _Contract.contract, event: "OwnershipTransferPrepared", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferPrepared is a free log subscription operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferPrepared(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferPrepared, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferPrepared", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferPrepared)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
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

// ParseOwnershipTransferPrepared is a log parse operation binding the contract event 0xdcf55418cee3220104fef63f979ff3c4097ad240c0c43dcb33ce837748983e62.
//
// Solidity: event OwnershipTransferPrepared(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) ParseOwnershipTransferPrepared(log types.Log) (*ContractOwnershipTransferPrepared, error) {
	event := new(ContractOwnershipTransferPrepared)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Contract contract.
type ContractOwnershipTransferredIterator struct {
	Event *ContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractOwnershipTransferred)
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
		it.Event = new(ContractOwnershipTransferred)
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
func (it *ContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractOwnershipTransferred represents a OwnershipTransferred event raised by the Contract contract.
type ContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractOwnershipTransferredIterator{contract: _Contract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Contract *ContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractOwnershipTransferred)
				if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Contract *ContractFilterer) ParseOwnershipTransferred(log types.Log) (*ContractOwnershipTransferred, error) {
	event := new(ContractOwnershipTransferred)
	if err := _Contract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractSellBaseTokenIterator is returned from FilterSellBaseToken and is used to iterate over the raw logs and unpacked data for SellBaseToken events raised by the Contract contract.
type ContractSellBaseTokenIterator struct {
	Event *ContractSellBaseToken // Event containing the contract specifics and raw log

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
func (it *ContractSellBaseTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSellBaseToken)
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
		it.Event = new(ContractSellBaseToken)
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
func (it *ContractSellBaseTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSellBaseTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSellBaseToken represents a SellBaseToken event raised by the Contract contract.
type ContractSellBaseToken struct {
	Seller       common.Address
	PayBase      *big.Int
	ReceiveQuote *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSellBaseToken is a free log retrieval operation binding the contract event 0xd8648b6ac54162763c86fd54bf2005af8ecd2f9cb273a5775921fd7f91e17b2d.
//
// Solidity: event SellBaseToken(address indexed seller, uint256 payBase, uint256 receiveQuote)
func (_Contract *ContractFilterer) FilterSellBaseToken(opts *bind.FilterOpts, seller []common.Address) (*ContractSellBaseTokenIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SellBaseToken", sellerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSellBaseTokenIterator{contract: _Contract.contract, event: "SellBaseToken", logs: logs, sub: sub}, nil
}

// WatchSellBaseToken is a free log subscription operation binding the contract event 0xd8648b6ac54162763c86fd54bf2005af8ecd2f9cb273a5775921fd7f91e17b2d.
//
// Solidity: event SellBaseToken(address indexed seller, uint256 payBase, uint256 receiveQuote)
func (_Contract *ContractFilterer) WatchSellBaseToken(opts *bind.WatchOpts, sink chan<- *ContractSellBaseToken, seller []common.Address) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SellBaseToken", sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSellBaseToken)
				if err := _Contract.contract.UnpackLog(event, "SellBaseToken", log); err != nil {
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

// ParseSellBaseToken is a log parse operation binding the contract event 0xd8648b6ac54162763c86fd54bf2005af8ecd2f9cb273a5775921fd7f91e17b2d.
//
// Solidity: event SellBaseToken(address indexed seller, uint256 payBase, uint256 receiveQuote)
func (_Contract *ContractFilterer) ParseSellBaseToken(log types.Log) (*ContractSellBaseToken, error) {
	event := new(ContractSellBaseToken)
	if err := _Contract.contract.UnpackLog(event, "SellBaseToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateGasPriceLimitIterator is returned from FilterUpdateGasPriceLimit and is used to iterate over the raw logs and unpacked data for UpdateGasPriceLimit events raised by the Contract contract.
type ContractUpdateGasPriceLimitIterator struct {
	Event *ContractUpdateGasPriceLimit // Event containing the contract specifics and raw log

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
func (it *ContractUpdateGasPriceLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateGasPriceLimit)
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
		it.Event = new(ContractUpdateGasPriceLimit)
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
func (it *ContractUpdateGasPriceLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateGasPriceLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateGasPriceLimit represents a UpdateGasPriceLimit event raised by the Contract contract.
type ContractUpdateGasPriceLimit struct {
	OldGasPriceLimit *big.Int
	NewGasPriceLimit *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterUpdateGasPriceLimit is a free log retrieval operation binding the contract event 0x808f99cfd15f1be6019f93dc76c81d5bd82e0b3e0b3d23a54f5a2e647a6cc3cc.
//
// Solidity: event UpdateGasPriceLimit(uint256 oldGasPriceLimit, uint256 newGasPriceLimit)
func (_Contract *ContractFilterer) FilterUpdateGasPriceLimit(opts *bind.FilterOpts) (*ContractUpdateGasPriceLimitIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateGasPriceLimit")
	if err != nil {
		return nil, err
	}
	return &ContractUpdateGasPriceLimitIterator{contract: _Contract.contract, event: "UpdateGasPriceLimit", logs: logs, sub: sub}, nil
}

// WatchUpdateGasPriceLimit is a free log subscription operation binding the contract event 0x808f99cfd15f1be6019f93dc76c81d5bd82e0b3e0b3d23a54f5a2e647a6cc3cc.
//
// Solidity: event UpdateGasPriceLimit(uint256 oldGasPriceLimit, uint256 newGasPriceLimit)
func (_Contract *ContractFilterer) WatchUpdateGasPriceLimit(opts *bind.WatchOpts, sink chan<- *ContractUpdateGasPriceLimit) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateGasPriceLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateGasPriceLimit)
				if err := _Contract.contract.UnpackLog(event, "UpdateGasPriceLimit", log); err != nil {
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

// ParseUpdateGasPriceLimit is a log parse operation binding the contract event 0x808f99cfd15f1be6019f93dc76c81d5bd82e0b3e0b3d23a54f5a2e647a6cc3cc.
//
// Solidity: event UpdateGasPriceLimit(uint256 oldGasPriceLimit, uint256 newGasPriceLimit)
func (_Contract *ContractFilterer) ParseUpdateGasPriceLimit(log types.Log) (*ContractUpdateGasPriceLimit, error) {
	event := new(ContractUpdateGasPriceLimit)
	if err := _Contract.contract.UnpackLog(event, "UpdateGasPriceLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateKIterator is returned from FilterUpdateK and is used to iterate over the raw logs and unpacked data for UpdateK events raised by the Contract contract.
type ContractUpdateKIterator struct {
	Event *ContractUpdateK // Event containing the contract specifics and raw log

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
func (it *ContractUpdateKIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateK)
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
		it.Event = new(ContractUpdateK)
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
func (it *ContractUpdateKIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateKIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateK represents a UpdateK event raised by the Contract contract.
type ContractUpdateK struct {
	OldK *big.Int
	NewK *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUpdateK is a free log retrieval operation binding the contract event 0x023a40bebf7ac113f81c3d628073246cf9e0bc49980a9d6a9531498ce9e3dd1c.
//
// Solidity: event UpdateK(uint256 oldK, uint256 newK)
func (_Contract *ContractFilterer) FilterUpdateK(opts *bind.FilterOpts) (*ContractUpdateKIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateK")
	if err != nil {
		return nil, err
	}
	return &ContractUpdateKIterator{contract: _Contract.contract, event: "UpdateK", logs: logs, sub: sub}, nil
}

// WatchUpdateK is a free log subscription operation binding the contract event 0x023a40bebf7ac113f81c3d628073246cf9e0bc49980a9d6a9531498ce9e3dd1c.
//
// Solidity: event UpdateK(uint256 oldK, uint256 newK)
func (_Contract *ContractFilterer) WatchUpdateK(opts *bind.WatchOpts, sink chan<- *ContractUpdateK) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateK")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateK)
				if err := _Contract.contract.UnpackLog(event, "UpdateK", log); err != nil {
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

// ParseUpdateK is a log parse operation binding the contract event 0x023a40bebf7ac113f81c3d628073246cf9e0bc49980a9d6a9531498ce9e3dd1c.
//
// Solidity: event UpdateK(uint256 oldK, uint256 newK)
func (_Contract *ContractFilterer) ParseUpdateK(log types.Log) (*ContractUpdateK, error) {
	event := new(ContractUpdateK)
	if err := _Contract.contract.UnpackLog(event, "UpdateK", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateLiquidityProviderFeeRateIterator is returned from FilterUpdateLiquidityProviderFeeRate and is used to iterate over the raw logs and unpacked data for UpdateLiquidityProviderFeeRate events raised by the Contract contract.
type ContractUpdateLiquidityProviderFeeRateIterator struct {
	Event *ContractUpdateLiquidityProviderFeeRate // Event containing the contract specifics and raw log

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
func (it *ContractUpdateLiquidityProviderFeeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateLiquidityProviderFeeRate)
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
		it.Event = new(ContractUpdateLiquidityProviderFeeRate)
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
func (it *ContractUpdateLiquidityProviderFeeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateLiquidityProviderFeeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateLiquidityProviderFeeRate represents a UpdateLiquidityProviderFeeRate event raised by the Contract contract.
type ContractUpdateLiquidityProviderFeeRate struct {
	OldLiquidityProviderFeeRate *big.Int
	NewLiquidityProviderFeeRate *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterUpdateLiquidityProviderFeeRate is a free log retrieval operation binding the contract event 0x3ce6ea91adda496b7a0546fa6558e5b52c3a509de6015820efb00ca4020e0a07.
//
// Solidity: event UpdateLiquidityProviderFeeRate(uint256 oldLiquidityProviderFeeRate, uint256 newLiquidityProviderFeeRate)
func (_Contract *ContractFilterer) FilterUpdateLiquidityProviderFeeRate(opts *bind.FilterOpts) (*ContractUpdateLiquidityProviderFeeRateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateLiquidityProviderFeeRate")
	if err != nil {
		return nil, err
	}
	return &ContractUpdateLiquidityProviderFeeRateIterator{contract: _Contract.contract, event: "UpdateLiquidityProviderFeeRate", logs: logs, sub: sub}, nil
}

// WatchUpdateLiquidityProviderFeeRate is a free log subscription operation binding the contract event 0x3ce6ea91adda496b7a0546fa6558e5b52c3a509de6015820efb00ca4020e0a07.
//
// Solidity: event UpdateLiquidityProviderFeeRate(uint256 oldLiquidityProviderFeeRate, uint256 newLiquidityProviderFeeRate)
func (_Contract *ContractFilterer) WatchUpdateLiquidityProviderFeeRate(opts *bind.WatchOpts, sink chan<- *ContractUpdateLiquidityProviderFeeRate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateLiquidityProviderFeeRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateLiquidityProviderFeeRate)
				if err := _Contract.contract.UnpackLog(event, "UpdateLiquidityProviderFeeRate", log); err != nil {
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

// ParseUpdateLiquidityProviderFeeRate is a log parse operation binding the contract event 0x3ce6ea91adda496b7a0546fa6558e5b52c3a509de6015820efb00ca4020e0a07.
//
// Solidity: event UpdateLiquidityProviderFeeRate(uint256 oldLiquidityProviderFeeRate, uint256 newLiquidityProviderFeeRate)
func (_Contract *ContractFilterer) ParseUpdateLiquidityProviderFeeRate(log types.Log) (*ContractUpdateLiquidityProviderFeeRate, error) {
	event := new(ContractUpdateLiquidityProviderFeeRate)
	if err := _Contract.contract.UnpackLog(event, "UpdateLiquidityProviderFeeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractUpdateMaintainerFeeRateIterator is returned from FilterUpdateMaintainerFeeRate and is used to iterate over the raw logs and unpacked data for UpdateMaintainerFeeRate events raised by the Contract contract.
type ContractUpdateMaintainerFeeRateIterator struct {
	Event *ContractUpdateMaintainerFeeRate // Event containing the contract specifics and raw log

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
func (it *ContractUpdateMaintainerFeeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateMaintainerFeeRate)
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
		it.Event = new(ContractUpdateMaintainerFeeRate)
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
func (it *ContractUpdateMaintainerFeeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateMaintainerFeeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateMaintainerFeeRate represents a UpdateMaintainerFeeRate event raised by the Contract contract.
type ContractUpdateMaintainerFeeRate struct {
	OldMaintainerFeeRate *big.Int
	NewMaintainerFeeRate *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterUpdateMaintainerFeeRate is a free log retrieval operation binding the contract event 0x6b04da3d58e4b37d99652babb3ea2bc25ce94379bfff3059f03d61b26c59e553.
//
// Solidity: event UpdateMaintainerFeeRate(uint256 oldMaintainerFeeRate, uint256 newMaintainerFeeRate)
func (_Contract *ContractFilterer) FilterUpdateMaintainerFeeRate(opts *bind.FilterOpts) (*ContractUpdateMaintainerFeeRateIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateMaintainerFeeRate")
	if err != nil {
		return nil, err
	}
	return &ContractUpdateMaintainerFeeRateIterator{contract: _Contract.contract, event: "UpdateMaintainerFeeRate", logs: logs, sub: sub}, nil
}

// WatchUpdateMaintainerFeeRate is a free log subscription operation binding the contract event 0x6b04da3d58e4b37d99652babb3ea2bc25ce94379bfff3059f03d61b26c59e553.
//
// Solidity: event UpdateMaintainerFeeRate(uint256 oldMaintainerFeeRate, uint256 newMaintainerFeeRate)
func (_Contract *ContractFilterer) WatchUpdateMaintainerFeeRate(opts *bind.WatchOpts, sink chan<- *ContractUpdateMaintainerFeeRate) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateMaintainerFeeRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateMaintainerFeeRate)
				if err := _Contract.contract.UnpackLog(event, "UpdateMaintainerFeeRate", log); err != nil {
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

// ParseUpdateMaintainerFeeRate is a log parse operation binding the contract event 0x6b04da3d58e4b37d99652babb3ea2bc25ce94379bfff3059f03d61b26c59e553.
//
// Solidity: event UpdateMaintainerFeeRate(uint256 oldMaintainerFeeRate, uint256 newMaintainerFeeRate)
func (_Contract *ContractFilterer) ParseUpdateMaintainerFeeRate(log types.Log) (*ContractUpdateMaintainerFeeRate, error) {
	event := new(ContractUpdateMaintainerFeeRate)
	if err := _Contract.contract.UnpackLog(event, "UpdateMaintainerFeeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Contract contract.
type ContractWithdrawIterator struct {
	Event *ContractWithdraw // Event containing the contract specifics and raw log

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
func (it *ContractWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWithdraw)
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
		it.Event = new(ContractWithdraw)
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
func (it *ContractWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWithdraw represents a Withdraw event raised by the Contract contract.
type ContractWithdraw struct {
	Payer         common.Address
	Receiver      common.Address
	IsBaseToken   bool
	Amount        *big.Int
	LpTokenAmount *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xe89c586bd81ee35a18f7eac22a732b56e589a2821497cce12a0208828540a36d.
//
// Solidity: event Withdraw(address indexed payer, address indexed receiver, bool isBaseToken, uint256 amount, uint256 lpTokenAmount)
func (_Contract *ContractFilterer) FilterWithdraw(opts *bind.FilterOpts, payer []common.Address, receiver []common.Address) (*ContractWithdrawIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Withdraw", payerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ContractWithdrawIterator{contract: _Contract.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xe89c586bd81ee35a18f7eac22a732b56e589a2821497cce12a0208828540a36d.
//
// Solidity: event Withdraw(address indexed payer, address indexed receiver, bool isBaseToken, uint256 amount, uint256 lpTokenAmount)
func (_Contract *ContractFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ContractWithdraw, payer []common.Address, receiver []common.Address) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Withdraw", payerRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWithdraw)
				if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xe89c586bd81ee35a18f7eac22a732b56e589a2821497cce12a0208828540a36d.
//
// Solidity: event Withdraw(address indexed payer, address indexed receiver, bool isBaseToken, uint256 amount, uint256 lpTokenAmount)
func (_Contract *ContractFilterer) ParseWithdraw(log types.Log) (*ContractWithdraw, error) {
	event := new(ContractWithdraw)
	if err := _Contract.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
