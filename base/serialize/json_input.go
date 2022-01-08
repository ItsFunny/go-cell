/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 12:56 下午
# @File : json_input.go
# @Description :
# @Attention :
*/
package serialize


var (
	_ IInputArchive=(*JSONInputArchive)(nil)
)
type JSONInputArchive struct {
	
}

func (J *JSONInputArchive) ReadByte() ([]byte, error) {
	panic("implement me")
}

func (J *JSONInputArchive) ReadBool() (bool, error) {
	panic("implement me")
}

func (J *JSONInputArchive) ReadBoolSlice() ([]bool, error) {
	panic("implement me")
}

func (J *JSONInputArchive) ReadInt32() (int32, error) {
	panic("implement me")
}

func (J *JSONInputArchive) ReadInt32Slice() ([]int32, error) {
	panic("implement me")
}

func (J *JSONInputArchive) ReadInt() (int, error) {
	panic("implement me")
}

func (J *JSONInputArchive) ReadIntSlice() ([]int, error) {
	panic("implement me")
}

func (J *JSONInputArchive) ReadInt64() (int64, error) {
	panic("implement me")
}

func (J *JSONInputArchive) ReadInt64Slice() ([]int64, error) {
	panic("implement me")
}

func (J *JSONInputArchive) ReadFloat32() (float32, error) {
	panic("implement me")
}

func (J *JSONInputArchive) ReadFloat32Slice() ([]float32, error) {
	panic("implement me")
}

func (J *JSONInputArchive) ReadString() (string, error) {
	panic("implement me")
}

func (J *JSONInputArchive) ReadStringSlice(strings []string, err error) {
	panic("implement me")
}
