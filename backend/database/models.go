package database

import (
	"time"
)

type Todo struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	Title          string    `json:"title"`
	Completed      bool      `json:"completed"`
	CreateDate     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createDate"`
	TimeToComplete *int      `json:"timeToComplete"` // в часах (optional)
	Priority       int       `json:"priority"`       // { 1 : 'low', 2: 'medium', : 3: 'high' }
}

var PRIORITY_ENUM = map[int]string{
	1: "low",
	2: "medium",
	3: "high",
}

func PriorityToString(priority int) *string {
	priorityMap := map[int]string{
		1: "low",
		2: "medium",
		3: "high",
	}

	if value, ok := priorityMap[priority]; ok {
		return &value
	}

	return nil
}

func PriorityToNumber(priority *string) int {
	if priority == nil {
		return 0
	}

	priorityMap := map[string]int{
		"low":    1,
		"medium": 2,
		"high":   3,
	}

	if value, ok := priorityMap[*priority]; ok {
		return value
	}

	return 0
}

func FormatTodo(todo Todo) map[string]interface{} {
	return map[string]interface{}{
		"id":             todo.ID,
		"title":          todo.Title,
		"completed":      todo.Completed,
		"createDate":     todo.CreateDate,
		"timeToComplete": todo.TimeToComplete,
		"priority":       PriorityToString(todo.Priority), // конвертация числа в строку
	}
}
