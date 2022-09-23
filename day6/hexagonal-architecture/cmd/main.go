package main

import (
	"hexagonal-architecture/internal/core/ports"
	"hexagonal-architecture/internal/core/services"
	"hexagonal-architecture/internal/handlers"
	"hexagonal-architecture/internal/repositories"
)

type Factory struct {
	UserRepository ports.UserRepository
}

func main() {
	// mongoConn := "secret🤫"
	//repositories
	userRepository := repositories.NewDbConn()

	//services
	userService := services.NewUserService(userRepository)
	//handlers
	userHandlers := handlers.NewUserHandler(userService)
	//server
	server.
}
