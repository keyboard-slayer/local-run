package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/keyboard-slayer/local-run/configs"
	"github.com/keyboard-slayer/local-run/internal/postgres"
	"github.com/keyboard-slayer/local-run/internal/server/router"
)

func run(ctx context.Context) error {
	notify, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	cfg, err := configs.LoadConfig()
	if err != nil {
		return err
	}

	pool, err := postgres.InitPool(cfg.Db)
	if err != nil {
		return err
	}

	var bind string
	if cfg.Http.Expose {
		bind = fmt.Sprintf("0.0.0.0:%d", cfg.Http.Port)
	} else {
		bind = fmt.Sprintf("127.0.0.1:%d", cfg.Http.Port)
	}

	app := &router.App{Pool: pool, Key: []byte(cfg.Security.JwtSecret)}
	router := router.CreateRouter(app)

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

			app.Pool.Close()

			if err := server.Shutdown(shutdownCtx); err != nil {
				return err
			}

			return nil
		}
	}
}

func main() {
	if err := run(context.Background()); err != nil {
		slog.Error(fmt.Sprintf("%s", err))
		os.Exit(1)
	}
}
