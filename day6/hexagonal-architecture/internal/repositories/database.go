package repositories

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbConn *gorm.DB
	once   sync.Once
)

type (
	dbConfig struct {
		Host string
		User string
		Pass string
		Port string
		Name string
	}

	mysqlConfig struct {
		dbConfig
	}
)

func (conf mysqlConfig) Connect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.Name,
	)

	var err error

	dbConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func Getenv(key, fallback string) string {
	var (
		val     string
		isExist bool
	)
	val, isExist = os.LookupEnv(key)
	if !isExist {
		val = fallback
	}
	return val
}

func CreateConnection() {
	conf := dbConfig{
		User: Getenv("DB_USER", "root"),
		Pass: Getenv("DB_PASS", "1234"),
		Host: Getenv("DB_HOST", "localhost"),
		Port: Getenv("DB_PORT", "3306"),
		Name: Getenv("DB_NAME", "db_agmc"),
	}

	mysql := mysqlConfig{dbConfig: conf}
	once.Do(func() {
		mysql.Connect()
	})
}

func GetConnection() *gorm.DB {
	if dbConn == nil {
		CreateConnection()
	}
	return dbConn
}
