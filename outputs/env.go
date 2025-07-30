package outputs

import (
	"fmt"
	"os"
	"strings"
)

func flatten(data map[string]interface{}) map[string]string {
	flat := make(map[string]string)

	var walker func(prefix string, val interface{})
	walker = func(prefix string, val interface{}) {
		switch v := val.(type) {
		case map[string]interface{}:
			for k, sub := range v {
				walker(prefix+"_"+strings.ToUpper(k), sub)
			}
		default:
			flat[strings.TrimPrefix(prefix, "_")] = fmt.Sprintf("%v", v)
		}
	}

	for k, v := range data {
		walker(strings.ToUpper(k), v)
	}
	return flat
}

func ToEnv(data map[string]interface{}, target string) error {
	switch target {
	case "GITHUB_ENV":
		path := os.Getenv("GITHUB_ENV")
		if path == "" {
			return fmt.Errorf("GITHUB_ENV path not set")
		}
		file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		for k, v := range flatten(data) {
			fmt.Fprintf(file, "%s=%s\n", k, v)
		}
	default:
		for k, v := range flatten(data) {
			_ = os.Setenv(k, v)
		}
	}
	return nil
}
