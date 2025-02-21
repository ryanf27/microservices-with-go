package routes

import (
	"github.com/gorilla/mux"

	"microservice/bookservice/pkg/controllers"
)

var SetupRouter = func(router *mux.Router){
    
    router.HandleFunc("/books", controllers.GetBook).Methods("GET")
    router.HandleFunc("/books/{id}", controllers.GetBookById).Methods("GET")
    router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
    router.HandleFunc("/books/{id}", controllers.UpdateBook).Methods("PUT")
    router.HandleFunc("/books/{id}", controllers.DeleteBook).Methods("DELETE")
    
}