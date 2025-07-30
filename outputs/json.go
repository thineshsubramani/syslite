package outputs

import (
	"encoding/json"
	"log"
	"os"
)

func ToJSON(data map[string]interface{}, path string) {
	file, err := os.Create(path)
	if err != nil {
		log.Printf("failed to write JSON: %v", err)
		return
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	err = enc.Encode(data)
	if err != nil {
		log.Printf("failed to encode JSON: %v", err)
	}
}
