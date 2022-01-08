/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/25 9:31 下午
# @File : couple.go
# @Description :
# @Attention :
*/
package couple

type IServerRequest interface {
	ContentLength() int
	GetHeader(name string)string
}
type IServerResponse interface {
	SetHeader(name, value string)
	SetStatus(status int)
	AddHeader(name, value string)
	FireResult(ret interface{})
	FireError(e error)
}
