package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)
// get database from .env file
func DatabaseURI() string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Unable to retrieve databasef from .env file")
		log.Fatal(err)
	}

	return os.Getenv("DATABASEURI")
}