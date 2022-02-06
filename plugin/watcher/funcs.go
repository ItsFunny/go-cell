/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/10 7:08 上午
# @File : funcs.go
# @Description :
# @Attention :
*/
package watcher

import "github.com/itsfunny/go-cell/structure/channel"

type DataConsumer interface {
	Async() bool
	Handle(channel.IData)
}

type defaultFuncConsumer struct {
	f     func(channel.IData)
	async bool
}

func NewFuncConsumer(f func(channel.IData)) DataConsumer {
	r := &defaultFuncConsumer{f: f}
	return r
}
func (d *defaultFuncConsumer) Async() bool {
	return d.async
}
func (d *defaultFuncConsumer) Handle(i channel.IData) {
	d.f(i)
}
