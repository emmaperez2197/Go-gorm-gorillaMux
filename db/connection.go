package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS = "host=localhost user=emmago password=12345 dbname=gorm port=5432"
var DB *gorm.DB

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DNS), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connection")
	}

}
