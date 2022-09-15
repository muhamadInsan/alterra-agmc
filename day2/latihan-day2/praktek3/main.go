package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id    uint
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GetUser(c echo.Context) error {
	user := User{Name: "Ismail", Email: "ismail@alterra.id"}
	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	// get data from value
	name := c.FormValue("name")
	email := c.FormValue("email")

	var user User
	user.Name = name
	user.Email = email

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// User Controller
func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{Id: uint(id), Name: "Ismail", Email: "ismail@alterra.id"}
	// Render Data - JSON Response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func UserSearchController(c echo.Context) error {
	// get data from query param
	match := c.QueryParam("match")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"match":  match,
		"result": []string{"adi", "aan", "asif"}, // hardcode data
	})
}

func main() {
	// create a new echo instance
	e := echo.New()
	// Route / to handler function
	e.GET("/", HelloController)
	e.GET("/users", GetUser)
	// e.GET("/users", UserSearchController)
	e.GET("/users/:id", GetUserController)
	// start the server, and log if it fails

	e.POST("/users", CreateUser)
	e.Start(":8080")
}

// handler - Simple handler to make sure environment is setup
func HelloController(c echo.Context) error {
	// return the string "Hello World" as the response body
	// with an http.StatusOK (200) status
	return c.String(http.StatusOK, "Hello World")
}
