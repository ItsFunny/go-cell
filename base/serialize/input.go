/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 12:52 下午
# @File : input.go
# @Description :
# @Attention :
*/
package serialize

type IInputArchive interface {
	ReadByte() ([]byte,error)
	ReadBool() (bool,error)
	ReadBoolSlice() ([]bool,error)

	ReadInt32()(int32,error)
	ReadInt32Slice() ([]int32,error)

	ReadInt()(int,error)
	ReadIntSlice()([]int,error)

	ReadInt64()(int64,error)
	ReadInt64Slice()([]int64,error)

	ReadFloat32()(float32,error)
	ReadFloat32Slice()([]float32,error)

	ReadString()(string,error)
	ReadStringSlice([]string,error)

}
