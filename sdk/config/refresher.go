package config

import "os"

type IConfigRefresher interface {
	start()
}

type ConfigRefresher struct {
	checkIntervalSeconds int
	running              bool
	configuration        *Configuration
}
type ConfigRefreshData struct {
	listeners    []IConfigListener
	lastEditDate int64
	file         os.File
	parser       IConfigurationParser
}

func newConfigRefresher(interval int, cfg *Configuration) *ConfigRefresher {
	ret := &ConfigRefresher{}
	ret.checkIntervalSeconds = interval
	ret.configuration = cfg
	return ret
}
func(this *ConfigRefresher)start(){

}
