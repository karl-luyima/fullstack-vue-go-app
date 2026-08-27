package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/services"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/ws"
)

type AuthController struct {
	authService services.AuthService
	hub         *ws.Hub
}

func NewAuthController(authService services.AuthService, hub *ws.Hub) *AuthController {
	return &AuthController{authService: authService, hub: hub}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var input services.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.authService.Register(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Tell every connected browser that a new user just joined.
	ctrl.hub.Broadcast <- []byte(`{"event":"new_signup","name":"` + result.User.Name + `"}`)

	c.JSON(http.StatusCreated, result)
}

func (ctrl *AuthController) Login(c *gin.Context) {
	var input services.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ctrl.authService.Login(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}