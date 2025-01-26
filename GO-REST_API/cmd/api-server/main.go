package main

import (
	"log/slog"
	"net/http"
	"new-api/internal/router"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))
	r := router.New()
	if err := http.ListenAndServe(":8080", r); err != nil {
		logger.Error("failed to start server", "error", err)
	}
}
