/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/26 5:59 上午
# @File : listener.go
# @Description :
# @Attention :
*/
package listener

import "github.com/itsfunny/go-cell/component/base"

type IListenerComponent interface {
	base.IComponent
	RegisterListener(topic ...string) <-chan interface{}
	NotifyListener(data interface{}, listenerIds ...string)
}
