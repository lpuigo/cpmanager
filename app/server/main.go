package main

import (
	"context"
	"github.com/lpuig/cpmanager/config"
	"github.com/lpuig/cpmanager/http"
	"github.com/lpuig/cpmanager/log"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
)

func start(conf config.Config, log *log.Logger) error {
	log.InfoContext(nil, "Starting app")

	// We load environment variables from .env if it exists
	//_ = env.Load()

	// Catch signals to gracefully shut down the app
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	// Set up the database, which is injected as a dependency into the HTTP server
	// Here, the database is just a fake one.
	//db := sql.NewDatabase(sql.NewDatabaseOptions{
	//  Log: log,
	//})
	//if err := db.Connect(); err != nil {
	//  return err
	//}

	// Set up the HTTP server, injecting the database and logger
	s, err := http.NewServer(http.ServerOptions{
		Config: conf,
		Log:    log,
	})
	if err != nil {
		return err
	}

	// Use an errgroup to wait for separate goroutines which can error
	errGrp, ctx := errgroup.WithContext(ctx)

	// Start the server within the errgroup.
	// You can do this for other dependencies as well.
	errGrp.Go(func() error {
		return s.Start()
	})

	// Wait for the context to be done, which happens when a signal is caught
	<-ctx.Done()
	log.InfoContext(nil, "Stopping app")

	// Stop the server gracefully
	errGrp.Go(func() error {
		return s.Stop()
	})

	// Wait for the server to stop
	if err := errGrp.Wait(); err != nil {
		return err
	}

	log.InfoContext(nil, "Stopped app")

	return nil
}

func main() {
	// Set up a logger that is used throughout the app
	log := log.New()

	// set config
	conf := config.Set()

	// Start the app, exit with a non-zero exit code on errors
	if err := start(conf, log); err != nil {
		log.ErrorContext(nil, "Error starting app", "error", err)
		os.Exit(1)
	}
}
