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

type message struct {
	l                logger.Logger
	tgMessageService *service.TGMessage
}

func newMessage(l logger.Logger, bot *telebot.Bot, tgMessageService *service.TGMessage) *message {
	m := &message{
		l:                l,
		tgMessageService: tgMessageService,
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
	msg := model.TGMessage{
		Text:   ctx.Message().Text,
		ChatID: chatID,
	}

	err := m.tgMessageService.Text(context.Background(), &user, &msg)
	if err != nil {
		m.l.Error(fmt.Errorf("%s: %w", op, err))
		return ctx.Send("Что-то пошло не так (")
	}

	options := telebot.SendOptions{
		ParseMode: telebot.ModeMarkdown,
		ReplyTo:   ctx.Message(),
	}

	return ctx.Send("[Google Calendar]("+msg.Url+")", &options)
}
