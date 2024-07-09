package routers

import (
	"microservice/bookservice/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/books", handlers.GetBooks).Methods("GET")
    // router.HandleFunc("/books/{id}", handlers.GetBook).Methods("GET")
    // router.HandleFunc("/books", handlers.CreateBook).Methods("POST")
    // router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
    // router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
    return router
}