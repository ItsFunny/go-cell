/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/22 12:45 下午
# @File : di.go
# @Description :
# @Attention :
*/
package di

import (
	"github.com/itsfunny/go-cell/base/reactor"
	"go.uber.org/fx"
)

const (
	FxExtension          = "extension"
	FxServer             = "server"
	FxProxy              = "proxy"
	FxDispatcher         = "dispatcher"
	FxCommand            = "command"
	FxSelectorGroup      = "httpSelector"
	FxHttpCommandHandler = "httpCommandHandler"
)

type OptionBuilder func() fx.Option

func RegisterExtension(constructor interface{}) fx.Option {
	return fx.Provide(fx.Annotated{
		Group:  FxExtension,
		Target: constructor,
	})
}

func RegisterServer(constructor interface{}) fx.Option {
	// return fx.Provide(fx.Annotated{
	// 	Group:  FxServer,
	// 	Target: constructor,
	// })
	return fx.Provide(constructor)
}
func RegisterProxy(constructor interface{}) fx.Option {
	// return fx.Provide(fx.Annotated{
	// 	Group:  FxProxy,
	// 	Target: constructor,
	// })
	return fx.Provide(constructor)
}
func RegisterDispatcher(constructor interface{}) fx.Option {
	// return fx.Provide(fx.Annotated{
	// 	Group:  FxDispatcher,
	// 	Target: constructor,
	// })
	return fx.Provide(constructor)
}

func RegisterCommandConstructor(constructor ...interface{}) fx.Option {
	ops := make([]fx.Option, 0)
	for _, con := range constructor {
		ops = append(ops, fx.Provide(
			fx.Annotated{
				Group:  FxCommand,
				Target: con,
			},
		))
	}
	return fx.Options(ops...)
}
func RegisterCommand(cmd reactor.ICommand) fx.Option {
	return fx.Provide(fx.Annotated{
		Group: FxCommand,
		Target: func() reactor.ICommand {
			return cmd
		},
	})
}

func CommandOptionBuilder(cmds ...reactor.ICommand) OptionBuilder {
	return func() fx.Option {
		ops := make([]fx.Option, 0)
		for _, cmd := range cmds {
			ops = append(ops, RegisterCommand(cmd))
		}
		return fx.Options(ops...)
	}
}

func RegisterHttpSelector(constructor interface{}) fx.Option {
	return fx.Options(
		fx.Provide(fx.Annotated{
			Group:  FxSelectorGroup,
			Target: constructor,
		}),
	)
}

func RegisterHttpCommandHandler(constructor interface{}) fx.Option {
	return fx.Options(
		fx.Provide(fx.Annotated{
			Group:  FxHttpCommandHandler,
			Target: constructor,
		}),
	)
}
