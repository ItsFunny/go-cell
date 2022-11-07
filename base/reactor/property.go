/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 1:20 下午
# @File : property.go
# @Description :
# @Attention :
*/
package reactor

import "github.com/itsfunny/go-cell/base/serialize"

type CommandProperty struct {
	Async bool
	//RequestType AliasRequestType
	//ResponseType AliasResponseType
	RequestDataCreateF         func() ICommandSerialize
	GetInputArchiveFromCtxFunc func(ctx IBuzzContext) (serialize.IInputArchive, error)
}
