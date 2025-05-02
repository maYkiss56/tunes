package server

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/maYkiss56/tunes/internal/config"
)

type HTTPServer struct {
	server   *http.Server
	listener net.Listener
	cfg      *config.Config
}

func NewHTTPServer(cfg *config.Config) (*HTTPServer, error) {
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
	}, nil
}

func (s *HTTPServer) Start(serverErr chan<- error) {
	fmt.Printf("Starting server at http://%s\n", s.listener.Addr().String())
	if err := s.server.Serve(s.listener); err != nil && err != http.ErrServerClosed {
		serverErr <- fmt.Errorf("server error: %w", err)
	}
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	fmt.Println("Shutting down server...")
	return s.server.Shutdown(ctx)
}
