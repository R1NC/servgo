package go_web_service

import (
	"fmt"
	"net/http"
)

type handler struct{
	api *API
}

func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Request: %s From: %s\n", request.URL, request.RemoteAddr)	
	switch request.URL.Path {
		case "/api/book/search":
			h.api.SearchBook(writer, request)
		case "/api/movie/search":
			h.api.SearchMovie(writer, request)
		case "/api/music/search":
                        h.api.SearchMusic(writer, request)
		default:
			http.NotFound(writer, request)
	}
}

func Start() {
	err := http.ListenAndServe(":32768", &handler{ &API{} })
	CheckErr(err)
}
