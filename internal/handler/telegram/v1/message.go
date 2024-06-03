package v1

import (
	"context"
	"fmt"
	"strconv"

	"github.com/eerzho/event_manager/internal/model"
	"github.com/eerzho/event_manager/internal/service"
	"github.com/eerzho/event_manager/pkg/logger"
	"gopkg.in/telebot.v3"
)

type message struct {
	l                logger.Logger
	tgMessageService *service.TGMessage
	tgUserService    *service.TGUser
}

func newMessage(l logger.Logger, bot *telebot.Bot, tgMessageService *service.TGMessage, tgUserService *service.TGUser) *message {
	m := &message{
		l:                l,
		tgMessageService: tgMessageService,
		tgUserService:    tgUserService,
	}

	bot.Handle(telebot.OnText, m.text)

	return m
}

func (m *message) text(ctx telebot.Context) error {
	const op = "./internal/handler/telegram/v1/message::text"

	chatID := strconv.FormatInt(ctx.Sender().ID, 10)

	user := model.TGUser{
		Username: ctx.Sender().Username,
		ChatID:   chatID,
	}
	if err := m.tgUserService.Create(context.Background(), &user); err != nil {
		m.l.Error(fmt.Errorf("%s: %w", op, err))
	}

	msg := model.TGMessage{
		Text:   ctx.Message().Text,
		ChatID: chatID,
	}
	if err := m.tgMessageService.Text(context.Background(), &msg); err != nil {
		m.l.Error(fmt.Errorf("%s: %w", op, err))
		return ctx.Send("Что-то пошло не так ((")
	}

	options := telebot.SendOptions{
		ParseMode: telebot.ModeMarkdown,
		ReplyTo:   ctx.Message(),
	}

	return ctx.Send(msg.Answer, &options)
}
