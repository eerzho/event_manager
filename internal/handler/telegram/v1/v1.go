package v1

import (
	"event_manager/internal/service"
	"event_manager/pkg/logger"
	"gopkg.in/telebot.v3"
)

func NewHandler(l logger.Logger, bot *telebot.Bot, tgUserService *service.TGUser, tgMessageService *service.TGMessage) {
	newCommand(l, bot, tgUserService)
	newMessage(l, bot, tgMessageService)
}
