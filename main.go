package main

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/caioleone/go-user-crud/api"
)

type User struct {
	id         int
	first_name string
	last_name  string
	biography  string
}

func main() {
	if err := run(); err != nil {
		slog.Error("Failed to execute code", "error", err)
		return
	}
	slog.Info("All Systems Offline")
}

func run() error {
	db := make(map[string]string)

	handler := api.NewHandler(db)
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
