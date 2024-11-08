package main

import (
	"log"

	"github.com/fishmanDK/miet_project/internal/app"
	"github.com/fishmanDK/miet_project/internal/config"
	"github.com/fishmanDK/miet_project/pkg/logger"
)

const (
	envLocal = "local"
	envDev   = "dev"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil{
		log.Fatal(err)
	}

	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	app := app.NewApp(cfg, appLogger)
	app.Run()
}