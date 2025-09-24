package app

import (
	"context"
	"os/user"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/ward/config"
	"github.com/osamikoyo/ward/core"
	"github.com/osamikoyo/ward/entity/data"
	"github.com/osamikoyo/ward/entity/grand"
	"github.com/osamikoyo/ward/handler"
	"github.com/osamikoyo/ward/logger"
	"github.com/osamikoyo/ward/repository"
	"github.com/osamikoyo/ward/retrier"
	"github.com/osamikoyo/ward/searchbase"
	"github.com/osamikoyo/ward/server"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	server *server.Server
	logger *logger.Logger
}

func ConnectApp(cfg *config.Config, logger *logger.Logger) (*App, error) {
	logger.Info("starting connect app",
		zap.Any("config", cfg))

	db, err := retrier.Connect(7, func() (*gorm.DB, error) {
		db, err := gorm.Open(postgres.Open(cfg.DSN))
		if err != nil {
			return nil, err
		}

		return db, db.AutoMigrate(&grand.Grand{}, &user.User{}, &data.Data{})
	})
	if err != nil {
		logger.Error("failed connect to db",
			zap.String("dsn", cfg.DSN),
			zap.Error(err))

		return nil, err
	}

	logger.Info("successfully connected to database")
	logger.Info("connecting to elasticsearch",
		zap.String("addr", cfg.ElasticSearhUrl))

	elasticsearchConn, err := retrier.Connect(7, func() (*elasticsearch.Client, error) {
		return elasticsearch.NewClient(elasticsearch.Config{
			Addresses: []string{
				cfg.ElasticSearhUrl,
			},
		})
	})
	if err != nil {
		logger.Error("failed connect to elasticsearch",
			zap.String("addr", cfg.ElasticSearhUrl),
			zap.Error(err))

		return nil, err
	}

	repository := repository.NewRepository(db, logger)
	searchbase := searchbase.NewSearchBase(elasticsearchConn, logger)

	core := core.NewWardCore(repository, searchbase, logger, cfg, 10*time.Second)
	handler := handler.NewHandler(core)

	e := echo.New()
	server := server.NewServer(e, cfg, logger)

	server.SetRouters(handler)

	return &App{
		server: server,
		logger: logger,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	a.logger.Info("starting app...")

	go func() {
		<-ctx.Done()
		a.server.Close(ctx)
	}()

	errChan := make(chan error, 1)
	go func() {
		if err := a.server.Start(ctx); err != nil {
			a.logger.Error("failed start app",
				zap.Error(err))
			errChan <- err

		}
	}()

	return <-errChan
}
