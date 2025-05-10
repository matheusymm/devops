package repositories

import "database/sql"

type Repositories struct {
	UserRepository *UserRepository
	MoodRepository *MoodRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		UserRepository: NewUserRepository(db),
		MoodRepository: NewMoodRepository(db),
	}
}