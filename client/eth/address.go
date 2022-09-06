package eth

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

var ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

var Prec, _ = new(big.Float).SetString(new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil).String())
