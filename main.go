package main

import (
	"github.com/read-and-code/cache/cache"
	"github.com/read-and-code/cache/http"
)

func main() {
	cache := cache.NewCache("inmemory")

	http.NewServer(cache).Listen()
}
