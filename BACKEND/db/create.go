package db

import (
	"Todo/models"
	"Todo/utils"
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

func CreateTodo(editor models.TodoEditor, userID uint) (int64, error) {

	filas, err := Db.Exec("INSERT INTO `Todos`(`UserID`,`Name`,`Description`,`Priority`,`Done`,`Deadline`) VALUES(?,?,?,?,?,?);",
		userID,
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
