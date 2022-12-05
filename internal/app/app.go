package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ilyasbulat/rest_api/internal/config"
	v1 "github.com/ilyasbulat/rest_api/internal/controller/http/v1"
	"github.com/ilyasbulat/rest_api/internal/usecase"
	"github.com/ilyasbulat/rest_api/internal/usecase/repo"
	"github.com/ilyasbulat/rest_api/pkg/httpserver"
	"github.com/ilyasbulat/rest_api/pkg/logger"
	"github.com/ilyasbulat/rest_api/pkg/postgres"
	"github.com/julienschmidt/httprouter"
)

func Start(cfg *config.Config) {
	l := logger.GetLogger(cfg.Log)
	defer l.Sync()
	// Repository
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
	)
	pg, err := postgres.New(dsn)
	if err != nil {
		l.Fatal(fmt.Sprintf("app - Run - postgres.New: %v", err))
	}
	defer pg.Close()

	// Use case
	translationUseCase := usecase.New(
		repo.New(pg),
		// webapi.New(),
	)

	// RabbitMQ RPC Server
	// rmqRouter := amqprpc.NewRouter(translationUseCase)

	// rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	// if err != nil {
	// 	l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	// }

	// HTTP Server
	router := httprouter.New()

	// Register routes
	v1.NewRouter(router, l, translationUseCase)

	httpServer := httpserver.New(router, cfg, httpserver.Port(cfg.App.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Sprintf("app - Run - httpServer.Notify: %v", err))
		// case err = <-rmqServer.Notify():
		// 	l.Error(fmt.Sprintf("app - Run - rmqServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Sprintf("app - Run - httpServer.Shutdown: %v", err))
	}

	// err = rmqServer.Shutdown()
	// if err != nil {
	// 	l.Error(fmt.Sprintf("app - Run - rmqServer.Shutdown: %v", err))
	// }
}
