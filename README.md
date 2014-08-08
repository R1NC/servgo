go_web_service
===========

Web services written in Golang.

Usage
==========

* <CODE>git clone</CODE> this project into <CODE>$GOPATH/src</CODE>;

* Create a <CODE>api.go</CODE> file:
```go
package main
	
import (
	"go_web_service"
)
	
func main() {
	go_web_service.Run(8888)
}
```

* <CODE>go run api.go</CODE>;

* Client request:
```
http://localhost:8888/api/book/search?q=abc&tag=&start=0&count=20
http://localhost:8888/api/movie/search?q=abc&tag=&start=0&count=20
http://localhost:8888/api/music/search?q=abc&tag=&start=0&count=20
``` 
