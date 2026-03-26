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

// FIND FUNC
func (r *Repository) FindAll() []User {
	users := []User{}

	for _, user := range r.data {
		users = append(users, user)
	}

	return users
}

func (r *Repository) FindById(id uuid.UUID) (User, bool) {
	user, exists := r.data[id]
	return user, exists
}

//FUNC INSERT
func (r *Repository) Insert(user User) User {
	user.ID = uuid.New()
	r.data[user.ID] = user
	return user
}

//FUNC UPDATE
func (r *Repository) Update(id uuid.UUID, updated User) (User, bool) {
	_, exists := r.data[id]
	if !exists {
		return User{}, false
	}

	updated.ID = id
	r.data[id] = updated
	return updated, true
}

//DELETE
func (r *Repository) Delete(id uuid.UUID) (User, bool) {
	user, exists := r.data[id]
	if !exists {
		return User{}, false
	}
	delete(r.data, id)
	return user, true
}
