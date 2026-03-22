package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware logs HTTP requests with method, path, status code, and duration.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Create a response writer to capture the status code
		wr := newResponseWriter(w)

		next.ServeHTTP(wr, r)

		duration := time.Since(start)
		log.Printf("%s %s %d %s", r.Method, r.URL.Path, wr.status, duration)
	})
}

// responseWriter wraps http.ResponseWriter to capture status code
type responseWriter struct {
	http.ResponseWriter
	status int
}

func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w, status: http.StatusOK}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	 rw.ResponseWriter.WriteHeader(code)
}