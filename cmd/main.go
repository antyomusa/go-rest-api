package main

import (
	"fmt"
	"net/http"
	"github.com/antyomusa/go-rest-api/router"
)

func main() {
	fmt.Println("Server is running on port 8080")

	r := router.InitRouter()

	http.ListenAndServe(":8080", r)
}
