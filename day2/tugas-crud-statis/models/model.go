package models

type Book struct {
	Id     uint
	Title  string `json:"title"`
	Author string `json:"author"`
}

// func BookData() (interface{}, error) {
// 	// var users = models.User
// 	books := []Book{
// 		{Id: 1, Title: "Nasib Programmer di tahun 2045", Author: "penaCode"},
// 		{Id: 2, Title: "Anak Baru, Gaji Direktur", Author: "penaCode"},
// 	}

// 	return books, nil
// }

type User struct {
	Id    uint
	Name  string `json:"name"`
	Email string `json:"email"`
}
