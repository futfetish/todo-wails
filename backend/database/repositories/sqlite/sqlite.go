package sqlite

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
	"todo/backend/database"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() *Database {
	db := &Database{}
	var err error
	db.db, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("не удалось подключиться к базе данных")
	}
	db.db.AutoMigrate(&database.Todo{})
	return db
}

func (d *Database) AddTodo(title string, priority *string, timeToComplete *int) map[string]interface{} {
	todo := database.Todo{
		Title:     title,
		Completed: false,
		Priority:  database.PriorityToNumber(priority),
		TimeToComplete: func() *int {
			if timeToComplete != nil && *timeToComplete < 0 {
				return nil
			}
			return timeToComplete
		}(),
	}
	d.db.Create(&todo)
	return database.FormatTodo(todo)
}

func (d *Database) GetTodos(completed *bool) []map[string]interface{} {
	var todos []database.Todo
	query := d.db

	// если передан completed, фильтруем по нему
	if completed != nil {
		query = query.Where("completed = ?", *completed)
	}

	if err := query.Find(&todos).Error; err != nil {
		fmt.Println("ошибка при получении задач:", err)
		return nil
	}

	// если задач нет, вернуть пустой массив, а не nil
	if len(todos) == 0 {
		return []map[string]interface{}{}
	}

	// преобразуем список в формат с priority как строкой
	var result []map[string]interface{}
	for _, todo := range todos {
		result = append(result, database.FormatTodo(todo))
	}

	return result
}

func (d *Database) ToggleTodo(id uint) {
	var todo database.Todo
	d.db.First(&todo, id)
	todo.Completed = !todo.Completed
	d.db.Save(&todo)
}

func (d *Database) DeleteTodo(id uint) {
	d.db.Delete(&database.Todo{}, id)
}

func (d *Database) UpdateTodo(id uint, title string, priority *string, timeToComplete *int) (map[string]interface{}, error) {
	// Находим задачу по ID
	var todo database.Todo
	if err := d.db.First(&todo, id).Error; err != nil {
		// Если задача не найдена, возвращаем ошибку
		return nil, fmt.Errorf("задача с id %d не найдена", id)
	}
	todo.CreateDate = time.Now()
	// Обновляем поля задачи, если они не равны nil
	todo.Title = title

	if priority != nil {
		todo.Priority = database.PriorityToNumber(priority)
	}
	if timeToComplete != nil && *timeToComplete >= 0 {
		todo.TimeToComplete = timeToComplete

	}

	// Сохраняем обновленную задачу
	if err := d.db.Save(&todo).Error; err != nil {
		log.Println("Ошибка при обновлении задачи:", err)
		return nil, err
	}

	// Возвращаем обновленную задачу в формате map
	return database.FormatTodo(todo), nil
}
