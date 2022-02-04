package config

import (
	"github.com/ChengjinWu/gojson"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/itsfunny/go-cell/base/common/utils"
	"io/ioutil"
	"path/filepath"
	"sync"
)

var (
	REFRESH_CHECK_INTERVAL_SECONDS = 10
)

type Configuration struct {
	mtx             sync.RWMutex
	parser          map[string]IConfigurationParser
	configModuleMap map[string][]string

	modules map[string]*ConfigModule

	configTypes []string
	repoRoot    string
	configType  string
	initialized bool
	manager     *Manager
}

func NewConfiguration(m *Manager) *Configuration {
	ret := &Configuration{
		parser:          make(map[string]IConfigurationParser),
		configModuleMap: make(map[string][]string),
		modules:         make(map[string]*ConfigModule),
	}
	ret.manager = m
	ret.init()
	return ret
}
func (this *Configuration) flush() {

}
func (this *Configuration) init() {
	this.registerParser("json", newJSONParser())
}

func (this *Configuration) initialize(rootPath string, configType string) {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	if this.initialized {
		panic("asdd")
	}

	this.repoRoot = rootPath
	repos := make(map[string]*rootConfig)
	rootConfigPath := rootPath + string(filepath.Separator) + "root.json"
	bytes, err := ioutil.ReadFile(rootConfigPath)
	if nil != err {
		panic(err)
	}
	rootCfg := new(rootConfig)
	obj, err := gojson.FromBytes(bytes)
	//err = json.Unmarshal(bytes, rootCfg)
	if nil != err {
		panic(err)
	}
	rootCfg.Types = obj.GetJsonObject("types")
	rootCfg.DefaultType = obj.GetJsonObject("defaultType").GetString()
	rootCfg.Configs = obj.GetJsonObject("configs").GetJsonArray()
	rootCfg.Plugins = obj.GetJsonObject("plugins")
	this.configType = configType
	repos[this.repoRoot] = rootCfg

	types := rootCfg.getConfigTypes()
	if _, exist := types[this.configType]; !exist {
		panic("asd")
	}
	inheritance := this.buildInheritanceList(types)
	this.buildModulePathMap(repos, inheritance)

	for k, _ := range types {
		this.configTypes = append(this.configTypes, k)
	}
	this.initialized = true
}
func (this *Configuration) buildModulePathMap(repoMap map[string]*rootConfig, configTypeInheritance []string) {
	for k, v := range repoMap {
		pluginModules := v.getModules()
		for _, v2 := range pluginModules {
			moduleFilePathUnderConfigType := ""
			for _, typeDe := range configTypeInheritance {
				tmpPath := k + string(filepath.Separator) + typeDe + string(filepath.Separator) + v2.ModuleFullPath
				if !utils.CheckFileExists(tmpPath) {
					tmpPath = ""
				}
				if len(v2.ModuleDuePath) == 0 {
					v2.ModuleDuePath = tmpPath
				}
				if len(tmpPath) == 0 {
					continue
				}
				moduleFilePathUnderConfigType = tmpPath
			}
			v2.ModuleFullPath = moduleFilePathUnderConfigType
		}
		for k3, v3 := range pluginModules {
			this.modules[k3] = v3
		}
	}
}
func (this *Configuration) buildInheritanceList(configTypes map[string]string) []string {
	inheritance := arraylist.New()

	inheritance.Add(this.configType)
	for parent := configTypes[this.configType]; len(parent) > 0; parent = configTypes[parent] {
		inheritance.Insert(0, parent)
	}
	it := inheritance.Iterator()
	ret := make([]string, 0)
	for it.Next() {
		ret = append(ret, it.Value().(string))
	}
	return ret
}
func (this *Configuration) registerParser(schema string, parser IConfigurationParser) {
	this.parser[schema] = parser
}
func (this *Configuration) getConfigValue(moduleName string) IConfigValue {
	module := this.getModule(moduleName)
	if module.configValue != nil {
		return module.configValue
	}

	switch module.Schema {
	case "json":
		module.configValue = this.getConfigObject(moduleName)
	}
	return module.configValue
}

func (this *Configuration) getModule(moduleName string) *ConfigModule {
	m := this.modules[moduleName]
	if m == nil {
		panic("module not exists:" + moduleName)
	}
	return m
}

func (this *Configuration) getConfigObject(moduleName string) IConfigValue {
	module := this.getModule(moduleName)
	parser := this.parser[module.Schema]
	if parser == nil {
		panic("asd")
	}
	if len(module.ModuleFullPath) == 0 {
		panic("no valid config file for module [" + moduleName + "]")
	}
	from, err := parser.ParseFrom(this, new(ConfigValueJson), moduleName, module.ModuleFullPath, nil)
	if nil != err {
		panic(err)
	}

	return from
}
