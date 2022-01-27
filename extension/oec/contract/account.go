package contract

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
)

type Account struct {
	key      *ecdsa.PrivateKey
	address  common.Address
	Contract *Contract
	readyFlag    bool
	contractAddress map[string]common.Address
}

func (this Account) ready() bool {
	return this.readyFlag
}
