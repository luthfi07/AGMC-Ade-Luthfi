package database

import (
	"admc-day2-user/config"
	"admc-day2-user/models"
)

func GetBooks() (interface{}, error) {
	var books []models.Book
	if e := config.DB.Find(&books).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func GetBookById(id int) (interface{}, error) {
	var books []models.Book

	if e := config.DB.Find(&books, id).Error; e != nil {
		return nil, e
	}
	return books, nil
}

func CreateBook(book models.Book) (interface{}, error) {

	if e := config.DB.Create(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func UpdateBook(id int, book models.Book) (interface{}, error) {

	if e := config.DB.Model(models.Book{}).Where("id = ?", id).Updates(&book).Error; e != nil {
		return nil, e
	}
	return book, nil
}

func DeleteBoook(book models.Book) (models.Book, error) {

	if e := config.DB.Delete(&models.Book{}, book.Id).Error; e != nil {
		return book, e
	}
	return book, nil
}
