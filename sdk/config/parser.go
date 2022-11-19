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

	return j.newJsonValue(configuration, moduleName, obj, data)
}

func (j *JSONParser) newJsonValue(cfg *Configuration, moduleName string, jsonObject interface{}, originBytes []byte) (*ConfigValueJson, error) {
	key := &JSONKey{ModuleName: moduleName, Obj: jsonObject}
	v := j.jsonValues[key]
	if nil != v {
		return v, nil
	}

	switch vv := jsonObject.(type) {
	case *gojson.JsonObject:
		moduleObj := vv.GetJsonObject("module")
		if moduleObj != nil && moduleObj.Value != nil && len(moduleObj.GetString()) > 0 {
			ret, err := cfg.GetConfigValue(moduleObj.GetString())
			if nil != err {
				return nil, err
			}
			return ret.(*ConfigValueJson), nil
		}
	}
	jsonValue := newConfigValueJson(jsonObject, cfg, moduleName, originBytes)
	j.jsonValues[key] = jsonValue
	return jsonValue, nil
}
