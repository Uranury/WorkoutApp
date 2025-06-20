package main

import (
	"log"
	"os"

	"github.com/Uranury/WorkoutApp/config"
	"github.com/Uranury/WorkoutApp/internal/db"
	"github.com/Uranury/WorkoutApp/internal/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load()

	database, err := db.InitDB("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %v", err)
	}
	defer database.Close()

	router := gin.Default()

	UserHandler := InitUserHandler(database)
	ExerciseHandler := InitExerciseHandler(database)
	WorkoutHandler := InitWorkoutHandler(database)

	router.POST("/signup", UserHandler.Signup)
	router.POST("/login", UserHandler.Login)
	router.GET("/users", UserHandler.GetUsers)
	router.GET("/exercises", ExerciseHandler.GetExercises)

	protected := router.Group("/", middleware.JWTAuth())
	protected.POST("/exercises", ExerciseHandler.CreateExercise)
	protected.POST("/workouts", WorkoutHandler.CreateWorkout)
	protected.GET("/workouts", WorkoutHandler.GetWorkouts)
	protected.GET("/workouts/exercises", WorkoutHandler.GetFullWorkout)
	protected.POST("/workouts/:workoutID/exercises", WorkoutHandler.AddExerciseToWorkout)

	log.Printf("Listening on port %s...", os.Getenv("LISTEN_ADDR"))
	if err := router.Run(os.Getenv("LISTEN_ADDR")); err != nil {
		log.Fatal("Couldn't start server")
	}
}
