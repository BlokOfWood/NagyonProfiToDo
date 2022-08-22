package controllers

import (
	"ToDo/db"
	"ToDo/models"
	"net/http"
)

func Register_Controller(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		// Send back register page
		SendResponse(w, models.ResponseMessage{Message: "Send back registration page"})

	case http.MethodPost:
		// Create a new instance of RegistrationInfo
		var registrationInfo models.RegistrationInfo

		// Decode the request body into the RegistrationInfo instance
		if !DecodeRequest(w, r, &registrationInfo) {
			http.Error(w, "Invalid body", http.StatusInternalServerError)
			return
		}

		// Check if the username is already taken
		if !db.GetUsernameAvailable(registrationInfo.Username) {
			http.Error(w, "Username already taken", http.StatusInternalServerError)
			return
		}

		// Check that we can create user in the database
		if !db.CreateUser(&registrationInfo) {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
			return
		}

		// Send back a success message
		SendResponse(w, models.ResponseMessage{Message: "Registration successful"})
	}
}
