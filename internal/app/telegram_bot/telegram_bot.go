package telegram_bot

import (
	"log/slog"

	"event_manager/internal/ai"
	"event_manager/internal/app_log"
	"event_manager/internal/config"
	"event_manager/internal/database"
	commandH "event_manager/internal/handler/telegram/command"
	textH "event_manager/internal/handler/telegram/text"
	"event_manager/internal/repository/mongo/user_request"
	aiS "event_manager/internal/service/ai/chat_gpt"
	calendarS "event_manager/internal/service/calendar/google"
	textS "event_manager/internal/service/text"
	userRequestS "event_manager/internal/service/user_request"
	"gopkg.in/telebot.v3"
)

func Run() error {
	const op = "app.telegram_bot.Run"

	log := app_log.Logger().With(slog.String("op", op))

	log.Info("configuring telegram bot")
	bot, err := telebot.NewBot(telebot.Settings{
		Token: config.Cfg().Telegram.Token,
		Poller: &telebot.Webhook{
			Listen: "0.0.0.0:60",
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
	userRequestRepository := user_request.New(database.Db().Mongo)

	aiService := aiS.New(ai.Client().GPT, config.Cfg().GPT.Prompt)
	calendarService := calendarS.New(config.Cfg().GoogleCalendar.Url)
	userRequestService := userRequestS.New(userRequestRepository)
	textService := textS.New(aiService, calendarService, userRequestService)

	commandHandler := commandH.New()
	textHandler := textH.New(textService)

	bot.Handle("/start", commandHandler.Start)
	bot.Handle(telebot.OnText, textHandler.Process)
}
