package config

import (
	"os"

	models "ToDo/models"

	"github.com/joho/godotenv"
)

func LoadEnv(databaseParams *models.DatabaseParams) {
	err := godotenv.Load(".env")
	if err != nil {
		databaseParams.Driver = "mysql"
		databaseParams.Username = "MarijuanaPepsiJackson"
		databaseParams.Password = "Beezow Doo-Doo Zopittybop-Bop-Bop"
		databaseParams.Database = "todoapp"
		databaseParams.Address = ""
		// log.Fatal(err)
		return
	}
	// database
	databaseParams.Driver = os.Getenv("DatabaseDriver")
	databaseParams.Username = os.Getenv("DatabaseUsername")
	databaseParams.Password = os.Getenv("DatabasePassword")
	databaseParams.Database = os.Getenv("DatabaseName")
	databaseParams.Address = os.Getenv("Server")
}
