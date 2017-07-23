package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	// DB is the connection to the database
	DB *gorm.DB
)

// InitializeDb - creates a pointer to database server
func InitializeDb() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	DB, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	DB.LogMode(true)
}
