package controllers

import (
	"ToDo/db"
	"ToDo/utils"
	"fmt"
	"net/http"
)

func Login_Controller(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		w.Write([]byte("Send back login page"))

	case http.MethodPost:
		username := r.Header.Get("username")

		salt, err := db.GetSaltFromDB(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		hash := utils.EncodePassword(r.Header.Get("password"), salt)
		dbHash, err := db.GetHashFromDB(username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if hash != dbHash {
			http.Error(w, "Invalid sesssionId", http.StatusForbidden)
		}

		if !db.UpdateSessionID(username) {
			fmt.Println("Nem sikerült a sessionID update, bár ez nem tudom miért baj ")
		}

		w.Write([]byte("Success login"))
	}
}
