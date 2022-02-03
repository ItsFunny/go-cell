package config

import (
	"fmt"
	"testing"
)

func TestInalize(t *testing.T) {
	cfg := NewConfiguration()
	cfg.initialize("/Users/lvcong/go/src/github.com/itsfunny/go-cell/sdk/config/demo/config", "test-asd")
	type A struct {
		ServerAddr string `json:"serverAddr"`
	}
	a := &A{}
	err := cfg.getConfigValue("nacos.properties").AsObject(a)
	fmt.Println(err)
	fmt.Println(a)
}

// 会读取最新的配置文件
func TestIndentent(t *testing.T) {
	cfg := NewConfiguration()
	cfg.initialize("/Users/lvcong/go/src/github.com/itsfunny/go-cell/sdk/config/demo/config", "test-asd2")
	type A struct {
		ServerAddr string `json:"serverAddr"`
	}
	a := &A{}
	err := cfg.getConfigValue("nacos.properties").AsObject(a)
	fmt.Println(err)
	fmt.Println(a)
}
