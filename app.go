package main

import (
	"context"
	"fmt"
)

type App struct {
	ctx context.Context
	db  *Database
}

func NewApp() *App {
	return &App{db: NewDatabase()}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("hello %s, it's show time!", name)
}

// API для работы с задачами
func (a *App) AddTodo(title string, priority *string, timeToComplete *int) map[string]interface{} {
	return a.db.AddTodo(title, priority, timeToComplete)
}

func (a *App) GetTodos(completed *bool) []map[string]interface{} {
	return a.db.GetTodos(completed)
}

func (a *App) ToggleTodo(id uint) {
	a.db.ToggleTodo(id)
}

func (a *App) DeleteTodo(id uint) {
	a.db.DeleteTodo(id)
}
