package cache

type CacheStatus struct {
	KeyCount int64

	KeySize int64

	ValueSize int64
}

func (cacheStatus *CacheStatus) add(key string, value []byte) {
	cacheStatus.KeyCount += 1
	cacheStatus.KeySize += int64(len(key))
	cacheStatus.ValueSize += int64(len(value))
}

func (cacheStatus *CacheStatus) delete(key string, value []byte) {
	cacheStatus.KeyCount -= 1
	cacheStatus.KeySize -= int64(len(key))
	cacheStatus.ValueSize -= int64(len(value))
}
