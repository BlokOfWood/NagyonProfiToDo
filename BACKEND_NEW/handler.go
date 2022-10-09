package main

import (
	"errors"
	"fmt"
	"log"

	Server "github.com/CodeFoxHu/go-serverlib"
)

func Register_Controller(ctx *Server.FoxContext) error {
	// Create a new instance of RegistrationInfo
	var registrationInfo RegistrationInfo

	err := ctx.UnmarshalBody(&registrationInfo)
	if err != nil {
		log.Println("error: ", err)
		return err
	}

	//TODO GetUsernameAvailable

	err = Server.MySQL.Begin(ctx, false)
	if err != nil {
		log.Println("error: ", err)
		return err
	}
	qry := Server.MySQL.NewQuery(ctx)
	var usernames map[string]struct{}
	err = qry.OpenSQL("SELECT * FROM USERS")
	if err != nil {
		log.Println("error: ", err)
		return err
	}
	asd := qry.FetchAll(&usernames)
	fmt.Println(asd)
	fmt.Println(usernames)
	if _, found := usernames[registrationInfo.Username]; found {
		return errors.New("Username is already taken")
	}

	//TODO CreateUser with Salt and save into the DB

	err = Server.MySQL.Begin(ctx, false)
	if err != nil {
		log.Println("error: ", err)
		return err
	}
	qry = Server.MySQL.NewQuery(ctx)
	var s Server.SqlStringList
	s = append(s, registrationInfo.Username)

	//TODO registrationInfo.Password + salt => password

	salt := GenerateSalt()
	password := EncodePassword(registrationInfo.Password, salt)
	s = append(s, password)
	s = append(s, salt)
	s = append(s, registrationInfo.Email)
	err = qry.ExecSQL("INSERT INTO `users` (`username`, `password`, `salt`, `email`) VALUES (" + Server.SQLIn(s) + ")")
	if err != nil {
		log.Println("error: ", err)
		return err
	}
	// Send back a success message
	ctx.RespondWithJson([]byte("Registration successful"), "Message")
	return nil
}

func Login_Controller(ctx *Server.FoxContext) error {
	// Create a new instance of LoginInfo
	var loginInfo LoginInfo

	err := ctx.UnmarshalBody(&loginInfo)
	if err != nil {
		return err
	}

	//TODO GetSaltFromDB

	//TODO EncodePassword

	//TODO GetHashFromDB

	//TODO hash != dbHash

	//TODO UpdateSessionID

	// Send back a success message
	ctx.RespondWithJson([]byte("Login successful"), "Message")

	return nil
}

func Todo_Controller(ctx *Server.FoxContext) error {
	return nil
}

func TodoID_Controller(ctx *Server.FoxContext) error {
	return nil
}
