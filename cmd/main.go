package main

import (
	"github.com/antyomusa/go-rest-api/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	router.InitRouter(r)

	r.Run(":8080")
}
