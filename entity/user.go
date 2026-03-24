package entity

type id uuid.UUID

type user struct {
	FirstName string
	LastName  string
	Biography string
}

type application struct {
	data map[id]user
}
