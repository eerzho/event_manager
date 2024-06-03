package v1

import (
	"net/http"

	"event_manager/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewHandler(l logger.Logger, router *gin.Engine) {
	h := &handler{l: l}

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", h.Ping)
	}
}

type handler struct {
	l logger.Logger
}

func (h *handler) Ping(c *gin.Context) {
	h.l.Info("pong pong")
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
