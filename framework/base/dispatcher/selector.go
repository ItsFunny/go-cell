/*
# -*- coding: utf-8 -*-
# @Author : joker
# @Time : 2022/1/16 8:36 上午
# @File : selector.go
# @Description :
# @Attention :
*/
package dispatcher

import (
	"github.com/itsfunny/go-cell/base/common/utils"
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/base/couple"
	"github.com/itsfunny/go-cell/base/reactor"
	couple2 "github.com/itsfunny/go-cell/framework/http/couple"
	"path/filepath"
)

type ICommandSelector interface {
	Select(req *CommandSelectReq)
	OnRegisterCommand(wrapper *CommandWrapper)
}

type CommandSelectReq struct {
	Commands map[reactor.ProtocolID]*CommandWrapper
	Request  couple.IServerRequest
	Promise  *promise.Promise
}

type CommandAddReq struct {
	Command *CommandWrapper
}

var (
	_ ICommandSelector = (*UriSelector)(nil)
	_ ICommandSelector = (*antPathSelector)(nil)
)

type UriSelector struct {
	commands map[reactor.ProtocolID]*CommandWrapper
}

func NewUriSelector() ICommandSelector {
	ret := &UriSelector{
		commands: make(map[reactor.ProtocolID]*CommandWrapper),
	}
	return ret
}

func (u *UriSelector) OnRegisterCommand(wrapper *CommandWrapper) {
	u.commands[wrapper.Command.ID()] = wrapper
}

func (u *UriSelector) Select(req *CommandSelectReq) {
	httpReq := req.Request.(*couple2.HttpServerRequest)
	uri := httpReq.Request.URL.Path
	ret := u.commands[reactor.ProtocolIDFromString(uri)]
	if nil != ret {
		req.Promise.Send(ret)
	}
}

// ///

type antPathSelector struct {
	wrappers map[string]*CommandWrapper
}

func NewAntPathSelector() ICommandSelector {
	ret := &antPathSelector{
		wrappers: make(map[string]*CommandWrapper),
	}
	return ret
}
func (a *antPathSelector) Select(req *CommandSelectReq) {
	httpReq := req.Request.(*couple2.HttpServerRequest)
	uri := httpReq.Request.URL.Path
	for path, v := range a.wrappers {
		ok, err := filepath.Match(path, uri)
		if !ok || err != nil {
			continue
		}
		if ok {
			req.Promise.Send(v)
		}
	}
}

func (a *antPathSelector) OnRegisterCommand(wrapper *CommandWrapper) {
	id := wrapper.Command.ID()
	if utils.IsPattern(id.String()) {
		a.wrappers[id.String()] = wrapper
	}
}
