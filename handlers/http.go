package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-api/models"
	"user-api/service"

	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	result, _ := h.service.CreateUser(user)
	w.WriteHeader(http.StatusCreated)
	
	json.NewEncoder(w).Encode(result)
	w.Write([]byte("user created"))
}
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, _ := h.service.GetUsers()
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
func (h *UserHandler) GetUsersByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(data)
	user, error := h.service.GetUserByID(id)
	if error != nil {
		http.Error(w, "User Not foud", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(data)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Wrong format", http.StatusBadRequest)
		return
	}
	newUser, err := h.service.UpdateUser(id, user)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(data)
	err := h.service.DeleteUser(id)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
