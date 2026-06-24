// @title Job Radar API
// @version 1.0
// @description API for Job Radar
// @host localhost:8080
// @BasePath /

package main

import (
	"log"

	_ "github.com/anggasspm/job-radar/backend/docs"

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
