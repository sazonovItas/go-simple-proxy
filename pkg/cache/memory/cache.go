package memcache

import (
	"context"
	"errors"
	"sync"
	"time"
)

var ErrCacheMissed = errors.New("cache missed")

// Cache is memory storage for cache data of type T
type cache[T any] struct {
	// ctx stops cache garbage collector (that clean expired data)
	ctx context.Context

	// mu makes storage thread-safe
	mu      sync.RWMutex
	storage map[string]item[T]

	// defaultExpiration deafult expiration time for data in storage
	// if defaultExpiration == 0, data will not have expiration time by default
	defaultExpiration time.Duration

	// cleanupInterval interval for checking expired data in storage
	// if cleanupInterval == 0, cache garbage collector will not clean expired data
	cleanupInterval time.Duration
}

// item represents data of type T in storage
type item[T any] struct {
	value      T
	createdAt  int64
	expiration int64
}

// New creates new cache of type T
// if cleanupInterval == 0, cache garbage collector would not clean expired data
// if defaultExpiration == 0, data will not have expiration time by default
func New[T any](
	ctx context.Context,
	defaultExpiration, cleanupInterval time.Duration,
) *cache[T] {
	cache := &cache[T]{
		ctx:               ctx,
		storage:           make(map[string]item[T]),
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}

	cache.startGC()

	return cache
}

// Get returns value under a key, or ErrCacheMissed if key is not found
func (c *cache[T]) Get(key string) (value T, err error) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.storage[key]
	if !found {
		return value, ErrCacheMissed
	}

	return item.value, nil
}

// Set sets value with a key with expiration
// if duration == 0, defaultExpiration will be used
func (c *cache[T]) Set(key string, value T, duration time.Duration) {
	if duration == 0 {
		duration = c.defaultExpiration
	}

	var (
		expiration int64
		now        time.Time = time.Now()
	)
	if duration > 0 {
		expiration = now.Add(duration).UnixNano()
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	c.storage[key] = item[T]{
		value:      value,
		expiration: expiration,
		createdAt:  now.UnixNano(),
	}
}

// Delete deletes data under a keys
func (c *cache[T]) Delete(keys ...string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, key := range keys {
		delete(c.storage, key)
	}
}

// startGC starts cache garbage collector
func (c *cache[T]) startGC() {
	if c.cleanupInterval > 0 {
		go c.startGCWorker()
	}
}

// startGCWorker starts cache garbage collector worker
func (c *cache[T]) startGCWorker() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-time.After(c.cleanupInterval):
			if keys := c.expiredItems(); len(keys) != 0 {
				c.cleanItems(keys)
			}
		}
	}
}

// expiredItems returns keys of expired data
func (c *cache[T]) expiredItems() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	keys := []string{}
	for key, value := range c.storage {
		if value.expiration > 0 && value.expiration < time.Now().UnixNano() {
			keys = append(keys, key)
		}
	}

	return keys
}

// cleanItems cleans items under a keys
func (c *cache[T]) cleanItems(keys []string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for _, key := range keys {
		delete(c.storage, key)
	}
}
