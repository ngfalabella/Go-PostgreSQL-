package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ngfalabella/Go-PostgreSQL-/db"
	"github.com/ngfalabella/Go-PostgreSQL-/models"
)

func HandleTask(w http.ResponseWriter , r *http.Request) {
	w.Write([]byte("Aguante boca"))
}

func GetTasksHandler ( w http.ResponseWriter , r *http.Request ) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}

func CreateTaskHandler  ( w http.ResponseWriter , r *http.Request ) {

	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)

	createdTask := db.DB.Create(&task)

	err := createdTask.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))		
	}

	json.NewEncoder(w).Encode(&task)

}

func GetTaskHandler ( w http.ResponseWriter , r *http.Request ){
	var task models.Task 
	params := mux.Vars(r)

	db.DB.First(&task , params["id"])

	if  task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No se encontro tarea"))
		return
	}

	json.NewEncoder(w).Encode(&task)

}

func DeleteTaskHandler ( w http.ResponseWriter , r *http.Request ){
	var task models.Task 
	params := mux.Vars(r)

	db.DB.First(&task , params["id"])

	if  task.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No se encontro tarea"))
		return
	}

	db.DB.Unscoped().Delete(&task)

	w.WriteHeader(http.StatusNoContent)
	
}