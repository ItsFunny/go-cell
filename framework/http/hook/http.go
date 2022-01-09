/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/9 9:32 上午
# @File : http.go
# @Description :
# @Attention :
*/
package hook

import "github.com/itsfunny/go-cell/base/reactor"

func HttpFinalHook(ctx reactor.IBuzzContext) {
	cmdCtx := ctx.GetCommandContext()
	cmd := cmdCtx.Command
	cmd.Execute(ctx)
}
