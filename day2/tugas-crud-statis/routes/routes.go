package routes

import (
	"day2/praktek5-homework/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello !")
	})
	e.GET("v1/books", controllers.GetAllBooks)
	e.GET("v1/books/:id", controllers.GetBookById)
	e.POST("v1/books", controllers.CreateBook)
	e.DELETE("v1/delete/:id", controllers.DeleteBook)
	e.PUT("v1/put/:id", controllers.UpdateBook)

	return e
}
