package utils

import (
	"sync"
	"sync/atomic"
)

// Cache is a simple in-memory thread-safe cache.
type Cache struct {
	store sync.Map
	hKey  atomic.Uint64

	// maxItems is the maximum number of items in the cache.
	maxItems uint64
}

// NewCache creates a new Cache instance.
func NewCache(maxCount uint64) *Cache {
	return &Cache{
		store:    sync.Map{},
		hKey:     atomic.Uint64{},
		maxItems: maxCount,
	}
}

// Set sets the value for the given key.
func (c *Cache) Set(key uint64, value interface{}) {
	c.store.Store(key, value)

	if c.hKey.Load() < key {
		c.hKey.Store(key)
	}

	c.store.Range(func(key, value interface{}) bool {
		if key.(uint64) < c.hKey.Load()-c.maxItems {
			c.store.Delete(key)
		}
		return true
	})
}

// Get returns the value for the given key.
func (c *Cache) Get(key uint64) (interface{}, bool) {
	return c.store.Load(key)
}

// GetHeadKey returns the head key.
func (c *Cache) GetHeadKey() uint64 {
	return c.hKey.Load()
}
