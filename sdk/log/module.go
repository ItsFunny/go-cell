/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/7 10:28 下午
# @File : module.go
# @Description :
# @Attention :
*/
package log

import (
	"fmt"
	"strings"
)

type Module interface {
	fmt.Stringer
	Index() uint16
	LogLevel() Level
}

// FIXME name 需要提供compare
type module struct {
	name  string
	index uint16
	level Level
}

func NewModuleWithLevel(name string, index uint16, level Level) module {
	name = strings.ToUpper(name)
	return module{
		index: index,
		name:  name,
		level: level,
	}
}

func (m module) String() string {
	return m.name
}

func (m module) Index() uint16 {
	return m.index
}
func (m module) LogLevel() Level {
	return m.level
}
