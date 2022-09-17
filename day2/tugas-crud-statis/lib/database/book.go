package database

import "day2/tugas-crud-statis/models"

func GetBook() (interface{}, error) {
	books := []models.Book{
		{Id: 1, Title: "Nasib Programmer di tahun 2045", Author: "penaCode"},
		{Id: 2, Title: "Anak Baru, Gaji Direktur", Author: "penaCode"},
	}

	return books, nil
}

// func GetBookById(id uint) (interface{}, error) {
// 	books := []models.Book{
// 		{Id: id, Title: "Nasib Programmer di tahun 2045", Author: "penaCode"},
// 		{Id: id, Title: "Anak Baru, Gaji Direktur", Author: "penaCode"},
// 	}
// 	return books, nil
// }

func CreateBook(title, author string) (interface{}, error) {
	// var book []models.Book

	book := []models.Book{
		{Title: title, Author: author},
	}
	// book.Title = title
	// book.Author = author
	return book, nil
}
