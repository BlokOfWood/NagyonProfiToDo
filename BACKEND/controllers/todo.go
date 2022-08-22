package controllers

import (
	"ToDo/db"
	"ToDo/models"
	"ToDo/utils"
	"net/http"
)

func ToDo_Controller(w http.ResponseWriter, r *http.Request) {
	// Get SessionID from header
	SessionID := r.Header.Get("sessionID")
	// Validate SessionID
	if !utils.ValidateSessionID(SessionID) {
		http.Error(w, "Invalid sessionID", http.StatusForbidden)
		return
	}
	// Get UserID by SessionID
	userID, err := db.GetUserIDBySessionID(SessionID)
	if err != nil {
		http.Error(w, "Invalid sessionID", http.StatusForbidden)
		return
	}
	switch r.Method {

	case http.MethodGet:

		// Get todos by username
		result, err := db.GetToDosFromDB(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// Send response back
		SendResponse(w, result)

	case http.MethodPost:
		var request models.TodoEditor

		// Validate data
		if DecodeRequest(w, r, &request) {
			return
		}
		request.UserID = userID

		//
		id, err := db.CreateToDo(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		SendResponse(w, id)
	}
}
