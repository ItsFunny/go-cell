package config

import "github.com/ChengjinWu/gojson"

type rootConfig struct {
	Types         *gojson.JsonObject       `json:"types"`
	DefaultType   string                   `json:"defaultType"`
	Configs       []*gojson.JsonObject     `json:"configs"`
	Plugins       *gojson.JsonObject       `json:"plugins"`
	ConfigTypes   map[string]string        `json:"configTypes"`
	ConfigModules map[string]*ConfigModule `json:"configModules"`
	ConfigPlugins []string                 `json:"configPlugins"`
}

func (this *rootConfig) getModules() map[string]*ConfigModule {
	if len(this.ConfigModules) == 0 {

	}
	this.ConfigModules = make(map[string]*ConfigModule)

	for i := 0; i < len(this.Configs); i++ {
		obj := this.Configs[i]
		module := obj.GetJsonObject("modules")
		schemaJsonObj := obj.GetJsonObject("schema")
		schema := "json"
		if schemaJsonObj.Value != nil && len(schemaJsonObj.GetString()) != 0 {
			schema = schemaJsonObj.GetString()
		}

		for k, v := range module.Attributes {
			this.ConfigModules[k] = newConfigModule(v.GetString(), nil, schema)
		}
	}
	return this.ConfigModules
}

func (this *rootConfig) getConfigTypes() map[string]string {
	if this.ConfigTypes != nil {
		return this.ConfigTypes
	}
	this.ConfigTypes = make(map[string]string)

	attr := this.Types.Attributes
	for k, v := range attr {
		p := v.GetJsonObject("parent")
		if p == nil || p.Value == nil {
			if k == "Default" {
				this.ConfigTypes[k] = ""
			}
			continue
		}
		parent := p.GetString()
		this.ConfigTypes[k] = parent
	}
	return this.ConfigTypes
}

// TODO
func (this *rootConfig) getPluginPathes(filePath string) []string {
	return nil
}
