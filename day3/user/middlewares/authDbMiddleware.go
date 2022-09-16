package middlewares

import (
	"admc-day2-user/config"
	"admc-day2-user/models"
	"fmt"

	"github.com/labstack/echo/v4"
)

var db = config.DB

func BasicAuthDB(username, password string, c echo.Context) (bool, error) {
	var db = config.DB
	var user models.User
	fmt.Println(username)
	fmt.Println(password)

	if err := db.Where("email = ? AND password = ?", username, password).First(&user).Error; err != nil {
		fmt.Println(err)

		return false, nil
	}

	return true, nil
}
