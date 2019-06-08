package http

import (
	"github.com/read-and-code/cache/cache"
	"net/http"
	"strconv"
)

type Server struct {
	cache cache.Cache
}

func (server *Server) Listen(port int) {
	http.Handle("/cache/", server.cacheHandler())
	http.Handle("/cacheStatus/", server.cacheStatusHandler())

	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)

	if err != nil {
		panic("Server error")
	}
}

func NewServer(cache cache.Cache) *Server {
	return &Server{cache}
}
