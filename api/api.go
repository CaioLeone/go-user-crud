package api

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type PostBody struct {
	URL string `json:url`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewHandler(db map[string]string) http.Handler {
	route := chi.NewMux()

	route.Use(middleware.Recoverer)
	route.Use(middleware.RequestID)
	route.Use(middleware.Logger)

	route.Post("/api/", handlePost(db))
	route.Get("/{code}", handleGet(db))
	route.Delete("/del/{code}", handleDelete(db))
	route.Put("/api/put/{code}", handlePut(db))

	return route
}

func sendJson(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("Failed to marshal json data", "error", err)
		sendJson(
			w,
			Response{Error: "Something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("Failed To Write Response To Client", "error", err)
		return
	}
}

func handlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJson(
				w,
				Response{Error: "Invalid Body"},
				http.StatusUnprocessableEntity)
			return
		}

		//Validar
		if _, err := url.Parse(body.URL); err != nil {
			sendJson(
				w,
				Response{Error: "Invalid Body"},
				http.StatusBadRequest,
			)
		}
		sendJson(w, Response{}, http.StatusCreated)
	}
}

func handleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		data, ok := db[code]
		if !ok {
			http.Error(w, "url nao encontrada", http.StatusNotFound)
			return
		}
		http.Redirect(w, r, data, http.StatusPermanentRedirect)
	}
}

func handlePut(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func handleDelete(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
