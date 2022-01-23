/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/27 4:11 下午
# @File : types.go
# @Description :
# @Attention :
*/
package reactor

import (
	"github.com/itsfunny/go-cell/base/couple"
)

type Function func(ctx IBuzzContext,reqData interface{}) error

//func FunctionWithBuz(ctx IBuzzContext, bo interface{}) Function {
//	return func(ctx IBuzzContext) error {
//		req:=ctx.GetCommandContext().ServerRequest
//
//	}
//}

//func FuncWithBuz(GetInputArchiveFromCtxFunc func(ctx IBuzzContext) serialize.IInputArchive,
//	factory func() serialize.ISerialize) Function {
//	return func(ctx IBuzzContext) error {
//		archive := GetInputArchiveFromCtxFunc(ctx)
//		req := factory()
//
//	}
//}

type PreRun func(req IBuzzContext) error

type PostRunMap map[RunType]func(response couple.IServerResponse) error
