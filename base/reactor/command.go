/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/27 2:04 下午
# @File : command.go
# @Description :
# @Attention :
*/
package reactor

var (
	_ ICommand = (*Command)(nil)
)

type ICommand interface {
	execute(ctx IBuzzContext) error
}

type Command struct {
	Run   Function
	Async bool
}

func (c *Command) execute(ctx IBuzzContext) error {
	c.Run(ctx)
	if c.Async{

	}
}
