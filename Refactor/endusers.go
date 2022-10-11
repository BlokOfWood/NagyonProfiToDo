package main

import (
	"errors"

	Server "github.com/CodeFoxHu/go-serverlib"
)

func GetUserFromDatabase() error {
	return nil
}

func RegisterCustomer(ctx *Server.FoxContext, customerInfo CustomerInfo) (int, error) {
	return SaveCustomer(ctx, customerInfo)
}

func SaveCustomer(ctx *Server.FoxContext, customerInfo CustomerInfo, id ...int) (int, error) {

	// validate
	err := ctx.ValidateStruct(customerInfo)
	if err != nil {
		return 0, err
	}

	isCreate := len(id) == 0

	// Check user exist in database
	qry := Server.MySQL.NewQuery(ctx)
	qry.SQL = "SELECT CustomerId FROM CUSTOMERS WHERE CustomerEmail = :customerEmail "
	qry.BindValue("customerEmail", customerInfo.Email)

	if !isCreate {
		qry.SQL += "AND CustomerId != :customerId"
		qry.BindValue("customerId", id[0])

	}
	err = qry.Open()
	if err != nil {
		return 0, err
	}
	if !qry.IsEmpty {
		return 0, errors.New("Email occupied")
	}

	// Register user in database
	if !isCreate {
		qry.BuildSQLFromData(customerInfo, "CUSTOMERS", Server.UPDATE_QUERY, "CustomerId = :customerId")
		qry.BindValue("customerId", id[0])
	} else {
		qry.BuildSQLFromData(customerInfo, "CUSTOMERS", Server.INSERT_QUERY)
	}
	err = qry.Exec()
	if err != nil {
		return 0, err
	}

	return qry.LastInsertId, nil

}

func UpdateCustomer(ctx *Server.FoxContext, customerInfo CustomerInfo, id int) error {
	_, err := SaveCustomer(ctx, customerInfo, id)
	return err
}
