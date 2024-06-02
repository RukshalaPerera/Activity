package configs

import (
	"github.com/joho/godotenv"
	"log"
)

func init() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	var err error
	DB, err = ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
}
