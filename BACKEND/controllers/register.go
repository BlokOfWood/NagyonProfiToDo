package controllers

import (
	"ToDo/db"
	"ToDo/utils"
	"fmt"
	"net/http"
)

func Register_Controller(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		w.Write([]byte("Send back registration form"))

	case http.MethodPost:
		username := r.Header.Get("username")
		email := r.Header.Get("email")
		password := r.Header.Get("password")
		passwordRepeated := r.Header.Get("passwordRepeated")

		//Validate Passwords
		if !utils.CheckPasswords(password, passwordRepeated) {
			fmt.Println("Sanyi szar a password")
			w.Write([]byte("Sanyi szar a password"))
			return
		}

		if !db.GetUsernameAvailable(username) {
			fmt.Println("Sanyi szar a username")
			w.Write([]byte("Sanyi szar a username"))
			return
		}

		if !db.CreateUser(&username, &email, &password) {
			fmt.Println("HEH?")
			return
		}

		w.Write([]byte("Success register"))
	}
}
