package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Port           string `env:"PORT" env-default:"8080"`
	Mongo          Mongo
	Logger         Logger
	GPT            GPT
	Telegram       Telegram
	GoogleCalendar GoogleCalendar
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

type Mongo struct {
	URL string `env:"MONGO_URL"`
	DB  string `env:"MONGO_DB"`
}

var cfg Config

func Parse() error {
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		return err
	}

	return nil
}

func Cfg() *Config {
	return &cfg
}
