package main

import (
	"net/http"

	Server "github.com/CodeFoxHu/go-serverlib"
)

type CustomerInfo struct {
	Username string `valid:"(<= 100) and (>= 4)"`
	Password string `valid:"(<= 100) and (>= 6)"`
	Email    string `valid:"email"`
}

func Controller_Customers(ctx *Server.FoxContext) error {
	switch ctx.HttpRequest.Method {
	case http.MethodGet:
	case http.MethodPost:
		var customerInfo CustomerInfo

		// Get Data from Body
		err := ctx.UnmarshalBody(&customerInfo)
		if err != nil {
			return err
		}

		result, err := RegisterCustomer(ctx, customerInfo)
		if err != nil {
			return err
		}

		ctx.Respond(result, "id")

	}
	return nil
}

func Controller_Customers_Id(ctx *Server.FoxContext) error {
	switch ctx.HttpRequest.Method {
	case http.MethodGet:
	case http.MethodPatch:
		var customerInfo CustomerInfo

		// Get Data from Body
		err := ctx.UnmarshalBody(&customerInfo)
		if err != nil {
			return err
		}

		var customerId int
		err = ctx.UnmarshalQueryParam("id", customerId)
		if err != nil {
			return err
		}

		err = UpdateCustomer(ctx, customerInfo, customerId)
		if err != nil {
			return err
		}

		ctx.Respond(Server.EmptyResponse)
	}
	return nil
}
