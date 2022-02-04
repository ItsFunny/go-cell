package config

import (
	"fmt"
	"testing"
	"time"
)

func TestManager_RegisterListener(t *testing.T) {
	mo := "nacos.properties"
	m := NewManager("/Users/lvcong/go/src/github.com/itsfunny/go-cell/sdk/config/demo/config", "test-asd2")
	m.Initialize()
	type A struct {
		ServerAddr string `json:"serverAddr"`
	}
	m.RegisterListener(mo, func() {
		fmt.Println("module update")
	})
	a := &A{}
	err := m.GetCurrentConfiguration().GetConfigValue(mo).AsObject(a)
	fmt.Println(err)
	fmt.Println(a)
	time.Sleep(time.Minute * 360)
}
