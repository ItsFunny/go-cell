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
	"github.com/itsfunny/go-cell/base/core/event"
	"github.com/itsfunny/go-cell/base/core/eventbus"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/base/node/core/extension"
	"github.com/itsfunny/go-cell/component/codec"
	"github.com/itsfunny/go-cell/di"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
	"go.uber.org/fx"
	"sync"
)

type CellApplication struct {
	*services.BaseService
	*application
	app  *fx.App
	stop func() error
}

type application struct {
	Event extension.IApplicationEventBus
}

// TODO,remove duplicate option
func New(ctx context.Context, builders ...di.OptionBuilder) *CellApplication {
	apl := &application{}
	ret := &CellApplication{}
	ret.BaseService = services.NewBaseService(ctx, nil, logsdk.NewModule("APPLICATION", 1), ret,
		services.BaseServiceWithCtx(ctx))
	ops := make([]fx.Option, 0)
	ops = append(ops, extension.ExtensionManagerModule)
	ops = append(ops, eventbus.DefaultEventBusModule)
	ops = append(ops, extension.ApplicationEventBusModule)
	ops = append(ops, codec.CodecModule)
	ops = append(ops, fx.Provide(func() context.Context { return ctx }))
	for _, b := range builders {
		ops = append(ops, b())
	}
	ops = append(ops, fx.Extract(apl))
	app := fx.New(
		ops...,
	)
	ret.application = apl
	var once sync.Once
	var stopErr error
	ret.stop = func() error {
		once.Do(func() {
			stopErr = app.Stop(context.Background())
			if stopErr != nil {
				logrusplugin.Error("failure on stop: ", stopErr)
			}
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
		}
	}()
	if app.Err() != nil {
		panic(app.Err())
	}
	ret.app = app
	return ret
}

func (app *CellApplication) Run(args []string) {
	if err := app.app.Start(app.GetContext()); err != nil {
		panic(err)
	}
	go app.step0(args)
	<-app.Quit()
}

func (app *CellApplication) step0(args []string) {
	wait := make(chan struct{})
	go func() {
		app.Event.FireApplicationEvents(app.GetContext(),
			extension.ApplicationEnvironmentPreparedEvent{
				ICallBack: event.CallBack{
					CB: func() {
						close(wait)
					},
				},
				Args: args,
				Ctx:  app.GetContext(),
			},
		)
	}()
	<-wait
	go app.step1()
}
func (app *CellApplication) step1() {
	wait := make(chan struct{})
	go func() {
		app.Event.FireApplicationEvents(app.GetContext(),
			extension.ApplicationInitEvent{
				ICallBack: event.CallBack{
					CB: func() {
						close(wait)
					},
				},
			},
		)
	}()
	<-wait
	go app.step2()
}

func (app *CellApplication) step2() {
	go func() {
		wait := make(chan struct{})
		app.Event.FireApplicationEvents(app.GetContext(), extension.ApplicationStartedEvent{
			ICallBack: event.CallBack{CB: func() {
				close(wait)
			}},
		})
		<-wait
		go app.step3()
	}()
}

func (app *CellApplication) step3() {
	go func() {
		wait := make(chan struct{})
		app.Event.FireApplicationEvents(app.GetContext(), extension.ApplicationReadyEvent{
			ICallBack: event.CallBack{
				CB: func() {
					close(wait)
				},
			},
		})
		<-wait
	}()
}

func (app *CellApplication) GetApplicationBus() extension.IApplicationEventBus {
	return app.application.Event
}

func (app *CellApplication) ExportDefaultGenesis(rootPath string) {
	cb := make(chan struct{})
	app.Event.FireApplicationEvents(app.GetContext(), extension.ApplicationExportEvent{
		ICallBack: event.CallBack{CB: func() {
			close(cb)
		}},
		Path: rootPath,
	})
	<-cb
}
