package main

import (
	"log"

	"github.com/anggasspm/job-radar/backend/config"
	"github.com/anggasspm/job-radar/backend/internal/api"
)

func main() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("Config file is not loaded properly %v\n", err)
	}

	log.Println("Starting application...")

	api.StartServer(cfg)

}
