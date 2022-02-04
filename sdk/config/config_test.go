package config

import (
	"fmt"
	"testing"
	"time"
)

func TestInalize(t *testing.T) {
	m := NewManager("/Users/lvcong/go/src/github.com/itsfunny/go-cell/sdk/config/demo/config", "test-asd")
	m.Initialize()
	type A struct {
		ServerAddr string `json:"serverAddr"`
	}
	a := &A{}
	err := m.GetCurrentConfiguration().GetConfigValue("nacos.properties").AsObject(a)
	fmt.Println(err)
	fmt.Println(a)
}

// 会读取最新的配置文件
func TestIndentent(t *testing.T) {
	m := NewManager("/Users/lvcong/go/src/github.com/itsfunny/go-cell/sdk/config/demo/config", "test-asd2")
	m.Initialize()
	type A struct {
		ServerAddr string `json:"serverAddr"`
	}
	a := &A{}
	err := m.GetCurrentConfiguration().GetConfigValue("nacos.properties").AsObject(a)
	fmt.Println(err)
	fmt.Println(a)
	time.Sleep(time.Minute * 100)
}
