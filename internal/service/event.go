package service

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"event_manager/internal/model"
	"github.com/go-playground/validator/v10"
	"github.com/sashabaranov/go-openai"
)

type Event struct {
	openai *openai.Client
	prompt string
}

func NewEvent(token, prompt string) *Event {
	return &Event{
		openai: openai.NewClient(token),
		prompt: prompt,
	}
}

func (e *Event) CreateFromText(ctx context.Context, text string) (*model.Event, error) {
	const op = "./internal/service/event::CreateFromText"

	messages := []openai.ChatCompletionMessage{
		{Role: openai.ChatMessageRoleSystem, Content: e.prompt + time.Now().Format("20060102T150405Z")},
		{Role: openai.ChatMessageRoleUser, Content: text},
	}
	req := openai.ChatCompletionRequest{
		Model:    openai.GPT4o,
		Messages: messages,
	}
	resp, err := e.openai.CreateChatCompletion(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if len(resp.Choices) < 1 {
		return nil, fmt.Errorf("%s: choices is empty", op)
	}

	choice := resp.Choices[0]
	jsonString := strings.Replace(strings.Replace(choice.Message.Content, "json", "", -1), "`", "", -1)

	var event model.Event
	if err = json.Unmarshal([]byte(jsonString), &event); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	if err = validator.New().Struct(event); err != nil {
		message := "AI не понимает вас, сформулируйте запрос более точно."
		if event.Message != "" {
			message = event.Message
		}
		return nil, fmt.Errorf("%s: %s", op, message)
	}

	return &event, nil
}
