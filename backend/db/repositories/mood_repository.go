package repositories

import (
	"database/sql"
	"example/backend/api/models"
)

type MoodRepositoryInterface interface {
	Create(*models.Mood) (bool, error)
	GetById(string) (*models.Mood, error)
	GetByUserId(string) ([]*models.Mood, error)
	GetAll() ([]*models.Mood, error)
	Update(*models.Mood) (bool, error)
	Delete(string) (bool, error)
}

type MoodRepository struct {
	db *sql.DB
}

func NewMoodRepository(db *sql.DB) *MoodRepository {
	return &MoodRepository{
		db: db,
	}
}

func (mr *MoodRepository) Create(mood *models.Mood) (bool, error) {
	query := "INSERT INTO mood (user_id, mood, description) VALUES ($1, $2, $3) RETURNING id"

	err := mr.db.QueryRow(query, mood.UserId, mood.Mood, mood.Description).Scan(&mood.Id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (mr *MoodRepository) GetById(id string) (*models.Mood, error) {
	query := "SELECT id, user_id, mood, description, created_at FROM mood WHERE id = $1"
	mood := &models.Mood{}

	err := mr.db.QueryRow(query, id).Scan(&mood.Id, &mood.UserId, &mood.Mood, &mood.Description, &mood.CreatedAt)

	if err != nil {
		return nil, err
	}
	return mood, nil
}

func (mr *MoodRepository) GetByUserId(userId string) ([]*models.Mood, error) {
	query := "SELECT id, user_id, mood, description, created_at FROM mood WHERE user_id = $1"
	rows, err := mr.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moods []*models.Mood

	for rows.Next() {
		mood := &models.Mood{}
		err := rows.Scan(&mood.Id, &mood.UserId, &mood.Mood, &mood.Description, &mood.CreatedAt)
		if err != nil {
			return nil, err
		}
		moods = append(moods, mood)
	}

	return moods, nil
}

func (mr *MoodRepository) GetAll() ([]*models.Mood, error) {
	query := "SELECT id, user_id, mood, description, created_at FROM mood"
	rows, err := mr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var moods []*models.Mood

	for rows.Next() {
		mood := &models.Mood{}
		err := rows.Scan(&mood.Id, &mood.UserId, &mood.Mood, &mood.Description, &mood.CreatedAt)
		if err != nil {
			return nil, err
		}
		moods = append(moods, mood)
	}

	return moods, nil
}

func (mr *MoodRepository) Update(mood *models.Mood) (bool, error) {
	query := "UPDATE mood SET user_id = $1, mood = $2, description = $3m created_at = $4 WHERE id = $5"

	_, err := mr.db.Exec(query, mood.UserId, mood.Mood, mood.Description, mood.CreatedAt, mood.Id)
	
	if err != nil {
		return false, err
	}
	return true, nil
}

func (mr *MoodRepository) Delete(id string) (bool, error) {
	query := "DELETE FROM mood WHERE id = $1"

	_, err := mr.db.Exec(query, id)

	if err != nil {
		return false, err
	}
	return true, nil
}