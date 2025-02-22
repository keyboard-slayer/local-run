package router

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/keyboard-slayer/local-run/api/types"
	"github.com/keyboard-slayer/local-run/internal/server/middleware"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	j, err := json.Marshal(types.NewPing())

	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't deliver ping to %s : %s", r.RemoteAddr, err))
		return
	}

	w.Write(j)
}

func auth(w http.ResponseWriter, r *http.Request) {
	var req types.AuthRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		slog.Error(fmt.Sprintf("Couldn't decode user authentication request: %s", err))
		return
	}

	slog.Info(fmt.Sprintf("Username: %s, Password: %s", req.Username, req.Password))

	http.Error(w, "Ok", http.StatusOK)
}

func CreateRouter() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/ping", middleware.DefaultMiddlewares(ping))
	r.HandleFunc("/auth", middleware.EnforceMethod(middleware.DefaultMiddlewares(auth), "POST"))

	return r
}
