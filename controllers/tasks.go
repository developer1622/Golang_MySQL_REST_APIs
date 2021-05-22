package controllers

import (
	"fmt"
	"log"
	"net/http"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request) {
	log.Println("In GetAllTasks !")
	fmt.Fprint(w, "Gophers are ready to work! !")

	log.Println("Existing GetAllTaks !")
}
