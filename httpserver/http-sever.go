package httpserver

import (
	"context"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

func StartServer(host *string, port *string) {
	ctx := context.Background()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	srv := &http.Server{
		Addr:              net.JoinHostPort(*host, *port),
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
		Handler:           http.TimeoutHandler(NewServer(logger), 2*time.Minute, ""),
	}
	l, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		log.Fatalln(err)
	}
	go func() {
		err := srv.Serve(l)
		if err != http.ErrServerClosed {
			log.Fatalln(err)
		}
	}()
	logger.Info("Server started listening", slog.String("address", srv.Addr))
	<-ctx.Done()
}
