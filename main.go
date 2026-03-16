package main

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	fmt.Println("Crud usando CHI")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
}
