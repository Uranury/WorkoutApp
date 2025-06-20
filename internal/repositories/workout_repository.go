package repositories

import (
	"database/sql"

	"github.com/Uranury/WorkoutApp/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type WorkoutRepository interface {
	CreateWorkout(workout *models.Workout) error
	GetWorkoutsByUserID(userID uuid.UUID) ([]models.Workout, error)
	UpdateWorkout(workout *models.Workout) error
	DeleteWorkout(workout *models.Workout) error
	GetUpcomingWorkouts(userID uuid.UUID) ([]models.Workout, error)
	GetExistingWorkout(name string, userID uuid.UUID) (*models.Workout, error)
}

type workoutRepository struct {
	database *sqlx.DB
}

func NewWorkoutRepository(db *sqlx.DB) *workoutRepository {
	return &workoutRepository{database: db}
}

func (r *workoutRepository) CreateWorkout(workout *models.Workout) error {
	_, err := r.database.Exec(
		`INSERT INTO workouts (id, user_id, name, scheduled_at, comment, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		workout.ID, workout.UserID, workout.Name, workout.ScheduledAt, workout.Comment, workout.CreatedAt, workout.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *workoutRepository) GetWorkoutsByUserID(userID uuid.UUID) ([]models.Workout, error) {
	var workouts []models.Workout
	err := r.database.Select(&workouts, "SELECT * FROM workouts WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	return workouts, nil
}

func (r *workoutRepository) GetExistingWorkout(name string, userID uuid.UUID) (*models.Workout, error) {
	var workout models.Workout
	if err := r.database.Get(&workout, "SELECT * FROM workouts WHERE name ILIKE $1 AND user_id = $2", name, userID); err != nil {
		return nil, err
	}
	return &workout, nil
}

func (r *workoutRepository) UpdateWorkout(workout *models.Workout) error {
	_, err := r.database.Exec(`
		UPDATE workouts
		SET name = $1, scheduled_at = $2, comment = $3, updated_at = NOW()
		WHERE id = $4
	`,
		workout.Name,
		workout.ScheduledAt,
		workout.Comment,
		workout.ID,
	)
	return err
}

func (r *workoutRepository) DeleteWorkout(workout *models.Workout) error {
	result, err := r.database.Exec(
		`DELETE FROM workouts WHERE id = $1`,
		workout.ID,
	)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *workoutRepository) GetUpcomingWorkouts(userID uuid.UUID) ([]models.Workout, error) {
	var workouts []models.Workout
	err := r.database.Select(
		&workouts,
		`SELECT * FROM workouts 
		WHERE user_id = $1 AND scheduled_at >= NOW()
		ORDER BY scheduled_at ASC`,
		userID,
	)
	if err != nil {
		return nil, err
	}

	return workouts, nil
}
