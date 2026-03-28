package main

import (
	"log/slog"
	"net/http"
	"time"

	newApi "github.com/caioleone/go-user-crud/api"
)

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}
	slog.Info("All Systems Offline")
}

func run() error {
	repo := newApi.NewRepository()

	handler := newApi.NewHandler(repo)

	s := http.Server{
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
		Addr:         ":8080",
		Handler:      handler,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
