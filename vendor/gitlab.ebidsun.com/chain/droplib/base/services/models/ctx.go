/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/5/16 1:34 下午
# @File : ctx.go
# @Description :
# @Attention :
*/
package models

import (
	"context"
	"gitlab.ebidsun.com/chain/droplib/base/services/constants"
)

// FIXME ,错误位置
type StartCTX struct {
	Ctx  context.Context
	Flag constants.START_FLAG
	// 用于临时传递变量
	Value     map[string]interface{}
	PreStart  func()
	PostStart func()
}

func (this *StartCTX) With(ops ...StartOption) StartOption {
	for _, opt := range ops {
		opt(this)
	}
	return startOption(this)
}

var startOption = func(copy *StartCTX) StartOption {
	return func(c *StartCTX) {
		c.Ctx = copy.Ctx
		c.Value = copy.Value
		c.PostStart = copy.PostStart
		c.PreStart = copy.PreStart
		c.Flag = copy.Flag
	}
}

type ResetCTX struct {
	value map[string]interface{}
}
func NewResetCTX()*ResetCTX{
	r:=&ResetCTX{
		value: make(map[string]interface{}),
	}
	return r
}

func (this *ResetCTX) GetValue(key string) interface{} {
	if this.value == nil {
		return nil
	}
	return this.value[key]
}

type ReadyCTX struct {
	Ctx   context.Context
	value map[string]interface{}

	ReadyFlag constants.READY_FALG

	PreReady  func()
	PostReady func()
}

func (this *ReadyCTX) With(ops ...ReadyOption) ReadyOption {
	for _, opt := range ops {
		opt(this)
	}
	return readyOption(this)
}
func readyOption(copy *ReadyCTX) ReadyOption {
	return func(c *ReadyCTX) {
		c.Ctx = copy.Ctx
		c.value = copy.value
		c.PostReady = copy.PostReady
		c.PreReady = copy.PreReady
		c.ReadyFlag = copy.ReadyFlag
	}
}
func (this *ReadyCTX) GetValue(key string) interface{} {
	if nil == this.value {
		return nil
	}
	return this.value[key]
}

type StopCTX struct {
	Force bool
	Value map[string]interface{}
}

// @deprected
func (this *StartCTX) GetValue(key string) interface{} {
	// if nil == this.Ctx {
	// 	return nil
	// }
	// return this.Ctx.Value(key)
	if nil == this.Value {
		return nil
	}
	return this.Value[key]
}

func (this *StartCTX) GetValueFromMap(key string) interface{} {
	if nil == this.Value {
		return nil
	}
	return this.Value[key]
}
