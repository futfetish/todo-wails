package jsondb

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
	"todo/backend/database"
)

// Структура для работы с JSON БД
type Database struct {
	filePath string
	mu       sync.Mutex
}

func NewDatabase() *Database {
	// Получаем путь к корню проекта (или текущей директории)
	rootDir, err := os.Getwd()
	if err != nil {
		panic("Не удалось получить рабочую директорию")
	}

	// Формируем путь к файлу базы данных в корне
	filePath := filepath.Join(rootDir, "jsondatabase.json")

	// Проверяем, существует ли файл
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			panic("Не удалось создать файл базы данных")
		}
		defer file.Close()

		// Записываем объект с ключом "todos" в файл
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(map[string]interface{}{"todos": []database.Todo{}}); err != nil {
			panic(fmt.Sprintf("Ошибка при записи в файл: %v", err))
		}
	}

	return &Database{
		filePath: filePath,
	}
}

// Чтение всех задач из файла JSON
func (d *Database) readTodos() ([]database.Todo, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	file, err := os.ReadFile(d.filePath)
	if err != nil {
		return nil, err
	}

	if len(file) == 0 {
		return []database.Todo{}, nil
	}

	var data struct {
		Todos []database.Todo `json:"todos"`
	}

	if err := json.Unmarshal(file, &data); err != nil {
		return nil, err
	}

	return data.Todos, nil
}

// Запись задач в файл JSON
func (d *Database) writeTodos(todos []database.Todo) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	data := map[string]interface{}{
		"todos": todos,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(d.filePath, jsonData, 0644)
}

// Добавление новой задачи
func (d *Database) AddTodo(title string, priority *string, timeToComplete *int) map[string]interface{} {
	todos, err := d.readTodos()
	if err != nil {
		fmt.Println("Ошибка при чтении задач:", err)
		return nil
	}

	var newID uint
	if len(todos) > 0 {
		newID = todos[len(todos)-1].ID + 1
	}

	priorityValue := 0
	if priority != nil {
		priorityValue = database.PriorityToNumber(priority)
	}

	todo := database.Todo{
		ID:             newID,
		Title:          title,
		Completed:      false,
		TimeToComplete: timeToComplete,
		Priority:       priorityValue,
		CreateDate:     time.Now(),
	}

	todos = append(todos, todo)

	err = d.writeTodos(todos)
	if err != nil {
		fmt.Println("Ошибка при записи задач:", err)
		return nil
	}

	return database.FormatTodo(todo)
}

// Получение всех задач
func (d *Database) GetTodos(completed *bool) []map[string]interface{} {
	todos, err := d.readTodos()
	if err != nil {
		fmt.Println("Ошибка при чтении задач:", err)
		return []map[string]interface{}{}
	}

	var result []map[string]interface{}
	for _, todo := range todos {
		// Фильтрация по completed
		if completed != nil && todo.Completed != *completed {
			continue
		}
		result = append(result, database.FormatTodo(todo))
	}

	return result
}

// Переключение статуса задачи
func (d *Database) ToggleTodo(id uint) {
	todos, err := d.readTodos()
	if err != nil {
		fmt.Println("Ошибка при чтении задач:", err)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Completed = !todo.Completed
			err := d.writeTodos(todos)
			if err != nil {
				fmt.Println("Ошибка при записи задач:", err)
			}
			return
		}
	}

	fmt.Println("Задача с таким ID не найдена")
}

// Изменение задачи
func (d *Database) UpdateTodo(id uint, title string, priority *string, timeToComplete *int) (map[string]interface{}, error) {
	todos, err := d.readTodos()
	if err != nil {
		return nil, err
	}

	for i, todo := range todos {
		if todo.ID == id {
			if title != "" {
				todo.Title = title
			}
			if priority != nil {
				todo.Priority = database.PriorityToNumber(priority)
			}
			if timeToComplete != nil {
				todo.TimeToComplete = timeToComplete
				todo.CreateDate = time.Now()
			}

			todos[i] = todo

			err := d.writeTodos(todos)
			if err != nil {
				return nil, err
			}

			return database.FormatTodo(todo), nil
		}
	}

	return nil, errors.New("задача с таким ID не найдена")
}

// Удаление задачи
func (d *Database) DeleteTodo(id uint) {
	todos, err := d.readTodos()
	if err != nil {
		fmt.Println("Ошибка при чтении задач:", err)
		return
	}

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			err := d.writeTodos(todos)
			if err != nil {
				fmt.Println("Ошибка при записи задач:", err)
			}
			return
		}
	}

	fmt.Println("Задача с таким ID не найдена")
}
