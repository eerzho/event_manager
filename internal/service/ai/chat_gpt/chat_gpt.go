package chat_gpt

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"event_manager/internal/app_log"
	"event_manager/internal/dto/command"
	"event_manager/internal/exception"
	"github.com/sashabaranov/go-openai"
)

type Service struct {
	client *openai.Client
	prompt string
}

func New(client *openai.Client, prompt string) *Service {
	return &Service{
		client: client,
		prompt: prompt,
	}
}

func (s *Service) Json(cmd command.AIJson) (string, error) {
	const op = "service.ai.chat_gpt.Json"

	log := app_log.Logger().With(
		slog.String("op", op),
		slog.Any("cmd", cmd),
	)

	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: s.prompt + " " + time.Now().Format("20060102T150405Z")},
		{Role: openai.ChatMessageRoleUser, Content: cmd.Prompt},
	}

	req := openai.ChatCompletionRequest{
		Model:    openai.GPT4o,
		Messages: messages,
	}

	log.Info("sending completion request")
	resp, err := s.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Error("failed to send completion request", slog.String("error", err.Error()))
		return "", err
	}

	if len(resp.Choices) > 0 {
		return strings.Replace(
			strings.Replace(
				resp.Choices[0].Message.Content,
				"json", "", -1,
			),
			"`", "", -1,
		), nil
	}

	return "", exception.ErrAIResponse
}
