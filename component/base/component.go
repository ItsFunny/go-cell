/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/26 5:59 上午
# @File : component.go
# @Description :
# @Attention :
*/
package base

import (
	"context"
	"github.com/itsfunny/go-cell/base/core/services"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	"strings"
)

type IComponent interface {
	services.IBaseService
}

type BaseComponent struct {
	*services.BaseService
}

func NewBaseComponent(ctx context.Context, m logsdk.Module, i services.IBaseService) *BaseComponent {
	r := &BaseComponent{
		BaseService: nil,
	}
	name := strings.ToUpper(m.String())
	if !strings.Contains(name, "_COMPONENT") {
		name = name + "_COMPONENT"
		m = logsdk.NewModule(name, m.Index())
	}
	r.BaseService = services.NewBaseService(ctx, nil, m, i)
	return r
}
