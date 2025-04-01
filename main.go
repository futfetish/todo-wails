package main

import (
	"embed"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"os"
	"todo/backend"
	"todo/backend/database"
	"todo/backend/database/repositories/jsondb"
	"todo/backend/database/repositories/sqlite"
)

//go:embed all:frontend/dist
var assets embed.FS

func GetApp() *backend.App {
	// функция создания app в соответсвии с DATABASE_TYPE

	godotenv.Load()

	dbType := os.Getenv("DATABASE_TYPE")

	fmt.Println("DATABASE_TYPE =", dbType)

	if dbType == "" {
		// Если переменная DATABASE_TYPE не установлена, используем по умолчанию sqlite
		dbType = "sqlite"
	}

	var db database.TodoRepository
	switch dbType {
	case "sqlite":
		db = sqlite.NewDatabase()
	case "json":
		db = jsondb.NewDatabase()
	default:
		db = sqlite.NewDatabase() // по умолчанию используем sqlite
	}

	// Создаем и возвращаем экземпляр приложения с выбранной базой данных
	return backend.NewApp(db)
}

func main() {
	// Create an instance of the app structure

	app := GetApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "todo",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
