package server

import (
	"hexagonal-architecture/internal/core/ports"
	"hexagonal-architecture/internal/repositories"

	"github.com/labstack/echo/v4"
)

type Server struct {
	//We will add every new Handler here
	userHandlers ports.UserHandler
	//middlewares ports.IMiddlewares
	//paymentHandlers ports.IPaymentHandlers
}

func NewServer(userHandlers ports.UserHandler) *Server {
	return &Server{
		userHandlers: userHandlers,
		//paymentHandlers: pHandlers
	}
}

func (s *Server) Initialize() {
	repositories.InitDB()

	app := echo.New()
	v1 := app.Group("/v1")

	userRoutes := v1.Group("/user")
	userRoutes.GET("/users", s.userHandlers.GetUser)
	// userRoutes.Post("/register", s.userHandlers.Register)

	app.Logger.Fatal(app.Start(":8080"))
}
