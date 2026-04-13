package http

import (
	"github.com/gin-gonic/gin"
	"github.com/antyomusa/go-rest-api/internal/usecases"
)

type AuthHandler struct {
	Service *usecases.AuthService
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	token, refreshToken, err := h.Service.Register(
		req.Name,
		req.Email,
		req.Password,
	)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"token":         token,
		"refresh_token": refreshToken,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	token, refreshToken, err := h.Service.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"token":         token,
		"refresh_token": refreshToken,
	})
}