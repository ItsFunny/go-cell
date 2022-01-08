/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 12:51 下午
# @File : serialize.go
# @Description :
# @Attention :
*/
package serialize

type ISerialize interface {
	Read(archive IInputArchive)error
	ToBytes() []byte
	FromBytes([]byte)
}