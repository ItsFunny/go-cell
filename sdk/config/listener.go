package config

type IConfigListener interface {
	ConfigRefreshed(configModule string) error
}
