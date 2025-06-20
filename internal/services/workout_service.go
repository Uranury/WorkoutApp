package services

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/Uranury/WorkoutApp/internal/apperror"
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/Uranury/WorkoutApp/internal/repositories"
	"github.com/google/uuid"
)

type WorkoutService struct {
	workoutRepo repositories.WorkoutRepository
}

func NewWorkoutService(workoutRepo repositories.WorkoutRepository) *WorkoutService {
	return &WorkoutService{workoutRepo: workoutRepo}
}

func (s *WorkoutService) CreateWorkout(input models.WorkoutDTO, userID uuid.UUID) error {
	if input.Name == "" {
		return apperror.NewAppError("workout name must be provided", http.StatusBadRequest)
	}
	if input.ScheduledAt.IsZero() {
		return apperror.NewAppError("scheduled_at must be provided", http.StatusBadRequest)
	}

	existingWorkout, err := s.workoutRepo.GetExistingWorkout(input.Name, userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return apperror.ErrDatabaseError
	}
	if existingWorkout != nil {
		return apperror.ErrWorkoutAlreadyExists
	}

	now := time.Now().UTC()
	workout := &models.Workout{
		ID:          uuid.New(),
		UserID:      userID,
		Name:        input.Name,
		Comment:     input.Comment,
		ScheduledAt: input.ScheduledAt,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.workoutRepo.CreateWorkout(workout); err != nil {
		return apperror.ErrDatabaseError
	}

	return nil
}

func (s *WorkoutService) GetWorkouts(userID uuid.UUID) ([]models.Workout, error) {
	workouts, err := s.workoutRepo.GetWorkoutsByUserID(userID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, apperror.ErrDatabaseError
	}
	if len(workouts) == 0 {
		return []models.Workout{}, nil
	}
	return workouts, nil
}
