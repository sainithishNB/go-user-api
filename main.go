package main

import (
	"log"
	"net/http"
	"user-api/config"
	"user-api/database"
	"user-api/handlers"
	"user-api/logger"
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
	cfg := config.Load()
	log := logger.NewLogger(cfg)
	log.Debug("This is a debug log")
	log.Info("Application started")
	log.Warn("This is a warning")
	log.Error("This is an error")
	r := chi.NewRouter()
	db, err := database.Connect(cfg, log)
	repo := repository.NewMySQLRepository(db, log)
	service := service.NewUserService(repo)
	handler := handlers.NewUserHandler(service, log)
	r.Get("/users", handler.GetUsers)
	r.Post("/users", handler.CreateUser)
	r.Get("/users/{id}", handler.GetUsersByID)
	r.Put("/users/update/{id}", handler.Update)
	r.Delete("/users/delete/{id}", handler.Delete)
	http.ListenAndServe(":8080", r)

}
