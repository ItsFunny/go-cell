package extension

import "github.com/itsfunny/go-cell/base/core/options"

var (
	home       = "cellHome"
	configType = "configType"
	ip         = "ip"

	ipOption         = options.StringOption(ip, ip, "ip address").WithDefault("127.0.0.1")
	homeOption       = options.StringOption(home, home, "配置文件根路径").WithRequired(true)
	configTypeOption = options.StringOption(configType, configType, configType).WithDefault("Default")
)
