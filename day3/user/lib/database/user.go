package database

import (
	"admc-day2-user/config"
	"admc-day2-user/middlewares"
	"admc-day2-user/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func GetUserById(id int) (interface{}, error) {
	var users []models.User

	if e := config.DB.Find(&users, id).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func CreateUser(user models.User) (interface{}, error) {

	if e := config.DB.Create(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func UpdateUser(user models.User, id int) (interface{}, error) {

	if e := config.DB.Model(models.User{}).Where("id = ?", id).Updates(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func DeleteUser(user models.User, id int) (interface{}, error) {

	if e := config.DB.Delete(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func GetDetailUser(userId int) (interface{}, error) {
	var user models.User
	if e := config.DB.Find(&user, userId).Error; e != nil {
		return nil, e
	}
	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}
	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
