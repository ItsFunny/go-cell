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
type Data struct {
	ContentType string
	Data        []byte
}

func (d *Data) Render(writer RenderWriter) error {
	panic("implement me")
}

func (d *Data) WriteContentType(writer RenderWriter) {
	panic("implement me")
}
