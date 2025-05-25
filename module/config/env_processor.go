package config

import (
	"github.com/samber/lo"
	"os"
	"strconv"
	"strings"
)

func processProxiesEnv() map[string]*SpaProxy {
	envs := os.Environ()
	projectMap := map[string]*SpaProxy{} // key: PROJECT1, value: SpaProxy 指针
	prefix := "PROXITE_SPA_PROXIES_"

	lo.ForEach(envs, func(e string, index int) {
		if !strings.HasPrefix(e, prefix) {
			return
		}
		kv := strings.SplitN(e, "=", 2)
		if len(kv) != 2 {
			return
		}
		key, val := kv[0], kv[1]
		key = strings.TrimPrefix(key, prefix)
		parts := strings.SplitN(key, "_", 2)
		if len(parts) < 2 {
			return
		}
		projectName := parts[0]
		field := parts[1]

		sp := lo.IfF(projectMap[projectName] != nil, func() *SpaProxy {
			return projectMap[projectName]
		}).ElseF(func() *SpaProxy {
			projectMap[projectName] = &SpaProxy{}
			return projectMap[projectName]
		})

		switch {
		case field == "ROOT":
			sp.Root = val
		case field == "SPA_PATH":
			sp.SpaPath = val
		case strings.HasPrefix(field, "PROXY_"):
			parts := strings.SplitN(field, "_", 3)
			if len(parts) < 3 {
				return
			}
			idx, err := strconv.Atoi(parts[1])
			if err != nil {
				return
			}
			for len(sp.Proxy) < idx {
				sp.Proxy = append(sp.Proxy, ProxyRule{})
			}
			pr := &sp.Proxy[idx-1]
			switch parts[2] {
			case "PATH_PREFIX":
				pr.PathPrefix = val
			case "TARGET":
				pr.Target = val
			}
		}
	})

	return projectMap
}
