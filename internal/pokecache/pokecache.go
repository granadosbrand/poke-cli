package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntry map[string]cacheEntry
	mux        *sync.RWMutex
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()

	ce := cacheEntry{
		val:       val,
		createdAt: time.Now(),
	}

	c.cacheEntry[key] = ce

}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mux.RLock()
	defer c.mux.RUnlock()

	v, ok := c.cacheEntry[key]

	if !ok {
		return nil, false
	}
	return v.val, true

}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			now := time.Now()
			c.mux.Lock()
			for key, entry := range c.cacheEntry {
				if now.Sub(entry.createdAt) > interval {
					delete(c.cacheEntry, key) // Eliminar entrada si es más antigua que el intervalo
				}
			}
			c.mux.Unlock()
		}
	}

}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) *Cache {

	ch := &Cache{
		cacheEntry: map[string]cacheEntry{},
		mux: &sync.RWMutex{},
	}

	go ch.reapLoop(interval)

	return ch
}
