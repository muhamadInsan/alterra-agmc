package routes

import (
	c "day3/tugas-crud-dinamis/controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	// e.GET("/status", ShowStatus)
	e.GET("v1/users", c.GetUserController)
	e.GET("v1/users/:id", c.GetUserControllerById)
	e.POST("v1/users", c.CreateUserController)
	e.DELETE("v1/users/:id", c.DeleteUserController)
	e.PUT("v1/users/:id", c.UpdateUserController)

	e.POST("v1/login", c.LoginController)

	//JWT Group
	jwtAuth := e.Group("")
	jwtAuth.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	jwtAuth.GET("v1/users/:id", c.GetUserControllerById)

	// implement middleware with group rounting
	// eAuth := e.Group("")
	// eAuth.Use(echoMid.BasicAuth(m.BasicAuthDb))
	// eAuth.DELETE("v1/users/:id", c.GetUserControllerById)

	return e

}
