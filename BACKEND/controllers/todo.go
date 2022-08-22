package controllers

import (
	"ToDo/db"
	"ToDo/models"
	"ToDo/utils"
	"fmt"
	"net/http"
)

func ToDo_Controller(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		// Get SessionID from header
		SessionID := r.Header.Get("sessionID")
		// Validate SessionID
		if !utils.ValidateSessionID(SessionID) {
			fmt.Println("Invalid sessionID")
			http.Error(w, "Invalid sessionID", http.StatusNotAcceptable)
		}

		// Get username by SessionID
		username := db.GetUserBySessionID(SessionID)
		// Get todos by username
		result, err := db.GetToDosFromDB(username)
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
		id, err := db.CreateToDo(request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		SendResponse(w, id)
	}
}
