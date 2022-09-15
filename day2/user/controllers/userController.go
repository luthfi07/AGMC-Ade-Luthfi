package controllers

import (
	"admc-day2-user/lib/database"
	"admc-day2-user/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUsersControllers(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserByIdControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	users, e := database.GetUserById(id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func CreateUserControllers(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	users, e := database.CreateUser(user)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func UpdateUserControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	c.Bind(&user)
	users, e := database.UpdateUser(user, id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func DeleteUserControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	user.ID = uint(id)
	users, e := database.DeleteUser(user, id)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}
