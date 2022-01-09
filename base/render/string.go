/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/9 4:45 下午
# @File : string.go
# @Description :
# @Attention :
*/
package render

import (
	"fmt"
	"github.com/itsfunny/go-cell/base/common/utils"
)

// String contains the given interface object slice and its format.
type RenderString struct {
	Format string
	Data   []interface{}
}

func (r RenderString) WriteContentType(response RenderWriter) {
	response.WriteContentType("Content-Type", plainContentType)
}

var plainContentType = []string{"text/plain; charset=utf-8"}

func (r RenderString) Render(w RenderWriter) error {
	return WriteString(w, r.Format, r.Data)
}

func WriteString(w RenderWriter, format string, data []interface{}) (err error) {
	if len(data) > 0 {
		_, err = fmt.Fprintf(w, format, data...)
		return
	}
	_, err = w.Write(utils.StringToBytes(format))
	return
}
