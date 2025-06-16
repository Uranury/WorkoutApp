package services

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/Uranury/WorkoutApp/internal/apperror"
	"github.com/Uranury/WorkoutApp/internal/auth"
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/Uranury/WorkoutApp/internal/repositories"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo repositories.IUserRepository
}

func NewUserService(userrepo repositories.IUserRepository) *UserService {
	return &UserService{userRepo: userrepo}
}

func (s *UserService) CreateUser(user *models.User) error {
	existingUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return apperror.ErrDatabaseError
	}
	if existingUser != nil {
		return apperror.ErrEmailAlreadyInUse
	}

	existingUser, err = s.userRepo.GetUserByUsername(user.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return apperror.ErrDatabaseError
	}
	if existingUser != nil {
		return apperror.ErrUsernameAlreadyInUse
	}

	// Generate UUID for the user
	user.ID = uuid.New()

	plainPassword := strings.TrimSpace(user.PasswordHash)
	if plainPassword == "" {
		return apperror.ErrInvalidCredentials
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return apperror.ErrHashPassword
	}
	user.PasswordHash = string(hashedPassword)

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// Set default role if not set
	if user.Role == "" {
		user.Role = auth.User
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return apperror.ErrDatabaseError
	}

	return nil
}

func (s *UserService) LoginUser(username string, password string) (string, error) {
	existingUser, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", apperror.ErrUserNotFound
		}
		return "", apperror.ErrDatabaseError
	}

	if existingUser == nil {
		return "", apperror.ErrUserNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.PasswordHash), []byte(password)); err != nil {
		return "", apperror.ErrInvalidCredentials
	}

	token, err := auth.GenerateJWT(existingUser.ID, existingUser.Role)
	if err != nil {
		return "", apperror.ErrInternalServer
	}

	return token, nil
}

func (s *UserService) GetUsers() ([]models.User, error) {
	users, err := s.userRepo.GetUsers()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, apperror.ErrNoUsersFound
		}
		return nil, apperror.ErrDatabaseError
	}

	return users, nil
}
