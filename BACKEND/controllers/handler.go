package controllers

import (
	"encoding/json"
	"net/http"
)

func SendResponse(w http.ResponseWriter, input any) {
	data, err := json.Marshal(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
