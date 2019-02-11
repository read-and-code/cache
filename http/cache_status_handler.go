package http

import (
	"encoding/json"
	"log"
	"net/http"
)

type CacheStatusHandler struct {
	server *Server
}

func (cacheStatusHandler *CacheStatusHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodGet {
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	bytes, err := json.Marshal(cacheStatusHandler.server.cache.GetCacheStatus())

	if err != nil {
		log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = responseWriter.Write(bytes)

	if err != nil {
		log.Println(err)

		responseWriter.WriteHeader(http.StatusInternalServerError)

		return
	}
}

func (server *Server) cacheStatusHandler() http.Handler {
	return &CacheStatusHandler{server}
}
