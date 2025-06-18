package repositories

import (
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(userID uuid.UUID) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
	GetUsers() ([]models.User, error)
}

type userRepository struct {
	database *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *userRepository {
	return &userRepository{database: db}
}

func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.database.Get(&user, "SELECT * FROM users WHERE email = $1", email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := r.database.Get(&user, "SELECT * FROM users WHERE id = $1", userID); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.database.Get(&user, "SELECT * FROM users WHERE username = $1", username); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(user *models.User) error {
	_, err := r.database.Exec(
		`INSERT INTO users (id, username, email, password_hash, created_at) VALUES ($1, $2, $3, $4, $5)`,
		user.ID, user.Username, user.Email, user.PasswordHash, user.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := r.database.Select(&users, "SELECT * FROM users")
	return users, err
}
