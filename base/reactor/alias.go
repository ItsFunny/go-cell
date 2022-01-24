/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 12:50 下午
# @File : alias.go
# @Description :
# @Attention :
*/
package reactor

type AliasRequestType byte

type AliasResponseType byte


type ProtocolID string

func ProtocolIDFromString(str string) ProtocolID {
	return ProtocolID(str)
}

func(this ProtocolID)String()string{
	return string(this)
}