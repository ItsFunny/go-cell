package contract

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"sync"
)

type Account struct {
	key             *ecdsa.PrivateKey
	address         common.Address
	Contract        *Contract
	readyFlag       bool
	contractAddress map[string]common.Address
	moniker         string
	gasPrice        int64
}

func (this Account) ready() bool {
	return this.readyFlag
}
func newAccount() {}

type accountCache struct {
	mtx      sync.RWMutex
	accounts map[string]*Account
}

func newAccountCache() *accountCache {
	ret := &accountCache{
		accounts: make(map[string]*Account),
	}
	return ret
}
func (this *accountCache) addOne(moniker string, acc *Account) {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	this.accounts[moniker] = acc
}
func (this *accountCache) get(moniker string) *Account {
	this.mtx.RLock()
	defer this.mtx.RUnlock()
	return this.accounts[moniker]
}

func(this *accountCache)size()int{
	this.mtx.RLock()
	defer this.mtx.RUnlock()
	return len(this.accounts)
}

func(this *accountCache)getAccounts()map[string]*Account{
	ret:=make(map[string]*Account)
	this.mtx.Lock()
	defer this.mtx.Unlock()
	for k,v:=range this.accounts{
		ret[k]=v
	}
	return ret
}