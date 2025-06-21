package main

import (
	"log"

	"github.com/Uranury/WorkoutApp/config"
	"github.com/Uranury/WorkoutApp/internal/db"
	"github.com/Uranury/WorkoutApp/internal/handlers"
	"github.com/Uranury/WorkoutApp/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Server struct {
	router          *gin.Engine
	db              *sqlx.DB
	userHandler     *handlers.UserHandler
	exerciseHandler *handlers.ExerciseHandler
	workoutHandler  *handlers.WorkoutHandler
	config          *config.Config
}

func NewServer(cfg *config.Config) (*Server, error) {
	database, err := db.InitDB("postgres", cfg.DbUrl)
	if err != nil {
		return nil, err
	}

	server := &Server{
		router: gin.Default(),
		db:     database,
		config: cfg,
	}

	server.userHandler = InitUserHandler(database)
	server.exerciseHandler = InitExerciseHandler(database)
	server.workoutHandler = InitWorkoutHandler(database)

	server.setupRoutes()

	return server, nil
}

func (s *Server) setupRoutes() {
	// Public routes
	s.router.POST("/signup", s.userHandler.Signup)
	s.router.POST("/login", s.userHandler.Login)
	s.router.GET("/users", s.userHandler.GetUsers)
	s.router.GET("/exercises", s.exerciseHandler.GetExercises)

	// Protected routes
	protected := s.router.Group("/", middleware.JWTAuth())
	protected.POST("/exercises", s.exerciseHandler.CreateExercise)
	protected.POST("/workouts", s.workoutHandler.CreateWorkout)
	protected.GET("/workouts", s.workoutHandler.GetWorkouts)
	protected.GET("/workouts/exercises", s.workoutHandler.GetFullWorkout)
	protected.POST("/workouts/:workoutID/exercises", s.workoutHandler.AddExerciseToWorkout)
}

func (s *Server) Start(addr string) error {
	log.Printf("Listening on port %s...", addr)
	return s.router.Run(addr)
}

func (s *Server) Close() error {
	return s.db.Close()
}
