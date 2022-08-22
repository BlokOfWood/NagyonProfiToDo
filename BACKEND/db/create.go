package db

import (
	"ToDo/models"
	"ToDo/utils"
)

func CreateUser(registrationInfo *models.RegistrationInfo) bool {
	//Generate Salt
	salt := utils.GenerateSalt()

	//Generate SaltedHash
	saltedHash := utils.EncodePassword(registrationInfo.Password, salt)

	// INSERT INTO
	_, err := Db.Exec("INSERT INTO `Users`(`Name`,`Email`,`PasswordHash`,`Salt`) VALUES(?,?,?,?);", registrationInfo.Username, registrationInfo.Email, saltedHash, salt)

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
