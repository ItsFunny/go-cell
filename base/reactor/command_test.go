/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 4:06 下午
# @File : command_test.go.go
# @Description :
# @Attention :
*/
package reactor

import (
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/itsfunny/go-cell/base/common/protocol"
	"github.com/itsfunny/go-cell/base/core/options"
	"testing"
)

func TestRun(t *testing.T) {
}

var dummyCommand = &Command{
	ProtocolID: "/dummy",
	property:   CommandProperty{},
	RunType:    RunTypeHttpPost,
	Options: []options.Option{
		options.StringOption("id").WithDefault("123").WithRequired(true),
		options.BoolOption("bo").WithDefault(false).WithRequired(true),
	},
	Description: "",
	MetaData: MetaData{
		Description: "swagger demo",
		Produces: []string{
			protocol.HttpApplicationJson,
		},
		Tags: []string{
			"demo接口",
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

func TestCommand_ToSwagger(t *testing.T) {
	node := dummyCommand.ToSwaggerPath()
	fmt.Println(node)
}
