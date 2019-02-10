package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type CacheHandler struct {
	server *Server
}

func (cacheHandler *CacheHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	key := strings.Split(request.URL.EscapedPath(), "/")[2]

	if len(key) == 0 {
		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	method := request.Method

	if method == http.MethodPut {
		bytes, err := ioutil.ReadAll(request.Body)

		if err != nil {
			log.Println(err)

			responseWriter.WriteHeader(http.StatusInternalServerError)

			return
		}

		if len(bytes) != 0 {
			err := cacheHandler.server.cache.Set(key, bytes)

			if err != nil {
				log.Println(err)

				responseWriter.WriteHeader(http.StatusInternalServerError)
			}
		}

		return
	} else if method == http.MethodGet {
		bytes, err := cacheHandler.server.cache.Get(key)

		if err != nil {
			log.Println(err)

			responseWriter.WriteHeader(http.StatusInternalServerError)

			return
		}

		if len(bytes) == 0 {
			responseWriter.WriteHeader(http.StatusNotFound)

			return
		}

		_, err = responseWriter.Write(bytes)

		if err != nil {
			log.Println(err)

			responseWriter.WriteHeader(http.StatusInternalServerError)
		}

		return
	} else if method == http.MethodDelete {
		err := cacheHandler.server.cache.Delete(key)

		if err != nil {
			log.Println(err)

			responseWriter.WriteHeader(http.StatusInternalServerError)
		}

		return
	}

	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (server *Server) cacheHandler() http.Handler {
	return &CacheHandler{server}
}
