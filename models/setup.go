package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/gin-gorm"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	DB = db
}
