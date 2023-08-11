package cache

import (
	"sync"
	"time"
)

type Cache struct {
	store sync.Map
	key   string
	ttl   time.Duration
}

func NewCache(ttl time.Duration, key string) *Cache {
	return &Cache{
		ttl: ttl,
		key: key,
	}
}

func (c *Cache) Set(value interface{}) {
	c.store.Store(c.key, value)
	time.AfterFunc(c.ttl, func() {
		c.store.Delete(c.key)
	})
}

func (c *Cache) Get() (interface{}, bool) {
	return c.store.Load(c.key)
}
