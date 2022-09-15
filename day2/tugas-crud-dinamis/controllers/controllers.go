package controllers

import (
	"day2/praktek6-crud-dinamis/config"
	"day2/praktek6-crud-dinamis/lib/database"
	"day2/praktek6-crud-dinamis/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all user functino
func GetUserController(c echo.Context) error {
	users, err := database.GetUser()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "success get all user",
		"user": users,
	})
}

// find user by id
func GetUserControllerById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	users, err := database.GetUserId(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "success get user by id",
		"user": users,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	// DB.Save for save to db
	if err := config.DB.Save(&user).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

// delete function
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := database.DeleteUser(uint(id))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg": "success delete user by id",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {

	id := c.Param("id")
	email := c.FormValue("email")
	name := c.FormValue("name")

	users := models.User{
		Name:  name,
		Email: email,
	}

	if err := config.DB.Where("id = ?", id).Updates(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// 	// if err != nil {
	// 	// 	return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	// 	// }

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "success update user by id",
		"user": users,
	})
}