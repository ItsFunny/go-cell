/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/11 3:25 下午
# @File : funcs.go
# @Description :
# @Attention :
*/
package pipeline

type HandlerFunc func(ctx *Context)
type HandlersChain []HandlerFunc

func (c HandlersChain) Last() HandlerFunc {
	if length := len(c); length > 0 {
		return c[length-1]
	}
	return nil
}

