package routes

import (
	"admc-day2-user/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controllers.GetUsersControllers)
	e.GET("/users/:id", controllers.GetUserByIdControllers)
	e.POST("/users", controllers.CreateUserControllers)
	e.PUT("/users/:id", controllers.UpdateUserControllers)
	e.DELETE("/users/:id", controllers.DeleteUserControllers)

	return e
}
