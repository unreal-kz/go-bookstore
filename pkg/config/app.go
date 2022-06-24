package config

import (
	"github.com/jinzhu/grom"
	_ "github.com/jinzhu/grom/dialects/mysql"
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
