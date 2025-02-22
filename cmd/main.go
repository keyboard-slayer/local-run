package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/keyboard-slayer/local-run/internal/postgres"
	"github.com/keyboard-slayer/local-run/internal/server/router"
)

func run(ctx context.Context, port uint16, expose bool, dbname string) error {
	notify, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	postgres.InitPool(dbname)

	err := godotenv.Load()
	if err != nil {
		return err
	}

	var bind string
	if expose {
		bind = fmt.Sprintf("0.0.0.0:%d", port)
	} else {
		bind = fmt.Sprintf("127.0.0.1:%d", port)
	}

	router := router.CreateRouter()

	server := http.Server{
		Addr:    bind,
		Handler: router,
	}

	slog.Info(fmt.Sprintf("Serving HTTP on http://%s", bind))

	errChan := make(chan error)
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	for {
		var err error
		select {
		case err = <-errChan:
			return err
		case <-notify.Done():
			shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
			defer shutdownRelease()

			slog.Info("Closing server...")

			if err := server.Shutdown(shutdownCtx); err != nil {
				return err
			}

			return nil
		}
	}
}

func main() {
	portPtr := flag.Int("port", 8080, "The port to use for the HTTP server")
	exposePtr := flag.Bool("expose", false, "Whatever to expose the service to the network or not")
	dbnamePtr := flag.String("dbname", "local_run", "The name of the database to use")
	flag.Parse()

	if err := run(context.Background(), uint16(*portPtr), *exposePtr, *dbnamePtr); err != nil {
		slog.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
}
