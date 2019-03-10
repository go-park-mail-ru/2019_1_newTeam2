package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	fmt.Println("starting server at :8090")

	server := InitServer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	c := cors.New(cors.Options{
		AllowedHeaders:     []string{"Access-Control-Allow-Origin", "Charset", "Content-Type"},
		AllowedOrigins:     []string{"http://localhost:3000", "https://thawing-gorge-14317.herokuapp.com/", "http://localhost:8090"},
		AllowCredentials:   true,
		AllowedMethods:     []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"},
		OptionsPassthrough: true,
		Debug:              true,
	})

	handler := c.Handler(server.Router)
	http.ListenAndServe(":"+port, handler)
}
