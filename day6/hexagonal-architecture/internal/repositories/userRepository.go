package repositories

import (
	"hexagonal-architecture/internal/core/domain"

	"gorm.io/gorm"
)

type user struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *user {
	return &user{db}
}

// funtion for auto create table base on define struct
// func (c *DbConn) InitialMigration() {
// 	c.db.AutoMigrate(&domain.User{})
// 	c.db.AutoMigrate(&domain.Book{})
// }

// func init() {
// 	a :=
// 	repositories.InitialMigration()
// }

// get all user functino
func (u *user) GetUser() (interface{}, error) {
	var users []domain.User

	if err := u.Db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
