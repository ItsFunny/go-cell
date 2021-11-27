/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/25 9:33 下午
# @File : proxu.go
# @Description :
# @Attention :
*/
package proxy

import "github.com/itsfunny/go-cell/sdk/event_driven/event"

type IProxy interface {
	proxy(event event.IProcessEvent) error
}

var (
	_ IProxy = (*BaseProxy)(nil)
)

type BaseProxy struct {
}

func (b *BaseProxy) proxy(event event.IProcessEvent) error {
	return b.OnProxy(event)
}
func (b *BaseProxy) OnProxy(event event.IProcessEvent) error {
	return nil
}
