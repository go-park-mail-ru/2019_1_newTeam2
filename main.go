package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := InitServer()
	fmt.Println("starting server at :8090")
	http.ListenAndServe(":8090", server.Router)
}
