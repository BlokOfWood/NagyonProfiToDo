package controllers

import (
	"Todo/db"
	"Todo/models"
	"Todo/utils"
	"fmt"
	"net/http"
)

func Todo_Controller(w http.ResponseWriter, r *http.Request) {
	var err error

	// Get userID from sessionID and validate that.
	userID := CheckSessionID(w, r)

	switch r.Method {

	case http.MethodGet:
		// Get todos by userID
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
		Todo.Deadline, err = utils.ValidateDate(Todo.Deadline)
		if err != nil {
			fmt.Println("ValidateDate failed")
			http.Error(w, "ValidateDate failed", http.StatusBadRequest)
			return
		}

		//	Create a new TodoItem
		id, err := db.CreateTodo(Todo, userID)
		if err != nil {
			fmt.Println("CreateTodo failed")
			fmt.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		SendResponse(w, id)
	}
}
