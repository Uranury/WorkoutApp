package repositories

import (
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type WorkoutExerciseRepository interface {
	AddExerciseToWorkout(exercise *models.WorkoutExercise) error
	GetExercisesByWorkoutID(workoutID uuid.UUID) ([]models.WorkoutExerciseDetail, error)
	GetWorkoutExercise(workoutID, exerciseID uuid.UUID) (*models.WorkoutExercise, error)
}

type workoutExerciseRepository struct {
	database *sqlx.DB
}

func NewWorkoutExerciseRepository(db *sqlx.DB) *workoutExerciseRepository {
	return &workoutExerciseRepository{database: db}
}

func (r *workoutExerciseRepository) AddExerciseToWorkout(exercise *models.WorkoutExercise) error {
	_, err := r.database.Exec(
		`INSERT INTO workoutexercises (id, workout_id, exercise_id, sets, reps, weight)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		exercise.ID, exercise.WorkoutID, exercise.ExerciseID, exercise.Sets, exercise.Reps, exercise.Weight,
	)
	return err
}

func (r *workoutExerciseRepository) GetExercisesByWorkoutID(workoutID uuid.UUID) ([]models.WorkoutExerciseDetail, error) {
	var workoutExercises []models.WorkoutExerciseDetail

	query := `
		SELECT 
			we.id,
			we.workout_id,
			we.exercise_id,
			we.sets,
			we.reps,
			we.weight,
			e.name,
			e.muscle_group,
			e.description
		FROM
			workoutexercises we
		JOIN
			exercises e on we.exercise_id = e.id
		WHERE
			we.workout_id = $1;
	`

	err := r.database.Select(&workoutExercises, query, workoutID)
	if err != nil {
		return nil, err
	}

	return workoutExercises, nil
}

func (r *workoutExerciseRepository) GetWorkoutExercise(workoutID, exerciseID uuid.UUID) (*models.WorkoutExercise, error) {
	var we models.WorkoutExercise
	err := r.database.Get(&we, `
        SELECT * FROM workoutexercises 
        WHERE workout_id = $1 AND exercise_id = $2
    `, workoutID, exerciseID)
	if err != nil {
		return nil, err
	}
	return &we, nil
}
