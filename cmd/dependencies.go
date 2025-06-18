package main

import (
	"github.com/Uranury/WorkoutApp/internal/handlers"
	"github.com/Uranury/WorkoutApp/internal/repositories"
	"github.com/Uranury/WorkoutApp/internal/services"
	"github.com/jmoiron/sqlx"
)

func InitUserHandler(db *sqlx.DB) *handlers.UserHandler {
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	return handlers.NewUserHandler(userService)
}

func InitExerciseHandler(db *sqlx.DB) *handlers.ExerciseHandler {
	exerciseRepo := repositories.NewExerciseRepository(db)
	exerciseService := services.NewExerciseService(exerciseRepo)
	return handlers.NewExerciseHandler(exerciseService)
}
