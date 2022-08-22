package db

import (
	"ToDo/models"
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

func GetToDosFromDB(userID uint) ([]models.Todo, error) {
	result := make([]models.Todo, 0)

	// SELECT * FROM Users WHERE Username
	rows, err := Db.Query("SELECT * FROM `todos` WHERE `UserID`=? ;", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var item models.Todo
	for rows.Next() {
		rows.Scan(&item)
		result = append(result, item)
	}

	return result, nil
}

// func GetTaskFromDB(username string, id int) (*models.TodoItemEditor, error) {
// 	result := &models.TodoItemEditor{}

// 	// SELECT * FROM Users WHERE Username
// 	rows, err := Db.Query("SELECT * FROM `Users` WHERE `Name`=?  AND `Task`;", username)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var item models.TodoItem
// 	for rows.Next() {
// 		rows.Scan(&item)
// 		result = append(result, item)
// 	}

// 	return result, nil
// }
