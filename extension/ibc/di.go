package ibc

import (
	"github.com/itsfunny/go-cell/di"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"go.uber.org/fx"
)

var (
	IBCModule di.OptionBuilder = func() fx.Option {
		return fx.Options(di.RegisterExtension(newIBCExtension))
	}

	module = logsdk.NewModule("ibc", 1)
)
