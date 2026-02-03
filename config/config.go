package config

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type AppConfig struct {
	Host string `env:"HOST, default="localhost"`
	Port string `env:"PORT, default=8080"`
}

func InitConfig() *AppConfig {
	// Загружаем переменные из .env файла
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	ctx := context.Background()
	cfg := new(AppConfig)

	if err := envconfig.Process(ctx, cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
