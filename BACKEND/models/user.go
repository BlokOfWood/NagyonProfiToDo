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
	ToDoID      uint         `json:"toDoID"`
	UserID      uint         `json:"userID"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Priority    PriorityEnum `json:"priority"`
	Done        bool         `json:"done"`
	Deadline    string       `json:"deadline"`
	CreatedAt   string       `json:"createdAt"`
}
