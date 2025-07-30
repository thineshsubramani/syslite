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

	data := extract.CollectAll(cfg.Extract)

	for _, format := range cfg.Output.Formats {
		switch format {
		case "json":
			outputs.ToJSON(data, cfg.Output.JSONPath)
		case "yaml":
			outputs.ToYAML(data, cfg.Output.YAMLPath)
		case "env":
			outputs.ToEnv(data, cfg.Output.EnvTarget)
		}
	}
}
