package db

import (
	"ToDo/utils"
)

func UpdateSessionID(username string) bool {
	//Generate SessionID
	sessionID := utils.GenerateSessionID()

	// INSERT INTO
	_, err := Db.Exec("UPDATE `Users` SET `SessionID` = ? WHERE `Name` = ?;", sessionID, username)
	return err == nil
}
