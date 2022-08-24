package controllers

import (
	"Todo/db"
	"Todo/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func TodoID_Controller(w http.ResponseWriter, r *http.Request) {

	// TODO: Check if the request contains a valid ID
	todoID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		fmt.Println("Error getting id")
		http.Error(w, "invalid id", http.StatusBadRequest)
	}

	userID := CheckSessionID(w, r)

	switch r.Method {

	case http.MethodGet:

		response, err := db.GetTodoFromDB(userID, todoID)
		if err != nil {
			fmt.Println("Error getting todo item")
			http.Error(w, http.StatusText(418), http.StatusTeapot)
		}

		SendResponse(w, &response)

	case http.MethodPatch:
		fmt.Println("sdfsdf")
		var request models.TodoEditor

		if !DecodeRequest(w, r, &request) {
			fmt.Println("Error decoding request")
			return
		}

		err := db.UpdateTodoItem(request, todoID)
		if err != nil {
			fmt.Println("Error updating todo item")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		SendResponse(w, models.ResponseMessage{Message: "TodoItem updated"})

	case http.MethodDelete:
		// TODO: Check the user is the owner of the todo
		// TODO: Delete the todo
		err := db.DeleteTodoFromDB(userID, todoID)
		if err != nil {
			fmt.Println("Error deleting todo item")
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// Send response back
		SendResponse(w, models.ResponseMessage{Message: "TodoItem deleted"})
	}

}
