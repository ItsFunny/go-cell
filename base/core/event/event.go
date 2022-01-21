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

type CallBack struct {
	CB func()
}

func (c CallBack) CallBack() {
	if nil != c.CB {
		c.CB()
	}
}
