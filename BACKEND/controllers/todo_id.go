package controllers

import "net/http"

func ToDoID_Controller(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		w.Write([]byte("Send back task page"))
	case http.MethodPatch:
		w.Write([]byte("Patch TaskID_Controller"))
	case http.MethodDelete:
		w.Write([]byte("Delete TaskID_Controller"))
	}
}
