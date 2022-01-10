/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/10 10:13 下午
# @File : option.go
# @Description :
# @Attention :
*/
package promise

import (
	"context"
	"time"
)

type PromiseOntion func(promise *Promise)

func WithTimeOut(tt time.Duration) PromiseOntion {
	return func(promise *Promise) {
		promise.flag |= timeout
		promise.ctx, promise.cancel = context.WithTimeout(promise.ctx, tt)
	}
}

func WithCancel() PromiseOntion {
	return func(promise *Promise) {
		promise.flag |= enableCancel
		promise.ctx, promise.cancel = context.WithCancel(promise.ctx)
	}
}
