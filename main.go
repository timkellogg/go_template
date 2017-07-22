package main

import (
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	router := NewRouter()
	handler := cors.Default().Handler(router)

	port := ":" + os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	log.Println("Server up on " + port)
	log.Fatal(http.ListenAndServe(port, handler))
}
