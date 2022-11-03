/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/11/25 9:35 下午
# @File : error.go
# @Description :
# @Attention :
*/
package common

type ErrorCode int64

type CellError interface {
	error
	Code() ErrorCode
}

type WrappedError struct {
	code ErrorCode
	err  error
}

func NewWrappedError(code ErrorCode, err error) CellError {
	return &WrappedError{code: code, err: err}
}

func (w WrappedError) Error() string {
	return w.err.Error()
}

func (w WrappedError) Code() ErrorCode {
	return w.code
}
