package models

import (
	"time"

	"github.com/Uranury/WorkoutApp/internal/auth"
	"github.com/google/uuid"
)

type AddExerciseToWorkoutDTO struct {
	ExerciseID uuid.UUID `json:"exercise_id" example:"b5c3a6b6-2a5f-4d0d-9f39-2df5df45c2a3"`
	Sets       int       `json:"sets" example:"3"`
	Reps       int       `json:"reps" example:"10"`
	Weight     int       `json:"weight" example:"50"`
}

type WorkoutDTO struct {
	Name        string    `json:"name" example:"Leg Day"`
	Comment     string    `json:"comment" example:"Focus on squats and lunges"`
	ScheduledAt time.Time `json:"scheduled_at" example:"2025-06-27T18:00:00Z"`
}

type UserCreateRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"` // Plain password
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	Message string    `json:"message" example:"User created successfully"`
	User    UserInner `json:"user"`
}

type UserInner struct {
	ID        uuid.UUID `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	Username  string    `json:"username" example:"johndoe"`
	Email     string    `json:"email" example:"johndoe@example.com"`
	CreatedAt time.Time `json:"created_at" example:"2025-06-27T18:00:00Z"`
	Role      auth.Role `json:"role" example:"user"`
}

type ExerciseCreateRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	MuscleGroup string `json:"muscle_group" binding:"required"`
}

func (req *ExerciseCreateRequest) ToExercise() *Exercise {
	return &Exercise{
		Name:        req.Name,
		Description: req.Description,
		MuscleGroup: req.MuscleGroup,
	}
}
