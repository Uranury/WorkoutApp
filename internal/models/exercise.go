package models

import "github.com/google/uuid"

type Exercise struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	MuscleGroup string    `db:"muscle_group" json:"muscle_group"`
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
