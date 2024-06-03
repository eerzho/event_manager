package v1

import (
	"context"
	"fmt"
	"strconv"

	"event_manager/internal/model"
	"event_manager/internal/service"
	"event_manager/pkg/logger"
	"gopkg.in/telebot.v3"
)

type command struct {
	l             logger.Logger
	tgUserService *service.TGUser
}

func newCommand(l logger.Logger, bot *telebot.Bot, tgUserService *service.TGUser) *command {
	c := &command{
		l:             l,
		tgUserService: tgUserService,
	}

	bot.Handle("/start", c.start)

	return c
}

func (c *command) start(ctx telebot.Context) error {
	const op = "./internal/handler/telegram/v1/command::start"

	user := model.TGUser{
		Username: ctx.Sender().Username,
		ChatID:   strconv.FormatInt(ctx.Sender().ID, 10),
	}

	err := c.tgUserService.Create(context.Background(), &user)
	if err != nil {
		c.l.Error(fmt.Errorf("%s: %w", op, err))
	}

	return ctx.Send("Скорее пишите, пока AI не убежал от нас!")
}
