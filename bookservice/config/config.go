package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func LoadConfig(){
	var err error

	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres dbname=books password=postgres sslmode=disable"), &gorm.Config{})
	
	if err != nil{
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Connected to database")

}