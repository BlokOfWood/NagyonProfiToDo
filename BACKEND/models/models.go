package models

type LoginInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegistrationInfo struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type SessionInfo struct {
	SessionID string `json:"sessionID"`
}
