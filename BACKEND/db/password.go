package db

import (
	"log"
)

func GetUserHash(username string) string {
	var hash string

	// SELECT * FROM usuarios
	rows, err := Db.Query("SELECT `PasswordHash` FROM `Users` WHERE `Name` = ?;", username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&hash)
	}
	return hash
}
