package cache

import (
	externalCache "github.com/patrickmn/go-cache"
	"time"
)

type Cache interface {
	Set(k string, x interface{})
	Retrieve(k string) (interface{}, bool)
}

type cache struct {
	instance *externalCache.Cache
}

func (c *cache) Set(k string, x interface{}) {
	c.instance.Set(k, x, externalCache.NoExpiration)
}

func (c *cache) Retrieve(k string) (value interface{}, b bool) {
	value, b = c.instance.Get(k)
	return value, b
}

func New(defaultExpiration, cleanupInterval time.Duration) Cache {
	return &cache{
		instance: externalCache.New(defaultExpiration, cleanupInterval),
	}
}
