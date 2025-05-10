package repositories

import (
	"database/sql"
	"example/backend/api/models"
)

type UserRepositoryInterface interface {
	Create(*models.User) (string, error)
	GetById(string) (*models.User, error)
	GetAll() ([]*models.User, error)
	Update(*models.User) (bool, error)
	Delete(string) (bool, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user *models.User) (string, error) {
	query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"

	err := ur.db.QueryRow(query, user.Name, user.Email, user.Password).Scan(&user.Id)

	if err != nil {
		return "", err
	}
	return user.Id.String(), nil
}

func (ur *UserRepository) GetById(id string) (*models.User, error) {
	query := "SELECT id, name, email FROM users WHERE id = $1"
	user := &models.User{}

	err := ur.db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id, name, password FROM users WHERE email = $1"
	user := &models.User{}

	err := ur.db.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Password)

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) GetAll() ([]*models.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) Update(user *models.User) (bool, error) {
	query := "UPDATE users SET"

	if user.Name != "" {
		query += " name = $1,"
	}
	if user.Email != "" {
		query += " email = $2,"
	}
	if string(user.Password) != "" {
		query += " password = $3,"
	}

	query = query[:len(query)-1] + " WHERE id = $4"
	_, err := ur.db.Exec(query, user.Name, user.Email, user.Password, user.Id)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (ur *UserRepository) Delete(id string) (bool, error) {
	query := "DELETE FROM users WHERE id = $1"
	_, err := ur.db.Exec(query, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
