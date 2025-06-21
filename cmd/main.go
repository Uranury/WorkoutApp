package main

import (
	"log"
	"os"

	"github.com/Uranury/WorkoutApp/config"
)

func main() {
	cfg := config.Load()

	server, err := NewServer(cfg)
	if err != nil {
		log.Fatalf("Couldn't create server: %v", err)
	}
	defer server.Close()

	if err := server.Start(os.Getenv("LISTEN_ADDR")); err != nil {
		log.Fatal("Couldn't start server")
	}
}
