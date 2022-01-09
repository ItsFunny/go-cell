/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/11 4:16 下午
# @File : error.go
# @Description :
# @Attention :
*/
package pipeline

import (
	"errors"
	"github.com/itsfunny/go-cell/base/common/errs"
)

var (
	COMMAND_NOT_EXIST  = errors.New("as")
	ERROR_FALL_THROUGH = &errs.Error{
		Err:  errors.New("asd"),
		Type: ErrorTypeNext,
	}
)

const (
	ErrorTypeNext    errs.ErrorType = 1 << 0
	ErrorTypePrivate errs.ErrorType = 1 << 1
)

type errorMsgs []*errs.Error
