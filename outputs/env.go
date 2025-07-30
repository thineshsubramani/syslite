package outputs

import (
	"fmt"
	"os"
	"strings"
)

func writeENV(data map[string]interface{}, path string) error {
	var builder strings.Builder
	for k, v := range data {
		builder.WriteString(fmt.Sprintf("%s=%v\n", strings.ToUpper(k), v))
	}
	return os.WriteFile(path, []byte(builder.String()), 0644)
}

func printENV(data map[string]interface{}) error {
	for k, v := range data {
		fmt.Printf("%s=%v\n", strings.ToUpper(k), v)
	}
	return nil
}
