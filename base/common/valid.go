/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/8 4:59 下午
# @File : valid.go
# @Description :
# @Attention :
*/
package common

type IMessage interface {
	ValidateBasic() error
}
