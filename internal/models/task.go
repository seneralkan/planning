package models

// Task represents a to-do task
type Task struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Duration   int    `json:"duration"`   // in hours
	Difficulty int    `json:"difficulty"` // 1 to 5 scale
}
