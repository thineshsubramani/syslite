package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Extract []string `yaml:"extract"`
	Output  struct {
		Formats   []string `yaml:"formats"`
		JSONPath  string   `yaml:"json_path"`
		YAMLPath  string   `yaml:"yaml_path"`
		EnvTarget string   `yaml:"env_target"`
	} `yaml:"output"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
