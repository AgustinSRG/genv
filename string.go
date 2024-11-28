// Environment variables as string

package genv

import (
	"os"
)

// Gets environment variable as string.
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the value of the variable.
// If the variable is not set, the default value will be returned.
func GetEnvString(variableName string, defaultValue string) string {
	v, exists := os.LookupEnv(variableName)

	if !exists {
		v = defaultValue
	}

	return v
}

// Gets environment variable as string.
// If the variable is set, but empty, the default value is returned.
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// Returns the value of the variable.
// If the variable is not set, or it is empty, the default value will be returned.
func GetEnvStringNonEmpty(variableName string, defaultValue string) string {
	v := os.Getenv(variableName)

	if v == "" {
		v = defaultValue
	}

	return v
}
