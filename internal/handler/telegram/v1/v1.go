package v1

import (
	"sync"

	"github.com/eerzho/event_manager/internal/service"
	"github.com/eerzho/event_manager/pkg/logger"
	"gopkg.in/telebot.v3"
)

func NewHandler(l logger.Logger, bot *telebot.Bot, tgUserService *service.TGUser, tgMessageService *service.TGMessage) {
	// middleware
	mv := newMiddleware(l)

	// handler
	newCommand(l, bot, tgUserService)
	newMessage(l, mv, bot, tgMessageService, tgUserService)
}

type middleware struct {
	l             logger.Logger
	mu            sync.Mutex
	activeRequest map[int64]struct{}
}

func newMiddleware(l logger.Logger) *middleware {
	return &middleware{
		l:             l,
		activeRequest: make(map[int64]struct{}),
	}
}

func (m *middleware) limiter(next telebot.HandlerFunc) telebot.HandlerFunc {
	return func(ctx telebot.Context) error {
		userId := ctx.Message().Chat.ID

		m.mu.Lock()
		if _, ok := m.activeRequest[userId]; ok {
			m.mu.Unlock()
			options := &telebot.SendOptions{ReplyTo: ctx.Message()}
			return ctx.Send("Вы отправляете сообщения слишком часто. Пожалуйста, подождите.", options)
		}
		m.activeRequest[userId] = struct{}{}
		m.mu.Unlock()

		defer func() {
			m.mu.Lock()
			delete(m.activeRequest, userId)
			m.mu.Unlock()
		}()

		return next(ctx)
	}
}
