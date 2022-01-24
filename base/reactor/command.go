/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/27 2:04 下午
# @File : command.go
# @Description :
# @Attention :
*/
package reactor

import (
	"errors"
	"github.com/itsfunny/go-cell/base/common"
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/serialize"
)

var (
	_ ICommand = (*Command)(nil)
)

type RunType int32

const (
	RunTypeHttp = 1 << 0
	RunTypeRpc  = 1 << 1
	RunTypeCli  = 1 << 2
)

type ICommand interface {
	ID() ProtocolID
	Execute(ctx IBuzzContext)
	SupportRunType() RunType
}

type ICommandSerialize interface {
	serialize.ISerialize
	common.IMessage
}

type Command struct {
	ProtocolID ProtocolID
	PreRun     PreRun
	Run        Function
	PostRun    PostRunMap

	property CommandProperty

	RunType RunType

	Options []options.Option
}

func (c *Command) ID() ProtocolID {
	return c.ProtocolID
}
func (c *Command) SupportRunType() RunType {
	return c.RunType
}
func (c *Command) Execute(ctx IBuzzContext) {
	if c.PreRun!=nil{
		if err := c.PreRun(ctx); nil != err {
			ctx.Response(c.CreateResponseWrapper().
				WithStatus(common.FAIL).WithError(err))
			return
		}
	}

	async := c.property.Async
	if async {
		panic("not supported yet")
	} else {
		c.fire(ctx)
	}
}

func (c *Command) fire(ctx IBuzzContext) {
	defer func() {
		if !ctx.Done() {
			ctx.Response(c.CreateResponseWrapper().WithError(errors.New("missing ret")))
		}
	}()
	req, err := c.newInstance(ctx)
	if nil != err {
		ctx.Error("获取参数失败", "err", err)
		return
	}
	if err := c.Run(ctx, req); nil != err {
		ctx.Error("调用失败", "err", err)
	}
	post := c.PostRun[ctx.PostRunType()]
	if nil != post {
		if err := post(ctx.GetCommandContext().ServerResponse); nil != err {
			ctx.Error("postRun失败", "err", err)
			ctx.Response(c.CreateResponseWrapper().WithError(err))
		}
	}
}

func (c *Command) newInstance(ctx IBuzzContext) (ICommandSerialize, error) {
	if nil == c.property.RequestDataCreateF {
		return nil, nil
	}
	if c.property.GetInputArchiveFromCtxFunc == nil {
		return nil, nil
	}

	reqBo := c.property.RequestDataCreateF()
	if err := reqBo.Read(c.property.GetInputArchiveFromCtxFunc(ctx)); nil != err {
		return nil, err
	}

	return reqBo, reqBo.ValidateBasic()
}

func (c *Command) CreateResponseWrapper() *ContextResponseWrapper {
	ret := &ContextResponseWrapper{Cmd: c}
	return ret
}
