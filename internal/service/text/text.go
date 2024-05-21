package text

import (
	"encoding/json"
	"log/slog"

	"event_manager/internal/app_log"
	"event_manager/internal/dto/command"
	"event_manager/internal/dto/model"
	"github.com/go-playground/validator/v10"
)

type aiService interface {
	Json(cmd command.AIJson) (string, error)
}

type calendarService interface {
	Url(cmd command.CalendarUrl) (string, error)
}

type Service struct {
	aiService       aiService
	calendarService calendarService
}

func New(aiService aiService, calendarService calendarService) *Service {
	return &Service{
		aiService:       aiService,
		calendarService: calendarService,
	}
}

func (s *Service) Process(cmd command.TextProcess) (string, error) {
	const op = "service.text.Process"

	log := app_log.Logger().With(
		slog.String("op", op),
		slog.Any("cmd", cmd),
	)

	log.Info("generating json")
	jsonString, err := s.aiService.Json(command.AIJson{Prompt: cmd.Content})
	if err != nil {
		log.Error("failed to generate json", slog.String("error", err.Error()))
		return "", err
	}

	log.Info("validating json")
	var event model.Event
	if err = json.Unmarshal([]byte(jsonString), &event); err != nil {
		log.Error("failed to validate json", slog.String("error", err.Error()))
		return "", err
	}
	if err = validator.New().Struct(event); err != nil {
		if event.Message != nil {
			return *event.Message, err
		}
		return "", err
	}

	log.Info("generating url")
	url, err := s.calendarService.Url(command.CalendarUrl{Event: event})
	if err != nil {
		log.Error("failed to generate url")
		return "", err
	}

	return "[Google Calendar](" + url + ")", nil
}
