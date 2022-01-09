/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2021/10/11 4:48 下午
# @File : engine_test.go.go
# @Description :
# @Attention :
*/
package pipeline

import (
	"fmt"
	"reflect"
	"testing"
)

type AA struct {
	Name string
}

func TestNew(t *testing.T) {
	eg := New()
	a := func(c *Context) {
		fmt.Println(c.Request, "zzzzz")
	}
	b := func(c *Context) {
		fmt.Println(c.Request, "--")
	}
	eg.RegisterFunc(reflect.TypeOf(AA{}), a, b)
	eg.Serve(AA{
		Name: "asd",
	})
}
