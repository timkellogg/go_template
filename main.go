package main

import (
	"log"
	"net/http"
	"os"

	"./config"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

func main() {
	config.InitializeDb()
	router := NewRouter()
	handler := cors.Default().Handler(router)

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":3000"
	}

	log.Println("Server up on " + port)
	log.Fatal(http.ListenAndServe(port, handler))
}
