package http

import (
	"github.com/read-and-code/cache/cache"
	"net/http"
)

type Server struct {
	cache cache.Cache
}

func (server *Server) Listen() {
	http.Handle("/cache/", server.cacheHandler())
	http.Handle("/cacheStatus/", server.cacheStatusHandler())

	err := http.ListenAndServe(":12345", nil)

	if err != nil {
		panic("Server error")
	}
}

func NewServer(cache cache.Cache) *Server {
	return &Server{cache}
}
