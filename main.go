package main

import (
	_ "fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ngfalabella/Go-PostgreSQL-/db"
	"github.com/ngfalabella/Go-PostgreSQL-/models"
	"github.com/ngfalabella/Go-PostgreSQL-/routes"
)



func main() {

	db.DBConection()

	db.DB.AutoMigrate( &models.Task{} )
	db.DB.AutoMigrate( &models.User{} )
	

	router := mux.NewRouter()

	router.HandleFunc("/" , routes.HomeHandler)
	router.HandleFunc("/task" , routes.HandleTask)

	router.HandleFunc("/users" , routes.GetUsersHandler).Methods("GET")
	router.HandleFunc("/users/{id}" , routes.GetUserHandler).Methods("GET")
	router.HandleFunc("/users" , routes.PostUserHandler).Methods("POST")
	router.HandleFunc("/users/{id}" , routes.DeleteUserHandler).Methods("DELETE")


	http.ListenAndServe(":8080" ,  router)

}