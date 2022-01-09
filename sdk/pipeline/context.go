/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/11 4:14 下午
# @File : context.go
# @Description :
# @Attention :
*/
package pipeline

import "github.com/itsfunny/go-cell/base/common/errs"

type Context struct {
	index    int8
	handlers HandlersChain
	Errors   errorMsgs
	Request  interface{}

}

func (c *Context) Next() {
	c.index++
	for ; c.index < int8(len(c.handlers));{
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Error(err error) *errs.Error {
	if err == nil {
		panic("err is nil")
	}

	parsedError, ok := err.(*errs.Error)
	if !ok {
		parsedError = &errs.Error{
			Err:  err,
			Type: 0,
		}
	}

	c.Errors = append(c.Errors, parsedError)
	return parsedError
}

func (c *Context) Abort() {
	c.index = abortIndex
}

func (c *Context) reset() {
	c.handlers = nil
	c.index = -1
	c.Errors = c.Errors[0:0]
}