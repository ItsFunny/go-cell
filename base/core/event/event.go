/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/22 6:19 上午
# @File : event.go
# @Description :
# @Attention :
*/
package event

type ICallBack interface {
	CallBack()
}

func NewCallBack(f func()) *CallBack {
	return &CallBack{CB: f}
}

type CallBack struct {
	CB func()
}

func (c CallBack) CallBack() {
	if nil != c.CB {
		c.CB()
	}
}
