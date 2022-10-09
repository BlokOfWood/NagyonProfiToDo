package main

import (
	Server "github.com/CodeFoxHu/go-serverlib"
)

type asd struct {
	Message string
}

func main() {

	Server.Initialize()

	Server.Router.Add("/register", "POST").Handler(Register_Controller).AllowWithoutToken(true)
	Server.Router.Add("/login", "POST").Handler(Login_Controller).AllowWithoutToken(true)
	Server.Router.Add("/todos", "GET", "POST").Handler(Todo_Controller).AllowWithoutToken(true)
	Server.Router.Add("/todos/{id:int}", "GET", "PATCH", "DELETE").Handler(TodoID_Controller).AllowWithoutToken(true)
	// Server.Router.Add("/checksession", "GET", "PATCH").Handler(controllers.CheckSession).AllowWithoutToken(true)
	Server.Router.Add("/", "GET").Handler(func(ctx *Server.FoxContext) error {
		ctx.RespondWithData([]byte("Server is running!"))
		return nil
	}).AllowWithoutToken(true)
	// Server.Router.Add("/cs", "GET").Handler(func(ctx *Server.FoxContext) error {
	// 	db.UpdateSessionItem(controllers.CheckSessionID(), controllers.DecodeSessionID(r))
	// 	controllers.SendResponse(w, models.ResponseMessage{Message: "SessionID updated"})
	// 	return nil
	// }())
	Server.Start()

	// header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "X-Content-Type-Options", "username", "password", "sessionID"})
	// methods := handlers.AllowedMethods([]string{"GET", "POST", "PATCH", "DELETE"})
	// origins := handlers.AllowedOrigins([]string{"*"})
	// origins := handlers.AllowedOrigins([]string{"localhost:4000"})
}

func Register_Controller(ctx *Server.FoxContext) error {
	return nil
}

func Login_Controller(ctx *Server.FoxContext) error {
	return nil
}

func Todo_Controller(ctx *Server.FoxContext) error {
	return nil
}

func TodoID_Controller(ctx *Server.FoxContext) error {
	return nil
}
