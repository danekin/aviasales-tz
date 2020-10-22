package dictionary

import (
	"sync"
)

type CachedDictionary struct {
	dict  []string
	mutex sync.RWMutex
}

func NewCachedDictionary() *CachedDictionary {
	return &CachedDictionary{
		mutex: sync.RWMutex{},
	}
}

func (c *CachedDictionary) Get() []string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	return c.dict
}

func (c *CachedDictionary) Set(dict []string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.dict = dict
}
