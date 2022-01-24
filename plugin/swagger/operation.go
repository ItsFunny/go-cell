/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/24 9:29 下午
# @File : operation.go
# @Description :
# @Attention :
*/
package swaggerss

type Operation struct {
	Tags        []string `json:"tags"`
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
	OperationId string   `json:"operationId"`

	Schemes []Scheme `json:"schemes"`
	Consumes []string
	Produces []string
	parameters []Parameter


}
