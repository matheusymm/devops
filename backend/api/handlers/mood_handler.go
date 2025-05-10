package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"example/backend/api/models"
	"example/backend/db/repositories"

	"github.com/go-chi/chi/v5"
)

type MoodHandler struct {
	MoodRepository *repositories.MoodRepository
}

func NewMoodHandler(ur *repositories.MoodRepository) *MoodHandler {
	return &MoodHandler{
		MoodRepository: ur,
	}
}

func (mh *MoodHandler) CreateMood(w http.ResponseWriter, r *http.Request) {
	var mood models.Mood
	if err := json.NewDecoder(r.Body).Decode(&mood); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		log.Println("Error decoding request payload:", err)
		return
	}

	result, err := mh.MoodRepository.Create(&mood)

	if err != nil || !result {
		http.Error(w, "failed to create mood", http.StatusInternalServerError)
		log.Println("Error creating mood:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (mh *MoodHandler) GetMoodsByUserId(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "missing user ID", http.StatusBadRequest)
		return
	}

	moods, err := mh.MoodRepository.GetByUserId(id)
	if err != nil {
		http.Error(w, "failed to get moods", http.StatusInternalServerError)
		log.Println("Error getting moods:", err)
		return
	}
	if moods == nil {
		http.Error(w, "moods not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(moods)
}

func (mh *MoodHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	moods, err := mh.MoodRepository.GetAll()
	if err != nil {
		http.Error(w, "failed to get moods", http.StatusInternalServerError)
		log.Println("Error getting moods:", err)
		return
	}
	if moods == nil {
		http.Error(w, "moods not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(moods)
}

func (mh *MoodHandler) UpdateMood(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing mood ID", http.StatusBadRequest)
		return
	}

	var mood models.Mood
	if err := json.NewDecoder(r.Body).Decode(&mood); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		log.Println("Error decoding request payload:", err)
		return
	}

	result, err := mh.MoodRepository.Update(&mood)
	if err != nil || !result {
		http.Error(w, "failed to update mood", http.StatusInternalServerError)
		log.Println("Error updating mood:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (mh *MoodHandler) DeleteMood(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "missing mood ID", http.StatusBadRequest)
		return
	}

	result, err := mh.MoodRepository.Delete(id)

	if err != nil || !result {
		http.Error(w, "failed to delete mood", http.StatusInternalServerError)
		log.Println("Error deleting mood:", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
