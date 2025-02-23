package router

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/jackc/pgx/v5"
	"github.com/keyboard-slayer/local-run/api/types"
)

func (app *App) auth(w http.ResponseWriter, r *http.Request) error {
	var req types.AuthRequest

	w.Header().Set("content-type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var hash string
	if err := app.Pool.QueryRow(ctx, "SELECT password FROM users WHERE username=$1", req.Username).Scan(&hash); err != nil {
		if err == pgx.ErrNoRows {
			return WriteJson(w, types.AuthResponse{Status: "error", Msg: "Invalid username or password"})
		}

		return err
	}

	match, err := argon2id.ComparePasswordAndHash(req.Password, hash)
	if err != nil {
		return err
	}

	if !match {
		return WriteJson(w, types.AuthResponse{Status: "error", Msg: "Invalid username or password"})
	}

	return WriteJson(w, types.AuthResponse{Status: "ok"})

}
