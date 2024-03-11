package cache

import (
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("not found value")
)

type mutexCache struct {
	mx      sync.RWMutex
	storage map[int64]any
}

func NewCache() *mutexCache {
	return &mutexCache{
		storage: make(map[int64]any),
	}
}

func (c *mutexCache) Set(key int64, value any) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.storage[key] = value
	return nil
}

func (c *mutexCache) Get(key int64) (any, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	value, ok := c.storage[key]
	if !ok {
		return nil, ErrNotFound
	}
	return value, nil
}

func (c *mutexCache) GetAllKeys() ([]int64, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	allKeys := make([]int64, 0, len(c.storage))
	for key, _ := range c.storage {
		allKeys = append(allKeys, key)
	}
	return allKeys, nil
}

func (c *mutexCache) GetLen() (int, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	return len(c.storage), nil
}

func (c *mutexCache) Delete(key int64) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.storage, key)
	return nil
}
