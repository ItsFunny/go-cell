/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/11 4:17 下午
# @File : error.go
# @Description :
# @Attention :
*/
package errs

type ErrorType uint64

type Error struct {
	Err  error
	Type ErrorType
	Meta interface{}
}

func (e Error) Error() string {
	return e.Err.Error()
}
