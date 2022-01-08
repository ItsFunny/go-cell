/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 11:01 上午
# @File : context.go
# @Description :
# @Attention :
*/
package context

type IContext interface {
	Discard()
	Done() bool

}
