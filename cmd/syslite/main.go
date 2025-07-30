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

	for _, format := range cfg.Output.Formats {
		switch format {
		case "json":
			if err := outputs.ToJSON(data, cfg.Output.JSONPath); err != nil {
				log.Printf("json output error: %v", err)
			}
		case "env":
			if err := outputs.ToEnv(data, cfg.Output.EnvTarget); err != nil {
				log.Printf("env output error: %v", err)
			}
		default:
			log.Printf("unsupported format: %s", format)
		}
	}
}
