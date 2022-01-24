/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/24 8:51 下午
# @File : swagger.go
# @Description :
# @Attention :
*/
package reactor

type SwaggerInfo struct {
	Description string      `json:"description"`
	Producers   []string    `json:"producers"`
	Tags        []string    `json:"tags"`
	Summary     string      `json:"summary"`
	Parameters  []Parameter `json:"parameters"`
}
type Parameter struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Name        string `json:"name"`
	In          string `json:"in"`
	Required    bool   `json:"required"`
}
