package models

import (
	"time"

	"github.com/google/uuid"
)

type Workout struct {
	ID          uuid.UUID `db:"id" json:"id"`
	UserID      uuid.UUID `db:"user_id" json:"user_id"` // FK to User
	Name        string    `db:"name" json:"name"`
	ScheduledAt time.Time `db:"scheduled_at" json:"scheduled_at"`
	Comment     string    `db:"comment" json:"comment"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type WorkoutDTO struct {
	Name        string    `json:"name" example:"Leg Day"`
	Comment     string    `json:"comment" example:"Focus on squats and lunges"`
	ScheduledAt time.Time `json:"scheduled_at" example:"2025-06-27T18:00:00Z"`
}

type WorkoutExercise struct {
	ID         uuid.UUID `db:"id" json:"id"`
	WorkoutID  uuid.UUID `db:"workout_id" json:"workout_id"`   // FK to Workout
	ExerciseID uuid.UUID `db:"exercise_id" json:"exercise_id"` // FK to Exercise
	Sets       int       `db:"sets" json:"sets"`
	Reps       int       `db:"reps" json:"reps"`
	Weight     int       `db:"weight" json:"weight"` // kg or lbs, you decide
}

type WorkoutExerciseDetail struct {
	ID          uuid.UUID `db:"id" json:"id"`
	WorkoutID   uuid.UUID `db:"workout_id" json:"workout_id"`
	ExerciseID  uuid.UUID `db:"exercise_id" json:"exercise_id"`
	Sets        int       `db:"sets" json:"sets"`
	Reps        int       `db:"reps" json:"reps"`
	Weight      int       `db:"weight" json:"weight"`
	Name        string    `db:"name" json:"name"`
	MuscleGroup string    `db:"muscle_group" json:"muscle_group"`
	Description string    `db:"description" json:"description"`
}

type FullWorkout struct {
	Workout
	Exercises []WorkoutExerciseDetail `json:"exercises"`
}
