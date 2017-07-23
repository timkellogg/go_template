package main

import (
	"log"
	"net/http"
	"os"

	"./config"
	"./services/workers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/rs/cors"
)

func main() {
	// initialize database connection
	config.InitializeDb()

	// start worker pool
	p := workers.NewPool(nil)
	p.Run()

	// initialize router
	router := NewRouter()
	handler := cors.Default().Handler(router)

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":3000"
	}

	// handle requests on port
	log.Println("Server up on " + port)
	log.Fatal(http.ListenAndServe(port, handler))
}
