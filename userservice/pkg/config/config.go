package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db * gorm.DB
)

func Connect(){
	var err error

	d, err := gorm.Open(postgres.Open("host=localhost port=5432 user=dev dbname=go_users password=qwerty sslmode=disable"), &gorm.Config{})
	
	if err != nil{
		log.Fatal("Failed to connect to database: ", err)
	}

	db = d

	log.Println("Connected to database")

}

func GetDB() *gorm.DB{
	return db
}