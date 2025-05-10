package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"example/backend/api/models"
	"example/backend/db/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	UserRepository *repositories.UserRepository
}

func NewUserHandler(ur *repositories.UserRepository) *UserHandler {
	return &UserHandler{
		UserRepository: ur,
	}
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("Error decoding request payload:", err)
        return
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
    if err != nil {
        http.Error(w, "Password hashing failed", http.StatusInternalServerError)
		log.Println("Error hashing password:", err)
        return
    }

    user.Password = hashedPassword

    result, err := uh.UserRepository.Create(&user)
    if err != nil || !result {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
		log.Println("Error creating user:", err)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func (uh *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing user ID", http.StatusBadRequest)
		return
	}

	user, err := uh.UserRepository.GetById(id)
	if err != nil {
		http.Error(w, "failed to get user", http.StatusInternalServerError)
		log.Println("Error getting user:", err)
		return
	}
	if user == nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "failed to encode user", http.StatusInternalServerError)
		log.Println("Error encoding user:", err)
		return
	}
}

func (uh *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	users, err := uh.UserRepository.GetAll()
	if err != nil {
		http.Error(w, "failed to get users", http.StatusInternalServerError)
		log.Println("Error getting users:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "failed to encode users", http.StatusInternalServerError)
		log.Println("Error encoding users:", err)
		return
	}
}

func (uh *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing user ID", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("Error decoding request payload:", err)
		return
	}

	result, err := uh.UserRepository.Update(&user)

	if err != nil || !result {
		http.Error(w, "failed to update user", http.StatusInternalServerError)
		log.Println("Error updating user:", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (uh *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing user ID", http.StatusBadRequest)
		return
	}

	result, err := uh.UserRepository.Delete(id)

	if err != nil || !result {
		http.Error(w, "failed to delete user", http.StatusInternalServerError)
		log.Println("Error deleting user:", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}