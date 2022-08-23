package db

import "Todo/models"

func UpdateTodoItem(editor models.TodoEditor, id int) (int64, error) {
	result, err := Db.Exec("UPDATE `todos` SET `Name`= ?, `Description`= ?, `Priority`= ?, `Done`= ?, `Deadline`= ? WHERE `ToDoID`= ?",
		editor.Name,
		editor.Description,
		editor.Priority,
		editor.Done,
		editor.Deadline,
		id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func DeleteTodoFromDB(userID uint, ToDoID uint) (int64, error) {
	result, err := Db.Exec("DELETE FROM `todos` WHERE `UserID`=? AND `ToDoID`=? ;", userID, ToDoID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func DeleteTodosFromDB(userID uint) (int64, error) {
	result, err := Db.Exec("DELETE FROM `todos` WHERE `UserID`=?", userID)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
