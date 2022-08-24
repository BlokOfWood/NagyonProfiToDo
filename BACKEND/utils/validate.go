package utils

import (
	"fmt"
	"log"
	"time"
)

func ValidateSessionID(sessionID string) bool {
	if len(sessionID) == 16 {
		return true
	}
	if sessionID == "" {
		log.Printf("Error\n\t")
		fmt.Println("sessionID is empty")
		return false
	}
	log.Printf("Error\n\t")
	fmt.Println("sessionID is not valid")
	return false
}

func CheckPasswords(password string, repeatedPassword string) bool {
	if password != repeatedPassword {
		return false
	}

	// Check password contains dfbvdkjfnlfgnsldfgsdklfg

	return true
}

func ValidateDate(date string) (string, error) {
	// Check date is valid
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, date)
	if err != nil {
		return "", err
	}
	return t.Format("2006-01-02 15:04:05"), nil
}
