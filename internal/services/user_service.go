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
	userRepo repositories.UserRepository
}

func NewUserService(userrepo repositories.UserRepository) *UserService {
	return &UserService{userRepo: userrepo}
}

func (s *UserService) CreateUser(username, email, plainPassword string) (*models.User, error) {
	username = strings.TrimSpace(username)
	email = strings.TrimSpace(email)
	plainPassword = strings.TrimSpace(plainPassword)

	if username == "" || email == "" || plainPassword == "" {
		return nil, apperror.ErrInvalidCredentials
	}

	existingUser, err := s.userRepo.GetUserByEmail(email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, apperror.ErrDatabaseError
	}
	if existingUser != nil {
		return nil, apperror.ErrEmailAlreadyInUse
	}

	existingUser, err = s.userRepo.GetUserByUsername(username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, apperror.ErrDatabaseError
	}
	if existingUser != nil {
		return nil, apperror.ErrUsernameAlreadyInUse
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperror.ErrHashPassword
	}

	now := time.Now()
	user := &models.User{
		ID:           uuid.New(),
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    now,
		UpdatedAt:    now,
		Role:         auth.User,
	}

	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, apperror.ErrDatabaseError
	}

	return user, nil
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
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, apperror.ErrDatabaseError
	}
	if len(users) == 0 {
		return []models.User{}, nil
	}

	return users, nil
}
