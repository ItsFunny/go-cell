package config

type HttpConfiguration struct {
	IP   string `json:"ip"`
	Port uint   `json:"port"`
}

func DefaultHttpConfiguration() *HttpConfiguration {
	return &HttpConfiguration{
		IP:   "0.0.0.0",
		Port: 8080,
	}
}
