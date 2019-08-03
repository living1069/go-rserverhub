package app

import (
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func CreateConnection() {
	var uri = "host=10.8.0.1 port=5432 user=dev password=dev dbname=rserverhub sslmode=disable"
	var err error
	DB, err = gorm.Open("postgres", uri)

	if err != nil {
		log.Fatal(err.Error())
	}

	DB.LogMode(true)
}
