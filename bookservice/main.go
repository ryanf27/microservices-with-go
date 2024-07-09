package main

import (
	"log"
	"microservice/bookservice/config"
	"microservice/bookservice/routers"
	"net/http"
)

func main () {

	config.LoadConfig()

	router := routers.SetupRouter()

	log.Println("Starting server on: 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}

