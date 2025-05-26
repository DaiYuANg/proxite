package config

import (
	"fmt"
	"github.com/samber/lo"
	"os"
	"strings"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
	"github.com/knadh/koanf/v2"
)

func load(k *koanf.Koanf, path string) (*Config, error) {
	// 默认值
	if err := k.Load(structs.Provider(defaultConfig, "default"), nil); err != nil {
		return nil, fmt.Errorf("load default config: %w", err)
	}

	// 加载文件
	if path != "" {
		if _, err := os.Stat(path); err == nil {
			if err := k.Load(file.Provider(path), toml.Parser()); err != nil {
				return nil, fmt.Errorf("load config file: %w", err)
			}
		} else if !os.IsNotExist(err) {
			return nil, fmt.Errorf("stat config file: %w", err)
		}
	}

	// 加载环境变量（非 SpaProxies 特殊处理）
	if err := k.Load(env.Provider("PROXITE_", ".", func(s string) string {
		return strings.Replace(
			strings.ToLower(
				strings.TrimPrefix(s, "PROXITE_"),
			),
			"_",
			".",
			-1,
		)
	}), nil); err != nil {
		return nil, fmt.Errorf("load env: %w", err)
	}
	// Unmarshal 除了 SpaProxies 的所有字段
	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	// ============ 自定义处理 SpaProxies 环境变量配置 ============
	projectMap := processProxiesEnv()
	// 把默认/文件配置里的 SpaProxies 与 env 中的合并（按 Root 匹配）
	envSpaProxies := lo.MapToSlice(projectMap, func(_ string, sp *SpaProxy) SpaProxy {
		return *sp
	})
	cfg.SpaProxies = mergeSpaProxies(cfg.SpaProxies, envSpaProxies)

	return &cfg, nil
}

func mergeSpaProxies(old []SpaProxy, env []SpaProxy) []SpaProxy {
	rootMap := make(map[string]SpaProxy)
	lo.ForEach(old, func(sp SpaProxy, index int) {
		rootMap[sp.Root] = sp
	})
	lo.ForEach(env, func(sp SpaProxy, index int) {
		rootMap[sp.Root] = sp
	})
	return lo.Values(rootMap)
}
