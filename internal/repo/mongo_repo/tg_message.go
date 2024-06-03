package mongo_repo

import (
	"context"
	"fmt"

	"github.com/eerzho/event_manager/internal/model"
	"github.com/eerzho/event_manager/pkg/mongo"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const TgMessageTable = "tg_messages"

type TGMessage struct {
	*mongo.Mongo
}

func NewTGMessage(m *mongo.Mongo) *TGMessage {
	return &TGMessage{m}
}

func (t *TGMessage) All(ctx context.Context) ([]model.TGMessage, error) {
	const op = "./internal/repo/mongo_repo/tg_user::All"

	var messages []model.TGMessage
	filter := bson.D{}
	opts := options.Find()

	cursor, err := t.DB.Collection(TgMessageTable).Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var message model.TGMessage
		if err := cursor.Decode(&message); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		messages = append(messages, message)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return messages, nil
}

func (t *TGMessage) Create(ctx context.Context, message *model.TGMessage) error {
	const op = "./internal/repo/mongo_repo/tg_message::Create"

	message.ID = uuid.New().String()

	result, err := t.DB.Collection(TgMessageTable).InsertOne(ctx, message)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, ok := result.InsertedID.(string); !ok {
		return fmt.Errorf("%s: document is nil", op)
	}

	return nil
}
