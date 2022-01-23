/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/23 10:47 上午
# @File : command.go
# @Description :
# @Attention :
*/
package swagger

import (
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/framework/http/couple"
	"github.com/swaggo/files"
)

var s = &reactor.Command{
	ProtocolID: "/swagger/*any",
	PreRun:     nil,
	Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
		req := ctx.GetCommandContext().ServerRequest.(*couple.HttpServerRequest)
		resp := ctx.GetCommandContext().ServerResponse.(*couple.HttpServerResponse)
		swaggerFiles.Handler.ServeHTTP(resp.Writer.GetInternalWriter(), req.Request)
		return nil
	},
	PostRun: nil,
	RunType: reactor.RunTypeHttp,
	Options: nil,
}
