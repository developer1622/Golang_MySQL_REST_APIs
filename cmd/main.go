package main

import (
	"log"
	"net/http"

	"github.com/developer1622/Golang_MySQL_REST_APIs/routes"
)

func main() {
	r := routes.InitializeRoutes()
	log.Println("Starting the server on 8080")
	http.ListenAndServe(":8443", r)
}
