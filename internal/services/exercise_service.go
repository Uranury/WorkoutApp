package services

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/Uranury/WorkoutApp/internal/apperror"
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/Uranury/WorkoutApp/internal/repositories"
	"github.com/google/uuid"
)

type ExerciseService struct {
	exerciseRepo repositories.ExerciseRepository
}

func NewExerciseService(exerciseRepo repositories.ExerciseRepository) *ExerciseService {
	return &ExerciseService{exerciseRepo: exerciseRepo}
}

func (s *ExerciseService) GetExercises() ([]models.Exercise, error) {
	exercises, err := s.exerciseRepo.GetExercises()
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, apperror.ErrDatabaseError
	}
	if len(exercises) == 0 {
		return []models.Exercise{}, nil
	}

	return exercises, nil
}

func (s *ExerciseService) CreateExercise(exercise *models.Exercise) error {
	exercise.Name = strings.TrimSpace(exercise.Name)
	exercise.MuscleGroup = strings.TrimSpace(exercise.MuscleGroup)

	existingExercise, err := s.exerciseRepo.GetExerciseByName(exercise.Name)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return apperror.ErrDatabaseError
	}
	if existingExercise != nil {
		return apperror.ErrExerciseAlreadyExists
	}
	exercise.ID = uuid.New()

	if err := s.exerciseRepo.CreateExercise(exercise); err != nil {
		return apperror.ErrDatabaseError
	}
	return nil
}
