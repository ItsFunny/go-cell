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
	_ IInputArchive = (*JSONInputArchive)(nil)
	_ IInputArchive = (*ByteJSONInputArchive)(nil)
)

type JSONInputArchive struct {
}

func (j *JSONInputArchive) ReadByte() ([]byte, error) {
	panic("implement me")
}

func (j *JSONInputArchive) ReadBool() (bool, error) {
	panic("implement me")
}

func (j *JSONInputArchive) ReadBoolSlice() ([]bool, error) {
	panic("implement me")
}

func (j *JSONInputArchive) ReadInt32() (int32, error) {
	panic("implement me")
}

func (j *JSONInputArchive) ReadInt32Slice() ([]int32, error) {
	panic("implement me")
}

func (j *JSONInputArchive) ReadInt() (int, error) {
	panic("implement me")
}

func (j *JSONInputArchive) ReadIntSlice() ([]int, error) {
	panic("implement me")
}

func (j *JSONInputArchive) ReadInt64() (int64, error) {
	panic("implement me")
}

func (j *JSONInputArchive) ReadInt64Slice() ([]int64, error) {
	panic("implement me")
}

func (j *JSONInputArchive) ReadFloat32() (float32, error) {
	panic("implement me")
}

func (j *JSONInputArchive) ReadFloat32Slice() ([]float32, error) {
	panic("implement me")
}

func (j *JSONInputArchive) ReadString() (string, error) {
	panic("implement me")
}

func (j *JSONInputArchive) ReadStringSlice(strings []string, err error) {
	panic("implement me")
}

/////

type ByteJSONInputArchive struct {
	data []byte
}

func NewByteJSONInputArchive(data []byte) *ByteJSONInputArchive {
	return &ByteJSONInputArchive{data: data}
}

func (b *ByteJSONInputArchive) ReadByte() ([]byte, error) {
	return b.data, nil
}

func (b *ByteJSONInputArchive) ReadBool() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (b *ByteJSONInputArchive) ReadBoolSlice() ([]bool, error) {
	//TODO implement me
	panic("implement me")
}

func (b *ByteJSONInputArchive) ReadInt32() (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (b *ByteJSONInputArchive) ReadInt32Slice() ([]int32, error) {
	//TODO implement me
	panic("implement me")
}

func (b *ByteJSONInputArchive) ReadInt() (int, error) {
	//TODO implement me
	panic("implement me")
}

func (b *ByteJSONInputArchive) ReadIntSlice() ([]int, error) {
	//TODO implement me
	panic("implement me")
}

func (b *ByteJSONInputArchive) ReadInt64() (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (b *ByteJSONInputArchive) ReadInt64Slice() ([]int64, error) {
	//TODO implement me
	panic("implement me")
}

func (b *ByteJSONInputArchive) ReadFloat32() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (b *ByteJSONInputArchive) ReadFloat32Slice() ([]float32, error) {
	//TODO implement me
	panic("implement me")
}

func (b *ByteJSONInputArchive) ReadString() (string, error) {
	//TODO implement me
	panic("implement me")
}

func (b *ByteJSONInputArchive) ReadStringSlice(strings []string, err error) {
	//TODO implement me
	panic("implement me")
}
