package main

import (
	"fmt"
	"net/http"

	transportTTP "github.com/Olu-wafemi/comment-rest-api/internal/transport/http"
)

// App - the struct which cotains thinsg like pointers to db connection
type App struct{}

func (app *App) Run() error {
	fmt.Println("Seetiing up app")
	handler := transportTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up server")
		return err
	}
	return nil
}

// Run  - sets up our app
func main() {
	fmt.Println("GO REST API")
	app := App{}

	if err := app.Run(); err != nil {
		{
			fmt.Println("Error starting the APi")
			fmt.Println(err)
		}
	}

}
