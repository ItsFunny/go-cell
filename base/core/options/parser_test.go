/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/18 8:39 下午
# @File : parser_test.go.go
# @Description :
# @Attention :
*/
package options

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParse(t *testing.T) {
	ops := []Option{
		StringOption("name", "n", "user name"),
		StringOption("address", "a", "addr"),
	}
	cmds := []string{
		"--name", "joker",
		"-a", "浙江",
	}
	m := make(map[string]Option)
	for _, v := range ops {
		for _, n := range v.Names() {
			m[n] = v
		}
	}
	res, err := Parse(cmds, m)
	require.Equal(t, res["name"], "joker")
	require.Equal(t, res["address"], "浙江")
	require.NoError(t, err)
}
