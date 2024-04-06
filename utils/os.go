package utils

import (
	"os"
	"strings"
)

var OsId = ""

func LoadOSInfo() {
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		panic(err)
	}

	releaseInfo := string(data)
	OsId = parseLines(releaseInfo)["ID"]
}

func parseLines(data string) map[string]string {
	lines := strings.Split(data, "\n")
	returnMap := make(map[string]string)

	for _, line := range lines {
		if line != "" {
			parts := strings.Split(line, "=")
			value := parts[1]
			if strings.Contains(value, "\"") {
				value = strings.ReplaceAll(value, "\"", "")
			}

			returnMap[parts[0]] = value
		}
	}

	return returnMap
}
