package routes

import (
	"github.com/gorilla/mux"

	"github.com/developer1622/Golang_MySQL_REST_APIs/controllers"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/books", controllers.GetAllTasks).Methods("GET")
	return r
}
