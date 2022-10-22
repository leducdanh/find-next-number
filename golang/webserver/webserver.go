package webserver

import (
	"fmt"
	"test-github/domain/repository"
	"test-github/routes"
)

func main() {
	config := repository.Config{}

	fmt.Println("Welcome to the webserver")
	e := routes.New()

	e.Start(config.Port)
}
