package models

import "time"

type Workout struct {
	ID          string    `db:"id" json:"id"`
	UserID      string    `db:"user_id" json:"user_id"` // FK to User
	Name        string    `db:"name" json:"name"`
	ScheduledAt time.Time `db:"scheduled_at" json:"scheduled_at"`
	Comment     string    `db:"comment" json:"comment"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type WorkoutExercise struct {
	ID         string `db:"id" json:"id"`
	WorkoutID  string `db:"workout_id" json:"workout_id"`   // FK to Workout
	ExerciseID string `db:"exercise_id" json:"exercise_id"` // FK to Exercise
	Sets       int    `db:"sets" json:"sets"`
	Reps       int    `db:"reps" json:"reps"`
	Weight     int    `db:"weight" json:"weight"` // kg or lbs, you decide
}
