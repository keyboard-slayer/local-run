package router

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
	var id uint

	if err := app.Pool.QueryRow(ctx, "SELECT id, password FROM users WHERE username=$1", req.Username).Scan(&id, &hash); err != nil {
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

	jwtId, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	claims := types.AuthClaims{
		Id: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        jwtId.String(),
			Subject:   req.Username,
			IssuedAt:  &jwt.NumericDate{time.Now()},
			ExpiresAt: &jwt.NumericDate{time.Now().Add(24 * time.Hour)},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(app.Key)
	if err != nil {
		return err
	}

	return WriteJson(w, types.AuthResponse{Status: "ok", Token: tokenStr})
}
