package main

import (
	"log"
	"net/http"
	"user-api/database"
	"user-api/handlers"
	"user-api/repository"
	"user-api/service"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env files")
	}
	
	r := chi.NewRouter()
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	repo := repository.NewMySQLRepository(db)
	service := service.NewUserService(repo)
	handler := handlers.NewUserHandler(service)
	r.Get("/users", handler.GetUsers)
	r.Post("/users", handler.CreateUser)
	r.Get("/users/{id}", handler.GetUsersByID)
	r.Put("/users/update/{id}", handler.Update)
	r.Delete("/users/delete/{id}", handler.Delete)
	http.ListenAndServe(":8080", r)

}
