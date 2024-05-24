package user_request

import (
	"event_manager/internal/dto/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRequest struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	SenderID      int64              `bson:"sender_id"`
	SenderMessage string             `bson:"sender_message"`
	SystemMessage string             `bson:"system_message"`
}

func toModel(entity *UserRequest) *model.UserRequest {
	return &model.UserRequest{
		ID:            entity.ID.String(),
		SenderID:      entity.SenderID,
		SenderMessage: entity.SenderMessage,
		SystemMessage: entity.SystemMessage,
	}
}
