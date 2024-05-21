package google

import (
	"fmt"
	"log/slog"
	"net/url"

	"event_manager/internal/app_log"
	"event_manager/internal/dto/command"
)

type Service struct {
	url string
}

func New(url string) *Service {
	return &Service{url: url}
}

func (s *Service) Url(cmd command.CalendarUrl) (string, error) {
	const op = "service.calendar.google.Url"

	log := app_log.Logger().With(
		slog.String("op", op),
		slog.Any("cmd", cmd),
	)

	log.Info("generating url")

	params := url.Values{}
	params.Add("action", "TEMPLATE")
	params.Add("text", cmd.Event.Text)
	params.Add("dates", fmt.Sprintf("%s/%s", cmd.Event.StartDate, cmd.Event.EndDate))
	if cmd.Event.CTZ != nil {
		params.Add("ctz", *cmd.Event.CTZ)
	}
	params.Add("details", cmd.Event.Details)
	if cmd.Event.Location != nil {
		params.Add("location", *cmd.Event.Location)
	}
	params.Add("crm", cmd.Event.CRM)
	if cmd.Event.TRP {
		params.Add("trp", "true")
	} else {
		params.Add("trp", "false")
	}
	if cmd.Event.Recur != nil {
		params.Add("recur", *cmd.Event.Recur)
	}

	return fmt.Sprintf("%s?%s", s.url, params.Encode()), nil
}
