package config

import (
	"encoding/json"
)

//type IUnmarshal interface {
//	Unmarshal(v IConfigValue) error
//}
type IConfigValue interface {
	GetArrayObject(index int) interface{}
	GetObjectKeys() []string
	GetObject() IConfigValue
	GetModuleName() string
	AsValueList() []IConfigValue
	// TODO
	AsObject(m interface{}) error
	AsBoolean() bool
	AsByte() byte
	AsInt32() int32
	AsInt() int
	AsInt64() int64
	AsFloat32() float32
	AsFloat64() float64
	AsBytes() []byte
	AsString() string
}

var (
	_ IConfigValue = (*ConfigValueJson)(nil)
)

type BaseConfigValue struct {
	impl                 IConfigValue
	ConfigurationManager *Configuration
	ModuleName           string
}

func (this *BaseConfigValue) GetModuleName() string {
	return this.ModuleName
}

type ConfigValueJson struct {
	*BaseConfigValue
	parseer IConfigurationParser

	originBytes []byte
	data        interface{}
}

func newConfigValueJson(data interface{}, cfg *Configuration, moduleName string, originBytes []byte) *ConfigValueJson {
	ret := &ConfigValueJson{}
	ret.BaseConfigValue = newBaseConfigValue(cfg, ret, moduleName)
	ret.data = data
	ret.originBytes = originBytes
	return ret
}

func newBaseConfigValue(cfg *Configuration, impl IConfigValue, moduleName string) *BaseConfigValue {
	ret := &BaseConfigValue{
		impl:                 impl,
		ConfigurationManager: cfg,
		ModuleName:           moduleName,
	}
	return ret
}

func (c *ConfigValueJson) GetArrayObject(index int) interface{} {
	panic("implement me")
}

func (c *ConfigValueJson) GetObjectKeys() []string {
	panic("implement me")
}

func (c *ConfigValueJson) GetObject() IConfigValue {
	panic("implement me")
}

func (c *ConfigValueJson) AsValueList() []IConfigValue {
	panic("implement me")
}

func (c *ConfigValueJson) AsObject(u interface{}) error {
	return json.Unmarshal(c.originBytes, u)
}

func (c *ConfigValueJson) AsBoolean() bool {
	panic("implement me")
}

func (c *ConfigValueJson) AsByte() byte {
	panic("implement me")
}

func (c *ConfigValueJson) AsInt32() int32 {
	panic("implement me")
}

func (c *ConfigValueJson) AsInt() int {
	panic("implement me")
}

func (c *ConfigValueJson) AsInt64() int64 {
	panic("implement me")
}

func (c *ConfigValueJson) AsFloat32() float32 {
	panic("implement me")
}

func (c *ConfigValueJson) AsFloat64() float64 {
	panic("implement me")
}

func (c *ConfigValueJson) AsBytes() []byte {
	return c.originBytes
}

func (c *ConfigValueJson) AsString() string {
	panic("implement me")
}
