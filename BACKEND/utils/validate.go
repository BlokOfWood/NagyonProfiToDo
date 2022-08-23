package utils

import (
	"time"
)

func ValidateSessionID(sessionID string) bool {
	return len(sessionID) == 16
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
