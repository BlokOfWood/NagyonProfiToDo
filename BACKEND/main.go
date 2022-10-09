package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"Todo/controllers"
	"Todo/db"
	"Todo/models"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var Address string = "0.0.0.0:4000"

const (
	defaultAddress string = "0.0.0.0"
	defaultPort    int    = 4000
)

var (
	address string
	port    int
)

func LoadFlags() string {
	flag.StringVar(&address, "address", defaultAddress, "Server's Address")
	flag.IntVar(&port, "port", defaultPort, "Server's Port")
	flag.Parse()
	return fmt.Sprintf("%s:%d", address, port)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("\n\n\nMiddleware\t", r.URL)
		log.Println("")
		fmt.Printf("Address:\t%s\n", r.RemoteAddr)
		fmt.Printf("Method:\t%s\n", r.Method)
		next.ServeHTTP(w, r)
	})
}

func main() {
	Address = LoadFlags()
	mux := mux.NewRouter()
	fmt.Println("Connecting to database")
	db.ConnectDatabase()
	fmt.Println("Set up CORS.")
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "X-Content-Type-Options", "username", "password", "sessionID"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	// origins := handlers.AllowedOrigins([]string{"localhost:4000"})

	mux.HandleFunc("/register", controllers.Register_Controller).Methods("POST")
	mux.HandleFunc("/login", controllers.Login_Controller).Methods("POST")
	mux.HandleFunc("/todos", controllers.Todo_Controller).Methods("GET", "POST")
	mux.HandleFunc("/todos/{id:[0-9]+}", controllers.TodoID_Controller).Methods("GET", "PATCH", "DELETE")

	mux.HandleFunc("/checksession", controllers.CheckSession).Methods("GET", "POST")
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	}).Methods("GET")
	mux.HandleFunc("/cs", func(w http.ResponseWriter, r *http.Request) {
		db.UpdateSessionItem(controllers.CheckSessionID(w, r), controllers.DecodeSessionID(r))
		controllers.SendResponse(w, models.ResponseMessage{Message: "SessionID updated"})
	}).Methods("GET")

	mux.Use(Middleware)
	fmt.Printf("The server is running on " + Address + "\n\n")
	err := http.ListenAndServe(Address, handlers.CORS(header, methods, origins)(mux))
	fmt.Println(err)
}
