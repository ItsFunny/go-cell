/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 4:07 下午
# @File : mock.go
# @Description :
# @Attention :
*/
package mock

import "github.com/itsfunny/go-cell/base/couple"

var (
	_ couple.IServerResponse = (*MockServerResponse)(nil)
	_ couple.IServerRequest  = (*MockServerRequest)(nil)
)

type MockServerRequest struct {
}

type MockServerResponse struct {
	header map[string]string
	status int
	ret    chan<- interface{}
	err    error
}

func (m *MockServerResponse) SetHeader(name, value string) {
	m.header[name] = value
}

func (m *MockServerResponse) SetStatus(status int) {
	m.status = status
}

func (m *MockServerResponse) AddHeader(name, value string) {
	m.header[name] = value
}

func (m *MockServerResponse) FireResult(ret interface{}) {
	m.ret <- ret
}

func (m *MockServerResponse) FireError(e error) {
	m.err = e
}
