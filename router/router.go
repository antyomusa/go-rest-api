package router

import (
	"github.com/antyomusa/go-rest-api/configs"
	userHttp "github.com/antyomusa/go-rest-api/internal/delivery/http"
	"github.com/antyomusa/go-rest-api/internal/repository"
	"github.com/antyomusa/go-rest-api/internal/usecases"
	"github.com/antyomusa/go-rest-api/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine){

	db, _ := configs.ConnectDB()

	userRepo := &repository.UserRepository{DB: db}
	userService := &usecases.UserService{Repo: userRepo}
	userHandler := &userHttp.UserHandler{Service: userService}

	authService := &usecases.AuthService{Repo: userRepo}
	authHandler := &userHttp.AuthHandler{Service: authService}

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)

	// protected route
	authGroup := r.Group("/")
	authGroup.Use(middleware.JWTAuthMiddleware())
	{
		authGroup.GET("/users", userHandler.GetUsers)
	}
}