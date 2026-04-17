package pokecache

import (
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte 
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	mu sync.Mutex
	interval time.Duration
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mu.Lock()
		for key, item := range c.cacheEntries {
			elapsed := time.Since(item.createdAt)
			if elapsed >  c.interval  {
				delete(c.cacheEntries, key)
			}
		}
		c.mu.Unlock()
	}	
}

func NewCache(interval time.Duration) *Cache {
	newCache := &Cache{
		cacheEntries: map[string]cacheEntry{},
		interval: interval,
	}
	go newCache.reapLoop()

	return newCache
} 

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, exists := c.cacheEntries[key]
	if !exists {
		return nil, exists
	}
	return value.val, exists
}