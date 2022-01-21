/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/21 10:55 下午
# @File : context.go
# @Description :
# @Attention :
*/
package application

import (
	"context"
	"time"
)

type valueContext struct {
	context.Context
}

func (valueContext) Deadline() (deadline time.Time, ok bool) { return }
func (valueContext) Done() <-chan struct{}                   { return nil }
func (valueContext) Err() error                              { return nil }

