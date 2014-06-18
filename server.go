package go_web_service

import (
	"fmt"
	"net/http"
)

type Handler struct{
	api *API
}

func (handler *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Request: %s From: %s\n", request.URL, request.RemoteAddr)	
	switch request.URL.Path {
		case "/api/book/search":
			handler.api.SearchBook(writer, request)
		case "/api/movie/search":
			handler.api.SearchMovie(writer, request)
		case "/api/music/search":
                        handler.api.SearchMusic(writer, request)
		default:
			http.NotFound(writer, request)
	}
}

func Start() {
	err := http.ListenAndServe(":32768", &Handler{ &API{} })
	CheckErr(err)
}
