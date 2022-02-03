package config

import (
	"fmt"
	"testing"
)

func TestInalize(t *testing.T) {
	cfg := NewConfiguration()
	cfg.initialize("/Users/lvcong/go/src/github.com/itsfunny/go-cell/sdk/config/demo/config", "test-asd")
	value := cfg.getConfigValue("nacos.properties")
	fmt.Println(value)
}
