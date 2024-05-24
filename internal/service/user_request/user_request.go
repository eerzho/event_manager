package user_request

import (
	"event_manager/internal/dto/command"
	"event_manager/internal/dto/model"
	"event_manager/internal/dto/query"
)

type userRequestRepository interface {
	Create(qry query.UserRequestCreate) (*model.UserRequest, error)
}

type Service struct {
	repository userRequestRepository
}

func New(repository userRequestRepository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Store(cmd command.UserRequestStore) (*model.UserRequest, error) {
	return s.repository.Create(query.UserRequestCreate{
		SenderID:      cmd.SenderID,
		SenderMessage: cmd.SenderMessage,
		SystemMessage: cmd.SystemMessage,
	})
}
