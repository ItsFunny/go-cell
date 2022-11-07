/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 12:51 下午
# @File : serialize.go
# @Description :
# @Attention :
*/
package serialize

import (
	"github.com/itsfunny/go-cell/component/codec/types"
)

type ISerialize interface {
	Read(archive IInputArchive, cdc types.Codec) error
	ToBytes(cdc types.Codec) ([]byte, error)
	//FromBytes([]byte)
}
