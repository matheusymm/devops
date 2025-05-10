package handlers

import "example/backend/db/repositories"

type Handlers struct {
	UserHandler *UserHandler
	MoodHandler *MoodHandler
}

func NewHandlers(ur *repositories.UserRepository, mr * repositories.MoodRepository) *Handlers {
	return &Handlers{
		UserHandler: NewUserHandler(ur),
		MoodHandler: NewMoodHandler(mr),
	}
}