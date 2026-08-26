package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/middleware"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

// Me returns the profile of whoever's token was sent — RequireAuth has
// already run by the time this executes, so the claims are guaranteed
// to be present.
func (ctrl *UserController) Me(c *gin.Context) {
	claims := middleware.MustGetClaims(c)

	user, err := ctrl.userService.GetByID(c.Request.Context(), claims.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}