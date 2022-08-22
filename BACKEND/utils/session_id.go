package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-0123456789")

const sessionIDLength uint8 = 16

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateSessionID() string {
	fmt.Println("GenerateSessionID")
	sessionID := make([]rune, sessionIDLength)
	for i := range sessionID {
		sessionID[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(sessionID)
}

func ValidateSessionID(sessionID string) bool {
	return len(sessionID) == 16
}
