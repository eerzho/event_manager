package service

import (
	"context"
	"fmt"

	"event_manager/internal/model"
)

type (
	TGMessageRepo interface {
		Create(ctx context.Context, message *model.TGMessage) error
	}

	TGMessage struct {
		repo                  TGMessageRepo
		tgUserService         *TGUser
		eventService          *Event
		googleCalendarService *GoogleCalendar
	}
)

func NewTGMessage(repo TGMessageRepo, tgUserService *TGUser, eventService *Event, googleCalendarService *GoogleCalendar) *TGMessage {
	return &TGMessage{
		repo:                  repo,
		tgUserService:         tgUserService,
		eventService:          eventService,
		googleCalendarService: googleCalendarService,
	}
}

func (t *TGMessage) Text(ctx context.Context, user *model.TGUser, message *model.TGMessage) error {
	const op = "./internal/service/tg_messag1e::Text"

	usr, err := t.tgUserService.ByChatID(ctx, message.ChatID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	if usr == nil {
		err = t.tgUserService.Create(ctx, user)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}
	}

	event, err := t.eventService.CreateFromText(ctx, message.Text)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	message.Url = t.googleCalendarService.CreateUrl(ctx, event)

	err = t.repo.Create(ctx, message)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
