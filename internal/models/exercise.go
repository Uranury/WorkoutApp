package models

type Exercise struct {
	ID          string `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
	MuscleGroup string `db:"muscle_group" json:"muscle_group"` // or muscle group
}
