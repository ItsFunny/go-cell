package contract

type TransferReq struct {
	From     string
	To       string
	AmountV  int64
	GasPrice int64
}

type OneToMoreReq struct {
	From           string
	ToAccountLimit int
}

type OneToMoreResp struct {
}

type BenchReq struct {
	TransactionLimit        int
	AccountLimit int

}


type RegisterAccountReq struct {
	Moniker string
	TransferFrom string
}