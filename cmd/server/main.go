package main

import "fmt"

// App - Struct which contains pointers to dependencies
type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up App")
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
