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
	"errors"
	"fmt"
	"github.com/itsfunny/go-cell/base/common"
	"github.com/itsfunny/go-cell/base/context"
	"github.com/itsfunny/go-cell/base/couple"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
	"strconv"
	"time"
)

var (
	_ IBuzzContext = (*BaseBuzzContext)(nil)
)

type FireResult func(response couple.IServerResponse, ret *ContextResponseWrapper)

type IBuzzContext interface {
	context.IContext
	logsdk.Logger
	Response(wrapper *ContextResponseWrapper)

	GetCommandContext() *CommandContext

	PostRunType() PostRunType

	Module() logsdk.Module
}

type BaseBuzzContext struct {
	*context.BaseContext
	CommandContext *CommandContext

	PostType PostRunType

	impl IBuzzContext
}

func NewBaseBuzzContext(commandContext *CommandContext,
	postType PostRunType,
	impl IBuzzContext) *BaseBuzzContext {
	ret := &BaseBuzzContext{
		CommandContext: commandContext,
		PostType:       postType,
		impl:           impl,
	}
	ret.BaseContext = context.NewBaseContext(commandContext.Ctx, ret)

	return ret
}

func (b *BaseBuzzContext) Module() logsdk.Module {
	return b.impl.Module()
}

func (b *BaseBuzzContext) PostRunType() PostRunType {
	return b.PostType
}

func (b *BaseBuzzContext) prefix(msg string) string {
	return fmt.Sprintf("sequenceId=%s,protocolId=%s,msg=%s",
		b.CommandContext.Summary.GetSequenceId(), b.CommandContext.Summary.GetProtocolId(), msg)
}

func (b *BaseBuzzContext) Info(msg string, keyvals ...interface{}) {
	logrusplugin.Info(b.prefix(msg), keyvals...)
}

func (b *BaseBuzzContext) Panicf(msg string, keyvals ...interface{}) {
	panic(msg)
}

func (b *BaseBuzzContext) Infof(template string, keyvals ...interface{}) {
	logrusplugin.Info(b.prefix(fmt.Sprintf(template, keyvals...)))
}

func (b *BaseBuzzContext) Debug(msg string, keyvals ...interface{}) {
	logrusplugin.Debug(b.prefix(msg), keyvals...)
}

func (b *BaseBuzzContext) Debugf(template string, keyvals ...interface{}) {
	logrusplugin.Debug(b.prefix(fmt.Sprintf(template, keyvals...)))
}

func (b *BaseBuzzContext) Warn(msg string, keyvals ...interface{}) {
	logrusplugin.Warn(b.prefix(msg), keyvals...)
}

func (b *BaseBuzzContext) Warningf(template string, keyvals ...interface{}) {
	logrusplugin.Warn(b.prefix(fmt.Sprintf(template, keyvals...)))
}

func (b *BaseBuzzContext) Error(msg string, keyvals ...interface{}) {
	logrusplugin.Error(b.prefix(msg), keyvals...)
}

func (b *BaseBuzzContext) Errorf(template string, keyvals ...interface{}) {
	logrusplugin.Error(b.prefix(fmt.Sprintf(template, keyvals...)))
}

func (b *BaseBuzzContext) With(fileds map[string]interface{}) logsdk.Logger {
	panic("not supprted")
}

func (b *BaseBuzzContext) GetCommandContext() *CommandContext {
	return b.CommandContext
}

func (b *BaseBuzzContext) Response(wrapper *ContextResponseWrapper) {
	now := time.Now().Unix()
	beginTime := b.CommandContext.Summary.GetReceiveTimeStamp()
	cost := now - beginTime
	seqId := b.CommandContext.Summary.GetSequenceId()
	logrusplugin.MInfo(b.impl.Module(), fmt.Sprintf("response,protocol=%s,ip=%s,sequenceId=%s,cost=%d,ret=%v",
		b.CommandContext.Command.ID(), b.CommandContext.Summary.GetRequestIp(), seqId, cost, wrapper.Ret))

	resp := b.CommandContext.ServerResponse
	if wrapper.Headers != nil {
		for k, v := range wrapper.Headers {
			resp.AddHeader(k, v)
		}
	}

	resp.AddHeader(common.RESPONSE_HEADER_CODE, strconv.Itoa(wrapper.Status))
	resp.AddHeader(common.RESPONSE_HEADER_MSG, wrapper.Msg)

	if wrapper.Error != nil {
		resp.FireError(wrapper.Error)
		return
	}

	if resp.SetOrExpired() {
		b.Error("duplicate result", "resp", resp)
		resp.FireError(errors.New("duplicate call response"))
		return
	}

	// TODO: other
	if common.IsTimeOut(wrapper.Status) {
		logrusplugin.MWarn(b.impl.Module(), "超时:xxx")
	}

	resp.FireResult(wrapper)
}

func FireResultWithSuccessOrFail(succ, fail FireResult) FireResult {
	return func(response couple.IServerResponse, ret *ContextResponseWrapper) {
		if common.IsSuccess(ret.Status) {
			succ(response, ret)
		} else {
			fail(response, ret)
		}
	}
}

func FireResultWithHook(fs ...FireResult) FireResult {
	return func(response couple.IServerResponse, ret *ContextResponseWrapper) {
		for _, f := range fs {
			f(response, ret)
		}
		if response.SetOrExpired() {
			return
		}
		response.FireResult(ret.Ret)
	}
}
