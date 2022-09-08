package bridge

//go:generate abigen --sol contract/Bridge.sol --pkg contract --out contract/bridge.go
// abigen v1.10.1
// solc 0.8.1

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/emberfarkas/pkg/contracts/bridge/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/go-kratos/kratos/v2/errors"
)

// Bridge is a Go wrapper around an on-chain checkpoint oracle contract.
type Bridge struct {
	address  common.Address
	contract *contract.Bridge
}

// NewBridge binds checkpoint contract and returns a registrar instance.
func NewBridge(contractAddr common.Address, backend bind.ContractBackend) (ctrt *Bridge, err error) {
	c, err := contract.NewBridge(contractAddr, backend)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	ctrt = &Bridge{address: contractAddr, contract: c}
	return
}

// ContractAddr returns the address of contract.
func (ctrct *Bridge) ContractAddr() common.Address {
	return ctrct.address
}

// Contract returns the underlying contract instance.
func (ctrct *Bridge) Contract() *contract.Bridge {
	return ctrct.contract
}

func (ctrct *Bridge) GetIn(ctx context.Context, from common.Address, offset *big.Int) (contract.BridgeTokenIn, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	in, err := ctrct.contract.GetIn(opts, offset)
	if err != nil {
		return in, errors.FromError(err)
	}
	return in, nil
}

func (ctrct *Bridge) GetOut(ctx context.Context, from common.Address, offset *big.Int) (contract.BridgeTokenOut, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	out, err := ctrct.contract.GetOut(opts, offset)
	if err != nil {
		return out, errors.FromError(err)
	}
	return out, nil
}

func (ctrct *Bridge) GetInLength(ctx context.Context, from common.Address) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	out, err := ctrct.contract.GetInLength(opts)
	if err != nil {
		return nil, errors.FromError(err)
	}
	return out, err
}

// GetBalance 获取主币余额
func (ctrct *Bridge) GetBalance(ctx context.Context, from common.Address) (*big.Int, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	out, err := ctrct.contract.GetBalance(opts)
	if err != nil {
		return nil, errors.FromError(err)
	}
	return out, nil
}

// Owner 获取主币余额
func (ctrct *Bridge) Owner(ctx context.Context, from common.Address) (common.Address, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	out, err := ctrct.contract.Owner(opts)
	if err != nil {
		return out, errors.FromError(err)
	}
	return out, nil
}

func (ctrct *Bridge) Operator1(ctx context.Context, from common.Address) (common.Address, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	out, err := ctrct.contract.Operator1(opts)
	if err != nil {
		return out, errors.FromError(err)
	}
	return out, nil
}

func (ctrct *Bridge) Operator2(ctx context.Context, from common.Address) (common.Address, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	out, err := ctrct.contract.Operator2(opts)
	if err != nil {
		return out, errors.FromError(err)
	}
	return out, nil
}

func (ctrct *Bridge) DepositToken(ctx context.Context, chainID *big.Int, fromPriv *ecdsa.PrivateKey, nonce *big.Int, token common.Address, value *big.Int) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyedTransactorWithChainID(fromPriv, chainID)
	if err != nil {
		return
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.GasLimit = 100000
	tx, err := ctrct.contract.DepositToken(opts, token, value)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	txHash = tx.Hash().Hex()
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	rawTx = hexutil.Encode(rawTxBytes)
	return
}

func (ctrct *Bridge) Withdraw(ctx context.Context, chainID *big.Int, fromPriv *ecdsa.PrivateKey, nonce *big.Int, token common.Address) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyedTransactorWithChainID(fromPriv, chainID)
	if err != nil {
		return
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.Value = big.NewInt(0)
	opts.GasLimit = 100000

	tx, err := ctrct.contract.Withdraw(opts, token)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	txHash = tx.Hash().Hex()
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	rawTx = hexutil.Encode(rawTxBytes)
	return
}

func (ctrct *Bridge) SendToken(ctx context.Context, chainID *big.Int, fromPriv *ecdsa.PrivateKey, nonce *big.Int, fromIndex *big.Int, token common.Address, to common.Address, value *big.Int) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyedTransactorWithChainID(fromPriv, chainID)
	if err != nil {
		return
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.Value = big.NewInt(0)
	opts.GasLimit = 400000

	tx, err := ctrct.contract.SendToken(opts, fromIndex, token, to, value)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	txHash = tx.Hash().Hex()
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	rawTx = hexutil.Encode(rawTxBytes)
	return
}

func (ctrct *Bridge) SetOperator(ctx context.Context, chainID *big.Int, fromPriv *ecdsa.PrivateKey, nonce *big.Int, operator common.Address, index *big.Int) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyedTransactorWithChainID(fromPriv, chainID)
	if err != nil {
		return
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.Value = big.NewInt(0)
	opts.GasLimit = 100000

	tx, err := ctrct.contract.SetOperator(opts, operator, index)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	txHash = tx.Hash().Hex()
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	rawTx = hexutil.Encode(rawTxBytes)
	return
}

func (ctrct *Bridge) PauseToken(ctx context.Context, chainID *big.Int, fromPriv *ecdsa.PrivateKey, nonce *big.Int, token common.Address) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyedTransactorWithChainID(fromPriv, chainID)
	if err != nil {
		return
	}
	opts.Context = ctx
	opts.Nonce = nonce
	opts.Value = big.NewInt(0)
	opts.GasLimit = 100000

	tx, err := ctrct.contract.PauseToken(opts, token)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	txHash = tx.Hash().Hex()
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	rawTx = hexutil.Encode(rawTxBytes)
	return
}

func (ctrct *Bridge) UnpauseToken(ctx context.Context, chainID *big.Int, fromPriv *ecdsa.PrivateKey, nonce *big.Int, token common.Address) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyedTransactorWithChainID(fromPriv, chainID)
	if err != nil {
		return
	}
	opts.Nonce = nonce
	opts.Value = big.NewInt(0)
	opts.GasLimit = 100000

	tx, err := ctrct.contract.UnpauseToken(opts, token)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	txHash = tx.Hash().Hex()
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	rawTx = hexutil.Encode(rawTxBytes)
	return
}
