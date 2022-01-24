/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/23 8:05 上午
# @File : di.go
# @Description :
# @Attention :
*/
package swagger

import (
	"github.com/itsfunny/go-cell/di"
	"go.uber.org/fx"
)

var (
	SwaggerModule di.OptionBuilder = func() fx.Option {
		return fx.Options(
			di.RegisterExtension(newSwaggerExtension),
			di.RegisterCommand(cmd),
		)
	}
)
