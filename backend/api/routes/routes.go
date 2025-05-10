package routes

import (
	"example/backend/api/handlers"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(r *chi.Mux, h *handlers.Handlers) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", h.UserHandler.CreateUser)
		r.Get("/{id}", h.UserHandler.GetUserByID)
		r.Get("/", h.UserHandler.GetAll)
		r.Put("/{id}", h.UserHandler.UpdateUser)
		r.Delete("/{id}", h.UserHandler.DeleteUser)
	})
	r.Route("/moods", func(r chi.Router) {
		r.Post("/", h.MoodHandler.CreateMood)
		r.Get("/user/{id}", h.MoodHandler.GetMoodByUserId)
		r.Get("/", h.MoodHandler.GetAll)
		r.Put("/{id}", h.MoodHandler.UpdateMood)
		r.Delete("/{id}", h.MoodHandler.DeleteMood)
	})
}