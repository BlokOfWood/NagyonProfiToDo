package models

type SID struct {
	SessionID string `json:"sessionID"`
}

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
