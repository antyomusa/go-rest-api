package http

import (
	"net/http"

	"github.com/antyomusa/go-rest-api/internal/usecases"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service *usecases.UserService
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.Service.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}