package middleware

import "net/http"

func SetHeaders(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := w.Header()

		// TODO: remove me
		headers.Set("Access-Control-Allow-Origin", "*")

		headers.Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		headers.Set("Access-Control-Allow-Credentials", "true")
		headers.Set("Access-Control-Allow-Headers", "Content-Type, X-LOCALRUN-CSRF-PROTECTION")
		headers.Set("Access-Control-Max-Age", "86400")
		headers.Set("X-Frame-Options", "DENY")
		headers.Set("Content-Security-Policy", "fault-src 'self'; script-src 'self'; object-src 'none'; frame-ancestors 'none'; base-uri 'self'; block-all-mixed-content; upgrade-insecure-requests; img-src 'self'; style-src 'self'; font-src 'self';")
		headers.Set("X-Content-Type-Options", "nosniff")
		headers.Set("Referrer-Policy", "strict-origin-when-cross-origin")
		headers.Set("Cross-Origin-Resource-Policy", "same-site")
		headers.Set("Permissions-Policy", "interest-cohort=()")

		h(w, r)
	})
}
