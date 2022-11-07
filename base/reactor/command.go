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
	"github.com/go-openapi/spec"
	"github.com/itsfunny/go-cell/base/common"
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/serialize"
)

var (
	_ ICommand = (*Command)(nil)
)

type RunType int32

const (
	RunTypeSwagger  = 1 << 0
	RunTypeHttp     = 1<<1 | RunTypeSwagger
	RunTypeRpc      = 1 << 2
	RunTypeCli      = 1 << 3
	RunTypeHttpPost = 1<<4 | RunTypeHttp
	RunTypeHttpGet  = 1<<5 | RunTypeHttp
)

var cmdTypeDesc = map[RunType]string{
	RunTypeHttpPost: "post",
	RunTypeHttpGet:  "get",
}

func getRunTypeDesc(runT RunType) string {
	return cmdTypeDesc[runT]
}

func (r RunType) SupportSwagger() bool {
	return r&RunTypeSwagger >= RunTypeSwagger
}

type ICommand interface {
	ID() ProtocolID
	Execute(ctx IBuzzContext)
	SupportRunType() RunType
	ToSwaggerPath() *PathItemWrapper
	GetOptions() []options.Option
}

type ICommandSerialize interface {
	serialize.ISerialize
	IMessage
}

type Command struct {
	ProtocolID ProtocolID
	PreRun     PreRun
	Run        Function
	PostRun    PostRunMap

	Property CommandProperty

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
func (c *Command) GetOptions() []options.Option {
	return c.Options
}
func (c *Command) Execute(ctx IBuzzContext) {
	if c.PreRun != nil {
		if err := c.PreRun(ctx); nil != err {
			ctx.Response(c.CreateResponseWrapper().
				WithStatus(common.FAIL).WithError(err))
			return
		}
	}

	async := c.Property.Async
	if async {
		panic("not supported yet")
	} else {
		c.fire(ctx)
	}
}

func (c *Command) fire(ctx IBuzzContext) {
	defer func() {
		if !ctx.Done() {
			ctx.Response(ctx.CreateResponseWrapper().WithRet(nil))
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
	if nil == c.Property.RequestDataCreateF {
		return nil, nil
	}
	if c.Property.GetInputArchiveFromCtxFunc == nil {
		return nil, nil
	}

	reqBo := c.Property.RequestDataCreateF()
	archive, err := c.Property.GetInputArchiveFromCtxFunc(ctx)
	if nil != err {
		return nil, err
	}
	if err := reqBo.Read(archive, ctx.GetCommandContext().Codec.GetCodec()); nil != err {
		return nil, err
	}

	return reqBo, reqBo.ValidateBasic(ctx)
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

type PathItemWrapper struct {
	ID       string
	PathItem spec.PathItem
}

func (c *Command) ToSwaggerPath() *PathItemWrapper {
	ret := &PathItemWrapper{}
	// map[string]PathItem
	item := spec.PathItem{}

	// p := swag.New()
	// ops := swag.NewOperation(p)
	op := spec.NewOperation(c.ID().String())
	op.Description = c.MetaData.Description
	op.Produces = c.MetaData.Produces
	op.Tags = c.MetaData.Tags
	op.Summary = c.MetaData.Summary
	//
	// op:=spec.NewOperation(c.ID().String())
	// op.
	for _, opt := range c.Options {
		name := opt.Name()
		param := spec.QueryParam(name)
		param.Description = opt.Description()
		param.Type = opt.Type().String()
		param.Required = opt.Required()
		op.AddParam(param)
	}
	op.Responses = &spec.Responses{
		ResponsesProps: spec.ResponsesProps{
			StatusCodeResponses: map[int]spec.Response{},
		},
	}
	for k, v := range c.MetaData.Response {
		resp := spec.NewResponse()
		resp.ResponseProps = v
		op.Responses.ResponsesProps.StatusCodeResponses[k] = *resp
	}

	if c.RunType == RunTypeHttpPost {
		item.Post = op
	} else {
		item.Get = op
	}
	ret.ID = c.ID().String()
	ret.PathItem = item
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
