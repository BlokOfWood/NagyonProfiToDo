package main

import (
	Server "github.com/CodeFoxHu/go-serverlib"
)

func main() {

	Server.Initialize()

	Server.Router.Add("/customers", "GET", "POST").Handler(Controller_Customers).AllowWithoutToken(true)
	Server.Router.Add("/customers/id:int", "GET", "PATCH").Handler(Controller_Customers_Id).AllowWithoutToken(true)

	Server.Start()

}
