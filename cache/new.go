package cache

import "log"

func NewCache(cacheType string) Cache {
	var cache Cache

	if cacheType == "inmemory" {
		cache = newInMemoryCache()
	}

	if cache == nil {
		panic("Unknown cache type " + cacheType)
	}

	log.Println(cacheType, "ready to serve")

	return cache
}