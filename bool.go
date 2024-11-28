// Environment variables as boolean

package genv

import (
	"os"
	"strings"
)

// Gets environment variable as boolean.
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// The value for the variable must be (case insensitive):
//   - "TRUE" or "YES" to be true
//   - "FALSE" or "NO" to be false
//
// Returns the parsed value of the variable.
// If empty, or the value does not match the format, the default value will be returned.
func GetEnvBool(variableName string, defaultValue bool) bool {
	v, _ := GetEnvBoolWithWarning(variableName, defaultValue)
	return v
}

// Gets environment variable as boolean.
// Also gets a warning flag if the value is invalid
//
// Parameters
//   - variableName: Name of the environment variable
//   - defaultValue: Default value
//
// The value for the variable must be (case insensitive):
//   - "TRUE" or "YES" to be true
//   - "FALSE" or "NO" to be false
//
// Returns
//   - val - The parsed value or the default value (if empty or cannot be parsed)
//   - warning - True only if the value is set, but has invalid format
func GetEnvBoolWithWarning(variableName string, defaultValue bool) (val bool, warning bool) {
	v := strings.ToUpper(os.Getenv(variableName))

	if v == "YES" || v == "TRUE" {
		return true, false
	} else if v == "NO" || v == "FALSE" {
		return false, false
	} else if v == "" {
		return defaultValue, false
	} else {
		return defaultValue, true
	}
}
