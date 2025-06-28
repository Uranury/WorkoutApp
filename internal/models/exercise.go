package models

import (
	"github.com/google/uuid"
)

type Exercise struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	MuscleGroup string    `db:"muscle_group" json:"muscle_group"`
}
