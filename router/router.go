package router

import (
	"net/http"
	"github.com/antyomusa/go-rest-api/configs"
	userHttp "github.com/antyomusa/go-rest-api/internal/delivery/http"
	"github.com/antyomusa/go-rest-api/internal/repository"
	"github.com/antyomusa/go-rest-api/internal/usecases"
)

func InitRouter() http.Handler {
	mux := http.NewServeMux()

	db, _ := configs.ConnectDB()

	userRepo := &repository.UserRepository{DB: db}
	userService := &usecases.UserService{Repo: userRepo}
	userHandler := &userHttp.UserHandler{Service: userService}

	mux.HandleFunc("/users", userHandler.GetUsers)

	return mux
}