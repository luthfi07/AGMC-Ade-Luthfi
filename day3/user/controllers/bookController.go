package controllers

import (
	"admc-day2-user/lib/database"
	"admc-day2-user/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	book, e := database.GetBooks()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  book,
	})
}

func GetBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	book, e := database.GetBookById(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"books":  book,
	})
}

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	books, e := database.CreateBook(book)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   books,
	})
}
func UpdateBookController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	bookEdit := models.Book{}
	c.Bind(&bookEdit)
	book, e := database.UpdateBook(id, bookEdit)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}

func DeleteBookController(c echo.Context) error {
	id := c.Param("id")
	book := models.Book{}
	book.Id, _ = strconv.Atoi(id)
	book, e := database.DeleteBoook(book)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"book":   book,
	})
}
