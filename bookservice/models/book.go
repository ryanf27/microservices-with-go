package models

type Book struct {
	ID     uint   `gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
