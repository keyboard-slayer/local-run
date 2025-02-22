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
