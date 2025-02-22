package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
)

func EnforceMethod(h http.HandlerFunc, m string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != m && r.Method != "OPTIONS" {
			slog.Error(fmt.Sprintf("%s tried to send a %s request (only allowing %s)", r.RemoteAddr, r.Method, m))
			http.Error(w, "Bad request", http.StatusBadRequest)
		}

		h(w, r)
	})
}

func DefaultMiddlewares(h http.HandlerFunc) http.HandlerFunc {
	return WithLogging(SetHeaders(h))
}
