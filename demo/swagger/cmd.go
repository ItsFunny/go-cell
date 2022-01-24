/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/24 10:42 下午
# @File : cmd.go
# @Description :
# @Attention :
*/
package main

import (
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/itsfunny/go-cell/base/common/protocol"
	"github.com/itsfunny/go-cell/base/core/options"
	"github.com/itsfunny/go-cell/base/reactor"
)

var demoCmd = &reactor.Command{
	ProtocolID: "/demo",
	PreRun: func(req reactor.IBuzzContext) error {
		fmt.Println("asd")
		return nil
	},
	Run: func(ctx reactor.IBuzzContext, reqData interface{}) error {
		ctx.Response(ctx.CreateResponseWrapper().WithRet(23))
		return nil
	},
	PostRun: nil,
	RunType: reactor.RunTypeHttpPost,
	Options: []options.Option{
		options.StringOption("id").WithRequired(true),
		options.BoolOption("bbbb"),
	},
	Description: "asdd",
	MetaData: reactor.MetaData{
		Description: "assss",
		Produces: []string{
			protocol.HttpApplicationJson,
		},
		Tags: []string{
			"tags",
		},
		Summary: "summary",
		Response: map[int]spec.ResponseProps{
			200: {
				Description: "ok",
				Schema:      nil,
				Headers:     nil,
				Examples:    nil,
			},
		},
	},
}
