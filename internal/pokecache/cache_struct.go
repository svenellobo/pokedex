package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheVault map[string]cacheEntry
	interval   time.Duration
	mu         sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
