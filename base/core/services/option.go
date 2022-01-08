/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/7 10:21 下午
# @File : option.go
# @Description :
# @Attention :
*/
package services

import (
	"context"
)

type StartOption func(c *StartCTX)
type ReadyOption func(c *ReadyCTX)
type StopOption func(c *StopCTX)
type ResetOption func(c *ResetCTX)

func ResetCTXWithKV(key string, value interface{}) ResetOption {
	return func(c *ResetCTX) {
		if c.value == nil {
			c.value = make(map[string]interface{})
		}
		c.value[key] = value
	}
}
func ResetCTXWithResetCtx(ctx *ResetCTX) ResetOption {
	return func(c *ResetCTX) {
		c.value = ctx.value
	}
}
func StartCTXWithKV(key string, value interface{}) StartOption {
	return func(c *StartCTX) {
		if c.Value == nil {
			c.Value = make(map[string]interface{})
		}
		c.Value[key] = value
	}
}
func CtxStartOpt(ctx context.Context) StartOption {
	return func(c *StartCTX) {
		c.Ctx = ctx
	}
}
func PreStartOpt(f func()) StartOption {
	return func(c *StartCTX) {
		c.PreStart = f
	}
}
func PostStartOpt(f func()) StartOption {
	return func(c *StartCTX) {
		c.PostStart = f
	}
}
func AsyncStartWaitReadyOpt(c *StartCTX) {
	c.Flag = ASYNC_START_WAIT_READY
}
func SyncStartWaitReadyOpt(c *StartCTX) {
	c.Flag = SYNC_START_WAIT_READY
}
func SyncStartOpt(c *StartCTX) {
	c.Flag = SYNC_START
}
func AsyncStartOpt(c *StartCTX) {
	c.Flag = ASYNC_START
}

func ReadyWaitStartOpt(c *ReadyCTX) {
	c.ReadyFlag = SYNC_READY_UNTIL_START
}
func ReadyPanicIfErrOpt(c *ReadyCTX) {
	c.ReadyFlag = READY_ERROR_IF_NOT_STARTED
}
func PreReadyOpt(f func()) ReadyOption {
	return func(c *ReadyCTX) {
		c.PreReady = f
	}
}
func PostReadyOpt(f func()) ReadyOption {
	return func(c *ReadyCTX) {
		c.PostReady = f
	}
}

func ReadyAsyncWithUtilStart(c *ReadyCTX) {
	c.ReadyFlag = ASYNC_READY_UTIL_START
}
func ReadyOptWithCtx(ctx context.Context) ReadyOption {
	return func(c *ReadyCTX) {
		c.Ctx = ctx
	}
}
func StartCTXFromReady(ctx *ReadyCTX) StartOption {
	return func(c *StartCTX) {
		c.Ctx = ctx.Ctx
		if nil == c.Value {
			c.Value = make(map[string]interface{})
		}
		if nil != ctx.value {
			for k, v := range ctx.value {
				c.Value[k] = v
			}
		}
	}
}
func ReadyCTXFromStart(start *StartCTX) ReadyOption {
	return func(c *ReadyCTX) {
		c.Ctx = start.Ctx
		if nil == c.value {
			c.value = make(map[string]interface{})
		}
		if nil != start.Value {
			for k, v := range start.Value {
				c.value[k] = v
			}
		}
	}
}
func ReadyCTXWithKV(k string, v interface{}) ReadyOption {
	return func(c *ReadyCTX) {
		if c.value == nil {
			c.value = make(map[string]interface{})
		}
		c.value[k] = v
	}
}

func StopCTXWithForce(c *StopCTX) {
	c.Force = true
}

func StopCTXWithKV(k string, v interface{}) StopOption {
	return func(c *StopCTX) {
		c.Value[k] = v
	}

}
func StopCTXAsChild(cc *StopCTX) StopOption {
	return func(c *StopCTX) {
		if nil != cc.Value {
			if nil == c.Value {
				c.Value = make(map[string]interface{})
				for _k, v := range c.Value {
					c.Value[_k] = v
				}
			}
			c.Force = cc.Force
		}
	}

}
