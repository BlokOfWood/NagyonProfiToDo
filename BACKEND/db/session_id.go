package db

import (
	"Todo/utils"
)

func UpdateSessionID(username string) error {
	// INSERT INTO
	_, err := Db.Exec("UPDATE `Sessions` SET `ExpirationDate` = ADDDATE(current_timestamp(), INTERVAL 1 WEEK) WHERE `UserID` = ?;", GetUserIDByName(username))
	return err
}

func InitalizeSessionID(username string) error {
	//Generate SessionID
	sessionID := utils.GenerateSessionID()

	// INSERT INTO
	_, err := Db.Exec("INSERT INTO `Sessions` VALUES (?,?,ADDDATE(current_timestamp(), INTERVAL 1 WEEK));", GetUserIDByName(username), sessionID)
	return err
}
