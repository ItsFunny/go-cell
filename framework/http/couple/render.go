/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/10 9:40 下午
# @File : render.go
# @Description :
# @Attention :
*/
package couple

import (
	"github.com/itsfunny/go-cell/base/common"
	"net/http"
)

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header[common.CONTENT_TYPE]; len(val) == 0 {
		header[common.CONTENT_TYPE] = value
	}
}
