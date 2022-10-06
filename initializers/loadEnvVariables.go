package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVariables() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading local.env file ", err)
	}
}
