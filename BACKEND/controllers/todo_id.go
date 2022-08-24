package controllers

import (
	"Todo/db"
	"Todo/models"
	"Todo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func TodoID_Controller(w http.ResponseWriter, r *http.Request) {

	// TODO: Check if the request contains a valid ID
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println("Error getting id")
		http.Error(w, "invalid id", http.StatusBadRequest)
	}
	fmt.Println(id)

	// TODO: Check if the user is logged in
	// Get SessionID from request header
	sessionID := DecodeSessionID(r)

	// Validate SessionID
	if !utils.ValidateSessionID(sessionID) {
		http.Error(w, "Invalid sessionID", http.StatusForbidden)
		return
	}

	// TODO: Check if the session belongs to a user
	// Get UserID by SessionID
	userID, err := db.GetUserIDBySessionID(sessionID)
	if err != nil {
		fmt.Println("Error getting userID by sessionID")
		http.Error(w, "Get UserID by SessionID failed", http.StatusForbidden)
		return
	}

	switch r.Method {

	case http.MethodGet:

		response, err := db.GetTodoFromDB(userID, id)
		if err != nil {
			fmt.Println("Error getting todo item")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		SendResponse(w, &response)

	case http.MethodPatch:
		fmt.Println("sdfsdf")
		var request models.TodoEditor

		if !DecodeRequest(w, r, &request) {
			fmt.Println("Error decoding request")
			return
		}

		err := db.UpdateTodoItem(request, id)
		if err != nil {
			fmt.Println("Error updating todo item")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		SendResponse(w, models.ResponseMessage{Message: "TodoItem updated"})

	case http.MethodDelete:
		// TODO: Check the user is the owner of the todo
		// TODO: Delete the todo
		err := db.DeleteTodoFromDB(userID, id)
		if err != nil {
			fmt.Println("Error deleting todo item")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Send response back
		SendResponse(w, models.ResponseMessage{Message: "TodoItem deleted"})
	}

	// // Decode the request body into the SessionInfo instance
	// DecodeRequest(w, r, &SessionID)

	// // Get UserID by SessionID
	// _, err := db.GetUserIDBySessionID(SessionID.SessionID)
	// if err != nil {
	// 	http.Error(w, "Invalid sessionID", http.StatusForbidden)
	// 	return
	// }

	//418

	// var Todo models.TodoEditor
	// // Validate data
	// if !DecodeRequest(w, r, &Todo) {
	// 	http.Error(w, "Invalid body", http.StatusBadRequest)
	// 	return
	// }
	// // Check date format
	// asd, err := utils.ValidateDate(Todo.Deadline)
	// println(asd)
	// if err != nil {
	// 	http.Error(w, "Invalid date", http.StatusBadRequest)
	// 	return
	// }
	// // Create a new instance of TodoEditor
	// var TodoEditor models.TodoEditor
	// // Decode the request body into the TodoEditor instance
	// DecodeRequest(w, r, &TodoEditor)
	// // Insert the new TodoEditor instance into the database
	// err = db.UpdateTodoItem(TodoEditor)
	// if err != nil {
	// 	http.Error(w, "InsertTodo failed", http.StatusInternalServerError)
	// 	return
	// }
	// // Send response back
	// SendResponse(w, TodoEditor)
}

// 	if rows.Next() {
// 		rows.Scan(&result.Text, &result.Done, &result.Description)
// 	} else {
// 		return nil, errors.New("item not found")
// 	}

// 	// Get todos by username
// 	result, err := db.GetToDoFromDB(userID)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}
// 	// Send response back
// 	SendResponse(w, result)
// case http.MethodPatch:

// case http.MethodDelete:

// 	// SessionID
// 	// ID
// 	var deleteInfos struct {
// 		models.SessionInfo `json:"session"`
// 		models.TodoEditor  `json:"editor"`
// 	}

// 	DecodeRequest(w, r, &DeleteInfos)

//
// 	}
// }
