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
		case "distro":
			for k, v := range GetDistro() {
				result[k] = v
			}
		default:
			result[key] = fmt.Sprintf("mocked_%s_value", key)
		}
	}

	return result, nil
}
