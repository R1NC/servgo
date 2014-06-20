go_web_service
===========

Web services written in Golang.

Usage
==========

* Git clone this project into $GOPATH/src;
* Create a .go file to start service:
```go
package main {
	
	import (
		"go_web_service"
	)
	
	func main() {
		go_web_service.Start(8888)
	}
		
}
```
