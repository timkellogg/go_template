package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	bootstrap()

	router := NewRouter()
	handler := cors.Default().Handler(router)

	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":3000"
	}

	log.Println("Server up on " + port)
	log.Fatal(http.ListenAndServe(port, handler))
}

func bootstrap() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}
}
