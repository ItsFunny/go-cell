/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/11 3:24 下午
# @File : router.go
# @Description :
# @Attention :
*/
package pipeline

import (
	"math"
	"reflect"
	"runtime"
)


const abortIndex int8 = math.MaxInt8 / 2

type IRoute interface {
}

type RouterGroup struct {
	Handlers HandlersChain
}


func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}
func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain {
	finalSize := len(group.Handlers) + len(handlers)
	if finalSize >= int(abortIndex) {
		panic("asd")
	}
	mergedHandlers := make(HandlersChain, finalSize)
	copy(mergedHandlers, group.Handlers)
	copy(mergedHandlers[len(group.Handlers):], handlers)
	return mergedHandlers
}

