/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/9 4:46 下午
# @File : data.go
# @Description :
# @Attention :
*/
package render

// Data contains ContentType and bytes data.
type RenderData struct {
	Data        []byte
}

func (r RenderData) Render(writer RenderWriter) (err error) {
	_, err = writer.Write(r.Data)
	return
}
func (r RenderData) WriteContentType(w RenderWriter) {
}
