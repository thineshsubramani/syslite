package extract

import (
	"fmt"
)

func CollectAll(keys []string) (map[string]interface{}, error) {
	result := make(map[string]interface{})

	for _, key := range keys {
		switch key {
		case "os":
			for k, v := range GetOS() {
				result[k] = v
			}
		default:
			// You can replace this mock with real logic later
			result[key] = fmt.Sprintf("mocked_%s_value", key)
		}
	}

	return result, nil
}
