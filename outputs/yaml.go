package outputs

import (
	"os"

	"gopkg.in/yaml.v3"
)

func writeYAML(data map[string]interface{}, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := yaml.NewEncoder(f)
	return encoder.Encode(data)
}

func printYAML(data map[string]interface{}) error {
	encoder := yaml.NewEncoder(os.Stdout)
	return encoder.Encode(data)
}
