package repositories

import (
	"fmt"
	"hexagonal-architecture/internal/core/domain"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConn struct {
	db *gorm.DB
}

// NewAdapter creates a new Adapter
func NewDbConn() *DbConn {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		"utf8mb4",
		"True")

	// connect
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("db connection failur: %v", err)
	}
	return &DbConn{db: db}
}

// funtion for auto create table base on define struct
func (c *DbConn) InitialMigration() {
	c.db.AutoMigrate(&domain.User{})
	c.db.AutoMigrate(&domain.Book{})
}

// func init() {
// 	a :=
// 	repositories.InitialMigration()
// }

// get all user functino
func (c *DbConn) GetUser() (interface{}, error) {
	var users []domain.User

	if err := c.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
