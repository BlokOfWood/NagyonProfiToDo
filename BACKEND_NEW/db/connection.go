package db

import (
	"Todo/config"
	"Todo/models"

	"database/sql"
	"fmt"
	"log"
)

var Db *sql.DB

func ConnectDatabase() {
	var err error
	var databaseParams models.DatabaseParams
	config.LoadEnv(&databaseParams)
	Db, err = sql.Open(databaseParams.Driver, databaseParams.Username+":"+databaseParams.Password+"@tcp("+databaseParams.Address+")/"+databaseParams.Database)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to database.")
	// defer Db.Close()
}
