/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/16 5:29 下午
# @File : summary.go
# @Description :
# @Attention :
*/
package summary

import "github.com/itsfunny/go-cell/base/reactor"

var (
	_ reactor.ISummary = (*HttpSummary)(nil)
)

type HttpSummary struct {
	reactor.BaseSummary
}
