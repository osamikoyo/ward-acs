package server

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osamikoyo/ward/config"
	"github.com/osamikoyo/ward/logger"
	"go.uber.org/zap"
)

type (
	Routers interface {
		RegisterRouters(*echo.Echo)
	}
	Server struct {
		server *echo.Echo
		cfg    *config.Config
		logger *logger.Logger
	}
)

func (s *Server) NewServer(echo *echo.Echo, cfg *config.Config, logger *logger.Logger) *Server {
	return &Server{
		server: echo,
		cfg:    cfg,
		logger: logger,
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("starting server...",
		zap.String("addr", s.cfg.Addr))

	if err := s.server.Start(s.cfg.Addr); err != nil && err != http.ErrServerClosed {
		s.logger.Error("failed start server",
			zap.Error(err))

		return err
	}

	return nil
}

func (s *Server) SetRouters(routers Routers) {
	routers.RegisterRouters(s.server)
}

func (s *Server) Close(ctx context.Context) error {
	s.logger.Error("closing server...")

	if err := s.server.Shutdown(ctx); err != nil {
		s.logger.Error("failed shutdown server",
			zap.Error(err))

		return s.server.Close()
	}

	return nil
}
