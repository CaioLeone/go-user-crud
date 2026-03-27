package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"

	entity "github.com/caioleone/go-user-crud/user"
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

	//ROTA PADRAO
	route.Route("/api/users", func(r chi.Router) {

		//Sub Rota Post
		route.Post("/", handlePost(db))

		//Sub Rota Get
		route.Get("/", handleGetAll(db))
		route.Get("/id", handleGet(db)) //IMPLEMENTAR

		//Sub Rota Put
		route.Put("/id", handlePut(db))

		//Sub Rota Delete
		route.Delete("/id", handleDelete(db))
	})

	return route
}

// CREATE USER
func handleCreateUser(db map[string]entity.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user entity.User

		if err := json.Decoder(r.Body).Decode(&user); err != nil {
			sendJson(w, Response{Error: "Invalid Body"}, http.StatusBadRequest)
			return
		}

		//VALIDACAO
		if len(user.FirstName) < 2 || len(user.LastName) < 2 || len(user.Biography) < 20 {
			sendJson(w, Response{Error: "Invalid fields"}, http.StatusBadRequest)
			return
		}

		user.ID = uuid.New().String()
		db[user.ID] = user

		sendJson(w, Response{Data: user}, http.StatusCreated)
	}
}

// GET USER
func handleGetAllUsers(db map[string]entity.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []entity.User

		for _, u := range db{
			users, append(users, u)
		}

		sendJson(w, Response{Data: users}, http.StatusOK)
	}
}

// GET BY ID
func handleGetById(db map[string]entity.User) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		user, ok := db[id]
		if !ok {
			sendJson(w, Response{Error: "User not found"}, htpp.StatusNotFound)
			return
		}

		sendJson(w, Response{Data: user}, http.StatusOK)
	}
}

//UPDATE
func handleUpdateUser(db map[string]entity.User) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		_, ok := db[id]
		if !ok {
			sendJson(w, Response{Error: "User not found"}, htpp.StatusNotFound)
			return
		}

		var updated entity.User
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			sendJson(w, Response{Error: "Invalid Body"}, http.StatusBadRequest)
			return 
		}

		updated.ID = id
		db[id] = updated

		sendJson(w, Response{Data: updated}, http.StatusOK)
	}
}

//DELETE
func handleDeleteUser(db map[string]entity.User) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		user, ok := db[id]
		if !ok{
			sendJson(w, Response{Error: "User not found"}, http.StatusNotFound)
			return
		}

		delete(db, id)

		sendJson(w, Response{Data: user}, http.StatusOK)
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

// POST
func handlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var user entity.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			sendJson(
				w,
				Response{Error: "Invalid Body"},
				http.StatusUnprocessableEntity)
			return
		}

		//Validar
		if len(user.FirstName) < 2 || len(user.LastName) < 2 || len(user.Biography) < 20 {
			sendJson(
				w,
				Response{Error: "Invalid Fields"},
				http.StatusBadRequest)
			return
		}
		entity.User.ID = uuid.New().String()
		db[user.ID] = user

		sendJson(w, Response{Data: user}, http.StatusCreated)
	}
}

// GET
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

func handleGetAll(db map[string]string) http.HandlerFunc {
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

// PUT
func handlePut(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// DELETE
func handleDelete(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
