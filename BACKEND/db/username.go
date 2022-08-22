package db

import "log"

func GetUsernameAvailable(username string) bool {
	result, err := Db.Exec("SELECT `Name` FROM `Users` WHERE `Name` = ?;", username)
	if err != nil {
		return false
	}
	_, err = result.LastInsertId()
	return err == nil
}

func GetUserBySessionID(sessionID string) string {
	var username string
	rows, err := Db.Query("SELECT `Name` FROM `Users` WHERE `SessionID` = ?;", sessionID)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&username)
	}
	return username
}
