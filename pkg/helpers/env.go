package helpers

import "os"

func EnvWithDefault(key, defaultValue string) string {

	strVal := os.Getenv(key)
	if strVal == "" {
		return defaultValue
	}
	return strVal
}
