package models

type User struct {
	UserID       uint   `json:"userID"`
	SessionID    string `json:"sessionID"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	Salt         string `json:"salt"`
}

type UserEditor struct {
	UserID       uint   `json:"userID"`
	SessionID    string `json:"sessionID"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	Salt         string `json:"salt"`
}
