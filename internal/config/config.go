package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Logger         Logger
	GPT            GPT
	GoogleCalendar GoogleCalendar
	Telegram       Telegram
}

type Logger struct {
	Level string `env:"LOG_LEVEL"`
}

type GPT struct {
	Token  string `env:"GPT_TOKEN"`
	Prompt string `env:"GPT_PROMPT"`
}

type GoogleCalendar struct {
	Url string `env:"GOOGLE_CALENDAR_URL"`
}

type Telegram struct {
	Token      string `env:"TELEGRAM_TOKEN"`
	WebhookUrl string `env:"TELEGRAM_WEBHOOK_URL"`
}

var cfg Config

func Parse() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return err
	}

	return nil
}

func Cfg() *Config {
	return &cfg
}
