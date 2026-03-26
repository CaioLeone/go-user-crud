package user

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, `{"error":"Invalid Body"}`, http.StatusBadRequest)
		return
	}

	if len(user.FirstName) < 2 || len(user.LastName) < 2 || len(user.Biography) < 20 {
		http.Error(w, `{"error":"Invalid Body"}`, http.StatusBadRequest)
		return
	}

	created := h.repo.Insert(user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
