package text

import (
	"log/slog"

	"event_manager/internal/app_log"
	"event_manager/internal/dto/command"
	"gopkg.in/telebot.v3"
)

type textService interface {
	Process(cmd command.TextProcess) (string, error)
}

type Handler struct {
	textService textService
}

func New(textService textService) *Handler {
	return &Handler{
		textService: textService,
	}
}

func (h *Handler) Process(ctx telebot.Context) error {
	const op = "handler.telegram.text.Process"

	log := app_log.Logger().With(
		slog.String("op", op),
		slog.Int("update_id", ctx.Update().ID),
	)

	log.Info("starting process text")
	result, err := h.textService.Process(command.TextProcess{Content: ctx.Message().Text})
	if err != nil {
		log.Error("failed to process text", slog.String("error", err.Error()))
		if result == "" {
			result = "Я не понял ваш запрос("
		}
	}

	log.Info("sending message")
	options := telebot.SendOptions{
		ParseMode: telebot.ModeMarkdown,
		ReplyTo:   ctx.Message(),
	}
	if _, err = ctx.Bot().Send(ctx.Message().Sender, result, &options); err != nil {
		log.Error("failed to send message", slog.String("error", err.Error()))
		return err
	}

	return nil
}
