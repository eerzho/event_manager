package service

import (
	"context"
	"fmt"

	"event_manager/internal/model"
)

type (
	TGUserRepo interface {
		All(ctx context.Context) ([]model.TGUser, error)
		ByChatID(ctx context.Context, chatID string) (*model.TGUser, error)
		Create(ctx context.Context, user *model.TGUser) error
	}

	TGUser struct {
		repo TGUserRepo
	}
)

func NewTGUser(repo TGUserRepo) *TGUser {
	return &TGUser{repo: repo}
}

func (t *TGUser) All(ctx context.Context) ([]model.TGUser, error) {
	const op = "./internal/service/tg_user::All"

	users, err := t.repo.All(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}

func (t *TGUser) ByChatID(ctx context.Context, chatID string) (*model.TGUser, error) {
	const op = "./internal/service/tg_user::ByChatID"

	user, err := t.repo.ByChatID(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (t *TGUser) Create(ctx context.Context, user *model.TGUser) error {
	const op = "./internal/service/tg_user::Create"

	err := t.repo.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
