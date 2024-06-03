package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type TGMessage struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	ChatID string             `bson:"chat_id"`
	Text   string             `bson:"text"`
	Url    string             `bson:"url,omitempty"`
	File   string             `bson:"file,omitempty"`
}
