package routes

import (
	"admc-day2-book/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/books", controllers.GetBooksController)
	e.GET("/books/:id", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)

	return e
}
