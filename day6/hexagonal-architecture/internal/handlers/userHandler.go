package handlers

import (
	"hexagonal-architecture/internal/core/ports"

	"github.com/labstack/echo/v4"
)

type UserHandlers struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (h *UserHandlers) GetUser(c echo.Context) {

}
