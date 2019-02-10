package cache

type Cache interface {
	Set(string, []byte) error

	Get(string) ([]byte, error)

	Delete(string) error

	GetCacheStatus() CacheStatus
}