package model

type UserRequest struct {
	ID            string `json:"id"`
	SenderID      int64  `json:"sender_id"`
	SenderMessage string `json:"sender_message"`
	SystemMessage string `json:"system_message"`
}
