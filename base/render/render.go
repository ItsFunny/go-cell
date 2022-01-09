/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/9 4:44 下午
# @File : render.go
# @Description :
# @Attention :
*/
package render

import (
	"io"
)

var (
	_ Render = (*Data)(nil)
	_ Render = (*RenderString)(nil)
)

type RenderWriter interface {
	io.Writer
	WriteContentType(h string, v []string)
}
type Render interface {
	Render(response RenderWriter) error
	WriteContentType(response RenderWriter)
}
