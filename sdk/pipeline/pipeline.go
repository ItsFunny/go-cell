/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/17 9:26 下午
# @File : pipeline.go
# @Description :
# @Attention :
*/
package pipeline

import "reflect"

var (
	_ Pipeline = (*SingleEngine)(nil)
	_ Pipeline = (*Engine)(nil)
)

type Pipeline interface {
	Serve(data interface{})
	RegisterFunc(d reflect.Type, fs ...HandlerFunc)
}
