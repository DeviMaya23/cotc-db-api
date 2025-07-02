package helpers

import (
	"os"
	"strings"
)

func EnvWithDefault(key, defaultValue string) string {

	strVal := os.Getenv(key)
	if strVal == "" {
		return defaultValue
	}
	return strVal
}

func EnvWithDefaultBool(key string, defaultValue bool) bool {

	strVal := os.Getenv(key)
	if strVal == "" {
		return defaultValue
	}

	return strings.EqualFold(strVal, "true")
}
