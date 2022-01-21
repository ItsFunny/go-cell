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
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
	"go.uber.org/fx"
	"sync"
)

type CellApplication struct {
	Event extension.IApplicationEventBus

	stop func() error
}

// TODO,remove duplicate option
func Start(ctx context.Context, args []string, moduleOps ...fx.Option) *CellApplication {
	lctx := ctx
	ctx, cancel := context.WithCancel(valueContext{lctx})

	ret := &CellApplication{}
	ops := make([]fx.Option, 0)
	ops = append(ops, extension.ExtensionManagerModule)
	ops = append(ops, fx.Provide(eventbus.NewCommonEventBusComponentImpl))
	ops = append(ops, extension.ApplicationEventBusModule)
	ops = append(ops, moduleOps...)
	ops = append(ops, fx.Extract(ret))
	app := fx.New(
		ops...,
	)
	var once sync.Once
	var stopErr error
	ret.stop = func() error {
		once.Do(func() {
			stopErr = app.Stop(context.Background())
			if stopErr != nil {
				logrusplugin.Error("failure on stop: ", stopErr)
			}
			// Cancel the context _after_ the app has stopped.
			cancel()
		})
		return stopErr
	}

	go func() {
		select {
		case <-ctx.Done():
			err := ret.stop()
			if err != nil {
				logrusplugin.Error("failure on stop: ", err)
			}
		case <-ctx.Done():
		}
	}()
	if app.Err() != nil {
		panic(app.Err())
	}
	if err := app.Start(ctx); err != nil {
		panic(err)
	}
	select {}
	return ret
}
