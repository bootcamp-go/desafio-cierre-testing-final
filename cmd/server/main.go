package main

import (
	"app/cmd/server/dependencies"
	"os"
)

func main() {
	// env
	// ...

	// app
	app := dependencies.NewApp(os.Getenv("SERVER_ADDR"))

	// run
	err := app.Run()
	if err != nil {
		panic(err)
	}
}