/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/19 10:34 下午
# @File : cell.go
# @Description :
# @Attention :
*/
package application

import (
	"context"
	"github.com/itsfunny/go-cell/base/core/eventbus"
	"github.com/itsfunny/go-cell/base/node/core/extension"
	"go.uber.org/fx"
)

var (
	cellApplicationModule = fx.Provide(newCellApplication)
)

type CellApplication struct {
	event eventbus.ICommonEventBus
}

func newCellApplication(event eventbus.ICommonEventBus) *CellApplication {
	ret := &CellApplication{event: event}
	return ret
}

func Start(moduleOps ...fx.Option) {
	ops := make([]fx.Option, 0)
	ops = append(ops, extension.ExtensionManagerModule, CellApplicationOption())
	ops = append(ops, moduleOps...)
	ops = append(ops, extension.ApplicationEventBusModule)
	ops = append(ops, fx.Provide(eventbus.NewCommonEventBusComponentImpl))

	ret := &CellApplication{}
	ops = append(ops, fx.Extract(ret))
	app := fx.New(
		ops...,
	)
	app.Run()
	// err:=app.Start(context.Background())
	// fmt.Println(err)
}

func CellApplicationOption() fx.Option {
	return fx.Options(
		fx.Provide(newCellApplication),
	)
}

func (this *CellApplication) Start(lifecycle fx.Lifecycle) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	})
}
