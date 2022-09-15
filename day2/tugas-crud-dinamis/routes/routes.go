package routes

import (
	"day2/tugas-crud-dinamis/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()
	// e.GET("/status", ShowStatus)
	e.GET("v1/users", controllers.GetUserController)
	e.GET("v1/users/:id", controllers.GetUserControllerById)
	e.POST("v1/users", controllers.CreateUserController)
	e.DELETE("v1/users/:id", controllers.DeleteUserController)
	e.PUT("v1/users/:id", controllers.UpdateUserController)

	return e

}
