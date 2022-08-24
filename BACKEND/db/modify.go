package db

import (
	"Todo/models"
	"log"
)

func UpdateTodoItem(editor models.TodoEditor, todoID int) error {
	_, err := Db.Exec("UPDATE `todos` SET `Name`= ?, `Description`= ?, `Priority`= ?, `Done`= ?, `Deadline`= ? WHERE `TodoID`= ?",
		editor.Name,
		editor.Description,
		editor.Priority,
		editor.Done,
		editor.Deadline,
		todoID)

	if err != nil {
		log.Println("Failed to update TodoItem")
		return err
	}

	return nil
}

func UpdateSessionItem(userID uint, sessionID string) error {
	_, err := Db.Exec("UPDATE `sessions` SET `expirationDate`=ADDDATE(current_timestamp(), INTERVAL 1 WEEK) WHERE `UserID`= ? AND `SessionID`= ?", userID, sessionID)

	if err != nil {
		log.Println("Failed to update Session")
		return err
	}
	return nil
}

func DeleteTodoFromDB(userID uint, ToDoID int) error {
	_, err := Db.Exec("DELETE FROM `todos` WHERE `UserID`=? AND `TodoID`=? ;", userID, ToDoID)
	if err != nil {
		log.Println("Failed to delete TodoItem")
		return err
	}
	return nil
}

func DeleteTodosFromDB(userID uint) (int64, error) {
	result, err := Db.Exec("DELETE FROM `todos` WHERE `UserID`=?", userID)
	if err != nil {
		log.Println("Failed to delete TodoItems")
		return 0, err
	}
	return result.RowsAffected()
}
