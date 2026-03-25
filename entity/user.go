package entity

type id uuid.uuid

type User struct {
	FirstName string
	LastName  string
	Biography string
}

type application struct {
	data map[id]User
}
