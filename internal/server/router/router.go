package router

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/keyboard-slayer/local-run/api/types"
	"github.com/keyboard-slayer/local-run/internal/server/middleware"
)

type App struct {
	Pool *pgxpool.Pool
}

func (app *App) ping(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("content-type", "application/json")
	return WriteJson(w, types.PingSchema{Response: "ping"})
}

func CreateRouter(app *App) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/ping", middleware.DefaultMiddlewares(app.ping))
	r.HandleFunc("/auth", middleware.EnforceMethod(middleware.DefaultMiddlewares(app.auth), "POST"))

	return r
}
