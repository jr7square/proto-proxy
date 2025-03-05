package httpserver

import (
	"log/slog"
	"net/http"
	"time"
)

func NewServer(logger *slog.Logger) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux)
	var handler http.Handler = mux
	handler = middleWare(handler, logger)

	return handler
}

func middleWare(h http.Handler, logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		duration := time.Since(start)
		logger.Info("request", slog.String("method", r.Method), slog.String("url", r.URL.String()), slog.Duration("duration", duration),
			slog.String("remote_addr", r.RemoteAddr))
	})
}
