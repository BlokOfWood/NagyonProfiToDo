package models

type PriorityEnum byte

const (
	Critical   PriorityEnum = 0
	Urgent     PriorityEnum = 1
	Important  PriorityEnum = 2
	Normal     PriorityEnum = 3
	Eventually PriorityEnum = 4
)

type Todo struct {
	ToDoID      uint         `json:"toDoID"`
	UserID      uint         `json:"userID"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Priority    PriorityEnum `json:"priority"`
	Done        bool         `json:"done"`
	Deadline    string       `json:"deadline"`
	CreatedAt   string       `json:"createdAt"`
}

type TodoEditor struct {
	ToDoID      uint         `json:"toDoID"`
	UserID      uint         `json:"userID"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Priority    PriorityEnum `json:"priority"`
	Done        bool         `json:"done"`
	Deadline    string       `json:"deadline"`
	CreatedAt   string       `json:"createdAt"`
}
