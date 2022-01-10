/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/10 9:40 下午
# @File : render.go
# @Description :
# @Attention :
*/
package couple

import "net/http"

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}
