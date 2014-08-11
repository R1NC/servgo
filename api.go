package servgo

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const (
	url_douban_api = "https://api.douban.com/v2/"
)

type Api struct{}

func (api *Api) SearchBook(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	q := params.Get("q")
	tag := params.Get("tag")
	start := params.Get("start")
	count := params.Get("count")
	url := url_douban_api + "book/search?q=" + q + "&tag=" + tag + "&start=" + start + "&count=" + count
	writeFromHttpGet(url, writer)
}

func (api *Api) SearchMovie(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	q := params.Get("q")
	tag := params.Get("tag")
	start := params.Get("start")
	count := params.Get("count")
	url := url_douban_api + "movie/search?q=" + q + "&tag=" + tag + "&start=" + start + "&count=" + count 
	writeFromHttpGet(url, writer)
}

func (api *Api) SearchMusic(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	q := params.Get("q")
	tag := params.Get("tag")
	start := params.Get("start")
	count := params.Get("count")
	url := url_douban_api + "music/search?q=" + q + "&tag=" + tag + "&start=" + start + "&count=" + count
	writeFromHttpGet(url, writer)
}

func writeFromHttpGet(url string, writer http.ResponseWriter) {
	response, err := http.Get(url)
	CheckErr(err)
	defer response.Body.Close()	
	contents, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	fmt.Fprintf(writer, string(contents))
}
