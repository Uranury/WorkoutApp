package handlers

import (
	"net/http"

	"github.com/Uranury/WorkoutApp/internal/apperror"
	"github.com/Uranury/WorkoutApp/internal/middleware"
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/Uranury/WorkoutApp/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WorkoutHandler struct {
	workoutService *services.WorkoutService
}

func NewWorkoutHandler(workoutService *services.WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{workoutService: workoutService}
}

// CreateWorkout godoc
// @Summary Create a new workout
// @Description Creates a workout for the authenticated user
// @Tags workouts
// @Accept json
// @Produce json
// @Param workout body models.WorkoutDTO true "Workout input"
// @Success 201 {object} map[string]string
// @Failure 400 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Security BearerAuth
// @Router /workouts [post]
func (h *WorkoutHandler) CreateWorkout(c *gin.Context) {
	var workoutInput models.WorkoutDTO

	userID, _ := middleware.GetUserID(c)

	if err := c.ShouldBindJSON(&workoutInput); err != nil {
		HandleError(c, apperror.ErrBadRequest)
		return
	}

	if err := h.workoutService.CreateWorkout(workoutInput, userID); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Workout created"})
}

// GetWorkouts godoc
// @Summary Get all workouts of the user
// @Description Returns a list of all workouts a user has
// @Tags workouts
// @Produce json
// @Success 200 {array} models.Workout
// @Failure 404 {object} apperror.AppError
// @Failure 401 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Security BearerAuth
// @Router /workouts [get]
func (h *WorkoutHandler) GetWorkouts(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	workouts, err := h.workoutService.GetWorkouts(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, workouts)
}

// GetFullWorkout godoc
// @Summary Get a workout with all its details
// @Description Returns detailed information about a specific workout by ID
// @Tags workouts
// @Produce json
// @Param id query string true "Workout ID (UUID)"
// @Success 200 {object} models.FullWorkout
// @Failure 400 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Security BearerAuth
// @Router /workouts/exercises [get]
func (h *WorkoutHandler) GetFullWorkout(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)
	workoutIDParam := c.Query("id")
	workoutID, err := uuid.Parse(workoutIDParam)
	if err != nil {
		HandleError(c, &apperror.AppError{Message: "failed to parse workoutID", StatusCode: http.StatusBadRequest})
		return
	}

	fullWorkout, err := h.workoutService.GetFullWorkout(workoutID, userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, fullWorkout)
}

// AddExerciseToWorkout godoc
// @Summary Add an exercise to an existing workout
// @Tags workouts
// @Accept json
// @Produce json
// @Param workoutID path string true "Workout ID (UUID)"
// @Param exercise body models.AddExerciseToWorkoutDTO true "Exercise details"
// @Success 201 {object} map[string]string
// @Failure 400 {object} apperror.AppError
// @Failure 404 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Failure 409 {object} apperror.AppError
// @Security BearerAuth
// @Router /workouts/{workoutID}/exercises [post]
func (h *WorkoutHandler) AddExerciseToWorkout(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	workoutIDParam := c.Param("workoutID")
	workoutID, err := uuid.Parse(workoutIDParam)
	if err != nil {
		HandleError(c, apperror.NewAppError("invalid workout ID", http.StatusBadRequest))
		return
	}

	// Only bind the exercise-specific stuff
	var req models.AddExerciseToWorkoutDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		HandleError(c, apperror.ErrBadRequest)
		return
	}

	// Construct the full model internally
	workoutExercise := &models.WorkoutExercise{
		ID:         uuid.New(),
		WorkoutID:  workoutID,
		ExerciseID: req.ExerciseID,
		Sets:       req.Sets,
		Reps:       req.Reps,
		Weight:     req.Weight,
	}

	if err := h.workoutService.AddExerciseToWorkout(userID, workoutExercise); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "exercise added"})
}
