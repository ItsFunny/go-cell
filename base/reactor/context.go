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
	"fmt"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
)

var (
	_ IBuzzContext = (*BaseBuzzContext)(nil)
)

type IBuzzContext interface {
	logsdk.Logger
	Response(wrapper *ContextResponseWrapper)

	GetCommandContext() *CommandContext

	PostRunType() PostRunType
}

type BaseBuzzContext struct {
	CommandContext *CommandContext

	PostType PostRunType
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
	panic("implement me")
}
