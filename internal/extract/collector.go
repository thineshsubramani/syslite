package extract

import (
	"strings"

	"github.com/thineshsubramani/syslite/internal/config"
)

func CollectAll(cfg config.Config) map[string]string {
	result := make(map[string]string)

	for _, section := range cfg.Extract {
		switch strings.ToLower(section) {
		case "os":
			for k, v := range GetOS() {
				result[k] = v
			}
			// Add more later like:
			// case "kernel":
			// case "distro":
		}
	}

	return result
}
