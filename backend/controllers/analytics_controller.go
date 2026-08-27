package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/services"
)

type AnalyticsController struct {
	analyticsService services.AnalyticsService
}

func NewAnalyticsController(analyticsService services.AnalyticsService) *AnalyticsController {
	return &AnalyticsController{analyticsService: analyticsService}
}

func (ctrl *AnalyticsController) SignupsByDay(c *gin.Context) {
	days, err := strconv.Atoi(c.DefaultQuery("days", "30"))
	if err != nil || days <= 0 {
		days = 30
	}

	data, err := ctrl.analyticsService.SignupsByDay(c.Request.Context(), days)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}