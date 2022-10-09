package main

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

type DatabaseParams struct {
	Driver   string
	Username string
	Password string
	Database string
	Address  string
}

type PriorityEnum string

const (
	Critical   PriorityEnum = "Critical"
	Urgent     PriorityEnum = "Urgent"
	Important  PriorityEnum = "Important"
	Normal     PriorityEnum = "Normal"
	Eventually PriorityEnum = "Eventually"
)

type Todo struct {
	TodoID      uint         `json:"todoID"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Priority    PriorityEnum `json:"priority"`
	Done        bool         `json:"done"`
	Deadline    string       `json:"deadline"`
	CreatedAt   string       `json:"createdAt"`
}

type TodoEditor struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Priority    PriorityEnum `json:"priority"`
	Done        bool         `json:"done"`
	Deadline    string       `json:"deadline"`
}

type Username struct {
	UserId   uint
	Username string
}

type UserEditor struct {
	UserID       uint   `json:"userID"`
	SessionID    string `json:"sessionID"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
	Salt         string `json:"salt"`
}
