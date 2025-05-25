package config

var defaultConfig = &Config{
	Port:            9876,
	SpaProxies:      nil,
	Prefork:         false,
	WriteBufferSize: 4096,
}
