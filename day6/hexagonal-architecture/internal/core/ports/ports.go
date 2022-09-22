package ports

import "github.com/labstack/echo/v4"

type UserService interface {
	GetUser() error
	GetUserId(id uint) error
	CreateUser(name, email, password string) error
	DeleteUser(id uint) error
	UpdateUser(id uint) error
}

type UserRepository interface {
	GetUser() error
	GetUserId(id uint) error
	CreateUser(name, email, password string) error
	DeleteUser(id uint) error
	UpdateUser(id uint) error
}

type UserHandler interface {
	GetUser(c echo.Context) error
	GetUserId(id uint) error
	CreateUser(name, email, password string) error
	DeleteUser(id uint) error
	UpdateUser(id uint) error
}

type BookService interface {
	GetBook() error
	GetBookId(id uint) error
	CreateBook(title, auhtor string) error
	DeleteBook(id uint) error
	UpdateBook(id uint) error
}

type Server interface {
	Initialize()
}
