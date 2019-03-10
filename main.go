package main

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	fmt.Println("starting server at :8090")

	server := InitServer()
	// headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	// originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000/"}) //os.Getenv("ORIGIN_ALLOWED")})
	// methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// http.ListenAndServe(":8090", handlers.CORS(originsOk, headersOk, methodsOk)(server.Router))

	// r := mux.NewRouter()
	// r.HandleFunc("/signup/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("lolkek4eburek")
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.Write([]byte("{\"hello\": \"world\"}"))
	// })

	c := cors.New(cors.Options{
		AllowedHeaders:     []string{"Access-Control-Allow-Origin", "Charset", "Content-Type"},
		AllowedOrigins:     []string{"http://localhost:3000", "https://thawing-gorge-14317.herokuapp.com/"},
		AllowCredentials:   true,
		AllowedMethods:     []string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"},
		OptionsPassthrough: true,
		Debug:              true,
	})

	// Use default options
	// handler := cors.AllowAll().Handler(server.Router)
	handler := c.Handler(server.Router)
	http.ListenAndServe(":0", handler)
}
