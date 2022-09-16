package config

import (
	"fmt"

	"admc-day2-user/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	config := map[string]string{
		"DB_Username": "root",
		"DB_Password": "root",
		"DB_Port":     "3306",
		"DB_Host":     "127.0.0.1",
		"DB_Name":     "test",
	}

	connectionStirng := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionStirng), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}
func InitMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Book{})
}
