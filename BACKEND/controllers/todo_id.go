package controllers

import (
	"Todo/db"
	"Todo/models"
	"Todo/utils"
	"net/http"
)

func TodoID_Controller(w http.ResponseWriter, r *http.Request) {
	var SessionID models.SessionInfo

	// Decode the request body into the SessionInfo instance
	DecodeRequest(w, r, &SessionID)

	// Validate SessionID
	if !utils.ValidateSessionID(SessionID.SessionID) {
		http.Error(w, "Invalid sessionID", http.StatusForbidden)
		return
	}
	// Get UserID by SessionID
	_, err := db.GetUserIDBySessionID(SessionID.SessionID)
	if err != nil {
		http.Error(w, "Invalid sessionID", http.StatusForbidden)
		return
	}

	// switch r.Method {

	// case http.MethodGet:

	// if rows.Next() {
	// 	rows.Scan(&result.Text, &result.Done, &result.Description)
	// } else {
	// 	return nil, errors.New("item not found")
	// }

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

	// 	db.DeleteToDoFromDB(DeleteInfos.Editor.ID)

	// 	// Get todos by username
	// 	result, err := db.GetToDosFromDB(userID)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// 	// Send response back
	// 	SendResponse(w, result)
	// }
}
