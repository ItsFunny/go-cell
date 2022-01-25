package contract

import (
	"bytes"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
)

type Contract struct {
	Name     string
	Address  string
	Addr     common.Address
	Abi      abi.ABI
	ByteCode []byte
}

func NewContract(name, address string, abiFileHex string, binHexCodes string) (*Contract, error) {
	abiBytes, err := hex.DecodeString(abiFileHex)
	if nil != err {
		return nil, err
	}
	binBytes, err := hex.DecodeString(binHexCodes)
	if nil != err {
		return nil, err
	}

	ret := &Contract{}
	ret.Name=name
	ret.ByteCode = binBytes
	ret.Abi, err = abi.JSON(bytes.NewReader(abiBytes))
	if nil != err {
		return nil, err
	}

	if len(address) > 0 {
		ret.Addr = common.HexToAddress(address)
		logrusplugin.Info("new contract", "addr", address)
	}

	return ret, nil
}
