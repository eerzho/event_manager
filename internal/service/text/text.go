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

type userRequestService interface {
	Store(cmd command.UserRequestStore) (*model.UserRequest, error)
}

type Service struct {
	aiService          aiService
	calendarService    calendarService
	userRequestService userRequestService
}

func New(aiService aiService, calendarService calendarService, userRequestService userRequestService) *Service {
	return &Service{
		aiService:          aiService,
		calendarService:    calendarService,
		userRequestService: userRequestService,
	}
}

func (s *Service) Process(cmd command.TextProcess) (string, error) {
	const op = "service.text.Process"

	log := app_log.Logger().With(
		slog.String("op", op),
		slog.Any("cmd", cmd),
	)

	var systemMessage string

	//defer func() {
	//	log.Info("saving user request")
	//	if _, err := s.userRequestService.Store(command.UserRequestStore{
	//		SenderID:      cmd.UserID,
	//		SenderMessage: cmd.Content,
	//		SystemMessage: systemMessage,
	//	}); err != nil {
	//		log.Error("failed to save user request", slog.String("error", err.Error()))
	//	}
	//}()

	log.Info("generating json")
	jsonString, err := s.aiService.Json(command.AIJson{Prompt: cmd.Content})
	if err != nil {
		log.Error("failed to generate json", slog.String("error", err.Error()))
		systemMessage = "AI взбунтовался, попробуйте пожалуйста еще раз."
		return systemMessage, err
	}

	log.Info("validating json")
	var event model.Event
	if err = json.Unmarshal([]byte(jsonString), &event); err != nil {
		log.Error("failed to validate json", slog.String("error", err.Error()))
		systemMessage = "AI отказывается разговаривать с вами, попробуйте пожалуйста еще раз."
		return systemMessage, err
	}
	if err = validator.New().Struct(event); err != nil {
		systemMessage = "AI не понимает вас, сформулируйте запрос более точно."
		if event.Message != nil {
			systemMessage = *event.Message
		}
		return systemMessage, err
	}

	log.Info("generating url")
	url, err := s.calendarService.Url(command.CalendarUrl{Event: event})
	if err != nil {
		log.Error("failed to generate url")
		systemMessage = "AI и Google Calendar не поняли друг друга, сформулируйте запрос более точно."
		return systemMessage, err
	}

	systemMessage = "[Google Calendar](" + url + ")"
	return systemMessage, nil
}
