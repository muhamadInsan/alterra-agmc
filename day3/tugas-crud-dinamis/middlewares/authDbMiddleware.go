package middlewares

import (
	"day3/tugas-crud-dinamis/config"
	"day3/tugas-crud-dinamis/models"

	"github.com/labstack/echo/v4"
)

var db = config.DB

func BasicAuthDb(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.User

	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return false, nil
	}
	return true, nil
}
