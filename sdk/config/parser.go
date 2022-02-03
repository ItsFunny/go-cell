package config

import (
	"github.com/ChengjinWu/gojson"
	"io/ioutil"
)

var (
	_ IConfigurationParser = (*JSONParser)(nil)
)

type JSONKey struct {
	ModuleName string
	Obj        interface{}
}
type IConfigurationParser interface {
	//parseFrom(Configuration configuration, Class<T> clazz, String moduleName, String filePath, Object userData) throws IOException;
	ParseFrom(configuration *Configuration, value interface{}, moduleName string, filePath string, userData interface{}) (IConfigValue, error)
}

type JSONParser struct {
	jsonValues map[*JSONKey]*ConfigValueJson
}

func newJSONParser() *JSONParser {
	ret := &JSONParser{
		jsonValues: make(map[*JSONKey]*ConfigValueJson),
	}

	return ret
}

func (j *JSONParser) ParseFrom(configuration *Configuration, value interface{}, moduleName string, filePath string, userData interface{}) (IConfigValue, error) {
	_, ok := value.(*ConfigValueJson)
	if !ok {
		return nil, nil
	}
	data, err := ioutil.ReadFile(filePath)
	if nil != err {
		return nil, err
	}
	obj, err := gojson.FromBytes(data)
	if nil != err {
		return nil, err
	}

	jsonValue := j.newJsonValue(configuration, moduleName, obj)
	return jsonValue, nil
}

func (j *JSONParser) newJsonValue(cfg *Configuration, moduleName string, jsonObject interface{}) *ConfigValueJson {
	key := &JSONKey{ModuleName: moduleName, Obj: jsonObject}
	v := j.jsonValues[key]
	if nil != v {
		return v
	}

	switch vv := jsonObject.(type) {
	case *gojson.JsonObject:
		subModule := vv.GetJsonObject("module").GetString()
		if len(subModule) > 0 {
			return cfg.getConfigValue(subModule).(*ConfigValueJson)
		}
	}
	jsonValue := newConfigValueJson(jsonObject, cfg, moduleName)
	j.jsonValues[key] = jsonValue
	return jsonValue
}
