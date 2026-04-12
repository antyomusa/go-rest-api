package router

import (
	"net/http"
	userHttp "github.com/antyomusa/go-rest-api/internal/delivery/http"
)

func InitRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", userHttp.GetUsersHandler)

	return mux
}