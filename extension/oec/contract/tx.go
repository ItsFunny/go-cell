package contract

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/itsfunny/go-cell/base/core/promise"
	"sync"
)

type txCache struct {
	mtx sync.RWMutex
	txs map[common.Hash]*promise.Promise
}

func newTxCache() *txCache {
	ret := &txCache{}
	ret.txs = make(map[common.Hash]*promise.Promise)
	return ret
}

func (this *txCache) notify(hash common.Hash) {
	this.mtx.RLock()
	defer this.mtx.RUnlock()
	p := this.txs[hash]
	if p == nil {
		return
	}
	p.EmptyDone()
}

func (this *txCache) registerListener(ctx context.Context, hash common.Hash) *promise.Promise {
	p := promise.NewPromise(ctx)
	this.mtx.Lock()
	defer this.mtx.Unlock()
	_, exist := this.txs[hash]
	if exist {
		panic("asd")
	}
	this.txs[hash] = p
	return p
}

func(this *txCache)removeListener(hash common.Hash){
	this.mtx.Lock()
	defer this.mtx.Unlock()
	delete(this.txs,hash)
}