package handlers

import (
	"net/http"

	"github.com/Uranury/WorkoutApp/internal/apperror"
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/Uranury/WorkoutApp/internal/services"
	"github.com/gin-gonic/gin"
)

type ExerciseHandler struct {
	exerciseService *services.ExerciseService
}

func NewExerciseHandler(exerciseService *services.ExerciseService) *ExerciseHandler {
	return &ExerciseHandler{exerciseService: exerciseService}
}

// GetExercises godoc
// @Summary Retrieve all existing exercises, optionally filter by muscle group
// @Tags exercises
// @Produce json
// @Param muscle_group query string false "Filter by muscle group"
// @Success 200 {array} models.Exercise
// @Failure 500 {object} apperror.AppError
// @Router /exercises [get]
func (h *ExerciseHandler) GetExercises(c *gin.Context) {
	muscleGroup := c.Query("muscle_group")
	exercises, err := h.exerciseService.GetExercises(muscleGroup)
	if err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, exercises)
}

func (h *ExerciseHandler) CreateExercise(c *gin.Context) {

	/*
		userRole, _ := middleware.GetUserRole(c)
		if userRole != auth.Admin {
			HandleError(c, apperror.ErrUnauthorized)
			return
		}
	*/

	var exerciseRequest models.ExerciseCreateRequest
	if err := c.ShouldBindJSON(&exerciseRequest); err != nil {
		HandleError(c, apperror.ErrBadRequest)
		return
	}
	exercise := exerciseRequest.ToExercise()
	if err := h.exerciseService.CreateExercise(exercise); err != nil {
		HandleError(c, err)
		return
	}
	c.JSON(http.StatusCreated, exercise)
}
