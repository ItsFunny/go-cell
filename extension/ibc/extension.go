package ibc

import "github.com/itsfunny/go-cell/base/node/core/extension"

type ibcExtension struct {
	*extension.BaseExtension
}

func newIBCExtension() extension.INodeExtension {
	ret := &ibcExtension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	return ret
}
