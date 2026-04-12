package http

import (
	"encoding/json"
	"net/http"
	"github.com/antyomusa/go-rest-api/internal/usecases"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request){
	users := usecases.GetUsers()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}