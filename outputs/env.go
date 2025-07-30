package outputs

import (
	"fmt"
	"os"
)

func ToEnv(data map[string]interface{}, targetPath string) error {
	f, err := os.Create(targetPath)
	if err != nil {
		return err
	}
	defer f.Close()

	for k, v := range data {
		_, err := fmt.Fprintf(f, "export %s=%v\n", k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
