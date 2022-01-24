/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/24 7:24 下午
# @File : config.go
# @Description :
# @Attention :
*/
package swagger

import "html/template"

type swaggerConfig struct {
	URL                      string
	DeepLinking              bool
	DocExpansion             string
	DefaultModelsExpandDepth int
	Oauth2RedirectURL        template.JS
	Title                    string
}
type Config struct {
	//The url pointing to API definition (normally swagger.json or swagger.yaml). Default is `doc.json`.
	URL                      string
	DeepLinking              bool
	DocExpansion             string
	DefaultModelsExpandDepth int
	InstanceName             string
	Title                    string
}

func (c Config) ToSwaggerConfig() swaggerConfig {
	return swaggerConfig{
		URL:                      c.URL,
		DeepLinking:              c.DeepLinking,
		DocExpansion:             c.DocExpansion,
		DefaultModelsExpandDepth: c.DefaultModelsExpandDepth,
		Oauth2RedirectURL: template.JS(
			"`${window.location.protocol}//${window.location.host}$" +
				"{window.location.pathname.split('/').slice(0, window.location.pathname.split('/').length - 1).join('/')}" +
				"/oauth2-redirect.html`",
		),
		Title: c.Title,
	}
}
