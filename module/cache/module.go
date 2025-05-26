package cache

import (
	"github.com/hashicorp/golang-lru"
	"go.uber.org/fx"
)

var Module = fx.Module("cache", fx.Provide(newLru))

func newLru() (*lru.Cache, error) {
	return lru.New(100)
}
