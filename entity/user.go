package entity

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Biography string    `json:"biography"`
}

type Repository struct {
	data map[uuid.UUID]User
}

func NewRepository() *Repository {
	return &Repository{
		data: make(map[uuid.UUID]User),
	}
}
