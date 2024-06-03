package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type TGUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	ChatID   string             `bson:"chat_id"`
	Username string             `bson:"username"`
}
