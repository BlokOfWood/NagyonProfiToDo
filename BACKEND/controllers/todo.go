package controllers

import (
	"Todo/db"
	"Todo/models"
	"Todo/utils"
	"fmt"
	"net/http"
)

func Todo_Controller(w http.ResponseWriter, r *http.Request) {

	// Get SessionID from request header
	sessionID := DecodeSessionID(r)

	fmt.Println("SessionID: ", sessionID)

	// Validate SessionID
	if !utils.ValidateSessionID(sessionID) {
		fmt.Println("Validate sessionID failed")
		http.Error(w, "Validate sessionID failed", http.StatusForbidden)
		return
	}

	// Get UserID by SessionID
	userID, err := db.GetUserIDBySessionID(sessionID)
	if err != nil {
		fmt.Println("Get UserID by SessionID failed")
		http.Error(w, "Get UserID by SessionID failed", http.StatusForbidden)
		return
	}

	//
	switch r.Method {

	case http.MethodGet:

		// Get todos by username
		result, err := db.GetTodosFromDB(userID)
		if err != nil {
			fmt.Println("Get User by userID failed")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Send response back
		SendResponse(w, result)

	case http.MethodPost:

		var Todo models.TodoEditor

		// Validate data
		if !DecodeRequest(w, r, &Todo) {
			fmt.Println("DecodeRequest failed")
			http.Error(w, "DecodeRequest failed", http.StatusBadRequest)
			return
		}

		// Check date format
		asd, err := utils.ValidateDate(Todo.Deadline)
		println(asd)
		if err != nil {
			fmt.Println("ValidateDate failed")
			http.Error(w, "ValidateDate failed", http.StatusBadRequest)
			return
		}

		Todo.Deadline = asd

		//
		id, err := db.CreateTodo(Todo, userID)
		if err != nil {
			fmt.Println("CreateTodo failed")
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		SendResponse(w, id)
	}
}
