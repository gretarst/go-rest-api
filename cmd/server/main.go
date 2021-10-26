package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/gretarst/go-rest-api/internal/transport/http"
)

// App - Struct which contains pointers to dependencies
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up App")

	handler := transportHTTP.NewHandler()
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
