package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	log.Printf("Listening on port %s...", os.Getenv("LISTEN_ADDR"))
	if err := router.Run(os.Getenv("LISTEN_ADDR")); err != nil {
		log.Fatal("Couldn't start server")
	}
}
