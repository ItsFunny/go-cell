package config

type ConfigModule struct {
	ModuleFullPath string
	ModuleDuePath  string

	configValue IConfigValue

	Schema string
}

func newConfigModule(moduleFullPath string,configV IConfigValue,schema string) *ConfigModule {
	ret := &ConfigModule{
		ModuleFullPath: moduleFullPath,
		ModuleDuePath:  "",
		configValue:    configV,
		Schema:         schema,
	}
	return ret
}
