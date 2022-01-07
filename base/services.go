/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/23 5:43 下午
# @File : services.go
# @Description :
# @Attention :
*/
package base

import v2 "gitlab.ebidsun.com/chain/droplib/base/log/v2"

type IBaseService interface {
	BStart(ctx ...StartOption) error
	OnStart(ctx *StartCTX) error


	BStop(ctx ...StopOption) error
	OnStop(ctx *StopCTX)

	BReady(ctx ...ReadyOption) error
	OnReady(ctx *ReadyCTX) error

	Reset(ctx ...ResetOption) error
	OnReset(cts *ResetCTX) error

	IsRunning() bool

	Quit() <-chan struct{}

	String() string

	//SetLogger(logger v2.Logger)
}
