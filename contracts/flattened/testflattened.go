package flattened

import (
	"context"
	"encoding/hex"
	"math/big"

	"bls/pkg/contracts/flattened/contract"
	"bls/pkg/ecode"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

func (ctrct *Media) MintKs(ctx context.Context, chainID *big.Int, ks *keystore.KeyStore, acc accounts.Account, nonce *big.Int, tokenId *big.Int, tokenURI string, hash string) (txHash, rawTx string, err error) {
	opts, err := bind.NewKeyStoreTransactor(ks, acc)
	if err != nil {
		return
	}
	opts.Nonce = nonce
	opts.GasLimit = 100000
	opts.Signer = func(a common.Address, t *types.Transaction) (*types.Transaction, error) {
		return ks.SignTx(acc, t, chainID)
	}

	data := contract.IMediaMediaData{
		TokenURI:    tokenURI,
		ContentHash: common.HexToHash(hash),
	}
	tx, err := ctrct.contract.Mint(opts, tokenId, data)
	if err != nil {
		err = ecode.WrapError(err)
		return
	}
	txHash = tx.Hash().Hex()
	rawTxBytes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		err = ecode.WrapError(err)
		return
	}
	rawTxHex := hex.EncodeToString(rawTxBytes)
	rawTx = "0x" + rawTxHex
	return
}
