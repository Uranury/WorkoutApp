package main

import (
	"log"
	"os"

	"github.com/Uranury/WorkoutApp/config"
	"github.com/Uranury/WorkoutApp/internal/api"
	"github.com/Uranury/WorkoutApp/internal/db"
	"github.com/Uranury/WorkoutApp/internal/repositories"
	"github.com/Uranury/WorkoutApp/internal/services"
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

	newUserRepo := repositories.NewUserRepository(database)
	newUserService := services.NewUserService(newUserRepo)
	newUserHandler := api.NewUserHandler(newUserService)

	router.POST("/signup", newUserHandler.Signup)
	router.POST("/login", newUserHandler.Login)

	log.Printf("Listening on port %s...", os.Getenv("LISTEN_ADDR"))
	if err := router.Run(os.Getenv("LISTEN_ADDR")); err != nil {
		log.Fatal("Couldn't start server")
	}
}
