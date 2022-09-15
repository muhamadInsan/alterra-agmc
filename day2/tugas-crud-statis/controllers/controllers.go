package controllers

import (
	"day2/praktek5-homework/lib/database"
	"day2/praktek5-homework/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllBooks(c echo.Context) error {
	books, err := database.GetBook()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "success get all user",
		"user": books,
	})
}

func GetBookById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	books := []models.Book{
		{Id: uint(id), Title: "Nasib Programmer di tahun 2045", Author: "penaCode"},
		{Id: uint(id), Title: "Anak Baru, Gaji Direktur", Author: "penaCode"},
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"msg":  "success user by id",
		"user": books,
	})
}

func CreateBook(c echo.Context) error {
	title := c.FormValue("title")
	author := c.FormValue("author")

	books, err := database.CreateBook(title, author)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"data":     books,
	})
}

func DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	books := []models.Book{
		{Id: uint(id), Title: "Nasib Programmer di tahun 2045", Author: "penaCode"},
		{Id: uint(id), Title: "Anak Baru, Gaji Direktur", Author: "penaCode"},
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "delete success",
		"data":     books,
	})
}

func UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	title := c.Param("title")
	author := c.Param("author")

	books := []models.Book{
		{Id: uint(id), Title: title, Author: author},
		{Id: uint(id), Title: title, Author: author},
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "update success",
		"data":     books,
	})
}
