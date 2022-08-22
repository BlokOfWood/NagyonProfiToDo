package controllers

import (
	"encoding/json"
	"net/http"
)

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
	return decoder.Decode(input) == nil
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
