package main

import (
	"fmt"
	"net/http"

	"github.com/gretarst/go-rest-api/internal/comment"
	"github.com/gretarst/go-rest-api/internal/database"
	transportHTTP "github.com/gretarst/go-rest-api/internal/transport/http"
)

// App - Struct which contains pointers to dependencies
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up App")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	commentService := comment.NewService(db)

	handler := transportHTTP.NewHandler(commentService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("Go REST API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting REST API")
		fmt.Println(err)
	}
}
