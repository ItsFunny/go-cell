/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 11:08 上午
# @File : constants.go
# @Description :
# @Attention :
*/
package common

import (
	"strconv"
)

const (
	RESPONSE_HEADER_CODE = "code"
	RESPONSE_HEADER_MSG  = "msg"
)

const (
	CONTENT_TYPE = "Content-Type"
)

const (
	SUCCESS = int(1 << 0)
	FAIL    = 1 << 1
	TIMEOUT = 1<<2 | FAIL
)

var (
	STRING_FAIL = strconv.Itoa(FAIL)
)

func IsTimeOut(v int) bool {
	return v&TIMEOUT >= TIMEOUT
}

func IsSuccess(v int) bool {
	return v&SUCCESS >= SUCCESS
}
