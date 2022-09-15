package main

import (
	// "day2/praktek6-crud-dinamis/config"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// init connect to mysql db
func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		"admin",
		"12345",
		"localhost",
		"3306",
		"test",
		"utf8mb4",
		"True",
		"Local",
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Success!")
}

// init function execute before main function
func init() {
	InitDB()
	InitialMigration()
}

// funtion for auto migrate base on define struct
func InitialMigration() {
	DB.AutoMigrate(&User{})
}

func ShowStatus(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}

// struct user model
type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// get all user functino
func GetUserController(c echo.Context) error {
	var users []User

	if err := DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "success get all user",
		"user": users,
	})
}

// find user by id
func GetUserControllerById(c echo.Context) error {
	id := c.Param("id")
	var users []User

	if err := DB.Find(&users, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "success get user by id",
		"user": users,
	})
}

// delete function
func DeleteUserControllerId(c echo.Context) error {
	id := c.Param("id")
	var users []User

	if err := DB.Delete(&users, id).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "delete success",
		"user": users,
	})

}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	// DB.Save for save to db
	if err := DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {

	id := c.Param("id")
	email := c.FormValue("email")
	name := c.FormValue("name")
	password := c.FormValue("password")

	users := User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	if err := DB.Where("id = ?", id).Updates(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "success update user by id",
		"user": users,
	})
}

func main() {
	e := echo.New()

	e.GET("/status", ShowStatus)
	e.GET("/users", GetUserController)
	// challenge
	e.GET("/users/:id", GetUserControllerById)
	e.DELETE("/users/:id", DeleteUserControllerId)
	e.PUT("/users/:id", UpdateUserController)
	// ==============
	e.POST("/users", CreateUserController)

	e.Start(":8080")
}
