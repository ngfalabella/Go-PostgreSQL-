package routes

import (
	"encoding/json"
	"net/http"

	"github.com/ngfalabella/Go-PostgreSQL-/db"
	"github.com/ngfalabella/Go-PostgreSQL-/models"
)

func GetUsersHandler(w http.ResponseWriter , r *http.Request ) {
	w.Write([]byte("Listado de usuarios"))
}

func GetUserHandler(w http.ResponseWriter , r *http.Request ) {
	w.Write([]byte("Usuario individual"))
}

func PostUserHandler(w http.ResponseWriter , r *http.Request ) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)

	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)
}
func DeleteUserHandler(w http.ResponseWriter , r *http.Request ) {
	w.Write([]byte("Usuario eliminado"))
}