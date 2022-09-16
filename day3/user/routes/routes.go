package routes

import (
	"admc-day2-user/constants"
	c "admc-day2-user/controllers"
	m "admc-day2-user/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoMid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	// e.GET("/users", c.GetUsersControllers)
	// e.GET("/users/:id", c.GetUserByIdControllers)
	// e.POST("/users", c.CreateUserControllers)
	// e.PUT("/users/:id", c.UpdateUserControllers)
	// e.DELETE("/users/:id", c.DeleteUserControllers)

	//implement middleware with group routing
	e.POST("/login", c.LoginUserController)
	e.POST("/users", c.CreateUserControllers)

	Auth := e.Group("")
	Auth.Use(echoMid.BasicAuth(m.BasicAuthDB))

	// eAuth.DELETE("/users/:id", c.DeleteUserControllers)
	// eAuth.PUT("/users/:id", c.UpdateUserControllers)
	// //JWT Group
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT_KEY)))

	r.GET("/users", c.GetUsersControllers)
	r.GET("/users/:id", c.GetUserByIdControllers)
	r.PUT("/users/:id", c.UpdateUserControllers)
	r.DELETE("/users/:id", c.DeleteUserControllers)

	e.GET("/books", c.GetBooksController)
	e.GET("/books/:id", c.GetBookController)
	r.POST("/books", c.CreateBookController)
	r.PUT("/books/:id", c.UpdateBookController)
	r.DELETE("/books/:id", c.DeleteBookController)

	return e
}
