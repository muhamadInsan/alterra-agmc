package handlers

import (
	"hexagonal-architecture/internal/core/ports"
)

type UserHandlers struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (h *UserHandlers) GetUser() error {
	err := h.userService.GetUser()
	if err != nil {
		return err
	}

	return nil
}
