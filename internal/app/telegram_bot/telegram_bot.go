package telegram_bot

import (
	"log/slog"

	"event_manager/internal/ai"
	"event_manager/internal/app_log"
	"event_manager/internal/config"
	commandH "event_manager/internal/handler/telegram/command"
	textH "event_manager/internal/handler/telegram/text"
	aiS "event_manager/internal/service/ai/chat_gpt"
	calendarS "event_manager/internal/service/calendar/google"
	textS "event_manager/internal/service/text"
	"gopkg.in/telebot.v3"
)

func Run() error {
	const op = "app.telegram_bot.Run"

	log := app_log.Logger().With(slog.String("op", op))

	log.Info("configuring telegram bot")
	bot, err := telebot.NewBot(telebot.Settings{
		Token: config.Cfg().Telegram.Token,
		Poller: &telebot.Webhook{
			Listen: "0.0.0.0:8080",
			Endpoint: &telebot.WebhookEndpoint{
				PublicURL: config.Cfg().Telegram.WebhookUrl,
			},
		},
	})
	if err != nil {
		log.Error("failed to create telegram bot", slog.String("error", err.Error()))
		return err
	}

	setupHandlers(bot)

	log.Info("starting telegram bot")
	bot.Start()

	return nil
}

func setupHandlers(bot *telebot.Bot) {

	calendarService := calendarS.New(config.Cfg().GoogleCalendar.Url)
	aiService := aiS.New(ai.Client().GPT, config.Cfg().GPT.Prompt)
	textService := textS.New(aiService, calendarService)

	textHandler := textH.New(textService)
	commandHandler := commandH.New()

	bot.Handle("/start", commandHandler.Start)
	bot.Handle(telebot.OnText, textHandler.Process)
}
