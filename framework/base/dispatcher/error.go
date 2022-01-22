/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/22 9:18 上午
# @File : error.go
# @Description :
# @Attention :
*/
package dispatcher

import "errors"

var (
	duplicateCommand = errors.New("duplicate cmd")
)
