/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/26 6:10 上午
# @File : internal.go
# @Description :
# @Attention :
*/
package extension


type internalExtension struct {
	*BaseExtension
}

func newInternalExtension() INodeExtension {
	ret := &internalExtension{}
	ret.BaseExtension = NewBaseExtension(ret)
	return ret
}