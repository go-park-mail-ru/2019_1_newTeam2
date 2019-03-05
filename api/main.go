package main

import (
	"fmt"
	"net/http"
)

func main() {
	var server Server
	server.init()
	fmt.Println("starting server at :8090")
	http.ListenAndServe(":8090", server)
}
