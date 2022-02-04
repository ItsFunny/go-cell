package config

import (
	logsdk "github.com/itsfunny/go-cell/sdk/log"
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
	"sync"
)

var (
	managerModule = logsdk.NewModule("configuration_manager", 1)
)

type Manager struct {
	mtx        sync.RWMutex
	rootPath   string
	configType string

	newCfg *Configuration
	cur    *Configuration

	refresher IConfigRefresher

	OnFileCreateOrModified func(f string)
	OnDirCreate            func(dir string)
}

func (this *Manager) GetCurrentConfiguration() *Configuration {
	if this.newCfg != nil {
		this.mtx.Lock()
		defer this.mtx.Unlock()
		if nil != this.newCfg {
			prev := this.cur
			this.cur = this.newCfg
			this.newCfg = nil
			prev.flush()
		}
	}
	return this.cur
}

func NewManager(rootPath string, configType string) *Manager {
	ret := &Manager{newCfg: nil}
	ret.cur = NewConfiguration(ret)
	ret.rootPath = rootPath
	ret.configType = configType
	ret.OnFileCreateOrModified = ret.onFile
	ret.OnDirCreate = ret.onDir

	return ret
}
func (this *Manager) initialize() {
	this.mtx.Lock()
	defer this.mtx.Unlock()

	this.refresher = newConfigRefresher(REFRESH_CHECK_INTERVAL_SECONDS,
		this.rootPath,
		this.OnFileCreateOrModified,
		this.OnDirCreate)
	cur := this.GetCurrentConfiguration()
	if cur.initialized {
		panic("asdd")
	}
	cur.initialize(this.rootPath, this.configType)
	this.refresher.start()
}

func (this *Manager) onFile(f string) {
	logrusplugin.MInfo(managerModule, "detected new file,begin refresh configuration ", "name", f)
	this.refreshConfiguration()
	cfg := this.GetCurrentConfiguration()
	for k, v := range cfg.modules {
		if len(v.ModuleDuePath) == 0 {
			continue
		}
		if v.ModuleFullPath==f || v.ModuleDuePath==f{
			this.refresher.OnModuleChanged(k)
		}
	}
}

func (this *Manager) refreshConfiguration() {
	newCfg := NewConfiguration(this)
	newCfg.initialize(this.rootPath, this.configType)
	this.mtx.Lock()
	defer this.mtx.Unlock()
	if this.newCfg != nil {
		return
	}
	this.newCfg = newCfg
}
func (this *Manager) onDir(dir string) {
	logrusplugin.MInfo(managerModule, "detected new directory", "name", dir)
}
