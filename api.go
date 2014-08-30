package servgo

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const (
	url_douban_api = "https://api.douban.com/v2/"
	url_cn_weather = "http://m.weather.com.cn/data/"
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

func (api *Api) QueryWeather(writer http.ResponseWriter, request *http.Request) {
	params := request.URL.Query()
	cityName := params.Get("city")
	rds := &Redis{}
	cityId := rds.QueryCityIdByName(cityName)
	url := url_cn_weather + cityId + ".html"
	writeFromHttpGet(url, writer)
}

func writeFromHttpGet(url string, writer http.ResponseWriter) {
	response, err := http.Get(url)
	checkErr(err)
	defer response.Body.Close()	
	contents, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	fmt.Fprintf(writer, string(contents))
}
