package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Loading .ENV file")
	}

	//port := dotenv.GetString("PORT")
}
