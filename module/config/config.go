package config

type ProxyRule struct {
	PathPrefix string `koanf:"path_prefix"`
	Target     string `koanf:"target"`
}

type SpaProxy struct {
	Root    string      `koanf:"root"`     // URL 路径前缀，比如 "/project1"
	SpaPath string      `koanf:"spa_path"` // SPA 静态文件目录
	Proxy   []ProxyRule `koanf:"proxy"`    // 对应反向代理规则
}

type Config struct {
	Port            int        `koanf:"port"`
	SpaProxies      []SpaProxy `koanf:"spa_proxies"`
	Prefork         bool       `koanf:"prefork"`
	WriteBufferSize int        `koanf:"write.buffer.size"`
	AccessLog       string     `koanf:"access.log.path"`
}
