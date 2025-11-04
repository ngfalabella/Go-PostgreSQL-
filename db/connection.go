package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//gormgolang
var DSN string = "host=localhost user=postgres password=123456 dbname=golanguno port=5432"

var DB *gorm.DB

func DBConection() {
	var error error
	DB , error = gorm.Open( postgres.Open(DSN) , &gorm.Config{}) 
	if error != nil {
		log.Fatal(error)
	}else{
		log.Println("Succes se conecto insta")
	}
}