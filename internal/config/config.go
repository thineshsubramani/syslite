package config

import (
	"os"

	"github.com/thineshsubramani/syslite/outputs"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Extract []string     `yaml:"extract"`
	Output  OutputConfig `yaml:"output"`
}

type OutputConfig struct {
	Formats map[string]outputs.OutputFormat `yaml:"formats"`
}

type OutputFormat struct {
	Path   string `yaml:"path"`
	Stdout bool   `yaml:"stdout"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
