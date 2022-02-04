package config

import (
	"sync"
)

type IConfigRefresher interface {
	start()
	OnModuleChanged(m string)
	RegisterListener(moduleName string, listener func(), parser IConfigurationParser)
}

type ConfigRefresher struct {
	checkIntervalSeconds int
	running              bool
	m                    *Manager

	mtx                    sync.RWMutex
	configurationListeners map[string]*ConfigRefreshData

	w *RecursiveWatcher
}

type ConfigRefreshData struct {
	listeners    []func()
	lastEditDate int64
	//file         *os.File
	parser IConfigurationParser
}

func newConfigRefresher(interval int, path string, f func(f string), d func(dir string)) *ConfigRefresher {
	ret := &ConfigRefresher{}
	ret.checkIntervalSeconds = interval
	ret.configurationListeners=make(map[string]*ConfigRefreshData)
	w, err := NewRecursiveWatcher(path, f, d)
	if nil != err {
		panic(err)
	}
	ret.w = w
	return ret
}

// func (this *ConfigRefresher) RegisterListener(moduleName string, filePath string, listener func(), parser IConfigurationParser) {

func (this *ConfigRefresher) RegisterListener(moduleName string, listener func(), parser IConfigurationParser) {
	this.mtx.Lock()
	defer this.mtx.Unlock()
	data := this.configurationListeners[moduleName]
	if data == nil {
		data = &ConfigRefreshData{}
		this.configurationListeners[moduleName] = data
	}
	//file, err := os.Open(filePath)
	//if nil != err {
	//	panic(err)
	//}
	//data.file = file
	data.listeners = append(data.listeners, listener)
	//stat, err := data.file.Stat()
	//if nil != err {
	//	panic(err)
	//}
	//data.lastEditDate = stat.ModTime().Unix()
	data.parser = parser
}

func (this *ConfigRefresher) start() {
	if this.running {
		return
	}
	go this.run()
}
func (this *ConfigRefresher) run() {
	this.running = true
	this.w.run(false)
}

func (this *ConfigRefresher) OnModuleChanged(m string) {
	li := this.configurationListeners[m]
	if nil == li {
		return
	}
	for _, l := range li.listeners {
		l()
	}
}
