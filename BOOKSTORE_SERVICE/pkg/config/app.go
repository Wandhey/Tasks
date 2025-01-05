package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "wandhey:Root1234@tcp(127.0.0.1:3306)/devdatabase?charset=utf8&parseTime=True&loc=Local&timeout=10s")
	if err != nil {
		panic(err) //If the connection fails (err != nil), the program will terminate
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
