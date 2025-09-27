package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/osamikoyo/ward/app"
	"github.com/osamikoyo/ward/config"
	"github.com/osamikoyo/ward/logger"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	loggerCfg := logger.Config{
		AppName:   "ward",
		AddCaller: false,
		LogFile:   "ward.log",
		LogLevel:  "debug",
	}

	logger.Init(loggerCfg)

	logger := logger.Get()

	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error("failed get config",
			zap.Error(err))

		return
	}

	app, err := app.ConnectApp(cfg, logger)
	if err != nil {
		logger.Error("failed connect app",
			zap.Error(err))

		return
	}

	if err = app.Run(ctx); err != nil {
		logger.Fatal("failed run app",
			zap.Error(err))
	}
}
