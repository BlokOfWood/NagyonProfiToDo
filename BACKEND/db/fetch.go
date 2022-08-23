package db

import (
	"Todo/models"
	"fmt"
)

func GetSaltFromDB(username string) (string, error) {
	var result string

	// SELECT Salt FROM Users WHERE Username
	rows, err := Db.Query("SELECT `Salt` FROM `Users` WHERE `Name`=? ;", username)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&result)
	}

	return result, nil
}

func GetHashFromDB(username string) (string, error) {
	var result string

	// SELECT Hash FROM Users WHERE Username
	rows, err := Db.Query("SELECT `PasswordHash` FROM `Users` WHERE `Name`=? ;", username)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&result)
	}

	return result, nil
}

func GetTodosFromDB(userID uint) ([]models.Todo, error) {
	result := make([]models.Todo, 0)

	// SELECT * FROM Users WHERE Username
	rows, err := Db.Query("SELECT `TodoID`,`Name`,`Description`,`Priority`,`Done`,`Deadline`,`CreatedAt` FROM `todos` WHERE `UserID`=? ;", userID)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	var item models.Todo

	for rows.Next() {

		err := rows.Scan(&item.TodoID, &item.Name, &item.Description, &item.Priority, &item.Done, &item.Deadline, &item.CreatedAt)
		if err != nil {
			fmt.Println("Reading error:", err)
			return nil, err
		}
		result = append(result, item)
	}

	return result, nil
}
