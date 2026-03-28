package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
)

type PostBody struct {
	URL string `json:url`
}

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func NewHandler(repo *Repository) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	//ROTA PADRAO
	r.Route("/api/users", func(r chi.Router) {

		//Sub Rota Post
		r.Post("/", handleCreateUser(repo))

		//Sub Rota Get
		r.Get("/", handleGetAllUsers(repo))
		r.Get("/{id}", handleGetById(repo)) //IMPLEMENTAR

		//Sub Rota Put
		r.Put("/{id}", handleUpdateUser(repo))

		//Sub Rota Delete
		r.Delete("/{id}", handleDeleteUser(repo))
	})

	return r
}

// CREATE USER
func handleCreateUser(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			sendJson(w, Response{Error: "Invalid Body"}, http.StatusBadRequest)
			return
		}

		//VALIDACAO
		if len(user.FirstName) < 2 || len(user.LastName) < 2 || len(user.Biography) < 20 {
			sendJson(w, Response{Error: "Invalid fields"}, http.StatusBadRequest)
			return
		}

		created := repo.Insert(user)
		sendJson(w, Response{Data: created}, http.StatusCreated)
	}
}

// GET USER
func handleGetAllUsers(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		users := repo.FindAll()
		sendJson(w, Response{Data: users}, http.StatusOK)
	}
}

// GET BY ID
func handleGetById(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)
		if err != nil {
			sendJson(w, Response{Error: "Invalid ID"}, http.StatusBadRequest)
			return
		}

		user, ok := repo.FindById(id)
		if !ok {
			sendJson(w, Response{Error: "User not found"}, http.StatusNotFound)
			return
		}

		sendJson(w, Response{Data: user}, http.StatusOK)
	}
}

// UPDATE
func handleUpdateUser(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)
		if err != nil {
			sendJson(w, Response{Error: "Invalid ID"}, http.StatusBadRequest)
			return
		}

		var updated User
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			sendJson(w, Response{Error: "Invalid Body"}, http.StatusBadRequest)
			return
		}

		result, ok := repo.Update(id, updated)
		if !ok {
			sendJson(w, Response{Error: "User not found"}, http.StatusNotFound)
			return
		}

		sendJson(w, Response{Data: result}, http.StatusOK)
	}
}

// DELETE
func handleDeleteUser(repo *Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := chi.URLParam(r, "id")

		id, err := uuid.Parse(idStr)
		if err != nil {
			sendJson(w, Response{Error: "Invalid ID"}, http.StatusBadRequest)
			return
		}

		deleted, ok := repo.Delete(id)
		if !ok {
			sendJson(w, Response{Error: "User not found"}, http.StatusNotFound)
			return
		}

		sendJson(w, Response{Data: deleted}, http.StatusOK)
	}
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