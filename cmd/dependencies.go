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
