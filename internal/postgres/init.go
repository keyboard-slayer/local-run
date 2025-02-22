package postgres

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func createDb(pool *pgxpool.Pool) {
	// pool.Exec(schemas.UserSchema)
}

func InitPool(dbname string) error {
	ctx := context.Background()
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

	createDb(pool)

	return nil
}
