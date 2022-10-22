package main

import (
	"fmt"
	"test-github/domain/repository"
	"test-github/routes"
)

func main() {
	config := repository.Config{}

	fmt.Println("Welcome to the webserver111")
	e := routes.New()

	e.Logger.Fatal(e.Start(config.Port))
}
