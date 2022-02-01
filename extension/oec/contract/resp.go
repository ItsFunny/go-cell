package contract

import "github.com/itsfunny/go-cell/base/core/promise"

type TransferResp struct {
	Promise *promise.Promise
}

type BenchResp struct {
	Success    int32
	Fail       int32
	BeginBlock int64
	FinalBlock int64
}

type RegisterAccountResp struct {
	PrvHexString string
	Moniker      string
	Address string
}
