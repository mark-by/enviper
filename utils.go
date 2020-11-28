package enviper

import (
	"os"
	"strings"
)

// getFromEnv get env variable. If variable equals empty string, then returns defaultValue
func getFromEnv(name string, defaultValue string) string {
	value := os.Getenv(formatToEnv(name))
	if value != "" {
		return value
	} else {
		return defaultValue
	}
}

// formatToEnv format variable name "app.host" to "APP_HOST"
func formatToEnv(name string) string {
	return strings.ReplaceAll(strings.ToUpper(name), ".", "_")
}

