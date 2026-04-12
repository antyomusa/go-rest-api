package http

import (
	"encoding/json"
	"net/http"
	"github.com/antyomusa/go-rest-api/internal/usecases"
)

type UserHandler struct {
	Service *usecases.UserService
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.GetUsers()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(users)
}