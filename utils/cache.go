package utils

import (
	"sync"
	"sync/atomic"
	"time"
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
	for c.hKey.Load()+c.maxItems < key {
		// If the key is too large, we should block the set operation.
		time.Sleep(500 * time.Millisecond)
	}

	c.store.Store(key, value)

	c.store.Range(func(key, value interface{}) bool {
		if key.(uint64) < c.hKey.Load() {
			c.store.Delete(key)
		}
		return true
	})
}

// Get returns the value for the given key.
func (c *Cache) Get(key uint64) (interface{}, bool) {
	c.hKey.Store(key)
	return c.store.Load(key)
}

// GetHeadKey returns the head key.
func (c *Cache) GetHeadKey() uint64 {
	return c.hKey.Load()
}
