/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 12:57 下午
# @File : json_object.go
# @Description :
# @Attention :
*/
package serialize

import (
	"github.com/tidwall/gjson"
)

type IJSONArray interface {
	Size() int
	GetByteValue(index int) byte
	GetBoolValue(index int) bool
	GetIntValue(index int) int
	GetFloatValue(index int) float32
	GetString(index int) string
	GetObject(index int) IJSONObject
}
type IJSONObject interface {
	GetJSONArray(tag string) IJSONArray

	GetJSONObject(tag string) IJSONObject
	GetBuffer(tag string) []byte
	Keys() []string

	GetByteValue(tag string) byte
	GetByteSlice(tag string) []byte
	GetBoolValue(tag string) bool
	GetBoolSlice(tag string) []bool
	GetIntValue(tag string) int
	GetIntSlice(tag string) []int
	GetFloatValue(tag string) float32
	GetFloatSlice(tag string) []float32
	GetString(tag string) string
	GetStringSlice(tag string) []string
}

var (
	_ IJSONObject = (*FastJSONObject)(nil)
)

type FastJSONObject struct {
	j gjson.Result
}

func (f *FastJSONObject) GetJSONArray(tag string) IJSONArray {
	panic("implement me")
}

func (f *FastJSONObject) GetJSONObject(tag string) IJSONObject {
	panic("implement me")
}

func (f *FastJSONObject) GetBuffer(tag string) []byte {
	panic("implement me")
}

func (f *FastJSONObject) Keys() []string {
	panic("implement me")
}

func (f *FastJSONObject) GetByteValue(tag string) byte {
	panic("implement me")
}

func (f *FastJSONObject) GetByteSlice(tag string) []byte {
	panic("implement me")
}

func (f *FastJSONObject) GetBoolValue(tag string) bool {
	panic("implement me")
}

func (f *FastJSONObject) GetBoolSlice(tag string) []bool {
	panic("implement me")
}

func (f *FastJSONObject) GetIntValue(tag string) int {
	panic("implement me")
}

func (f *FastJSONObject) GetIntSlice(tag string) []int {
	panic("implement me")
}

func (f *FastJSONObject) GetFloatValue(tag string) float32 {
	panic("implement me")
}

func (f *FastJSONObject) GetFloatSlice(tag string) []float32 {
	panic("implement me")
}

func (f *FastJSONObject) GetString(tag string) string {
	panic("implement me")
}

func (f *FastJSONObject) GetStringSlice(tag string) []string {
	panic("implement me")
}
