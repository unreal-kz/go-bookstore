package config

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *grom.DB
)

func Connect() {
	d, err := grom.Open("mysql", "baubek:B@123/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *grom.DB {
	return db
}
