// Package flattened is a an on-chain light client checkpoint oracle.
package flattened

//go:generate abigen --sol contract/flattened.sol --pkg contract --out contract/flattened.go

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/emberfarkas/pkg/contracts/flattened/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/go-kratos/kratos/v2/errors"
)

// Media is a Go wrapper around an on-chain checkpoint oracle contract.
type Media struct {
	address  common.Address
	contract *contract.Media
}

// NewFlattened binds checkpoint contract and returns a registrar instance.
func NewMedia(contractAddr common.Address, backend bind.ContractBackend) (ctrt *Media, err error) {
	c, err := contract.NewMedia(contractAddr, backend)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	ctrt = &Media{address: contractAddr, contract: c}
	return
}

// ContractAddr returns the address of contract.
func (ctrct *Media) ContractAddr() common.Address {
	return ctrct.address
}

// Contract returns the underlying contract instance.
func (ctrct *Media) Contract() *contract.Media {
	return ctrct.contract
}

func (ctrct *Media) Name(ctx context.Context, from common.Address) (string, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	return ctrct.contract.Name(opts)
}

func (ctrct *Media) Symbol(ctx context.Context, from common.Address) (string, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	symbol, err := ctrct.contract.MediaCaller.Symbol(opts)
	return symbol, errors.FromError(err)
}

func (ctrct *Media) OwnerOf(ctx context.Context, from common.Address, tokenId *big.Int) (string, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	addr, err := ctrct.contract.OwnerOf(opts, tokenId)
	return addr.Hex(), errors.FromError(err)
}

func (ctrct *Media) TokenURI(ctx context.Context, from common.Address, tokenId *big.Int) (string, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	uri, err := ctrct.contract.TokenURI(opts, tokenId)
	return uri, errors.FromError(err)
}

func (ctrct *Media) TokenContentHashes(ctx context.Context, from common.Address, tokenId *big.Int) (string, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	_, err := ctrct.contract.TokenContentHashes(opts, tokenId)
	return "", errors.FromError(err)
}

func (ctrct *Media) TokenOfOwnerByIndex(ctx context.Context, from common.Address, owner common.Address, index *big.Int) (string, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	tokenId, err := ctrct.contract.TokenOfOwnerByIndex(opts, owner, index)
	return tokenId.String(), errors.FromError(err)
}

func (ctrct *Media) TokenByIndex(ctx context.Context, from common.Address, index *big.Int) (string, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	tokenId, err := ctrct.contract.TokenByIndex(opts, index)
	return tokenId.String(), errors.FromError(err)
}

func (ctrct *Media) GetTokenIdByContentHash(ctx context.Context, from common.Address, hash string) (string, error) {
	opts := &bind.CallOpts{
		Context: ctx,
		From:    from,
	}
	contentHash := common.BytesToHash([]byte(hash))
	tokenId, err := ctrct.contract.GetTokenIdByContentHash(opts, contentHash)
	return tokenId.String(), errors.FromError(err)
}

func (ctrct *Media) Mint(ctx context.Context, chainID *big.Int, from common.Address, fromPriv *ecdsa.PrivateKey, nonce *big.Int, tokenId *big.Int, tokenURI string, hash string) (txHash, rawTx string, err error) {
	opts := bind.NewKeyedTransactor(fromPriv)
	opts.Nonce = nonce
	opts.GasLimit = 1000000
	opts.Signer = func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(t, types.NewEIP155Signer(chainID), fromPriv)
	}
	data := contract.IMediaMediaData{
		TokenURI:    tokenURI,
		ContentHash: common.HexToHash(hash),
	}
	tx, err := ctrct.contract.Mint(opts, tokenId, data)
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

func (ctrct *Media) MintForCreator(ctx context.Context, chainID *big.Int, from common.Address, fromPriv *ecdsa.PrivateKey, to common.Address, nonce *big.Int, tokenId *big.Int, tokenURI string, hash string) (txHash, rawTx string, err error) {
	opts := bind.NewKeyedTransactor(fromPriv)
	opts.Nonce = nonce
	opts.GasLimit = 1000000
	opts.Signer = func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(t, types.NewEIP155Signer(chainID), fromPriv)
	}

	data := contract.IMediaMediaData{
		TokenURI:    tokenURI,
		ContentHash: common.HexToHash(hash),
	}
	tx, err := ctrct.contract.MintForCreator(opts, to, tokenId, data)
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

func (ctrct *Media) MintWithSig(ctx context.Context, chainID *big.Int, from common.Address, fromPriv *ecdsa.PrivateKey, nonce *big.Int, tokenId *big.Int, tokenURI string, hash string) (txHash, rawTx string, err error) {
	opts := bind.NewKeyedTransactor(fromPriv)
	opts.Nonce = nonce
	opts.GasLimit = 1000000
	opts.Signer = func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(t, types.NewEIP155Signer(chainID), fromPriv)
	}

	data := contract.IMediaMediaData{
		TokenURI:    tokenURI,
		ContentHash: common.HexToHash(hash),
	}
	sig := contract.IMediaEIP712Signature{}
	tx, err := ctrct.contract.MintWithSig(opts, from, tokenId, data, sig)
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

func (ctrct *Media) UpdateTokenURI(ctx context.Context, chainID *big.Int, from common.Address, fromPriv *ecdsa.PrivateKey, nonce *big.Int, tokenId *big.Int, tokenURI string) (txHash, rawTx string, err error) {
	opts := bind.NewKeyedTransactor(fromPriv)
	opts.Nonce = nonce
	opts.GasLimit = 1000000
	opts.Signer = func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(t, types.NewEIP155Signer(chainID), fromPriv)
	}

	tx, err := ctrct.contract.UpdateTokenURI(opts, tokenId, tokenURI)
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

func (ctrct *Media) TransferFrom(ctx context.Context, chainID *big.Int, from common.Address, fromPriv *ecdsa.PrivateKey, nonce *big.Int, fromx common.Address, to common.Address, tokenId *big.Int) (txHash, rawTx string, err error) {
	opts := bind.NewKeyedTransactor(fromPriv)
	opts.Nonce = nonce
	opts.GasLimit = 1000000
	opts.Signer = func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(t, types.NewEIP155Signer(chainID), fromPriv)
	}
	tx, err := ctrct.contract.TransferFrom(opts, fromx, to, tokenId)
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

func (ctrct *Media) Burn(ctx context.Context, chainID *big.Int, from common.Address, fromPriv *ecdsa.PrivateKey, nonce *big.Int, tokenId *big.Int) (txHash, rawTx string, err error) {
	opts := bind.NewKeyedTransactor(fromPriv)
	opts.Nonce = nonce
	opts.GasLimit = 1000000
	opts.Signer = func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
		return types.SignTx(t, types.NewEIP155Signer(chainID), fromPriv)
	}
	tx, err := ctrct.contract.Burn(opts, tokenId)
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
