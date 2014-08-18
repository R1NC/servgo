servgo
===========

Web services written in Golang.

Deploy service
==========

* <CODE>go get github.com/RincLiu/servgo</CODE>;

* Create a <CODE>test.go</CODE> file:
```go
package main
	
import (
	"github.com/RincLiu/servgo"
)
	
func main() {
	servgo.Run(8888)
}
```

* <CODE>go run test.go</CODE>;

Send request
==========
```
http://localhost:8888/api/book/search?q=abc&tag=def&start=0&count=20
http://localhost:8888/api/movie/search?q=abc&tag=def&start=0&count=20
http://localhost:8888/api/music/search?q=abc&tag=def&start=0&count=20
http://localhost:8888/api/weather/query?city=shanghai
``` 
