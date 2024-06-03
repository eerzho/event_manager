package telegram

import (
	"fmt"
	"log"
	"strings"

	"event_manager/config"
	v1 "event_manager/internal/handler/telegram/v1"
	"event_manager/internal/service"
	"event_manager/pkg/logger"
	"gopkg.in/telebot.v3"
)

type Bot struct {
	url string
	bot *telebot.Bot
}

func New(l logger.Logger, cfg *config.Config, tgUserService *service.TGUser, tgMessageService *service.TGMessage) (*Bot, error) {
	url := fmt.Sprintf("%s/wb", strings.Trim(cfg.Domain, "/"))
	settings := telebot.Settings{
		Token: cfg.Telegram.Token,
		Poller: &telebot.Webhook{
			Listen: "0.0.0.0:" + cfg.Telegram.Port,
			Endpoint: &telebot.WebhookEndpoint{
				PublicURL: url,
			},
		},
	}

	bot, err := telebot.NewBot(settings)
	if err != nil {
		return nil, fmt.Errorf("./internal/app/bot::New: %w", err)
	}

	v1.NewHandler(l, bot, tgUserService, tgMessageService)

	return &Bot{
		bot: bot,
		url: url,
	}, nil
}

func (t *Bot) Run() {
	const op = "./internal/app/telegram/bot::Run"

	log.Printf("%s: telegram bot listening at %s", op, t.url)
	t.bot.Start()
}

func (t *Bot) Shutdown() {
	const op = "./internal/app/telegram/bot::Shutdown"

	log.Printf("%s: telegram bot shutting down", op)
	err := t.bot.RemoveWebhook()
	if err != nil {
		log.Printf("%s: %v", op, err)
	}
}