package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() {
	_, err := gorm.Open(mysql.Open("root:123456@/api_test_db"), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}
}
