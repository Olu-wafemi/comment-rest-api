package main

import "fmt"

type App struct{}

func (app *App) Run() error {
	fmt.Println("Seetiing up app")

	return nil
}
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

//App - the struct which cotains thinsg like pointers to db connection
