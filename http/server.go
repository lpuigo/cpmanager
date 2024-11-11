package http

import (
	"context"
	"errors"
	"github.com/lpuig/cpmanager/config"
	"io"
	"log/slog"
	"net/http"
	"time"
)

type ServerOptions struct {
	Config *config.Config
	Log    *slog.Logger
}

type Server struct {
	config *config.Config
	log    *slog.Logger
	mux    *http.ServeMux
	server *http.Server
}

func NewServer(opts ServerOptions) *Server {
	if opts.Log == nil {
		opts.Log = slog.New(slog.NewTextHandler(io.Discard, nil))
	}

	mux := http.NewServeMux()

	return &Server{
		log:    opts.Log,
		config: opts.Config,
		mux:    mux,
		server: &http.Server{
			Addr:              ":8080",
			Handler:           mux,
			ReadTimeout:       5 * time.Second,
			ReadHeaderTimeout: 5 * time.Second,
			WriteTimeout:      5 * time.Second,
			IdleTimeout:       5 * time.Second,
		},
	}

}

// Start the server and set up routes.
func (s *Server) Start() error {
	s.log.Info("Starting http server", "address", s.server.Addr)

	s.setupRoutes()

	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop the server gracefully.
func (s *Server) Stop() error {
	s.log.Info("Stopping http server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	s.log.Info("Stopped http server")
	return nil
}
