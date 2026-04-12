package router

import (
	"github.com/antyomusa/go-rest-api/configs"
	userHttp "github.com/antyomusa/go-rest-api/internal/delivery/http"
	"github.com/antyomusa/go-rest-api/internal/repository"
	"github.com/antyomusa/go-rest-api/internal/usecases"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine){

	db, _ := configs.ConnectDB()

	userRepo := &repository.UserRepository{DB: db}
	userService := &usecases.UserService{Repo: userRepo}
	userHandler := &userHttp.UserHandler{Service: userService}

	r.GET("/users", userHandler.GetUsers)
}