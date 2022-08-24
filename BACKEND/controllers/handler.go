package controllers

import (
	"encoding/json"
	"net/http"
)

// func CheckSessionID(w http.ResponseWriter, r *http.Request) bool {
// 	var SessionID models.SessionInfo

// 	// Decode the request body into the SessionInfo instance
// 	DecodeRequest(w, r, &SessionID)

// 	// Validate SessionID
// 	if !utils.ValidateSessionID(SessionID.SessionID) {
// 		http.Error(w, "Invalid sessionID", http.StatusForbidden)
// 		return
// 	}
// 	// Get UserID by SessionID
// 	userID, err := db.GetUserIDBySessionID(SessionID.SessionID)
// 	if err != nil {
// 		http.Error(w, "Invalid sessionID", http.StatusForbidden)
// 		return
// 	}
// }

func SendResponse(w http.ResponseWriter, input any) {
	data, err := json.Marshal(input)
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func DecodeRequest(w http.ResponseWriter, r *http.Request, input any) bool {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}

	return true
}

func DecodeSessionID(r *http.Request) string {
	return r.Header.Get("sessionID")
}

func DecodeID(r *http.Request) string {
	return r.Header.Get("ID")
}

func SendResponseW(w http.ResponseWriter, i any, wrapper ...string) {
	data, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(wrapper) > 0 {
		data = append([]byte("{\""+wrapper[0]+"\":"), data...)
		data = append(data, []byte("}")...)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
