package database

import (
	"admc-day2-book/models"
	"strconv"
)

func GetBooks() []models.Book {
	books := []models.Book{
		{
			Id:     1,
			Name:   "Berhitung",
			Author: "Budi",
		}, {
			Id:     2,
			Name:   "Bernyanyi",
			Author: "Budi",
		},
	}
	return books
}

func GetBookById(Id string) models.Book {
	id, _ := strconv.Atoi(Id)
	book := models.Book{
		Id:     id,
		Name:   "Berhitung",
		Author: "Budi"}
	return book
}

func UpdateBookById(Id string, bookEdit models.Book) models.Book {
	id, _ := strconv.Atoi(Id)
	book := models.Book{}
	if id == bookEdit.Id {
		book = models.Book{
			Id:     bookEdit.Id,
			Name:   bookEdit.Name,
			Author: bookEdit.Author}
	} else {
		book = models.Book{
			Id:     bookEdit.Id,
			Name:   bookEdit.Name,
			Author: bookEdit.Author}
	}

	return book
}

func DeleteBookById(Id string) models.Book {
	id, _ := strconv.Atoi(Id)
	book := models.Book{}
	if id == book.Id {

		book = models.Book{}
	} else {

		book = models.Book{}
	}
	return book
}
