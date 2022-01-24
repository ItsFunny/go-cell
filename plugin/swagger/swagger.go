/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/24 9:24 下午
# @File : swagger.go
# @Description :
# @Attention :
*/
package swaggerss

type Swagger struct {
	Swag                string                `json:"swagger"`
	Info                Info                  `json:"info"`
	Host                string                `json:"host"`
	BasePath            string                `json:"basePath"`
	Scheme              Scheme                `json:"scheme"`
	Consumes            []string              `json:"consumes"`
	Produces            []string              `json:"produces"`
	SecurityRequirement []SecurityRequirement `json:"security"`
	// Paths map[string]
}
