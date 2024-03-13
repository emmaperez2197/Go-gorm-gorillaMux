package main

import (
	"fmt"
	"net/http"

	"github.com/emmaperez2197/api/db"
	"github.com/emmaperez2197/api/routes"
	"github.com/emmaperez2197/api/routes/models"
	"github.com/gorilla/mux"
)

func ecoDeLaMontana(mensaje string, iteraciones uint) {
	if iteraciones < 10 {
		ecoDeLaMontana(mensaje, iteraciones+1)
	}
	fmt.Println(mensaje, iteraciones)
}

func main() {

	ecoDeLaMontana("yodelayheehoo", 5)
	db.DBConnection()

	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomeHandler)

	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users/create", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/delete", routes.UpdateUserHandler).Methods("PUT")

	http.ListenAndServe(":3000", r)
}
