package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/utils"
	"github.com/karl-luyima/fullstack-vue-go-app/backend/ws"
)

type WSController struct {
	hub        *ws.Hub
	jwtManager *utils.JWTManager
}

func NewWSController(hub *ws.Hub, jwtManager *utils.JWTManager) *WSController {
	return &WSController{hub: hub, jwtManager: jwtManager}
}

var upgrader = websocket.Upgrader{
	// In production, restrict this to your real frontend origin instead
	// of allowing all — left open here to keep local dev simple.
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (ctrl *WSController) Handle(c *gin.Context) {
	// Browsers can't set custom headers when opening a WebSocket, so the
	// token travels as a query parameter instead: ws://.../ws?token=...
	token := c.Query("token")
	if _, err := ctrl.jwtManager.ParseToken(token); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing token"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := ws.NewClient(ctrl.hub, conn)
	client.Start()
}