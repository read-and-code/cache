package main

import (
	"cache/cache"
	"cache/http"
)

func main() {
	cache := cache.NewCache("inmemory")

	http.New(cache).Listen()
}
