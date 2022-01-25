package error

import "errors"

var (
	AccountNotExists     = errors.New("account not exists")
	AccountAlreadyExists = errors.New("account exists")
	AccountNotReady      = errors.New("asd")
)
