servgo
===========

Web services written in Golang.

Usage of Server
==========

* <CODE>git clone</CODE> this project into <CODE>$GOPATH/src/github.com/RincLiu/servgo</CODE>;

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

Usage of Client
==========
```
http://localhost:8888/api/book/search?q=abc&tag=&start=0&count=20
http://localhost:8888/api/movie/search?q=abc&tag=&start=0&count=20
http://localhost:8888/api/music/search?q=abc&tag=&start=0&count=20
``` 
