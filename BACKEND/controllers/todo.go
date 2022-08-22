package controllers

import (
	"ToDo/db"
	"ToDo/models"
	"fmt"
	"net/http"
)

func ToDo_Controller(w http.ResponseWriter, r *http.Request) {
	// var SessionID models.SessionInfo

	// // Decode the request body into the SessionInfo instance
	// DecodeRequest(w, r, &SessionID)

	// // Validate SessionID
	// if !utils.ValidateSessionID(SessionID.SessionID) {
	// 	http.Error(w, "Invalid sessionID", http.StatusForbidden)
	// 	return
	// }
	// // Get UserID by SessionID
	// userID, err := db.GetUserIDBySessionID(SessionID.SessionID)
	// if err != nil {
	// 	http.Error(w, "Invalid sessionID", http.StatusForbidden)
	// 	return
	// }

	//
	switch r.Method {

	case http.MethodGet:
		// sima item
		// Get todos by username
		result, err := db.GetToDosFromDB(0)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		// Send response back
		SendResponse(w, result)

	case http.MethodPost:
		/*

			{
				"session" : { "sessionID" : ""},
				"todo" : {
					"userID" : "NULL",
					"name" : "",
					"description" : "",
					"priority" : "",
					"done" : "",
					"deadline" : ""
				}
			}

		*/
		var Todo models.TodoEditor `json:"todo"`

		// var valami struct {
		// 	Session models.SessionInfo `json:"session"`
		// 	Todo    models.TodoEditor  `json:"todo"`
		// }

		fmt.Println(Todo.UserID)
		fmt.Println(Todo.Name)
		fmt.Println(Todo.Description)
		fmt.Println(Todo.Priority)
		fmt.Println(Todo.Done)
		fmt.Println(Todo.Deadline)

		// Validate data
		if !DecodeRequest(w, r, &Todo) {
			return
		}

		fmt.Println(Todo.UserID)
		fmt.Println(Todo.Name)
		fmt.Println(Todo.Description)
		fmt.Println(Todo.Priority)
		fmt.Println(Todo.Done)
		fmt.Println(Todo.Deadline)

		Todo.UserID = 3

		fmt.Println(Todo.UserID)
		fmt.Println(Todo.Name)
		fmt.Println(Todo.Description)
		fmt.Println(Todo.Priority)
		fmt.Println(Todo.Done)
		fmt.Println(Todo.Deadline)

		//
		id, err := db.CreateToDo(Todo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		SendResponse(w, id)
	}
}
