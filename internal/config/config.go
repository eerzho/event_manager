package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
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
	DB       string `env:"MONGO_DB"`
	Host     string `env:"MONGO_HOST"`
	Port     string `env:"MONGO_PORT"`
	UIPort   string `env:"MONGO_UI_PORT"`
	User     string `env:"MONGO_USER"`
	Password string `env:"MONGO_PASSWORD"`
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
