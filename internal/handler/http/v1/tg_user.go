package v1

import (
	"fmt"
	"net/http"

	"github.com/eerzho/event_manager/internal/service"
	"github.com/eerzho/event_manager/pkg/logger"
	"github.com/gin-gonic/gin"
)

type tgUser struct {
	l             logger.Logger
	tgUserService *service.TGUser
}

func newTGUser(l logger.Logger, router *gin.RouterGroup, tgUserService *service.TGUser) *tgUser {
	t := &tgUser{
		l:             l,
		tgUserService: tgUserService,
	}

	router.GET("/tg-users", t.all)

	return t
}

func (t *tgUser) all(ctx *gin.Context) {
	const op = "./internal/handler/http/v1/tg_user::all"

	users, err := t.tgUserService.All(ctx)
	if err != nil {
		t.l.Error(fmt.Errorf("%s: %w", op, err))
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, users)
}
