package routes

import (
	"github.com/gorilla/mux"

	"userservice/pkg/controllers"
)

var SetupRouter = func(router *mux.Router){
    
    router.HandleFunc("/users", controllers.GetUsers).Methods("GET")
    router.HandleFunc("/users/{id}", controllers.GetUserById).Methods("GET")
    router.HandleFunc("/users", controllers.CreateUser).Methods("POST")
    router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
    router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")
    
}