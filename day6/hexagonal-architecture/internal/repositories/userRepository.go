package repositories

import (
	"fmt"
	"hexagonal-architecture/internal/core/domain"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	e := godotenv.Load("lokal.env")
	if e != nil {
		log.Fatalf("Error env %s", e)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		// os.Getenv("DB_USER"),
		"root",
		// os.Getenv("DB_PASSWORD"),
		"123456",
		// os.Getenv("DB_HOST"),
		"localhost",
		// os.Getenv("DB_PORT"),
		"3306",
		// os.Getenv("DB_NAME"),
		"agmc",
		"utf8mb4",
		"True",
		"Local",
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Success!")
}

// init function execute before main function
func init() {
	InitDB()
	InitialMigration()
}

// funtion for auto create table base on define struct
func InitialMigration() {
	DB.AutoMigrate(&domain.User{})
	DB.AutoMigrate(&domain.Book{})
}

// get all user functino
func GetUser() (interface{}, error) {
	var users []domain.User

	if err := DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
