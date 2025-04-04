package main

import (
	"flag"
	"fmt"

	"log/slog"
	"net/http"
	"os"
)

type config struct {
	port int
}

type application struct {
	logger *slog.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	app := application{
		logger: logger,
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
	}

	slog.Info("starting server", "addr", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		slog.Error(err.Error())
	}

}
