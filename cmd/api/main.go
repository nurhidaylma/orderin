package main

import (
	"log"

	"github.com/nurhidaylma/orderin/internal/config"
)

func main() {
	cfg := config.Load()
	log.Println("starting goorder api on port", cfg.AppPort)

	// init db
	// init redis
	// init repositories
	// init services
	// init handlers
	// start server
}
