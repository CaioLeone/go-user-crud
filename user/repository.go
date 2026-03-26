package user

import "github.com/google/uuid"

type Repository struct {
	data map[uuid.UUID]User
}

func NewRepository() *Repository {
	return &Repository{
		data: make(map[uuid.UUID]User),
	}
}
