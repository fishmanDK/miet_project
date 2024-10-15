package config

import (
	"log"
	"os"
	"time"

	"github.com/fishmanDK/miet_project/internal/storage"
	"github.com/fishmanDK/miet_project/pkg/logger"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Logger      logger.Config  `yaml:"logger"`
		HTTP        HTTP           `yaml:"http"`  //TODO: возможно имеет смысл перенести структуры как в Logger, Mongo
		Postgres storage.Config 	`yaml:"postgres"`
}

	HTTP struct {
		Port         string        `yaml:"port"`
		ReadTimeout  time.Duration `yaml:"readTimeout"`
		WriteTimeout time.Duration `yaml:"writeTimeout"`
		// MaxHeaderBytes //TODO
	}
)

func InitConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("%w", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg, nil
}
