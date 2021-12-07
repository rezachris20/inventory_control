package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func NewDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/simple-pos?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
