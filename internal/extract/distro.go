package extract

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func GetDistro() map[string]interface{} {
	distroInfo := make(map[string]interface{})

	switch runtime.GOOS {
	case "linux":
		distroInfo["distro"] = getLinuxDistro()
	case "windows":
		distroInfo["distro"] = getWindowsDistro()
	default:
		distroInfo["distro"] = "unknown"
	}

	return distroInfo
}

func getLinuxDistro() string {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "Unknown Linux"
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "PRETTY_NAME=") {
			return strings.Trim(line[13:], `"`)
		}
	}
	return "Unknown Linux"
}

func getWindowsDistro() string {
	cmd := exec.Command("cmd", "/C", "ver")
	output, err := cmd.Output()
	if err != nil {
		return "Unknown Windows"
	}
	return strings.TrimSpace(string(output))
}
