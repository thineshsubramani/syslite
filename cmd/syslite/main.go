package main

import (
	"log"

	"github.com/thineshsubramani/syslite/internal/config"
	"github.com/thineshsubramani/syslite/internal/extract"
	"github.com/thineshsubramani/syslite/outputs"
)

func main() {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	data, err := extract.CollectAll(cfg.Extract)
	if err != nil {
		log.Fatalf("failed to extract data: %v", err)
	}

	if err := outputs.Render(data, cfg.Output.Formats); err != nil {
		log.Fatalf("failed to render outputs: %v", err)
	}
}
