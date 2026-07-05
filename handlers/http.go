package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"user-api/models"
	"user-api/service"

	"github.com/go-chi/chi/v5"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	result := service.CreateUser(user)
	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created"))
}
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := service.GetUsers()
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
func GetUsersByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(data)
	user, error := service.GetUsersByID(id)
	if error != nil {
		http.Error(w, "User Not foud", http.StatusNotFound)
		return
	}
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(data)
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Wrong format", http.StatusBadRequest)
		return
	}
	newUser, err := service.UpdateUser(id, user)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)

}
func Delete(w http.ResponseWriter, r *http.Request) {
	data := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(data)
	success, err := service.DeleteUser(id)
	if err != nil || success == false {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
