/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/17 10:30 下午
# @File : json.go
# @Description :
# @Attention :
*/
package render

type JSON struct {
	Data interface{}
}

var jsonContentType = []string{"application/json; charset=utf-8"}

func (J JSON) Render(w RenderWriter) (err error) {
	if err = WriteJSON(w, J.Data); err != nil {
		panic(err)
	}
	return
}

func (J JSON) WriteContentType(response RenderWriter) {
	writeContentType(response, jsonContentType)
}
