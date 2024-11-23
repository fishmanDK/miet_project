package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fishmanDK/miet_project/assets"
	"github.com/fishmanDK/miet_project/internal/config"
	"github.com/fishmanDK/miet_project/internal/handlers"
	"github.com/fishmanDK/miet_project/internal/service"
	"github.com/fishmanDK/miet_project/internal/storage"
	"github.com/fishmanDK/miet_project/pkg/logger"
	"github.com/gin-gonic/gin"

	c "github.com/fishmanDK/miet_project/internal/checker"
)

type app struct {
	log      logger.Logger
	cfg      *config.Config
	checker  *c.CheckerFirstReserveUsers
	handlers *handlers.Handlers
	service  *service.Service
	storage  *storage.Storage
	gin      *gin.Engine
	server *http.Server
}

func NewApp(cfg *config.Config, log logger.Logger) *app {
	return &app{cfg: cfg, log: log}
}

func (a *app) Run() {
	postgres, err := a.connectDB()
	if err != nil {
		a.log.Fatal(fmt.Sprintf("Failed connect postgres: %v", err))
	}

	a.checker = c.NewCheckerFirstReserveUsers(postgres, a.log)
	go a.checker.Start()

	a.storage = storage.NewStorage(postgres)
	a.service = service.NewSerivce(a.storage)

	tmpls := assets.NewTemplates(assets.Assets)

	a.handlers = handlers.NewHandlers(a.checker, a.service, tmpls, a.log)

	a.gin = a.handlers.InitRouts()

	a.server = a.createHttpServer()

	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			a.log.Fatal(fmt.Sprintf("Server error: %v"))
		}
	}()

	a.log.Info("server started")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	sig := <-stop
	fmt.Printf("Received signal: %v\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = a.checker.Stop()
	if err != nil{
		a.log.Info(err.Error())
	}

	if err := a.server.Shutdown(ctx); err != nil {
		a.log.Fatal(fmt.Sprintf("Server shutdown error: %v", err))
	}

	log.Println("Server gracefully stopped")
}
