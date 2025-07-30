package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type OutputConfig struct {
	Formats   []string `yaml:"formats"`
	JSONPath  string   `yaml:"json_path"`
	YAMLPath  string   `yaml:"yaml_path"`
	EnvTarget string   `yaml:"env_target"`
}

type Config struct {
	Extract []string     `yaml:"extract"`
	Output  OutputConfig `yaml:"output"`
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
