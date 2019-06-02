package tcp

import (
	"github.com/read-and-code/cache/cache"
	"net"
)

type Server struct {
	cache cache.Cache
}

func (server *Server) Listen() {
	listener, err := net.Listen("tcp", ":12346")

	if err != nil {
		panic(err)
	}

	for {
		connection, err := listener.Accept()

		if err != nil {
			panic(err)
		}

		go server.process(connection)
	}
}

func NewServer(cache cache.Cache) *Server {
	return &Server{cache}
}