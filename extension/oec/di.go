package oec

import (
	"github.com/itsfunny/go-cell/di"
	"github.com/itsfunny/go-cell/extension/oec/commands"
	"github.com/itsfunny/go-cell/extension/oec/contract"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"go.uber.org/fx"
)

var (
	OecModule di.OptionBuilder = func() fx.Option {
		return fx.Options(di.RegisterExtension(newOecExtension),
			fx.Provide(contract.NewContractServiceImpl),
			commands.OecCommands,
		)
	}

	module = logsdk.NewModule("oec", 1)
)
