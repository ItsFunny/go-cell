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
	"github.com/itsfunny/go-cell/base/common"
)

// String contains the given interface object slice and its format.
type RenderString struct {
	Data string
}

func (r RenderString) WriteContentType(response RenderWriter) {
	response.WriteContentType(common.CONTENT_TYPE, plainContentType)
}

var plainContentType = []string{"text/plain; charset=utf-8"}

func (r RenderString) Render(w RenderWriter) error {
	return WriteString(w, r.Data, nil)
}

func Write(w RenderWriter, data interface{}) error {
	var r Render
	switch v := data.(type) {
	case string:
		r = RenderString{Data: v}
	case []byte:
		r = RenderData{Data: v}
	default:
		r = &JSON{Data: data}
	}
	return r.Render(w)
}
