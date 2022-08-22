package controllers

import (
	"ToDo/db"
	"ToDo/models"
	"ToDo/utils"
	"fmt"
	"net/http"
)

func Login_Controller(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		w.Write([]byte("Send back login page"))

	case http.MethodPost:

		var loginInfo models.LoginInfo

		DecodeRequest(w, r, loginInfo)

		salt, err := db.GetSaltFromDB(loginInfo.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		hash := utils.EncodePassword(loginInfo.Password, salt)

		dbHash, err := db.GetHashFromDB(loginInfo.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if hash != dbHash {
			http.Error(w, "Bad password", http.StatusForbidden)
			return
		}

		sessionID, err := db.UpdateSessionID(loginInfo.Username)
		if err != nil {
			fmt.Println("Nem sikerült a sessionID update, bár ez nem tudom miért baj ")
		}

		var Sanyi models.SID
		Sanyi.SessionID = sessionID
		SendResponse(w, Sanyi)
	}
}
