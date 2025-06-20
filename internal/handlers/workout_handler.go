package handlers

import (
	"net/http"

	"github.com/Uranury/WorkoutApp/internal/apperror"
	"github.com/Uranury/WorkoutApp/internal/middleware"
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/Uranury/WorkoutApp/internal/services"
	"github.com/gin-gonic/gin"
)

type WorkoutHandler struct {
	workoutService *services.WorkoutService
}

func NewWorkoutHandler(workoutService *services.WorkoutService) *WorkoutHandler {
	return &WorkoutHandler{workoutService: workoutService}
}

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

func (h *WorkoutHandler) GetWorkouts(c *gin.Context) {
	userID, _ := middleware.GetUserID(c)

	workouts, err := h.workoutService.GetWorkouts(userID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, workouts)
}
