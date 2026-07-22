package util

import (
	"os"
	"strconv"
)

// Get environment variable with default value.
func GetEnv(name string, defaultValue string) string {
	value, found := os.LookupEnv(name)
	if !found {
		return defaultValue
	}
	return value
}

// Get integer environment variable with default value.
func GetEnvInt(name string, defaultValue int) int {
	value, found := os.LookupEnv(name)
	if !found {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return intValue
}
