package main

import (
	"fmt"
	"net/http"
	"user-api/handlers"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/users",handlers.GetUsers)
	r.Post("/users",handlers.CreateUser)
	r.Get("/users/{id}",handlers.GetUsersByID)
	r.Put("/users/update/{id}",handlers.Update)
	http.ListenAndServe(":8080", r)
	fmt.Println("Hi")
}
