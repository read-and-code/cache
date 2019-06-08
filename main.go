package main

import (
	"github.com/read-and-code/cache/cache"
	"github.com/read-and-code/cache/http"
	"github.com/read-and-code/cache/tcp"
)

func main() {
	cacheProvider := cache.NewCache("inmemory")

	go tcp.NewServer(cacheProvider).Listen(12346)
	http.NewServer(cacheProvider).Listen(12345)
}
