package helper

import (
	"os"
	"regexp"
)

// ReplaceEnvInString regex for this
func ReplaceEnvInString(text string) string {
	re := regexp.MustCompile(`\$[a-zA-Z\_]+`)
	stringWithEnvs := re.ReplaceAllStringFunc(text, func(envRegex string) string {
		osEnv := os.Getenv(envRegex[1:])

		if osEnv != "" {
			return osEnv
		}

		return envRegex
	})
	return stringWithEnvs
}
