package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"user-api/models"
	"user-api/service"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *service.UserService
	log     *slog.Logger
}

func NewUserHandler(service *service.UserService, log *slog.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		log:     log,
	}
}
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.log.Warn("Invalid Request Body", "method", r.Method, "path", r.URL.Path, "error", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	result, err := h.service.CreateUser(user)
	if err != nil {
		h.log.Error("Failed to create User", "method", r.Method, "path", r.URL.Path, "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	h.log.Info("user created successfull", "method", r.Method, "path", r.URL.Path, "user_id", result.ID)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created"))
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		h.log.Error(
			"Failed to encode response",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err,
		)
	}

}
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := h.service.GetUsers()
	if err != nil {
		h.log.Error("Failed to Get User", "method", r.Method, "path", r.URL.Path, "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		h.log.Error(
			"Failed to encode response",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err,
		)
	}

}
func (h *UserHandler) GetUsersByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(data)
	user, err := h.service.GetUserByID(id)
	if err != nil {
		h.log.Error("Failed to  Get User", "method", r.Method, "path", r.URL.Path, "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		h.log.Error(
			"Failed to encode response",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err,
		)
	}

}
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(data)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.log.Warn("Invalid Request Body", "method", r.Method, "path", r.URL.Path, "error", err)
		http.Error(w, "Wrong format", http.StatusBadRequest)
		return
	}
	newUser, err := h.service.UpdateUser(id, user)
	if err != nil {
		h.log.Error("Failed to  Update User", "method", r.Method, "path", r.URL.Path, "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(newUser)
	if err != nil {
		h.log.Error(
			"Failed to encode response",
			"method", r.Method,
			"path", r.URL.Path,
			"error", err,
		)
	}

}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(data)
	err := h.service.DeleteUser(id)
	if err != nil {
		h.log.Error("Failed to  Delete User", "method", r.Method, "path", r.URL.Path, "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
