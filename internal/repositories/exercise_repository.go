package repositories

import (
	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ExerciseRepository interface {
	GetExerciseByID(exerciseID uuid.UUID) (*models.Exercise, error)
	GetExercises() ([]models.Exercise, error)
	CreateExercise(exercise *models.Exercise) error
	GetExerciseByName(exerciseName string) (*models.Exercise, error)
	FilterExercisesByMuscleGroup(muscleGroup string) ([]models.Exercise, error)
}

type exerciseRepository struct {
	database *sqlx.DB
}

func NewExerciseRepository(db *sqlx.DB) *exerciseRepository {
	return &exerciseRepository{database: db}
}

func (r *exerciseRepository) GetExercises() ([]models.Exercise, error) {
	var exercises []models.Exercise
	if err := r.database.Select(&exercises, "SELECT * FROM exercises"); err != nil {
		return nil, err
	}
	return exercises, nil
}

func (r *exerciseRepository) GetExerciseByID(exerciseID uuid.UUID) (*models.Exercise, error) {
	var exercise models.Exercise
	if err := r.database.Get(&exercise, "SELECT * FROM exercises WHERE id = $1", exerciseID); err != nil {
		return nil, err
	}
	return &exercise, nil
}

func (r *exerciseRepository) CreateExercise(exercise *models.Exercise) error {
	_, err := r.database.Exec(
		`INSERT INTO exercises (id, name, description, muscle_group) VALUES ($1, $2, $3, $4)`,
		exercise.ID, exercise.Name, exercise.Description, exercise.MuscleGroup,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *exerciseRepository) GetExerciseByName(exerciseName string) (*models.Exercise, error) {
	var exercise models.Exercise
	if err := r.database.Get(&exercise, "SELECT * FROM exercises WHERE name = $1", exerciseName); err != nil {
		return nil, err
	}
	return &exercise, nil
}

func (r *exerciseRepository) FilterExercisesByMuscleGroup(muscleGroup string) ([]models.Exercise, error) {
	var exercises []models.Exercise
	err := r.database.Select(&exercises, "SELECT * FROM exercises WHERE muscle_group = $1", muscleGroup)
	if err != nil {
		return nil, err
	}
	return exercises, nil
}
