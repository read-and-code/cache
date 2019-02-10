package cache

import "sync"

type InMemoryCache struct {
	keyValuePairs map[string][]byte

	mutex sync.RWMutex

	cacheStatus CacheStatus
}

func (inMemoryCache *InMemoryCache) Set(key string, value []byte) error {
	inMemoryCache.mutex.Lock()

	defer inMemoryCache.mutex.Unlock()

	oldValue, exists := inMemoryCache.keyValuePairs[key]

	if exists {
		inMemoryCache.cacheStatus.delete(key, oldValue)
	}

	inMemoryCache.keyValuePairs[key] = value
	inMemoryCache.cacheStatus.add(key, value)

	return nil
}

func (inMemoryCache *InMemoryCache) Get(key string) ([]byte, error) {
	inMemoryCache.mutex.RLock()

	defer inMemoryCache.mutex.RUnlock()

	return inMemoryCache.keyValuePairs[key], nil
}

func (inMemoryCache *InMemoryCache) Delete(key string) error {
	inMemoryCache.mutex.Lock()

	defer inMemoryCache.mutex.Unlock()

	value, exists := inMemoryCache.keyValuePairs[key]

	if exists {
		delete(inMemoryCache.keyValuePairs, key)

		inMemoryCache.cacheStatus.delete(key, value)
	}

	return nil
}

func (inMemoryCache *InMemoryCache) GetCacheStatus() CacheStatus {
	return inMemoryCache.cacheStatus
}

func newInMemoryCache() *InMemoryCache {
	return &InMemoryCache{make(map[string][]byte), sync.RWMutex{}, CacheStatus{}}
}