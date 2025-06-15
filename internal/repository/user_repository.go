package repository

import (
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
}

type UserRepository struct {
	database *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{database: db}
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.database.Get(&user, "SELECT * FROM users WHERE email = $1", email); err != nil {
		return nil, err
	}
	return &user, nil
}
