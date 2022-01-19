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
	"go.uber.org/fx"
)

type CellApplication struct {
	event eventbus.ICommonEventBus
}

func NewCellApplication(event eventbus.ICommonEventBus) *CellApplication {
	ret := &CellApplication{event: event}
	return ret
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
