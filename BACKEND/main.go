package main

import (
	"fmt"
	"log"
	"net/http"

	"ToDo/controllers"
	"ToDo/db"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var Address string = "0.0.0.0:4000"

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware", r.URL)
		log.Printf("%s %s %s\n\n\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := mux.NewRouter()
	fmt.Println("Connecting to database")
	db.ConnectDatabase()
	fmt.Println("Set up CORS.")
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "X-Content-Type-Options", "username", "password", "sessionID"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	// origins := handlers.AllowedOrigins([]string{"localhost:4000"})

	mux.HandleFunc("/todos", controllers.ToDo_Controller).Methods("GET", "POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", controllers.ToDoID_Controller).Methods("GET", "PATCH", "DELETE")
	mux.HandleFunc("/login", controllers.Login_Controller).Methods("GET", "POST")
	mux.HandleFunc("/register", controllers.Register_Controller).Methods("GET", "POST")
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("asd"))
	}).Methods("GET")

	mux.Use(Middleware)
	fmt.Printf("The server is running on " + Address + "\n\n")
	err := http.ListenAndServe(Address, handlers.CORS(header, methods, origins)(mux))
	fmt.Println(err)
}
