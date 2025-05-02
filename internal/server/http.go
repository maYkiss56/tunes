package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/maYkiss56/tunes/internal/config"
	"github.com/maYkiss56/tunes/internal/logger"
)

type HTTPServer struct {
	server   *http.Server
	listener net.Listener
	cfg      *config.Config
	logger   *logger.Logger
}

func NewHTTPServer(cfg *config.Config, logger *logger.Logger) (*HTTPServer, error) {
	addr := net.JoinHostPort(cfg.HTTP.Host, cfg.HTTP.Port)

	listener, err := net.Listen(cfg.HTTP.Network, addr)
	if err != nil {
		return nil, fmt.Errorf("listen failed: %v", err)
	}

	//TODO: replaced nil on router
	server := &http.Server{
		Addr:         addr,
		Handler:      nil,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
	}

	return &HTTPServer{
		server:   server,
		listener: listener,
		cfg:      cfg,
		logger:   logger,
	}, nil
}

func (s *HTTPServer) Start(serverErr chan<- error) {
	s.logger.Info(
		"Starting HTTP server",
		"address",
		s.listener.Addr().String(),
		"timeout",
		s.cfg.HTTP.ReadTimeout,
	)
	if err := s.server.Serve(s.listener); err != nil && err != http.ErrServerClosed {
		serverErr <- fmt.Errorf("server error: %w", err)
	}
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	s.logger.Info("Shutting down HTTP server")
	return s.server.Shutdown(ctx)
}
