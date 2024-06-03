package service

import (
	"context"
	"fmt"

	"event_manager/internal/model"
	"event_manager/pkg/logger"
)

type (
	TGMessageRepo interface {
		Create(ctx context.Context, message *model.TGMessage) error
	}

	TGMessage struct {
		l                     logger.Logger
		repo                  TGMessageRepo
		tgUserService         *TGUser
		eventService          *Event
		googleCalendarService *GoogleCalendar
	}
)

func NewTGMessage(l logger.Logger, repo TGMessageRepo, tgUserService *TGUser, eventService *Event, googleCalendarService *GoogleCalendar) *TGMessage {
	return &TGMessage{
		l:                     l,
		repo:                  repo,
		tgUserService:         tgUserService,
		eventService:          eventService,
		googleCalendarService: googleCalendarService,
	}
}

func (t *TGMessage) Text(ctx context.Context, message *model.TGMessage) error {
	const op = "./internal/service/tg_message::Text"

	defer func() {
		if err := t.repo.Create(ctx, message); err != nil {
			t.l.Error(fmt.Errorf("%s: %w", op, err))
		}
	}()

	var event model.Event
	if err := t.eventService.CreateFromText(ctx, &event, message.Text); err != nil {
		t.l.Error(fmt.Errorf("%s: %w", op, err))
		return fmt.Errorf("%s: %w", op, err)
	}
	if event.Message != "" {
		message.Answer = event.Message
		return nil
	}

	url := t.googleCalendarService.CreateUrl(ctx, &event)
	message.Answer = "[Google Calendar](" + url + ")"

	return nil
}
