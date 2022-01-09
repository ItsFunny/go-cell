/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 4:07 下午
# @File : mock.go
# @Description :
# @Attention :
*/
package mock

//
//var (
//	_ couple.IServerResponse = (*MockServerResponse)(nil)
//	_ couple.IServerRequest  = (*MockServerRequest)(nil)
//)
//
//type MockServerRequest struct {
//}
//
//func (m *MockServerRequest) ContentLength() int {
//	return 0
//}
//
//func (m *MockServerRequest) GetHeader(name string) string {
//	return ""
//}
//
//type MockServerResponse struct {
//	header map[string]string
//	status int
//	ret    chan<- interface{}
//	err    error
//}
//
//func (m *MockServerResponse) SetHeader(name, value string) {
//	m.header[name] = value
//}
//
//func (m *MockServerResponse) SetStatus(status int) {
//	m.status = status
//}
//
//func (m *MockServerResponse) AddHeader(name, value string) {
//	m.header[name] = value
//}
//
//func (m *MockServerResponse) FireResult(ret interface{}) {
//	m.ret <- ret
//}
//
//func (m *MockServerResponse) FireError(e error) {
//	m.err = e
//}
//
//func(m *MockServerResponse)SetOrExpired() bool{
//	return false
//}