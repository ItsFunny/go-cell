/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/12/25 2:22 下午
# @File : types.go
# @Description :
# @Attention :
*/
package v2

type Mode byte

const (
	// NOT DDD
	MODE_DDD Mode = iota
	MODE_CONCURRENCY
)
