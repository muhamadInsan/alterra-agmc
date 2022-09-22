package main

import (
	"hexagonal-architecture/internal/core/services"
	"hexagonal-architecture/internal/handlers"
	"hexagonal-architecture/internal/repositories"
)

func main() {
	// mongoConn := "secret🤫"
	//repositories
	userRepository := repositories.InitDB()
	//services
	userService := services.NewUserService(userRepository)
	//handlers
	userHandlers := handlers.NewUserHandler(userService)
	//server
	server.Initialize(
		userHandlers,
	)
}
