package services

import (
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/Uranury/WorkoutApp/internal/repository"
)

type UserService struct {
	userRepo repository.IUserRepository
}

func NewUserService(userrepo repository.IUserRepository) *UserService {
	return &UserService{userRepo: userrepo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return nil
}

func (s *UserService) LoginUser(user *models.User) (string, error) {
	return "", nil
}

func (s *UserService) GetUsers() ([]models.User, error) {
	return nil, nil
}
