package controllers

import (
	"ToDo/db"
	"ToDo/models"
	"ToDo/utils"
	"net/http"
)

func Login_Controller(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		// Send back register page
		SendResponse(w, models.ResponseMessage{Message: "Send back login page"})

	case http.MethodPost:
		// Create a new instance of LoginInfo
		var loginInfo models.LoginInfo

		// Decode the request body into the LoginInfo instance
		if !DecodeRequest(w, r, &loginInfo) {
			http.Error(w, "Invalid body", http.StatusInternalServerError)
			return
		}

		// Get salt from database
		salt, err := db.GetSaltFromDB(loginInfo.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Hash the password with the salt
		hash := utils.EncodePassword(loginInfo.Password, salt)

		// Get hash from database
		dbHash, err := db.GetHashFromDB(loginInfo.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if the password is correct
		if hash != dbHash {
			http.Error(w, "Bad password", http.StatusForbidden)
			return
		}

		// Create a new session and update the database
		sessionID, err := db.UpdateSessionID(loginInfo.Username)
		if err != nil {
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}

		// Send the sessionID back to the client
		SendResponse(w, models.SessionInfo{SessionID: sessionID})
	}
}
