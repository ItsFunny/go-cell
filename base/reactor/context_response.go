/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/27 4:15 下午
# @File : context_response.go
# @Description :
# @Attention :
*/
package reactor

type ContextResponseWrapper struct {
	Status  int
	Msg     string
	Error   error
	Cmd     ICommand
	Ret     interface{}
	Headers map[string]string

	Other interface{}
}
func NewContextResponseWrapper(c ICommand)*ContextResponseWrapper{
	ret:=&ContextResponseWrapper{Cmd: c}
	return ret
}

func (c *ContextResponseWrapper) WithStatus(status int) *ContextResponseWrapper {
	c.Status = status
	return c
}
func (c *ContextResponseWrapper) WithError(err error) *ContextResponseWrapper {
	c.Error=err
	return c
}
func (c *ContextResponseWrapper) WithRet(ret interface{}) *ContextResponseWrapper {
	c.Ret=ret
	return c
}