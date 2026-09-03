package pokecache

import "time"

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		cacheVault: make(map[string]cacheEntry),
		interval:   interval,
	}
	go cache.reapLoop()
	return cache

}

func (c *Cache) CacheAdd(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cacheVault[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}

}

func (c *Cache) CacheGet(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, exists := c.cacheVault[key]
	if !exists {
		return nil, false
	}
	return entry.val, true

}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for k, v := range c.cacheVault {
			if time.Since(v.createdAt) > c.interval {
				delete(c.cacheVault, k)
			}

		}
		c.mu.Unlock()
	}
}
