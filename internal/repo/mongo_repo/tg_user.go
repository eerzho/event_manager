package mongo_repo

import (
	"context"
	"errors"
	"fmt"

	"event_manager/internal/failure"
	"event_manager/internal/model"
	"event_manager/pkg/mongo"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const TgUserTable = "tg_users"

type TGUser struct {
	*mongo.Mongo
}

func NewTGUser(m *mongo.Mongo) *TGUser {
	return &TGUser{m}
}

func (t *TGUser) All(ctx context.Context) ([]model.TGUser, error) {
	const op = "./internal/repo/mongo_repo/tg_user::All"

	var users []model.TGUser
	filter := bson.D{}
	opts := options.Find()

	cursor, err := t.DB.Collection(TgUserTable).Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.TGUser
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}

func (t *TGUser) ByChatID(ctx context.Context, chatID string) (*model.TGUser, error) {
	const op = "./internal/repo/mongo_repo/tg_user::ByChatID"

	var user model.TGUser

	filter := bson.D{{"chat_id", chatID}}

	err := t.DB.Collection(TgUserTable).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongoDriver.ErrNoDocuments) {
			return nil, fmt.Errorf("%s: %w", op, failure.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &user, nil
}

func (t *TGUser) Create(ctx context.Context, user *model.TGUser) error {
	const op = "./internal/repo/mongo_repo/tg_user::Create"

	user.ID = uuid.New().String()

	result, err := t.DB.Collection(TgUserTable).InsertOne(ctx, user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if _, ok := result.InsertedID.(primitive.ObjectID); !ok {
		return fmt.Errorf("%s: document is nil", op)
	}

	return nil
}
