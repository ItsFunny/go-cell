/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/23 8:04 上午
# @File : swagger.go
# @Description :
# @Attention :
*/
package swagger

import (
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/itsfunny/go-cell/base/node/core/extension"
	"github.com/swaggo/swag"
)

type swaggerExtension struct {
	*extension.BaseExtension
}

func newSwaggerExtension() extension.INodeExtension {
	ret := &swaggerExtension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	return ret
}

func (b *swaggerExtension) OnExtensionInit(ctx extension.INodeContext) error {
	cmds := ctx.GetCommands()
	p := swag.New()
	sg := p.GetSwagger()
	for _, cmd := range cmds {
		if cmd == swgCmd {
			continue
		}
		wrapper := cmd.ToSwaggerPath()
		sg.Paths.Paths[wrapper.ID] = wrapper.PathItem
	}
	sg.Swagger = "2.0"
	sg.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Description: "swagger",
			Title:       "title",
			Contact:     nil,
			License:     nil,
			Version:     "2.0",
		},
	}
	sg.Host=""
	ret, err := sg.MarshalJSON()
	if nil != err {
		return err
	}
	fmt.Println(string(ret))
	swgCmd.docJson = string(ret)
	swgCmd.ready = true
	return nil
}
func (b *swaggerExtension) OnExtensionReady(ctx extension.INodeContext) error {
	return nil
}
