package handlers

import (
	"microservice/bookservice/models"
	"microservice/bookservice/utils"
	"net/http"
)


func GetBooks(w http.ResponseWriter, r *http.Request) {
    var books []models.Book
    db.Find(&books) 
    utils.SendResponse(w, books)
}
