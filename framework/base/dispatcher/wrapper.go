/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 11:00 上午
# @File : wrapper.go
# @Description :
# @Attention :
*/
package dispatcher

import "github.com/itsfunny/go-cell/base/reactor"

type CommandWrapper struct {
	Command reactor.ICommand
}