package postgres

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/keyboard-slayer/local-run/configs"
	"github.com/keyboard-slayer/local-run/internal/schemas"
)

func createDb(ctx context.Context, pool *pgxpool.Pool) error {
	_, err := pool.Exec(ctx, schemas.UserSchema)
	if err != nil {
		return err
	}

	hash, err := argon2id.CreateHash("admin", argon2id.DefaultParams)
	if err != nil {
		return err
	}

	_, err = pool.Exec(ctx, "INSERT INTO users (username, password) VALUES ($1, $2)", "admin", hash)
	if err != nil {
		return err
	}

	return nil
}

func InitPool(db configs.Dbcfg) (*pgxpool.Pool, error) {
	ctxconn, cancelconn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelconn()

	uri := fmt.Sprintf("postgres://%s:%s@%s:%d", db.Username, db.Password, db.Host, db.Port)

	conn, err := pgx.Connect(ctxconn, fmt.Sprintf("%s/postgres", uri))

	if err != nil {
		return nil, err
	}

	defer conn.Close(ctxconn)

	var exists bool
	err = conn.QueryRow(ctxconn, "SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname=$1)", db.Dbname).Scan(&exists)
	if err != nil {
		return nil, err
	}

	if !exists {
		slog.Info(fmt.Sprintf("%s doesn't exists, creating...", db.Dbname))

		// NOTE: Since it's defined in a config file, it should be fine.
		_, err := conn.Exec(ctxconn, fmt.Sprintf("CREATE DATABASE %s", db.Dbname))

		if err != nil {
			return nil, err
		}
	}

	ctxpool, cancelpool := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelpool()

	pool, err := pgxpool.New(ctxpool, fmt.Sprintf("%s/%s", uri, db.Dbname))

	if !exists {
		err = createDb(ctxpool, pool)
		if err != nil {
			return nil, err
		}
	}

	return pool, nil
}
