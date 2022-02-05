/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/8/10 7:08 上午
# @File : funcs.go
# @Description :
# @Attention :
*/
package watcher


type DataConsumer interface {
	Async() bool
	Handle(IData)
}

type defaultFuncConsumer struct {
	f     func(IData)
	async bool
}

func NewFuncConsumer(f func(IData)) DataConsumer {
	r := &defaultFuncConsumer{f: f}
	return r
}
func (d *defaultFuncConsumer) Async() bool {
	return d.async
}
func (d *defaultFuncConsumer) Handle(i IData) {
	d.f(i)
}
