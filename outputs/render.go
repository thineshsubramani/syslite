package outputs

import (
	"fmt"
	"os"
)

type OutputFormat struct {
	Path   string
	Stdout bool
}

func Render(data map[string]interface{}, formats map[string]OutputFormat) error {
	for format, cfg := range formats {
		switch format {
		case "json":
			if cfg.Path != "" {
				if err := writeJSON(data, cfg.Path); err != nil {
					return err
				}
			}
			if cfg.Stdout {
				if err := printJSON(data); err != nil {
					return err
				}
			}

		case "yaml":
			if cfg.Path != "" {
				if err := writeYAML(data, cfg.Path); err != nil {
					return err
				}
			}
			if cfg.Stdout {
				if err := printYAML(data); err != nil {
					return err
				}
			}

		case "env":
			if cfg.Path != "" {
				if err := writeENV(data, cfg.Path); err != nil {
					return err
				}
			}
			if cfg.Stdout {
				if err := printENV(data); err != nil {
					return err
				}
			}

		default:
			fmt.Fprintf(os.Stderr, "unknown format: %s\n", format)
		}
	}
	return nil
}
