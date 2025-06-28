package services_test

import (
	"testing"
	"time"

	"github.com/Uranury/WorkoutApp/internal/apperror"
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/Uranury/WorkoutApp/internal/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockWorkoutRepo mocks the WorkoutRepository
type MockWorkoutRepo struct {
	mock.Mock
}

func (m *MockWorkoutRepo) CreateWorkout(workout *models.Workout) error {
	args := m.Called(workout)
	return args.Error(0)
}

func (m *MockWorkoutRepo) GetWorkoutsByUserID(userID uuid.UUID) ([]models.Workout, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Workout), args.Error(1)
}

func (m *MockWorkoutRepo) GetWorkoutByID(workoutID, userID uuid.UUID) (*models.Workout, error) {
	args := m.Called(workoutID, userID)
	return args.Get(0).(*models.Workout), args.Error(1)
}

func (m *MockWorkoutRepo) UpdateWorkout(workout *models.Workout) error {
	args := m.Called(workout)
	return args.Error(0)
}

func (m *MockWorkoutRepo) DeleteWorkout(workout *models.Workout) error {
	args := m.Called(workout)
	return args.Error(0)
}

func (m *MockWorkoutRepo) GetUpcomingWorkouts(userID uuid.UUID) ([]models.Workout, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.Workout), args.Error(1)
}

func (m *MockWorkoutRepo) GetExistingWorkout(name string, userID uuid.UUID) (*models.Workout, error) {
	args := m.Called(name, userID)
	result := args.Get(0)
	if result == nil {
		return nil, args.Error(1)
	}
	return result.(*models.Workout), args.Error(1)
}

func TestCreateWorkout_Success(t *testing.T) {
	mockRepo := new(MockWorkoutRepo)
	svc := services.NewWorkoutService(mockRepo, nil)

	userID := uuid.New()
	input := models.WorkoutDTO{
		Name:        "Leg Day",
		ScheduledAt: time.Now().Add(1 * time.Hour),
		Comment:     "Heavy squats",
	}

	mockRepo.On("GetExistingWorkout", input.Name, userID).Return(nil, nil)
	mockRepo.On("CreateWorkout", mock.AnythingOfType("*models.Workout")).Return(nil)

	err := svc.CreateWorkout(input, userID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateWorkout_MissingName(t *testing.T) {
	mockRepo := new(MockWorkoutRepo)
	svc := services.NewWorkoutService(mockRepo, nil)

	input := models.WorkoutDTO{
		Name:        "",
		ScheduledAt: time.Now().Add(1 * time.Hour),
	}

	err := svc.CreateWorkout(input, uuid.New())

	appErr, ok := err.(*apperror.AppError)
	require.True(t, ok)
	assert.Equal(t, 400, appErr.StatusCode)
	assert.Equal(t, "workout name must be provided", appErr.Message)

}

func TestCreateWorkout_ScheduledInPast(t *testing.T) {
	mockRepo := new(MockWorkoutRepo)
	svc := services.NewWorkoutService(mockRepo, nil)

	input := models.WorkoutDTO{
		Name:        "Chest Day",
		ScheduledAt: time.Now().Add(-1 * time.Hour),
	}

	err := svc.CreateWorkout(input, uuid.New())

	assert.Error(t, err)
	assert.Equal(t, "scheduled_at cannot be in the past", err.Error())
}

func TestCreateWorkout_AlreadyExists(t *testing.T) {
	mockRepo := new(MockWorkoutRepo)
	svc := services.NewWorkoutService(mockRepo, nil)

	userID := uuid.New()
	input := models.WorkoutDTO{
		Name:        "Push Day",
		ScheduledAt: time.Now().Add(2 * time.Hour),
	}

	mockRepo.On("GetExistingWorkout", input.Name, userID).
		Return(&models.Workout{}, nil)

	err := svc.CreateWorkout(input, userID)

	assert.Error(t, err)
	assert.Equal(t, apperror.ErrWorkoutAlreadyExists, err)
}
