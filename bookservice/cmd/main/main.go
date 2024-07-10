package main

import (
	"log"
	"microservice/bookservice/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main () {
	r :=mux.NewRouter()

	routes.SetupRouter(r)

	http.Handle("/", r)
	
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

