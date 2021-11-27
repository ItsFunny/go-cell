/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/27 2:05 下午
# @File : context.go
# @Description :
# @Attention :
*/
package reactor

import (
	"gitlab.ebidsun.com/chain/droplib/component/routine/services"
)
type IContext interface {
	discard()
	done() bool
}

type IBuzzContext interface {
	response(wrapper ContextResponseWrapper)
	GetRoutine() services.IRoutineComponent
}

var (
	_ IContext = (*EmptyContext)(nil)
)

type EmptyContext struct {
}

func (e *EmptyContext) discard() {
}

func (e *EmptyContext) done() bool {
	return true
}
