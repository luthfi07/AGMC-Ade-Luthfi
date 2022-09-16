package controllers

import (
	"admc-day2-user/constants"
	"admc-day2-user/lib/database"
	"admc-day2-user/models"
	"net/http"
	"strconv"
	"time"

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

	cookie, err := c.Cookie(constants.Cookie)
	if err != nil {
		return err
	}
	userid := strconv.Itoa(int(id))
	//cek authorized
	if cookie.Value == userid {
		users, e := database.UpdateUser(user, id)

		if e != nil {
			return echo.NewHTTPError(http.StatusBadRequest, e.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"users":  users,
			"cookie": cookie,
		})
	} else {
		//unauthorized
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

}

func DeleteUserControllers(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := models.User{}
	user.ID = uint(id)
	cookie, err := c.Cookie(constants.Cookie)
	if err != nil {
		return err
	}
	userid := strconv.Itoa(int(id))
	//cek authorized
	if cookie.Value == userid {
		users, e := database.DeleteUser(user, id)

		if e != nil {
			return echo.NewHTTPError(http.StatusBadRequest, e.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"users":  users,
		})
	} else {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
}

func LoginUserController(c echo.Context) error {
	user := models.User{}

	c.Bind(&user)

	users, e := database.LoginUsers(&user)
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	userid := strconv.Itoa(int(user.ID))
	cookie := new(http.Cookie)
	cookie.Name = constants.Cookie
	cookie.Value = userid
	cookie.Expires = time.Now().Add(1 * time.Hour)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
		"Cookie": cookie,
	})
}

func GetUserDetailControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	users, err := database.GetDetailUser((id))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}
