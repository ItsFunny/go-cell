package config

import "github.com/ChengjinWu/gojson"

type rootConfig struct {
	types         *gojson.JsonObject
	defaultType   string
	configs       []*gojson.JsonObject
	plugins       *gojson.JsonObject
	configTypes   map[string]string
	configModules map[string]*ConfigModule
	configPlugins []string
}

func (this *rootConfig) getModules() map[string]*ConfigModule {
	if len(this.configModules) == 0 {

	}
	this.configModules = make(map[string]*ConfigModule)

	for i := 0; i < len(this.configs); i++ {
		obj := this.configs[i]
		module := obj.GetJsonObject("modules")
		schema := module.GetJsonObject("schema").GetString()
		if len(schema) == 0 {
			schema = "json"
		}
		for k, v := range module.Attributes {
			this.configModules[k] = newConfigModule(v.GetString(), nil, schema)
		}
	}
	return this.configModules
}

func (this *rootConfig) getConfigTypes() map[string]string {
	ret := make(map[string]string)
	if this.configTypes != nil {
		return this.configTypes
	}
	this.configTypes = make(map[string]string)

	attr := this.types.Attributes
	for k, v := range attr {
		parent := v.GetJsonObject("parent").GetString()
		this.configTypes[k] = parent
	}
	return ret
}

// TODO
func(this *rootConfig)getPluginPathes(filePath string)[]string{
	return nil
}