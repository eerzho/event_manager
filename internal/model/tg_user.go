package model

type TGUser struct {
	ID       string `bson:"_id,omitempty"`
	ChatID   string `bson:"chat_id"`
	Username string `bson:"username"`
}
