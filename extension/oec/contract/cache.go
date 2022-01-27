package contract

import (
	"bytes"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type ContractCache struct {
	contracts map[string]*contractCacheNode
}

func newContractCache() *ContractCache {
	ret := &ContractCache{
		contracts: make(map[string]*contractCacheNode),
	}
	return ret
}

type contractCacheNode struct {
	Name     string
	Abi      abi.ABI
	ByteCode []byte
}

func (this *ContractCache) newNode(name string, abiFileHex string, binHexCodes string) (*contractCacheNode, error) {
	node, exist := this.contracts[name]
	if exist {
		return node, nil
	}
	abiBytes, err := hex.DecodeString(abiFileHex)
	if nil != err {
		return nil, err
	}
	binBytes, err := hex.DecodeString(binHexCodes)
	if nil != err {
		return nil, err
	}

	ret := &contractCacheNode{}
	ret.Name = name
	ret.ByteCode = binBytes
	ret.Abi, err = abi.JSON(bytes.NewReader(abiBytes))
	if nil != err {
		return nil, err
	}
	this.contracts[name] = ret
	return ret, nil
}

func(this *ContractCache)getNode(name string)*contractCacheNode{
	return this.contracts[name]
}