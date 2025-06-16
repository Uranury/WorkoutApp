package repositories

import (
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type IUserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(userid uuid.UUID) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
	GetUsers() ([]models.User, error)
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

func (r *UserRepository) GetUserById(userid uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.database.Get(&user, "SELECT * FROM users WHERE id = $1", userid); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.database.Get(&user, "SELECT * FROM users WHERE username = $1", username); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.database.Exec(
		`INSERT INTO users (id, username, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)`,
		user.ID, user.Username, user.Email, user.PasswordHash, user.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := r.database.Select(&users, "SELECT * FROM users")
	return users, err
}
