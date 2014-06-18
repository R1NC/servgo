package go_web_service

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const URL_ROOT = "https://api.douban.com/v2/"

type API struct{}

func (api *API) SearchBook(writer http.ResponseWriter, request *http.Request) {
        params := request.URL.Query()
        q := params.Get("q")
        tag := params.Get("tag")
        start := params.Get("start")
        count := params.Get("count")
        url := URL_ROOT + "book/search" + "?q=" + q + "&tag=" + tag + "&start=" + start + "&count=" + count
        httpGet(url, writer)
}

func (api *API) SearchMovie(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	q := params.Get("q")
	tag := params.Get("tag")
	start := params.Get("start")
	count := params.Get("count")
	url := URL_ROOT + "movie/search" + "?q=" + q + "&tag=" + tag + "&start=" + start + "&count=" + count 
	httpGet(url, writer)
}

func (api *API) SearchMusic(writer http.ResponseWriter, request *http.Request) {
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
	CheckErr(err)
	defer response.Body.Close()	
	contents, err := ioutil.ReadAll(response.Body)
	CheckErr(err)
	fmt.Fprintf(writer, string(contents))
}
