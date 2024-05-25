package main

import (
	"fmt"
	"log"

	"event_manager/internal/ai"
	"event_manager/internal/app/telegram_bot"
	"event_manager/internal/app_log"
	"event_manager/internal/config"
	"event_manager/internal/database"
)

func main() {
	log.Print("parsing configuration")
	if err := config.Parse(); err != nil {
		log.Fatalf("failed to parse configuration: %v", err)
	}

	log.Print("set upping logger")
	app_log.Setup(config.Cfg().Logger.Level)

	log.Print("connecting to database")
	if err := database.Connect(); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Print("connecting to ai")
	if err := ai.Connect(); err != nil {
		log.Fatalf("failed to connect to ai: %v", err)
	}

	fmt.Println(config.Cfg().Port)

	log.Print("running telegram bot")
	if err := telegram_bot.Run(); err != nil {
		log.Fatalf("failed to start telegram bot: %v", err)
	}
}
