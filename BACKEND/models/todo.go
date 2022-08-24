package models

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
