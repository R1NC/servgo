package go_web_service

import (
	"fmt"
	"net/http"
	"strconv"
)

type handler struct{
	api *Api
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

func Run(port int) {
	err := http.ListenAndServe(":" + strconv.Itoa(port), &handler{ &Api{} })
	CheckErr(err)
}
