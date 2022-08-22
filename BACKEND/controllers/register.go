package controllers

import (
	"ToDo/db"
	"ToDo/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Register_Controller(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		w.Write([]byte("Send back registration form"))

	case http.MethodPost:
		var registrationInfo models.RegistrationInfo
		beszarok := r.Body
		fmt.Println(beszarok)
		decoder := json.NewDecoder(beszarok)
		fmt.Println(beszarok)
		decoder.Decode(registrationInfo)
		fmt.Println(registrationInfo)
		DecodeRequest(w, r, registrationInfo)

		if !db.GetUsernameAvailable(registrationInfo.Username) {
			fmt.Println("Sanyi szar a username")
			w.Write([]byte("Sanyi szar a username"))
			return
		}

		if !db.CreateUser(&registrationInfo) {
			fmt.Println("HEH?")
			return
		}

		w.Write([]byte("Success register"))
	}
}
