package models

import (
	"userservice/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username  string `Json:"Username"`
	Email string `Json:"Email"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User {
	db.Create(&b)
	return b
}

func GetAllUsers() []User {
	var  Users []User
	db.Find(&Users)
	return Users
}

func GetUserById(Id int64) (*User, *gorm.DB) {
	var getUser User
	db := db.Where("ID=?", Id).Find(&getUser)
	return &getUser, db
}

func DeleteUser(ID int64)User{
	var user User
	db.Where("ID=?", ID).Delete(user)
	return user
}