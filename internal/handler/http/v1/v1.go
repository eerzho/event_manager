package v1

import (
	"github.com/eerzho/event_manager/internal/service"
	"github.com/eerzho/event_manager/pkg/logger"
	"github.com/gin-gonic/gin"
)

func NewHandler(l logger.Logger, router *gin.Engine, tgUserService *service.TGUser, tgMessageService *service.TGMessage) {
	v1 := router.Group("/api/v1")
	newTGUser(l, v1, tgUserService)
	newTGMessage(l, v1, tgMessageService)
}
