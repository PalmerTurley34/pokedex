package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]CacheEntry
	mu           *sync.Mutex
	Ticker       *time.Ticker
	reapInterval time.Duration
}

func NewCache(interval time.Duration) Cache {
	ticker := time.NewTicker(interval)
	returnCache := Cache{
		map[string]CacheEntry{},
		&sync.Mutex{},
		ticker,
		interval,
	}
	go returnCache.reapLoop()
	return returnCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	c.cacheEntries[key] = CacheEntry{time.Now(), val}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.cacheEntries[key]
	c.mu.Unlock()
	if !ok {
		return []byte{}, false
	}
	return entry.value, true
}

func (c *Cache) reapLoop() {
	for tickTime := range c.Ticker.C {
		reapMin := tickTime.Add(-c.reapInterval)
		for key, entry := range c.cacheEntries {
			if entry.createdAt.Before(reapMin) {
				c.mu.Lock()
				delete(c.cacheEntries, key)
				c.mu.Unlock()
			}
		}
	}
}

type CacheEntry struct {
	createdAt time.Time
	value     []byte
}
