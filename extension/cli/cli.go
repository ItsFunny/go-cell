package cli

import "github.com/itsfunny/go-cell/base/node/core/extension"

type CLIExtension struct {
	*extension.BaseExtension
}

func NewCLIExtension() extension.INodeExtension {
	ret := &CLIExtension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	return ret
}

func (cli *CLIExtension) OnExtensionStart(ctx extension.INodeContext) error {
	return nil
}
