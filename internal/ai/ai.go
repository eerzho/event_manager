package ai

import (
	"event_manager/internal/config"
	"github.com/sashabaranov/go-openai"
)

type AI struct {
	GPT *openai.Client
}

var ai AI

func Connect() error {
	ai.GPT = openai.NewClient(config.Cfg().GPT.Token)

	return nil
}

func Client() *AI {
	return &ai
}
