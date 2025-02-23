package middleware

import (
	"log"
	"net/http"
	"time"
)

func WithLogging(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h(w, r)
		log.Println(r.RemoteAddr, r.Method, r.URL.Path, time.Since(start))
	})
}

type HttpHandleFuncErr func(http.ResponseWriter, *http.Request) error

func ErrorCatcher(h HttpHandleFuncErr) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		err := h(w, r)

		if err != nil {
			log.Println(r.RemoteAddr, r.Method, r.URL.Path, time.Since(start))
			http.Error(w, "Internal error", http.StatusInternalServerError)
		}
	})
}
