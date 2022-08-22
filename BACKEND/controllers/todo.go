package controllers

import (
	"ToDo/db"
	"ToDo/models"
	"ToDo/utils"
	"fmt"
	"net/http"
)

func ToDo_Controller(w http.ResponseWriter, r *http.Request) {
	// Get SessionID from header
	SessionID := r.Header.Get("sessionID")
	// Validate SessionID
	if !utils.ValidateSessionID(SessionID) {
		fmt.Println("Invalid sessionID")
		http.Redirect(w, r, "/login", http.StatusForbidden)
	}
	var err error
	switch r.Method {

	case http.MethodGet:
		// Get username by SessionID
		userID, err := db.GetUserIDBySessionID(SessionID)
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusForbidden)
		}
		// Get todos by username
		result, err := db.GetToDosFromDB(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// Send response back
		SendResponse(w, result)

	case http.MethodPost:
		var request models.TodoEditor

		if DecodeRequest(w, r, &request) {
			return
		}

		//
		request.UserID, err = db.GetUserIDBySessionID(SessionID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		//
		id, err := db.CreateToDo(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		SendResponse(w, id)
	}
}
