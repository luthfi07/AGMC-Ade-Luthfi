package controllers

import (
	"admc-day2-book/lib/database"
	"admc-day2-book/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	book := database.GetBooks()
	return c.JSON(http.StatusOK, book)
}

func GetBookController(c echo.Context) error {
	id := c.Param("id")

	book := database.GetBookById(id)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"book": book,
	})
}

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Create Book",
		"book":    book,
	})
}
func UpdateBookController(c echo.Context) error {
	id := c.Param("id")
	bookEdit := models.Book{}
	c.Bind(&bookEdit)
	book := database.UpdateBookById(id, bookEdit)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Update Book",
		"book":    book,
	})
}

func DeleteBookController(c echo.Context) error {
	id := c.Param("id")
	book := database.DeleteBookById(id)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Delete Book",
		"book":    book.Id,
	})
}
