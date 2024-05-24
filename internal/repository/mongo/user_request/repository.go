package user_request

import (
	"context"

	"event_manager/internal/dto/model"
	"event_manager/internal/dto/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Database
}

func New(db *mongo.Database) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(qry query.UserRequestCreate) (*model.UserRequest, error) {
	entity := UserRequest{
		ID:            primitive.NewObjectID(),
		SenderID:      qry.SenderID,
		SenderMessage: qry.SenderMessage,
		SystemMessage: qry.SystemMessage,
	}
	result, err := r.db.Collection("user_requests").InsertOne(context.Background(), entity)
	if err != nil {
		return nil, err
	}

	if _, ok := result.InsertedID.(primitive.ObjectID); !ok {
		return nil, mongo.ErrNilDocument
	}

	return toModel(&entity), nil
}
