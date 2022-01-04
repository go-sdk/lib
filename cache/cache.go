package cache

import (
	"runtime"
	"sync"
	"time"
)

const DefaultExpiration time.Duration = 0

type MemoryCache struct {
	defaultExpiration time.Duration

	mu      sync.RWMutex
	items   map[string]MemoryCacheItem
	cleaner *MemoryCacheCleaner
}

func NewMemoryCache(defaultExpiration time.Duration, items map[string]MemoryCacheItem) *MemoryCache {
	return NewMemoryCacheWithCleaner(defaultExpiration, 0, items)
}

func NewMemoryCacheWithCleaner(defaultExpiration, cleanInterval time.Duration, items map[string]MemoryCacheItem) *MemoryCache {
	c := &MemoryCache{defaultExpiration: defaultExpiration}
	if len(items) > 0 {
		c.items = items
	} else {
		c.items = map[string]MemoryCacheItem{}
	}
	if cleanInterval > 0 {
		c.cleaner = &MemoryCacheCleaner{
			interval: cleanInterval,
			stop:     make(chan bool, 1),
		}
		go startMemoryCacheCleaner(c)
		runtime.SetFinalizer(c, stopMemoryCacheCleaner)
	}
	return c
}

type MemoryCacheItem struct {
	Object     interface{}
	Expiration int64
}

func (item MemoryCacheItem) Expired() bool {
	if item.Expiration == 0 {
		return false
	}
	return time.Now().UnixNano() > item.Expiration
}

type MemoryCacheCleaner struct {
	interval time.Duration
	stop     chan bool
}

func startMemoryCacheCleaner(c *MemoryCache) {
	ticker := time.NewTicker(c.cleaner.interval)
	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		case <-c.cleaner.stop:
			ticker.Stop()
			return
		}
	}
}

func stopMemoryCacheCleaner(c *MemoryCache) {
	c.cleaner.stop <- true
}

func (c *MemoryCache) Set(key string, value interface{}, expiration time.Duration) {
	if expiration == DefaultExpiration {
		expiration = c.defaultExpiration
	}
	item := MemoryCacheItem{Object: value}
	if expiration > 0 {
		item.Expiration = time.Now().Add(expiration).UnixNano()
	}
	c.mu.Lock()
	c.items[key] = item
	c.mu.Unlock()
}

func (c *MemoryCache) SetDefault(key string, value interface{}) {
	c.Set(key, value, DefaultExpiration)
}

func (c *MemoryCache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exist := c.items[key]
	if !exist || value.Expired() {
		return nil, false
	}
	return value.Object, true
}

func (c *MemoryCache) GetExpiration(key string) (interface{}, time.Time, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exist := c.items[key]
	if !exist || value.Expired() {
		return nil, time.Time{}, false
	}
	t := time.Time{}
	if value.Expiration > 0 {
		t = time.Unix(0, value.Expiration)
	}
	return value, t, true
}

func (c *MemoryCache) Delete(keys ...string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for i := 0; i < len(keys); i++ {
		delete(c.items, keys[i])
	}
}

func (c *MemoryCache) DeleteExpired() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, item := range c.items {
		if item.Expired() {
			delete(c.items, k)
		}
	}
}

func (c *MemoryCache) Items() map[string]interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	m := make(map[string]interface{}, len(c.items))
	for k, item := range c.items {
		if item.Expired() {
			continue
		}
		m[k] = item.Object
	}
	return m
}

func (c *MemoryCache) Size() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}

func (c *MemoryCache) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = map[string]MemoryCacheItem{}
}
