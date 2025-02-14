package main

import (
	"log"

	"github.com/ochinchind/docsproc/config"
	"github.com/ochinchind/docsproc/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
