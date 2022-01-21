/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/21 9:46 下午
# @File : module.go
# @Description :
# @Attention :
*/
package extension

import (
	"context"
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/core/services"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"go.uber.org/fx"
)

var (
	ipOption               = options.StringOption("ip", "i", "ip address").WithDefault("127.0.0.1")
	extensionManagerModule = logsdk.NewModule("manager", 1)

	ExtensionManagerModule = fx.Options(
		fx.Provide(NewExtensionManager),
		extensionModuleOption,
	)
	extensionModuleOption = fx.Options(
		fx.Invoke(start),
	)
)

func Register(constructor interface{}) fx.Option {
	return fx.Options(
		fx.Provide(fx.Annotated{
			Group:  "extension",
			Target: constructor,
		}),
	)
}

func start(lc fx.Lifecycle, m *NodeExtensionManager) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return m.BStart(services.SyncStartOpt)
		},
		OnStop: func(ctx context.Context) error {
			return m.BStop()
		},
	})
}

type Extensions struct {
	fx.In
	Extensions []INodeExtension `group:"extension"`
}
