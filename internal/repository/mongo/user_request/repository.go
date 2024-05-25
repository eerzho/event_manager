package user_request

import (
	"context"

	"event_manager/internal/dto/model"
	"event_manager/internal/dto/query"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const TABLE = "user_requests"

type Repository struct {
	col *mongo.Collection
}

func New(db *mongo.Database) *Repository {
	return &Repository{
		col: db.Collection(TABLE),
	}
}

func (r *Repository) Create(qry query.UserRequestCreate) (*model.UserRequest, error) {
	entity := UserRequest{
		ID:            primitive.NewObjectID(),
		SenderID:      qry.SenderID,
		SenderMessage: qry.SenderMessage,
		SystemMessage: qry.SystemMessage,
	}
	result, err := r.col.InsertOne(context.Background(), entity)
	if err != nil {
		return nil, err
	}

	if _, ok := result.InsertedID.(primitive.ObjectID); !ok {
		return nil, mongo.ErrNilDocument
	}

	return toModel(&entity), nil
}
