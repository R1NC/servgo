package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const URL_ROOT = "https://api.douban.com/v2/"

func searchBook(writer http.ResponseWriter, request *http.Request) {
        params := request.URL.Query()
        q := params.Get("q")
        tag := params.Get("tag")
        start := params.Get("start")
        count := params.Get("count")
        url := URL_ROOT + "book/search" + "?q=" + q + "&tag=" + tag + "&start=" + start + "&count=" + count
        httpGet(url, writer)
}

func searchMovie(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	q := params.Get("q")
	tag := params.Get("tag")
	start := params.Get("start")
	count := params.Get("count")
	url := URL_ROOT + "movie/search" + "?q=" + q + "&tag=" + tag + "&start=" + start + "&count=" + count 
	httpGet(url, writer)
}

func searchMusic(writer http.ResponseWriter, request *http.Request) {
        params := request.URL.Query()
        q := params.Get("q")
        tag := params.Get("tag")
        start := params.Get("start")
        count := params.Get("count")
        url := URL_ROOT + "music/search" + "?q=" + q + "&tag=" + tag + "&start=" + start + "&count=" + count
        httpGet(url, writer)
}

func httpGet(url string, writer http.ResponseWriter) {
	response, err := http.Get(url)
	checkErr(err)
	defer response.Body.Close()	
	contents, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	fmt.Fprintf(writer, string(contents))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type MyHttpHandler struct{}

func (handler *MyHttpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Request: %s From: %s\n", request.URL, request.RemoteAddr)	
	switch request.URL.Path {
		case "/api/book/search":
			searchBook(writer, request)
		case "/api/movie/search":
			searchMovie(writer, request)
		case "/api/music/search":
                        searchMusic(writer, request)
		default:
			http.NotFound(writer, request)
	}
}

func main() {
	err := http.ListenAndServe(":32768", &MyHttpHandler{})
	checkErr(err)
}
