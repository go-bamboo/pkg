// Package flattened is a an on-chain light client checkpoint oracle.
package v2

//go:generate abigen --abi contract/dodoAbi.json --pkg contract --out contract/dodo.go

import (
	"github.com/emberfarkas/pkg/contracts/dodoex/v2/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-kratos/kratos/v2/errors"
)

// Dodo is a Go wrapper around an on-chain checkpoint oracle contract.
type Dodo struct {
	address  common.Address
	contract *contract.Contract
}

// NewDodo binds checkpoint contract and returns a registrar instance.
func NewDodo(contractAddr common.Address, backend bind.ContractBackend) (ctrt *Dodo, err error) {
	c, err := contract.NewContract(contractAddr, backend)
	if err != nil {
		err = errors.FromError(err)
		return
	}
	ctrt = &Dodo{address: contractAddr, contract: c}
	return
}

// ContractAddr returns the address of contract.
func (ctrct *Dodo) ContractAddr() common.Address {
	return ctrct.address
}

// Contract returns the underlying contract instance.
func (ctrct *Dodo) Contract() *contract.Contract {
	return ctrct.contract
}
