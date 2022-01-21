/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/18 9:57 下午
# @File : eventbus.go
# @Description :
# @Attention :
*/
package eventbus

import (
	"context"
	"github.com/itsfunny/go-cell/base/core/services"
)

// IInternalPubSubComponent
type ICommonEventBus interface {
	services.IBaseService
	Subscribe(ctx context.Context, clientID string, query Query, outCapacity ...int) (Subscription, error)
	PublishWithEvents(ctx context.Context, msg interface{}, events map[string][]string) error
	SubscribeUnbuffered(ctx context.Context, clientID string, query Query) (*SubscriptionImpl, error)
	Unsubscribe(ctx context.Context, subscriber string, query Query) error
	UnsubscribeAll(ctx context.Context, subscriber string) error

	NumClients() int
	NumClientSubscriptions(clientID string) int
}
type Subscription interface {
	Out() <-chan PubSubMessage
	Canceled() <-chan struct{}
	Cancel(err error)
	Err() error
}
type Query interface {
	Matches(events map[string][]string) (bool, error)
	String() string
}

