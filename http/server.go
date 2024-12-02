package http

import (
	"context"
	"errors"
	"github.com/lpuig/cpmanager/config"
	"github.com/lpuig/cpmanager/http/session"
	"github.com/lpuig/cpmanager/log"
	"github.com/lpuig/cpmanager/model/manager"
	"net/http"
	"time"
)

type ServerOptions struct {
	Config *config.Config
	Log    *log.Logger
}

type Server struct {
	config *config.Config
	mux    *http.ServeMux
	server *http.Server

	Log      *log.Logger
	Sessions *session.Sessions
	Manager  *manager.Manager
}

func NewServer(opts ServerOptions) *Server {
	if opts.Log == nil {
		opts.Log = log.New()
	}

	mgr := manager.New(opts.Log)

	mux := http.NewServeMux()

	return &Server{
		Log:    opts.Log,
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

		Sessions: session.New(),
		Manager:  mgr,
	}

}

// Start the server and set up routes.
func (s *Server) Start() error {
	s.Log.Info("Starting manager")
	err := s.Manager.Init()
	if err != nil {
		return err
	}

	s.Log.Info("Starting http server", "address", s.server.Addr)

	s.setupRoutes()

	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop the server gracefully.
func (s *Server) Stop() error {
	s.Log.Info("Stopping http server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}
	s.Log.Info("Stopped http server")
	return nil
}
