package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &loggingResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)
		if wrapped.statusCode <= 400 {
			slog.Info(
				"Request Completed",
				"method", r.Method,
				"url", r.URL.String(),
				"status", wrapped.statusCode,
				"duration", time.Since(start),
			)
		} else {
			slog.Info(
				"Request Failed",
				"method", r.Method,
				"url", r.URL.String(),
				"status", wrapped.statusCode,
				"duration", time.Since(start),
			)
		}

	})
}
