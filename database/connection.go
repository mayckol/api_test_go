package database

import (
	"github.com/mayckol/api_test_go/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:123456@/api_test_db"), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})
}
