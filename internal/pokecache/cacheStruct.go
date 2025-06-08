package pokecache

import (
	"sync"
	"time"
)

type PokeCache struct {
	cache map[string]cacheEntry
	mu    *sync.Mutex
}

type cacheEntry struct {
	createdAT time.Time
	val       []byte
}

func NewCache(interval time.Duration) cacheEntry {
	return cacheEntry{
		createdAT: time.Now(),
		val:       nil,
	}
}
