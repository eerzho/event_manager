package model

type TGMessage struct {
	ID     string `bson:"_id,omitempty"`
	ChatID string `bson:"chat_id"`
	Text   string `bson:"text"`
	Answer string `bson:"answer,omitempty"`
	File   string `bson:"file,omitempty"`
}
