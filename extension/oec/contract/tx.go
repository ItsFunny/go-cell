package contract

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/itsfunny/go-cell/base/core/promise"
	"sync"
	"time"
)

type txCache struct {
	mtx sync.RWMutex
	txs map[common.Hash]*promiseWrapper
}

type promiseWrapper struct {
	p         *promise.Promise
	registerT time.Time
}

func newTxCache() *txCache {
	ret := &txCache{}
	ret.txs = make(map[common.Hash]*promiseWrapper)
	return ret
}

func (this *txCache) notify(hash common.Hash) {
	this.mtx.RLock()
	defer this.mtx.RUnlock()
	p := this.txs[hash]
	if p == nil {
		return
	}
	p.p.EmptyDone()
}

func (this *txCache) registerListener(ctx context.Context, hash common.Hash) *promise.Promise {
	p := promise.NewPromise(ctx)
	this.mtx.Lock()
	defer this.mtx.Unlock()
	_, exist := this.txs[hash]
	if exist {
		panic("asd")
	}
	this.txs[hash] = &promiseWrapper{
		p:         p,
		registerT: time.Now(),
	}
	return p
}

func (this *txCache) removeListener(hash common.Hash) {
	this.batchRemoveListener(hash)
}
func(this *txCache)batchRemoveListener(hash ...common.Hash){
	this.mtx.Lock()
	defer this.mtx.Unlock()
	for _,h:=range hash{
		delete(this.txs,h)
	}
}
