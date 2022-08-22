package db

import (
	"ToDo/models"
	"ToDo/utils"
)

func CreateUser(username *string, email *string, password *string) bool {
	//Generate Salt
	salt := utils.GenerateSalt()

	//Generate SaltedHash
	saltedHash := utils.EncodePassword(*password, salt)

	// INSERT INTO
	_, err := Db.Exec("INSERT INTO `Users`(`Name`,`Email`,`PasswordHash`,`Salt`) VALUES(?,?,?,?);", username, email, saltedHash, salt)

	return err == nil
}

func CreateToDo(editor models.TodoEditor) (int64, error) {

	filas, err := Db.Exec("INSERT INTO `ToDos`(`UserID`,`Name`,`Description`,`Priority`,`Done`,`Deadline`) VALUES(?,?,?,?,?,?);",
		editor.UserID,
		editor.Name,
		editor.Description,
		editor.Priority,
		editor.Done,
		editor.Deadline)
	if err != nil {
		return 0, err
	}
	return filas.LastInsertId()

}
