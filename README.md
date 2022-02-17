servgo
===========

Web services written in Golang.

Deploy service
==========

* `go get github.com/R1NC/servgo`;

* Create a `test.go` file:
```go
package main
	
import (
	"github.com/R1NC/servgo"
)
	
func main() {
	servgo.Run(8888)
}
```

* `go run test.go`;

Send request
==========
```
http://localhost:8888/api/book/search?q=abc&tag=def&start=0&count=20
http://localhost:8888/api/movie/search?q=abc&tag=def&start=0&count=20
http://localhost:8888/api/music/search?q=abc&tag=def&start=0&count=20
http://localhost:8888/api/weather/query?city=shanghai
``` 
