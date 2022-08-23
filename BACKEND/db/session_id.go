package db

import (
	"Todo/utils"
)

func UpdateSessionID(username string) (string, error) {
	//Generate SessionID
	sessionID := utils.GenerateSessionID()

	// INSERT INTO
	_, err := Db.Exec("UPDATE `Users` SET `SessionID` = ? WHERE `Name` = ?;", sessionID, username)
	return sessionID, err
}
