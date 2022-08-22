package db

import (
	_ "ToDo/models"
)

func GetUsernameAvailable(username string) bool {
	println(username)
	// var user models.User
	var asd string
	rows, err := Db.Query("SELECT `Name` FROM `Users` WHERE `Name` = ?;", username)
	if err != nil {
		return false
	}
	for rows.Next() {
		rows.Scan(&asd)
	}
	println(asd)
	return asd == ""
}

func GetUserBySessionID(sessionID string) (string, error) {
	var username string
	rows, err := Db.Query("SELECT `Name` FROM `Users` WHERE `SessionID` = ?;", sessionID)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&username)
	}
	return username, nil
}

func GetUserIDBySessionID(sessionID string) (uint, error) {
	var username uint
	rows, err := Db.Query("SELECT `UserID` FROM `Users` WHERE `SessionID` = ?;", sessionID)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&username)
	}
	return username, nil
}
