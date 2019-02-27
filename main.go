package main

import (
	"github.com/read-and-code/cache/cache"
	"github.com/read-and-code/cache/http"
)

func main() {
	cacheProvider := cache.NewCache("inmemory")

	http.NewServer(cacheProvider).Listen()
}
