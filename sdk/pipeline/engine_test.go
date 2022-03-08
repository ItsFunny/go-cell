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
	"encoding/hex"
	"fmt"
	"github.com/tendermint/tendermint/crypto/ed25519"
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

func TestA(t *testing.T) {
	a1 := "70080211340000000000000022480a202b7562bc264a103f11cad8af41704a001757efd623c0ffd1a0e800b57bbf906d122408011220b7e2183a8cb7e6a7c84c1e05afd32d9a1ada54960341c1c6c59bf6c82541c1d42a0c08c4d3f8900610c8f8ed9c01320b6578636861696e2d313031"
	a2 := "70080211340000000000000022480a20d18b4a3f496bac5b72139e36ff88ede8f6385612824c5ceeec0ec145a777eefa122408011220b7e2183a8cb7e6a7c84c1e05afd32d9a1ada54960341c1c6c59bf6c82541c1d42a0c08c4d3f8900610c8f8ed9c01320b6578636861696e2d313031"
	a1B, _ := hex.DecodeString(a1)
	a2B, _ := hex.DecodeString(a2)
	fmt.Println(string(a1B))
	fmt.Println(string(a2B))

}

func TestParse(t *testing.T) {
	pubB, _ := hex.DecodeString("498acbaa1a77f7d24392717e7a46aae734279911fd86ab73442d79d57065d70a")
	signB, _ := hex.DecodeString("6f1bf2001db0eae4aacdd4d131a9a649a233c615f64a274956b43887f76f22e9b1bce067b1ad8417a4156dd645278dd2d04be5cfeea5e3ff9e2f2d5957e49c00")
	msg1B, _ := hex.DecodeString("70080211130000000000000022480a205e0857b3b4d790069c9ce6a670b266cebca99ba1a31d4aab3e8632afebbcead512240801122007e8b3cef8564dbf64a777d41458299c95adb11050b8b1a0aca1bb8c8ff20be22a0c08ccc9fa900610e8a1898202320b6578636861696e2d313031")
	msg2B, _ := hex.DecodeString("70080211130000000000000022480a205e0857b3b4d790069c9ce6a670b266cebca99ba1a31d4aab3e8632afebbcead512240a2007e8b3cef8564dbf64a777d41458299c95adb11050b8b1a0aca1bb8c8ff20be210012a0c08ccc9fa900610e8a1898202320b6578636861696e2d313031")
	msg3B, _ := hex.DecodeString("70080211130000000000000022480a2067d223782b792f314f580dc136be052dbc8b7e72e26b252a2f91143bf1ac7c3512240a2007e8b3cef8564dbf64a777d41458299c95adb11050b8b1a0aca1bb8c8ff20be210012a0c08ccc9fa900610e8a1898202320b6578636861696e2d313031")
	var pubkeyBytes [ed25519.PubKeyEd25519Size]byte
	copy(pubkeyBytes[:], pubB[32:])
	pp := ed25519.PubKeyEd25519(pubkeyBytes)
	ok := pp.VerifyBytes(msg1B, signB)
	fmt.Println(ok)
	ok = pp.VerifyBytes(msg2B, signB)
	fmt.Println(ok)
	ok = pp.VerifyBytes(msg3B, signB)
	fmt.Println(ok)
}
