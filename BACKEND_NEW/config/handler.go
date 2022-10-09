// package main

// import (
// 	"Todo/db"
// 	"Todo/utils"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	Server "github.com/CodeFoxHu/go-serverlib"
// )

// func CheckSession(ctx *Server.FoxContext) error {

// 	// Get SessionID from request header
// 	sessionID := DecodeSessionID(r)
// 	var err error

// 	// Validate SessionID
// 	if !utils.ValidateSessionID(sessionID) {
// 		http.Error(w, "Validate sessionID failed", http.StatusForbidden)
// 	}
// 	fmt.Println("SessionID: ", sessionID)
// 	userID, err := db.GetUserIDBySessionID(sessionID)
// 	if err != nil {
// 		fmt.Println("Error getting userID by sessionID")
// 		http.Error(w, "Get UserID by SessionID failed", http.StatusForbidden)
// 	}
// }

// func CheckSessionID(w http.ResponseWriter, r *http.Request) uint {
// 	// Get SessionID from request header
// 	sessionID := DecodeSessionID(r)
// 	var err error

// 	// Validate SessionID
// 	if !utils.ValidateSessionID(sessionID) {
// 		http.Error(w, "Validate sessionID failed", http.StatusForbidden)
// 		return 0
// 	}
// 	fmt.Println("SessionID: ", sessionID)
// 	// Get UserID by SessionID
// 	userID, err := db.GetUserIDBySessionID(sessionID)
// 	if err != nil {
// 		fmt.Println("Error getting userID by sessionID")
// 		http.Error(w, "Get UserID by SessionID failed", http.StatusForbidden)
// 		return 0
// 	}
// 	return userID
// }

// func SendResponse(w http.ResponseWriter, input any) {
// 	data, err := json.Marshal(input)
// 	if err != nil {
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(data)
// }

// func DecodeRequest(w http.ResponseWriter, r *http.Request, input any) bool {
// 	decoder := json.NewDecoder(r.Body)

// 	err := decoder.Decode(input)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return false
// 	}

// 	return true
// }

// func DecodeSessionID(r *http.Request) string {
// 	return r.Header.Get("sessionID")
// }

// func DecodeID(r *http.Request) string {
// 	return r.Header.Get("ID")
// }

// func SendResponseW(w http.ResponseWriter, i any, wrapper ...string) {
// 	data, err := json.Marshal(i)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	if len(wrapper) > 0 {
// 		data = append([]byte("{\""+wrapper[0]+"\":"), data...)
// 		data = append(data, []byte("}")...)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(data)
// }