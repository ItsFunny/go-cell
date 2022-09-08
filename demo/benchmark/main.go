package main

import (
	"context"
	"github.com/itsfunny/go-cell/application"
	"github.com/itsfunny/go-cell/base/node/core/extension"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/di"
	"github.com/itsfunny/go-cell/extension/http"
	"go.uber.org/fx"
	"os"
	"strconv"
)

type benchMarkExtension struct {
	*extension.BaseExtension
}

var (
	benchModule di.OptionBuilder = func() fx.Option {
		return fx.Options(
			di.RegisterExtension(newBenchMarkExtension),
			commands(),
		)
	}
)

type benchCmd struct {
	path string
	*reactor.Command
}

func commands() fx.Option {
	fs := make([]fx.Option, 0)
	for i := 1; i <= 100; i++ {
		cmd := newBenchCmd("/bench/" + strconv.Itoa(i))
		fs = append(fs, di.RegisterCommand(cmd))
	}
	return fx.Options(fs...)
}
func newBenchCmd(path string) reactor.ICommand {
	ret := new(benchCmd)
	ret.path = path
	var pid reactor.ProtocolID
	pid = reactor.ProtocolID(path)
	ret.Command = &reactor.Command{
		ProtocolID: pid,
		Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
			ctx.Response(ctx.CreateResponseWrapper().WithRet(path))
			return nil
		},
		RunType: reactor.RunTypeHttpGet,
	}
	return ret
}
func newBenchMarkExtension() extension.INodeExtension {
	ret := new(benchMarkExtension)
	ret.BaseExtension = extension.NewBaseExtension(ret)
	return ret
}

func main() {
	app := application.New(context.Background(),
		benchModule, http.HttpModule)
	app.Run(os.Args)
}
