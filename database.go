package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
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

func priorityToString(priority int) *string {
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

func priorityToNumber(priority *string) int {
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

type Database struct {
	db *gorm.DB
}

func NewDatabase() *Database {
	database := &Database{}
	var err error
	database.db, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("не удалось подключиться к базе данных")
	}
	database.db.AutoMigrate(&Todo{})
	return database
}

func (d *Database) AddTodo(title string, priority *string, timeToComplete *int) Todo {
	todo := Todo{
		Title:          title,
		Completed:      false,
		Priority:       priorityToNumber(priority),
		TimeToComplete: timeToComplete,
	}
	d.db.Create(&todo)
	return todo
}

func (d *Database) GetTodos(completed *bool) []map[string]interface{} {
	var todos []Todo
	query := d.db

	// если передан completed, фильтруем по нему
	if completed != nil {
		query = query.Where("completed = ?", *completed)
	}

	query.Find(&todos)

	// преобразуем список в формат с priority как строкой
	var result []map[string]interface{}
	for _, todo := range todos {
		result = append(result, map[string]interface{}{
			"id":             todo.ID,
			"title":          todo.Title,
			"completed":      todo.Completed,
			"createDate":     todo.CreateDate,
			"timeToComplete": todo.TimeToComplete,
			"priority":       priorityToString(todo.Priority), // конвертация числа в строку
		})
	}

	return result
}

func (d *Database) ToggleTodo(id uint) {
	var todo Todo
	d.db.First(&todo, id)
	todo.Completed = !todo.Completed
	d.db.Save(&todo)
}

func (d *Database) DeleteTodo(id uint) {
	d.db.Delete(&Todo{}, id)
}
