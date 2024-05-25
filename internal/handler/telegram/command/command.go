package command

import (
	"log/slog"

	"event_manager/internal/app_log"
	"gopkg.in/telebot.v3"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (h *Handler) Start(ctx telebot.Context) error {
	const op = "handler.telegram.command.Start"

	log := app_log.Logger().With(
		slog.String("op", op),
		slog.Int("update_id", ctx.Update().ID),
	)

	log.Info("starting \"start\" command")

	log.Info("sending message")
	if _, err := ctx.Bot().Send(ctx.Message().Sender, "Скорее пишите, пока AI не убежал от нас!"); err != nil {
		log.Error("failed to send message", slog.String("error", err.Error()))
	}

	return nil
}
