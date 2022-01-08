/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 11:00 上午
# @File : command.go
# @Description :
# @Attention :
*/
package reactor

import "github.com/itsfunny/go-cell/base/context"

type ICommandReactor interface {
	Execute(ctx context.IContext)
}