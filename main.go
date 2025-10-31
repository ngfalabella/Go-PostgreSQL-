package main

import (
	_ "fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ngfalabella/Go-PostgreSQL-/db"
	"github.com/ngfalabella/Go-PostgreSQL-/routes"
)



func main() {

	db.DBConection()

	router := mux.NewRouter()

	router.HandleFunc("/" , routes.HomeHandler)
	router.HandleFunc("/task" , routes.HandleTask)

	http.ListenAndServe(":8080" ,  router)

}