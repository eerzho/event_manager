package v1

import (
	"event_manager/internal/service"
	"event_manager/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewHandler(l logger.Logger, router *gin.Engine, tgUserService *service.TGUser, tgMessageService *service.TGMessage) {
	v1 := router.Group("/api/v1")
	newTGUser(l, v1, tgUserService)
	newTGMessage(l, v1, tgMessageService)
}
