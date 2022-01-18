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
	"github.com/itsfunny/go-cell/base/core/eventbus"
)

const applicationEventTypeKey = "extension.event"
const applicationEvent = "applicationEvent"

var (
	_ IApplicationEventBus = (*applicationEventBus)(nil)
)

type IApplicationEventBus interface {
	eventbus.ICommonEventBus
	SubscribeApplicationEvent(ctx context.Context, clientId string) (eventbus.Subscription, error)
}
type applicationEventBus struct {
	eventbus.ICommonEventBus
}

func (a *applicationEventBus) SubscribeApplicationEvent(ctx context.Context, clientId string) (eventbus.Subscription, error) {
	return a.ICommonEventBus.Subscribe(ctx, clientId, eventbus.QueryForEvent(applicationEventTypeKey, applicationEventTypeKey), 1)
}
