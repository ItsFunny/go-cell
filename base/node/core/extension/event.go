/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/18 10:09 下午
# @File : event.go
# @Description :
# @Attention :
*/
package extension

import (
	"context"
	"github.com/itsfunny/go-cell/base/core/event"
	"github.com/itsfunny/go-cell/base/core/eventbus"
	"go.uber.org/fx"
	"sync/atomic"
)

const applicationEventTypeKey = "extension.event"
const applicationEvent = "applicationEvent"

var (
	_                         IApplicationEventBus = (*applicationEventBus)(nil)
	ApplicationEventBusModule                      = fx.Options(
		fx.Provide(NewApplicationEventBus),
	)
)

type IApplicationEventBus interface {
	eventbus.ICommonEventBus
	SubscribeApplicationEvent(ctx context.Context, clientId string) (eventbus.Subscription, error)
	FireApplicationEvents(ctx context.Context, data interface{}) error
}

func NewApplicationEventBus(bus eventbus.ICommonEventBus) IApplicationEventBus {
	ret := &applicationEventBus{ICommonEventBus: bus}
	return ret
}

type applicationEventBus struct {
	eventbus.ICommonEventBus

	applicationEventCount int32
}

func (a *applicationEventBus) SubscribeApplicationEvent(ctx context.Context, clientId string) (eventbus.Subscription, error) {
	atomic.AddInt32(&a.applicationEventCount, 1)
	return a.ICommonEventBus.Subscribe(ctx, clientId, eventbus.QueryForEvent(applicationEventTypeKey, applicationEvent), 1)
}
func (a *applicationEventBus) GetApplicationListenerCounts() int32 {
	return atomic.LoadInt32(&a.applicationEventCount)
}
func (a *applicationEventBus) FireApplicationEvents(ctx context.Context, data interface{}) error {
	return a.ICommonEventBus.PublishWithEvents(ctx, data, map[string][]string{
		applicationEventTypeKey: {applicationEvent},
	})
}

type ApplicationEnvironmentPreparedEvent struct {
	event.ICallBack
	Args []string
	Ctx  context.Context
}
type ApplicationInitEvent struct {
	event.ICallBack
}
type ApplicationStartedEvent struct {
	event.ICallBack
}

type ApplicationReadyEvent struct {
	event.ICallBack
}

type ExtensionLoadedEvent struct {
}

type ApplicationExportEvent struct {
	event.ICallBack
	Path string
}
