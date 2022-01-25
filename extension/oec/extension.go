package oec

import (
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/base/node/core/extension"
	"github.com/itsfunny/go-cell/extension/oec/config"
	"github.com/itsfunny/go-cell/extension/oec/contract"
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
)

type oecExtension struct {
	*extension.BaseExtension

	service contract.IContractService
}

func newOecExtension(s contract.IContractService) extension.INodeExtension {
	ret := &oecExtension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	ret.service = s
	return ret
}

func (b *oecExtension) OnExtensionInit(ctx extension.INodeContext) error {
	logrusplugin.MInfo(module, "extension init")

	return b.service.BStart(services.StartCTXWithKV("config", config.NewDefaultOECConfig()))
}
