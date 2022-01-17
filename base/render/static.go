/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/17 9:45 下午
# @File : static.go
# @Description :
# @Attention :
*/
package render

import (
	"encoding/json"
	"fmt"
	"github.com/itsfunny/go-cell/base/common"
	"io"
)

// WriteJSON marshals the given interface object and writes it with custom ContentType.
func WriteJSON(w RenderWriter, obj interface{}) error {
	writeContentType(w, jsonContentType)
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&obj)
	return err
}

// WriteString writes data according to its format and write custom ContentType.
func WriteString(w RenderWriter, format string, data []interface{}) (err error) {
	writeContentType(w, plainContentType)
	if len(data) > 0 {
		_, err = fmt.Fprintf(w, format, data...)
		return
	}
	_, err = io.WriteString(w, format)
	return
}

func writeContentType(w RenderWriter, value []string) {
	w.WriteContentType(common.CONTENT_TYPE, value)
}
