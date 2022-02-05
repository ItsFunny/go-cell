package contract

import "sync"

type transferCache struct {
	mtx     sync.RWMutex
	records map[string]*transferRecordWrapper
}

type transferRecordWrapper struct {
	records []transferRecord
}
type transferRecord struct {
	from   string
	to     string
	amount int64
}

func newTransferCache() *transferCache {
	ret := &transferCache{
		records: make(map[string]*transferRecordWrapper),
	}
	return ret
}
func (this *transferCache) clean() {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	this.records = make(map[string]*transferRecordWrapper)
}
func (this *transferCache) recordOne(from, to string, amount int64) {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	wp, exist := this.records[from]
	if !exist {
		wp = &transferRecordWrapper{}
		this.records[from] = wp
	}
	wp.records = append(wp.records, transferRecord{
		from:   from,
		to:     to,
		amount: amount,
	})
}
