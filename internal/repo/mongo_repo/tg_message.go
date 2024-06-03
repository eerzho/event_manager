package mongo_repo

import (
	"context"
	"fmt"

	"event_manager/internal/model"
	"event_manager/pkg/mongo"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const TgMessageTable = "tg_messages"

type TGMessage struct {
	*mongo.Mongo
}

func NewTGMessage(m *mongo.Mongo) *TGMessage {
	return &TGMessage{m}
}

func (t *TGMessage) Create(ctx context.Context, message *model.TGMessage) error {
	const op = "./internal/repo/mongo_repo/tg_message::Create"

	message.ID = uuid.New().String()

	result, err := t.DB.Collection(TgMessageTable).InsertOne(ctx, message)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, ok := result.InsertedID.(primitive.ObjectID); !ok {
		return fmt.Errorf("%s: document is nil", op)
	}

	return nil
}
