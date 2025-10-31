package routes

import (	
	"net/http"
)

func HandleTask(w http.ResponseWriter , r *http.Request) {
	w.Write([]byte("Aguante boca"))
}