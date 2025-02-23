package postgres

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/keyboard-slayer/local-run/internal/schemas"
)

func createDb(ctx context.Context, pool *pgxpool.Pool) {
	pool.Exec(ctx, schemas.UserSchema)
}

func InitPool(dbname string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, os.Getenv("POSTGRES_URI"))

	if err != nil {
		return err
	}

	defer pool.Close()

	var exists bool
	err = pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname=$1)", dbname).Scan(&exists)

	if !exists {
		slog.Info("Database doesn't exist, creating...")
		pool.Exec(ctx, "CREATE DATABASE $1", dbname)
	}

	createDb(ctx, pool)

	return nil
}
