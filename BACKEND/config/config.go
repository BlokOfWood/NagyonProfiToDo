package config

import (
	"fmt"
	"log"
	"os"

	"Todo/models"

	"github.com/joho/godotenv"
)

func LoadEnv(databaseParams *models.DatabaseParams) {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("fataaaaaal error")
		log.Fatal(err)
		return
	}
	// database
	databaseParams.Driver = os.Getenv("DatabaseDriver")
	databaseParams.Username = os.Getenv("DatabaseUsername")
	databaseParams.Password = os.Getenv("DatabasePassword")
	databaseParams.Database = os.Getenv("DatabaseName")
	databaseParams.Address = os.Getenv("Server")
}
