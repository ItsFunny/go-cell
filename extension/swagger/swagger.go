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
	"github.com/itsfunny/go-cell/base/node/core/extension"
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
	// cmds := ctx.GetCommands()
	// for _, cmd := range cmds {
	//
	// }
	// p:=swag.New()
	// op:=swag.NewOperation(p)
	// spec.QueryParam()
	// op.AddParam()
	return nil
}
func (b *swaggerExtension) OnExtensionReady(ctx extension.INodeContext) error {
	return nil
}
