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
	storage map[int]struct{}
}

func NewCache() *mutexCache {
	return &mutexCache{
		storage: make(map[int]struct{}),
	}
}

func (c *mutexCache) Set(key int) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.storage[key] = struct{}{}
	return nil
}

func (c *mutexCache) CheckKey(key int) error {
	_, ok := c.storage[key]
	if !ok {
		return ErrNotFound
	}
	return nil
}

func (c *mutexCache) GetAll() ([]int, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	allKeys := make([]int, 0, len(c.storage))
	for keys, _ := range c.storage {
		allKeys = append(allKeys, keys)
	}
	return allKeys, nil
}

func (c *mutexCache) Delete(key int) error {
	c.mx.Lock()
	defer c.mx.Unlock()
	delete(c.storage, key)
	return nil
}
