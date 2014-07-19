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
	go_web_service.Start(8888)
}
```

* <CODE>go run api.go</CODE>.
