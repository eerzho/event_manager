package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"event_manager/config"
	"event_manager/internal/app/http"
	"event_manager/internal/app/telegram"
	"event_manager/internal/repo/mongo_repo"
	"event_manager/internal/service"
	"event_manager/pkg/logger"
	"event_manager/pkg/mongo"
)

func main() {
	const op = "./cmd/app::main"

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}

	mg, err := mongo.New(cfg.Mongo.URL, cfg.Mongo.DB)
	if err != nil {
		log.Fatalf("%s: %v", op, err)
	}

	// repo
	tgUserRepo := mongo_repo.NewTGUser(mg)
	tgMessageRepo := mongo_repo.NewTGMessage(mg)

	// service
	tgUserService := service.NewTGUser(tgUserRepo)
	eventService := service.NewEvent(cfg.GPT.Token, cfg.GPT.Prompt)
	googleCalendarService := service.NewGoogleCalendar(cfg.Google.CalendarURL)
	tgMessageService := service.NewTGMessage(tgMessageRepo, tgUserService, eventService, googleCalendarService)

	l := logger.New(cfg.Level)
	httpServer := http.New(l, cfg)
	telegramBot, err := telegram.New(l, cfg, tgUserService, tgMessageService)
	if err != nil {
		log.Fatalf("%s: %s", op, err)
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		httpServer.Run()
	}()
	go func() {
		telegramBot.Run()
	}()

	log.Printf("%s: application started", op)
	<-stopChan

	log.Printf("%s: shutting down", op)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	httpServer.Shutdown(ctx)
	telegramBot.Shutdown()

	log.Printf("%s: application stopped", op)
}
