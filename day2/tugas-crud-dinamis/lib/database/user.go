package database

import (
	"day2/praktek6-crud-dinamis/config"
	"day2/praktek6-crud-dinamis/models"
	// "day2/praktek6-crud-dinamis/config"
)

func GetUser() (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil

}

func GetUserId(id uint) (interface{}, error) {
	var users []models.User

	if err := config.DB.Find(&users, id).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// func UpdateUsers(id int, users models.User) (interface{}, error) {
// var users []models.User

// DB.Save for save to db
// if err := config.DB.Where("id = ?", id).Updates(&users).Error; err != nil {
// 	return nil, err
// }
// if err := config.DB.Update(id, name, email).Error; err != nil {
// 	return nil, err
// }

// 	return users, nil
// }

func DeleteUser(id uint) (interface{}, error) {
	var users []models.User

	if err := config.DB.Delete(&users, id).Error; err != nil {
		return nil, err
	}

	return users, nil
}
