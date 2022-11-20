package config

import (
	"encoding/json"
	"github.com/itsfunny/go-cell/component/codec/types"
)

var (
	_ types.Unmarshaler = (*HttpConfiguration)(nil)
	_ types.Marshaller  = HttpConfiguration{}
)

type HttpConfiguration struct {
	IP       string `json:"ip"`
	Port     uint   `json:"port"`
	Protocol string `json:"protocol"`
}

func (h HttpConfiguration) Marshal() ([]byte, error) {
	return json.Marshal(h)
}

func (h *HttpConfiguration) Unmarshal(data []byte) error {
	return json.Unmarshal(data, h)
}

func DefaultHttpConfiguration() *HttpConfiguration {
	return &HttpConfiguration{
		IP:       "0.0.0.0",
		Port:     8080,
		Protocol: "http",
	}
}
