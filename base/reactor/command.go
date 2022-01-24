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
	"github.com/go-openapi/spec"
	"github.com/itsfunny/go-cell/base/common"
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/serialize"
	"github.com/swaggo/swag"
)

var (
	_ ICommand = (*Command)(nil)
)

type RunType int32

const (
	RunTypeHttp     = 1 << 0
	RunTypeRpc      = 1 << 1
	RunTypeCli      = 1 << 2
	RunTypeHttpPost = 1<<3 | RunTypeHttp
	RunTypeHttpGet  = 1<<4 | RunTypeHttp
)

var cmdTypeDesc = map[RunType]string{
	RunTypeHttpPost: "post",
	RunTypeHttpGet:  "get",
}

func getRunTypeDesc(runT RunType) string {
	return cmdTypeDesc[runT]
}

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

	Description string

	MetaData MetaData
}

func (c *Command) ID() ProtocolID {
	return c.ProtocolID
}
func (c *Command) SupportRunType() RunType {
	return c.RunType
}
func (c *Command) Execute(ctx IBuzzContext) {
	if c.PreRun != nil {
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

type MetaData struct {
	Description string
	Produces    []string
	Tags        []string
	Summary     string
	Response    map[int]spec.ResponseProps
}
type ParameterInfo struct {
	parameterType string
	description   string
	name          string
	in            string
	required      bool
}

type ResponseInfo struct {
	Code   string
	Detail ResponseDetail
}
type ResponseDetail struct {
	Description string
	Schema      Schema
}
type Schema struct {
	typeD string
}
type SwaggerNode struct {
	Path   string
	Method string
	Detail string
}

func (c *Command) ToSwagger() SwaggerNode {
	ret := SwaggerNode{}
	p := swag.New()
	op := swag.NewOperation(p)
	op.Description = c.MetaData.Description
	op.Produces = c.MetaData.Produces
	op.Tags = c.MetaData.Tags
	op.Summary = c.MetaData.Summary

	for _, opt := range c.Options {
		name := opt.Name()
		param := spec.QueryParam(name)
		param.Description = opt.Description()
		param.Type = opt.Type().String()
		param.Required = opt.Required()
		op.AddParam(param)
	}

	for k, v := range c.MetaData.Response {
		resp := spec.NewResponse()
		resp.ResponseProps = v
		op.AddResponse(k, resp)
	}

	ret.Path = c.ID().String()
	ret.Method = getRunTypeDesc(c.RunType)
	bs, err := op.MarshalJSON()
	if nil != err {
		panic(err)
	}
	ret.Detail = string(bs)
	return ret
}

// func (c *Command) ToSwagger() string {
// 	params := make([]ParameterInfo, 0)
// 	ret := make(map[string]string)
//
// 	for _, opt := range c.Options {
// 		// paths[c.ID().String()]=
// 		info := ParameterInfo{
// 			parameterType: opt.Type().String(),
// 			description:   opt.Description(),
// 			name:          opt.Name(),
// 			in:            opt.Name(),
// 			required:      opt.Required(),
// 		}
// 		params = append(params, info)
// 	}
// }
