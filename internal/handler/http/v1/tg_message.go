package v1

import (
	"fmt"
	"net/http"

	"event_manager/internal/service"
	"event_manager/pkg/logger"
	"github.com/gin-gonic/gin"
)

type tgMessage struct {
	l                logger.Logger
	tgMessageService *service.TGMessage
}

func newTGMessage(l logger.Logger, router *gin.RouterGroup, tgMessageService *service.TGMessage) *tgMessage {
	t := tgMessage{
		l:                l,
		tgMessageService: tgMessageService,
	}

	router.GET("/tg-messages", t.all)

	return &t
}

func (t *tgMessage) all(ctx *gin.Context) {
	const op = "./internal/handler/http/v1/tg_message::all"

	messages, err := t.tgMessageService.All(ctx)
	if err != nil {
		t.l.Error(fmt.Errorf("%s: %w", op, err))
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, messages)
}
